package service

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
	"errors"
)

func (s *TransactionService) VerifyDeposit(req *models.TransactionRequest) (*models.Transaction, error) {
	if req.CustomerReference == "" {
		return nil, errors.New("missing customer reference")
	}

	var transaction models.Transaction
	if err := s.Repo.FindTransactionByCustomerReference(req.CustomerReference, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}
