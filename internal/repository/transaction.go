package repository

import (
	"errors"

	"gorm.io/gorm"

	"dunsin-olubobokun/simple-payment-gateway/internal/models"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	// in order to prevent duplicate customer reference, as every request id or has to be treated uniquely from client
	var existingTransaction models.Transaction
	if err := r.DB.Where("customer_reference = ?", transaction.CustomerReference).First(&existingTransaction).Error; err == nil {
		return errors.New("transaction with the same customer reference already exists")
	}

	return r.DB.Create(transaction).Error
}

func (r *TransactionRepository) UpdateTransaction(transaction *models.Transaction) error {
	return r.DB.Save(transaction).Error
}

func (r *TransactionRepository) FindTransactionByCustomerReference(customerReference string, transaction *models.Transaction) error {
	return r.DB.Where("customer_reference = ?", customerReference).First(transaction).Error
}
