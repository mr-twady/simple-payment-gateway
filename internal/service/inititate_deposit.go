package service

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
	"errors"
)

func (s *TransactionService) InitiateDeposit(req *models.TransactionRequest) (*models.Transaction, error) {
	if req.Type != "deposit" || req.Amount <= 0 {
		return nil, errors.New("invalid deposit request")
	}

	transaction := models.Transaction{
		Type:              req.Type,
		Amount:            req.Amount,
		Currency:          req.Currency,
		Email:             req.Email,
		CustomerReference: req.CustomerReference,
		Status:            models.StatusPending, // initial transaction status
	}

	if err := s.Repo.CreateTransaction(&transaction); err != nil {
		return nil, err
	}

	// try run a retry worker logic or CB to attempt available gateways
	return &transaction, errors.New("all gateways failed to initiate deposit")
}
