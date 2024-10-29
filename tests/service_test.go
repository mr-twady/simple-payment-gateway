package tests

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/gateways"
	"dunsin-olubobokun/simple-payment-gateway/internal/models"
	"dunsin-olubobokun/simple-payment-gateway/internal/repository"
	"dunsin-olubobokun/simple-payment-gateway/internal/service"
	"dunsin-olubobokun/simple-payment-gateway/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupService(t *testing.T) (*service.TransactionService, func()) {
	// Set up an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	// Run migrations
	if err := db.AutoMigrate(&models.Transaction{}, &models.User{}); err != nil {
		t.Fatalf("failed to run migrations: %v", err)
	}

	// Use mock gateways
	gateways := map[string]gateways.PaymentGateway{
		"GatewayA": gateways.NewMockGatewayA(false),
		"GatewayB": gateways.NewMockGatewayB(false),
	}

	repo := repository.NewTransactionRepository(db)
	userRepo := repository.NewUserRepository(db)
	transactionService := service.NewTransactionService(repo, userRepo, gateways)

	// Return a cleanup function to close the database connection
	return transactionService, func() {
		// Here you can add any cleanup logic if needed
		db.Exec("PRAGMA writable_schema = 1;")
		db.Exec("DELETE FROM sqlite_sequence;")
	}
}

func TestInitiateDeposit(t *testing.T) {
	transactionService, cleanup := setupService(t)
	defer cleanup() // Ensure cleanup is called after the test finishes

	req := models.TransactionRequest{
		Type:              "deposit",
		Amount:            100.0,
		Currency:          "USD",
		Email:             "test@test.com",
		CustomerReference: utils.GenerateSimpleRandomString(),
	}

	transaction, err := transactionService.InitiateDeposit(&req)
	assert.NoError(t, err)
	assert.Equal(t, "deposit", transaction.Type)
	assert.Equal(t, 100.0, transaction.Amount)
	assert.Equal(t, "USD", transaction.Currency)
	assert.Equal(t, models.StatusProcessing, transaction.Status)
}

func TestVerifyDeposit(t *testing.T) {
	transactionService, cleanup := setupService(t)
	defer cleanup()

	req := models.TransactionRequest{
		Type:              "deposit",
		Amount:            100.0,
		Currency:          "USD",
		Email:             "test@test.com",
		CustomerReference: utils.GenerateSimpleRandomString(),
	}
	transaction, err := transactionService.InitiateDeposit(&req)
	assert.NoError(t, err)

	verifiedTransaction, err := transactionService.VerifyDeposit(&req)
	assert.NoError(t, err)
	assert.Equal(t, transaction.CustomerReference, verifiedTransaction.CustomerReference)
}

// COULD IMPROVE: if there was more time, add tests cases for ProcessWithdrawal and TestHandleCallback
// Add more test cases to assert withdrawal balance is updated as expected both after confirm deposit or withdrawal
