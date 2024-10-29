package gateways

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
)

/*
   In an ideal implementatiion, this MockGatewayB will be
   a detailed implementation of an API integration using SOAP/XML over HTTP and
   will return responses in a general data format applicable to this codebase e.g in JSON format.
*/

// MockGatewayB is a mock implementation of the PaymentGateway A interface for testing purposes.
type MockGatewayB struct {
}

func NewMockGatewayB(config bool) *MockGatewayB {
	return &MockGatewayB{}
}

func (m *MockGatewayB) InitiateDeposit(transaction *models.Transaction) error {
	// For testing, I assume it always succeeds.
	return nil
}

func (m *MockGatewayB) VerifyDeposit(req *models.TransactionRequest) (*models.Transaction, error) {
	transaction := models.Transaction{
		Email:             req.Email,
		CustomerReference: req.CustomerReference,
		Amount:            req.Amount,
		Currency:          req.Currency,
		Status:            models.StatusProcessing, // Mocked status
	}
	return &transaction, nil // Return the mocked transaction
}

func (m *MockGatewayB) ProcessWithdrawal(transaction *models.Transaction) error {
	// For testing, I assume it always succeeds.
	return nil
}

func (m *MockGatewayB) HandleCallback(data []byte) (*models.TransactionStatus, error) {
	status := models.TransactionStatus("completed")
	return &status, nil
}
