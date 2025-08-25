package models

import (
    "errors"
    "time"
)

type TransactionStatus string

const (
    StatusPending   TransactionStatus = "pending"
    StatusCompleted TransactionStatus = "completed"
    StatusFailed    TransactionStatus = "failed"
)

type Transaction struct {
    ID          int64             `json:"id"`
    FromUserID  int64             `json:"from_user_id"`
    ToUserID    int64             `json:"to_user_id"`
    Amount      float64           `json:"amount"`
    Type        string            `json:"type"`
    Status      TransactionStatus `json:"status"`
    CreatedAt   time.Time         `json:"created_at"`
}

// UpdateStatus changes the transaction status
func (t *Transaction) UpdateStatus(newStatus TransactionStatus) error {
    switch newStatus {
    case StatusPending, StatusCompleted, StatusFailed:
        t.Status = newStatus
        return nil
    default:
        return errors.New("invalid transaction status")
    }
}