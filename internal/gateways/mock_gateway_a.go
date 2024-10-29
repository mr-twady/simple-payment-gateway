package gateways

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
)

/*
   In an ideal implementatiion, this MockGatewayA will be
   a detailed implementation of an API integration uing using JSON over HTTP.
*/

// MockGatewayA is a mock implementation of the PaymentGateway A interface for testing purposes.
type MockGatewayA struct {
}

func NewMockGatewayA(config bool) *MockGatewayA {
	return &MockGatewayA{}
}

func (m *MockGatewayA) InitiateDeposit(transaction *models.Transaction) error {
	// For testing, I assume it always succeeds.
	return nil
}

func (m *MockGatewayA) VerifyDeposit(req *models.TransactionRequest) (*models.Transaction, error) {
	transaction := models.Transaction{
		Email:             req.Email,
		CustomerReference: req.CustomerReference,
		Amount:            req.Amount,
		Currency:          req.Currency,
		Status:            models.StatusProcessing, // Mocked status
	}
	return &transaction, nil // Return the mocked transaction
}

func (m *MockGatewayA) ProcessWithdrawal(transaction *models.Transaction) error {
	// For testing, I assume it always succeeds.
	return nil
}

func (m *MockGatewayA) HandleCallback(data []byte) (*models.TransactionStatus, error) {
	status := models.TransactionStatus("completed")
	return &status, nil
}
