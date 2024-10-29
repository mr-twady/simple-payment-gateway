package migrations

import (
	// "fmt"
	"gorm.io/gorm"

	"dunsin-olubobokun/simple-payment-gateway/internal/models"
	"dunsin-olubobokun/simple-payment-gateway/internal/utils"
)

// Migrate will create the necessary tables and seed initial data
func Migrate(db *gorm.DB) error {
	// Migrate the schema
	if err := db.AutoMigrate(&models.Transaction{}); err != nil {
		return err
	}

	if err := seedData(db); err != nil {
		return err
	}

	return nil
}

func seedData(db *gorm.DB) error {
	var transactions []models.Transaction
	if err := db.Find(&transactions).Error; err != nil {
		return err
	}

	// Only seed data if the table is empty
	if len(transactions) == 0 {
		initialTransactions := []models.Transaction{
			{Type: "deposit", Amount: 100.0, Currency: "USD", Email: "test@test.com", CustomerReference: utils.GenerateSimpleRandomString(), Status: models.StatusCompleted},
			{Type: "withdrawal", Amount: 50.0, Currency: "USD", Email: "test@test.com", CustomerReference: utils.GenerateSimpleRandomString(), Status: models.StatusCompleted},
		}

		if err := db.Create(&initialTransactions).Error; err != nil {
			return err
		}
	}

	return nil
}
