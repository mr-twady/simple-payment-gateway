package service

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
	"dunsin-olubobokun/simple-payment-gateway/internal/utils"
	"errors"
	// "dunsin-olubobokun/simple-payment-gateway/internal/repository"
)

func (s *TransactionService) HandleCallback(data map[string]interface{}) (models.TransactionStatus, error) {
	// Get customer reference and status from callback data
	transactionID, ok := data["customer_reference"].(string) // Check type assertion
	if !ok || transactionID == "" {
		return "", errors.New("missing or invalid customer_reference in callback data")
	}

	status, ok := data["status"].(models.TransactionStatus) // Check type assertion
	if !ok || status == "" {
		return "", errors.New("missing or invalid status in callback data")
	}

	var transaction models.Transaction
	if err := s.Repo.FindTransactionByCustomerReference(transactionID, &transaction); err != nil {
		return "", err
	}

	if transaction.Status == models.StatusCompleted {
		return "", errors.New("Transaction is already confirmed")
	}

	// Update the transaction status
	if err := utils.MapTransactionStatus(&transaction, status); err != nil {
		return "", err
	}

	// Update the transaction in the repository with callback status
	if err := s.Repo.UpdateTransaction(&transaction); err != nil {
		return "", err
	}

	// Update user balance if the transaction status is completed
	if transaction.Status == models.StatusCompleted {
		var user models.User
		if err := s.UserRepo.FindUserByEmail(transaction.Email, &user); err != nil {
			return "", err // Return error if user is not found
		}

		user.Balance += transaction.Amount // Assuming Amount is a float64
		if err := s.UserRepo.UpdateUser(&user); err != nil {
			return "", err
		}
	}

	return transaction.Status, nil
}

/* func (s *TransactionService) HandleCallbackV1(data map[string]interface{}) (models.TransactionStatus, error) {
    // get customer reference and status from callback data
    transactionID, ok := data["customer_reference"].(string) // Check type assertion
    if !ok || transactionID == "" {
        return "", errors.New("missing or invalid customer_reference in callback data")
    }
    status, ok := data["status"].(models.TransactionStatus) // Check type assertion
    if !ok || status == "" {
        return "", errors.New("missing or invalid status in callback data")
    }





    return transaction.Status, nil
} */
