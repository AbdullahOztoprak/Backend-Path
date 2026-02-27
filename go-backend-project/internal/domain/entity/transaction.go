package entity

import (
	"errors"
	"time"
)

type Transaction struct {
	ID            string    `json:"id"`
	FromUserID    string    `json:"from_user_id"`
	ToUserID      string    `json:"to_user_id"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Status        string    `json:"status"` // e.g., "pending", "completed", "failed"
	IdempotencyKey string   `json:"idempotency_key"` // for ensuring idempotent transactions
}

var (
	ErrInvalidTransactionAmount = errors.New("invalid transaction amount")
	ErrInvalidTransactionStatus  = errors.New("invalid transaction status")
)

func NewTransaction(fromUserID, toUserID string, amount float64, description, idempotencyKey string) (*Transaction, error) {
	if amount <= 0 {
		return nil, ErrInvalidTransactionAmount
	}

	return &Transaction{
		FromUserID:    fromUserID,
		ToUserID:      toUserID,
		Amount:        amount,
		Description:   description,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Status:        "pending",
		IdempotencyKey: idempotencyKey,
	}, nil
}

func (t *Transaction) UpdateStatus(status string) error {
	if status != "pending" && status != "completed" && status != "failed" {
		return ErrInvalidTransactionStatus
	}
	t.Status = status
	t.UpdatedAt = time.Now()
	return nil
}