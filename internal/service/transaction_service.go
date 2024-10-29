package service

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/repository"
)

type TransactionService struct {
	Repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {

	return &TransactionService{Repo: repo}
}
