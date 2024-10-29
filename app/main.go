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

	// Open the database connection
	db, err := gorm.Open(postgres.Open(conf.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// run migrations
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// to support some API resillence as highlighted in the project brief
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
	// COULD IMPROVE: ideally this should be a protected routes but limited cos of time constraints,
	apiRoutes.HandleFunc("/deposit", handler.InitiateDepositHandler).Methods("POST")
	apiRoutes.HandleFunc("/deposit/verify", handler.VerifyDepositHandler).Methods("POST")
	apiRoutes.HandleFunc("/withdrawal", handler.WithdrawalHandler).Methods("POST")
	apiRoutes.HandleFunc("/callback", handler.CallbackHandler).Methods("POST")

	// other routes
	router.HandleFunc("/health", handler.HealthCheck).Methods("GET")
	router.HandleFunc("/api/user/balance", handler.GetUserBalanceHandler).Methods("GET")
	// api documentation in OpenAPI format (Could alternatively use Swagger docs and annotations above my http handlers)
	router.HandleFunc("/docs", handler.ApiDocumentation).Methods("GET")

	// Start the server
	log.Printf("Server started on :%d", conf.HTTPPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.HTTPPort), router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
