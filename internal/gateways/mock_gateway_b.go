package gateways

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
)

/*
   COULD IMPROVE:
   This is a gateway Mock
   In an ideal implementatiion, this MockGatewayA will be
   a detailed implementation of an API integration using SOAP/XML over HTTP and
   will return responses in a general data format applicable to this codebase e.g in JSON format.
*/

// MockGatewayB is a mock implementation of the PaymentGateway interface for testing purposes.
type MockGatewayB struct {
	// You can add fields here if needed for testing scenarios.
}

// NewMockGatewayB creates a new instance of MockGatewayB.
func NewMockGatewayB(config bool) *MockGatewayB {
	return &MockGatewayB{}
}

// InitiateDeposit simulates processing a deposit transaction.
func (m *MockGatewayB) InitiateDeposit(transaction *models.Transaction) error {
	// Simulate success or failure logic here.
	// For testing, we can assume it always succeeds.
	return nil
}

// VerifyDeposit simulates processing a deposit transaction.
func (m *MockGatewayB) VerifyDeposit(req *models.TransactionRequest) (*models.Transaction, error) {
	// Simulate verifying the deposit
	transaction := models.Transaction{
		Email:             req.Email, // Assuming you set this
		CustomerReference: req.CustomerReference,
		Amount:            req.Amount,
		Currency:          req.Currency,
		Status:            models.StatusCompleted, // Mocked status
	}
	return &transaction, nil // Return the mocked transaction
}

// ProcessWithdrawal simulates processing a withdrawal transaction.
func (m *MockGatewayB) ProcessWithdrawal(transaction *models.Transaction) error {
	// Simulate success or failure logic here.
	// For testing, we can assume it always succeeds.
	return nil
}

// HandleCallback handles callbacks for MockGatewayB.
func (m *MockGatewayB) HandleCallback(data []byte) (*models.TransactionStatus, error) {
	// Mock logic for handling callback
	status := models.TransactionStatus("completed") // Example response
	return &status, nil
}
