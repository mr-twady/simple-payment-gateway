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

	// For the scope of this assessment and time, I would leave this like this

	// COULD IMPROVE: based on gateway, I could add here an external API/Gateway call to verify transaction status
	// i.e  based on some business logic and known gatewaysToTry e.g {"GatewayA", "GatewayB"}
	// I could do gatewayA.VerifyDeposit(req) to confirm transaction status
	//

	return &transaction, nil
}
