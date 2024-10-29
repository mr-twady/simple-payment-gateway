package api

import (
	"encoding/json"
	"net/http"

	"dunsin-olubobokun/simple-payment-gateway/internal/service"
	"dunsin-olubobokun/simple-payment-gateway/internal/utils"
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

	// COULD IMPROVE: In a roubust application, we could add some other system checks e.g ping database or
	// other necessary application specific dependencies to confirm application's health status
	response := map[string]string{"message": "Payment service is up and running!"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ApiDocumentation(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "docs/openapi.yaml")
}

func (h *Handler) InitiateDepositHandler(w http.ResponseWriter, r *http.Request) {
	req, err := ValidateTransactionRequest(w, r)
	if req == nil || err != nil {
		return
	}

	transaction, err := h.Service.InitiateDeposit(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func (h *Handler) VerifyDepositHandler(w http.ResponseWriter, r *http.Request) {
	req, err := ValidateTransactionRequest(w, r)
	if req == nil || err != nil {
		return
	}

	transaction, err := h.Service.VerifyDeposit(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

func (h *Handler) WithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	req, err := ValidateTransactionRequest(w, r)
	if req == nil || err != nil {
		return
	}

	transaction, err := h.Service.ProcessWithdrawal(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func (h *Handler) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	req, err := ValidateTransactionRequest(w, r)
	if req == nil || err != nil {
		return
	}
	callbackData := utils.ConvertTransactionRequestToMap(req)
	// fmt.Println("callbackData", callbackData)

	status, err := h.Service.HandleCallback(callbackData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": string(status)})
}

func (h *Handler) GetUserBalanceHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	balance, err := h.Service.GetUserBalance(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]float64{"balance": balance}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
