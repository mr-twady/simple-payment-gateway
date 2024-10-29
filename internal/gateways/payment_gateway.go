package gateways

import "dunsin-olubobokun/simple-payment-gateway/internal/models"

type PaymentGateway interface {
	InitiateDeposit(transaction *models.Transaction) error
	VerifyDeposit(req *models.TransactionRequest) (*models.Transaction, error)
	ProcessWithdrawal(transaction *models.Transaction) error
	HandleCallback(data []byte) (*models.TransactionStatus, error)
}
