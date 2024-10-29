package service

import (
	"dunsin-olubobokun/simple-payment-gateway/internal/gateways"
	"dunsin-olubobokun/simple-payment-gateway/internal/repository"
	"time"

	"github.com/sony/gobreaker"
)

type TransactionService struct {
	Repo     *repository.TransactionRepository
	Gateways map[string]gateways.PaymentGateway
	cb       *gobreaker.CircuitBreaker
}

// NewTransactionService initializes the TransactionService with a circuit breaker
func NewTransactionService(repo *repository.TransactionRepository, gateways map[string]gateways.PaymentGateway) *TransactionService {
	settings := gobreaker.Settings{
		Name:        "PaymentGateway",
		MaxRequests: 1,
		Interval:    5 * time.Second,
		Timeout:     30 * time.Second, // could be set in env
	}
	cb := gobreaker.NewCircuitBreaker(settings)

	return &TransactionService{Repo: repo, Gateways: gateways, cb: cb}
}
