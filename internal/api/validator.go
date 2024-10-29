package api

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
	"encoding/json"
	"net/http"
)

// COULD IMPROVE: if there was more time, there are definitely other polished alternatives to handle request validation

// Simple reusable validator I'm using to validate the request method and decode the transaction request
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

	if req.Type == "deposit" || req.Type == "withdrawal" {
		if req.Email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return nil, nil
		}
		if req.CustomerReference == "" {
			http.Error(w, "Customer reference is required", http.StatusBadRequest)
			return nil, nil
		}
		if req.Type == "" {
			http.Error(w, "Transaction type is required", http.StatusBadRequest)
			return nil, nil
		}
		if req.Amount <= 0 {
			http.Error(w, "Amount must be greater than zero", http.StatusBadRequest)
			return nil, nil
		}
		if req.Currency == "" {
			http.Error(w, "currency is required", http.StatusBadRequest)
			return nil, nil
		}
	}

	return &req, nil
}
