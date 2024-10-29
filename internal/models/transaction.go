package models

import (
	"time"
)

type TransactionStatus string

const (
	StatusPending    TransactionStatus = "pending"    // initial status of transaction
	StatusProcessing TransactionStatus = "processing" // when a handshake has been established with a payment gateway but we can't confirm the actual transaction status
	StatusCompleted  TransactionStatus = "completed"  // verfied a transaction as successful. For the scope of this assessment, I plan to allowed this to happen only via callback
	StatusFailed     TransactionStatus = "failed"     // transaction failed
)

type Transaction struct {
	ID                uint              `json:"-" gorm:"primaryKey;autoincrement"`
	Type              string            `json:"type"`
	Amount            float64           `json:"amount"`
	Currency          string            `json:"currency"`
	Status            TransactionStatus `json:"status"`
	CustomerReference string            `json:"customer_reference" gorm:"index;unique,not null"` // indexed and unique transaction reference from client per requests
	Email             string            `json:"email" gorm:"index,not null"`                     // a unique PI Data to identify user, indexed for efficient retrieval
	CreatedAt         time.Time         `json:"-"`
	UpdatedAt         time.Time         `json:"-"`
}

type TransactionRequest struct {
	Email             string            `json:"email"`
	CustomerReference string            `json:"customerReference"`
	Type              string            `json:"type"`
	Amount            float64           `json:"amount"`
	Currency          string            `json:"currency"`
	Status            TransactionStatus `json:"status"`
}
