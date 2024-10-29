package service

import (
	"context"
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
	"dunsin-olubobokun/simple-payment-gateway/internal/utils"
	"errors"
	"time"
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
		Status:            models.StatusPending,
	}

	if err := s.Repo.CreateTransaction(&transaction); err != nil {
		return nil, err
	}

	// attempt to process the deposit through multiple gateways
	gatewaysToTry := []string{"GatewayA", "GatewayB"}

	for _, gatewayKey := range gatewaysToTry {
		gateway, exists := s.Gateways[gatewayKey]
		if !exists {
			return nil, errors.New("unsupported gateway")
		}

		// use the circuit breaker and retry logic to call the payment gateway
		err := utils.CallGatewayWithRetry(context.Background(), 3, 2*time.Second, func(ctx context.Context) error {
			_, err := s.cb.Execute(func() (interface{}, error) {
				return nil, gateway.InitiateDeposit(&transaction)
			})
			return err
		})

		// For the scope of this assement, I would only update users baalnce after I confirm either by callback or verify,
		/// kindly refer to process_callback for details on user balance update
		if err == nil {
			transaction.Status = models.StatusProcessing // NB: Updates from pending to procesing
			s.Repo.UpdateTransaction(&transaction)
			return &transaction, nil
		}

		transaction.Status = models.StatusFailed
		s.Repo.UpdateTransaction(&transaction)
	}

	return &transaction, errors.New("all gateways failed to process the deposit")
}
