package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"dunsin-olubobokun/simple-payment-gateway/internal/api"
	"dunsin-olubobokun/simple-payment-gateway/internal/config"
	"dunsin-olubobokun/simple-payment-gateway/internal/repository"
	"dunsin-olubobokun/simple-payment-gateway/internal/service"
)

func main() {
	// load env configuration
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Unable to load config:", err)
	}

	// can do my DB connection here
	var db *gorm.DB

	// for this soltuion, I'm using 2 mock gateways for testing
	/* gateways := map the gateways

	such as PaymentGateway{
		"GatewayA": gateways.NewMockGatewayA(false), // could be in json
		"GatewayB": gateways.NewMockGatewayB(false), // could be in SOAP/XML
	}
	} */

	// set up repositories and services
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)

	// set up HTTP handlers
	handler := api.NewHandler(transactionService)
	router := mux.NewRouter()

	router.HandleFunc("/health", handler.HealthCheck).Methods("GET")
	router.HandleFunc("/deposit", handler.InitiateDepositHandler).Methods("POST")
	router.HandleFunc("/deposit/verify", handler.VerifyDepositHandler).Methods("POST")
	router.HandleFunc("/withdrawal", handler.WithdrawalHandler).Methods("POST")
	router.HandleFunc("/callback", handler.CallbackHandler).Methods("POST")

	// Start the server
	log.Printf("Server started on :%d", conf.HTTPPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.HTTPPort), router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
