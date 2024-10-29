package api

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *service.TransactionService
}

func NewHandler(service *service.TransactionService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{"message": "Payment service is up and running!"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// test that all services are up
func (h *Handler) InitiateDepositHandler(w http.ResponseWriter, r *http.Request) {
	req, err := ValidateTransactionRequest(w, r)
	if req == nil || err != nil {
		return
	}

	response := map[string]string{"message": "Deposit initiated!"}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) VerifyDepositHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Deposit verified!"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) WithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "withdrwala processed!"}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "callback handled!"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
