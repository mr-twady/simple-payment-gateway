package utils

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"dunsin-olubobokun/simple-payment-gateway/internal/models"
)

func GenerateSimpleRandomString() string {
	rand.Seed(time.Now().UnixNano())
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	randomChar := rand.Intn(100)

	return timestamp + "-" + strconv.Itoa(randomChar)
}

func MapTransactionStatus(transaction *models.Transaction, status models.TransactionStatus) error {
	switch status {
	case "completed":
		transaction.Status = models.StatusCompleted
	case "pending":
		transaction.Status = models.StatusPending
	case "processing":
		transaction.Status = models.StatusProcessing
	case "failed":
		transaction.Status = models.StatusFailed
	default:
		return errors.New("invalid status")
	}
	return nil
}

func ConvertTransactionRequestToMap(req *models.TransactionRequest) map[string]interface{} {
	return map[string]interface{}{
		"customer_reference": req.CustomerReference,
		"status":             req.Status,
	}
}
