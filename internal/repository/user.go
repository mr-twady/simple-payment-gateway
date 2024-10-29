package repository

import (
	// "errors"
	"dunsin-olubobokun/simple-payment-gateway/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindUserByEmail(email string, user *models.User) error {
	return r.DB.Where("email = ?", email).First(user).Error
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}
