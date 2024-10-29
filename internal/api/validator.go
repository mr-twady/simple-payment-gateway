package api

import (
	"encoding/json"
	"net/http"

	"dunsin-olubobokun/simple-payment-gateway/internal/models"
)

func ValidateTransactionRequest(w http.ResponseWriter, r *http.Request) (*models.TransactionRequest, error) {
	// since all relevant apis makes use of POST http request method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return nil, nil
	}

	var req models.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return nil, err
	}

	return &req, nil
}
