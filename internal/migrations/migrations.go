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
	if err := db.AutoMigrate(&models.User{}, &models.Transaction{}); err != nil {
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

	var users []models.User
	if err := db.Find(&transactions).Error; err != nil {
		return err
	}

	// Only seed users if the table is empty
	if len(users) == 0 {
		initialUsers := []models.User{
			{Name: "Dunsin Tester", Email: "test@test.com", Balance: 500.0, Address: "Dubai"},
		}

		// Check for existing user by email before inserting
		for _, user := range initialUsers {
			var existingUser models.User
			if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
				// If no user exists, create the new user
				if err == gorm.ErrRecordNotFound {
					if err := db.Create(&user).Error; err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
