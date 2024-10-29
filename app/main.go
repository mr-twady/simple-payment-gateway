package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sony/gobreaker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"dunsin-olubobokun/simple-payment-gateway/internal/api"
	"dunsin-olubobokun/simple-payment-gateway/internal/config"
	"dunsin-olubobokun/simple-payment-gateway/internal/gateways"
	"dunsin-olubobokun/simple-payment-gateway/internal/middleware"
	"dunsin-olubobokun/simple-payment-gateway/internal/migrations"
	"dunsin-olubobokun/simple-payment-gateway/internal/repository"
	"dunsin-olubobokun/simple-payment-gateway/internal/service"
)

func main() {
	// load configuration
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Unable to load config:", err)
	}

	// to support some degree of resillence, I implemented a simple retry logic on DB connection
	// sometimes it might take a few seconds for the DB to be ready. Hence, the reason I implemented a simple retry logic
	var db *gorm.DB
	var dbErrx error
	for retries := 0; retries < 3; retries++ {
		db, dbErrx = gorm.Open(postgres.Open(conf.DBUrl), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second) // wait before retrying
	}
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", dbErrx)
	}

	// run migrations
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// initialize the circuit breaker
	settings := gobreaker.Settings{
		Name:        "PaymentGateway",
		MaxRequests: 1,
		Interval:    5 * time.Second,
		Timeout:     30 * time.Second, // or we could configured in env if we want to have a general rule of thumb for timeouts
	}
	cb := gobreaker.NewCircuitBreaker(settings)

	// for this soltuion, I'm using 2 mock gateways for testing
	gateways := map[string]gateways.PaymentGateway{
		"GatewayA": gateways.NewMockGatewayA(false), // could be in json
		"GatewayB": gateways.NewMockGatewayB(false), // could be in SOAP/XML
	}

	// set up repositories and services
	transactionRepo := repository.NewTransactionRepository(db)
	userRepo := repository.NewUserRepository(db)
	transactionService := service.NewTransactionService(transactionRepo, userRepo, gateways)

	// set up HTTP handlers
	handler := api.NewHandler(transactionService)
	router := mux.NewRouter()

	// grouping relevant routes with "api" subrouter and midddleware
	apiRoutes := router.PathPrefix("/api").Subrouter()
	apiRoutes.Use(middleware.CircuitBreakerMiddleware(cb))
	// COULD IMPROVE: for time constraints, ideally this should be a protected route
	apiRoutes.HandleFunc("/deposit", handler.InitiateDepositHandler).Methods("POST")
	apiRoutes.HandleFunc("/deposit/verify", handler.VerifyDepositHandler).Methods("POST")
	apiRoutes.HandleFunc("/withdrawal", handler.WithdrawalHandler).Methods("POST")
	apiRoutes.HandleFunc("/user/balance", handler.GetUserBalanceHandler).Methods("GET")

	// other routes
	router.HandleFunc("/health", handler.HealthCheck).Methods("GET")
	router.HandleFunc("/callback", handler.CallbackHandler).Methods("POST")

	// Start the server
	log.Printf("Server started on :%d", conf.HTTPPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.HTTPPort), router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
