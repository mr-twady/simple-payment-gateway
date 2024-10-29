package service

import (
	"context"
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
	"dunsin-olubobokun/simple-payment-gateway/internal/repository"
	"dunsin-olubobokun/simple-payment-gateway/internal/utils"
	"errors"
	"time"
)

func (s *TransactionService) ProcessWithdrawal(req *models.TransactionRequest) (*models.Transaction, error) {
	if req.Type != "withdrawal" || req.Amount <= 0 {
		return nil, errors.New("invalid withdrawal request")
	}

	var user models.User
	userRepo := repository.NewUserRepository(s.Repo.DB) // Assuming Repo has a DB field for GORM
	if err := userRepo.FindUserByEmail(req.Email, &user); err != nil {
		return nil, errors.New("user not found")
	}

	// Check if user has sufficient balance
	if user.Balance < req.Amount {
		return nil, errors.New("insufficient balance")
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

	// attempt to process the withdrawal through multiple gateways
	gatewaysToTry := []string{"GatewayA", "GatewayB"}

	for _, gatewayKey := range gatewaysToTry {
		gateway, exists := s.Gateways[gatewayKey]
		if !exists {
			return nil, errors.New("unsupported gateway")
		}

		// Use the circuit breaker and retry logic to call the payment gateway
		err := utils.CallGatewayWithRetry(context.Background(), 3, 2*time.Second, func(ctx context.Context) error {
			_, err := s.cb.Execute(func() (interface{}, error) {
				return nil, gateway.ProcessWithdrawal(&transaction)
			})
			return err
		})

		if err == nil {
			transaction.Status = models.StatusCompleted
			s.Repo.UpdateTransaction(&transaction)

			// update user balance
			user.Balance = user.Balance - transaction.Amount
			if err := userRepo.UpdateUser(&user); err != nil {
				return nil, err
			}

			return &transaction, nil
		}

		transaction.Status = models.StatusFailed
		s.Repo.UpdateTransaction(&transaction)
	}

	return &transaction, errors.New("all gateways failed to process the withdrawal")
}
