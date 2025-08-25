package models

import (
    "sync"
    "time"
)

type Balance struct {
    UserID        int64     `json:"user_id"`
    Amount        float64   `json:"amount"`
    LastUpdatedAt time.Time `json:"last_updated_at"`
    mu            sync.RWMutex
}

// UpdateAmount safely updates the balance amount
func (b *Balance) UpdateAmount(delta float64) {
    b.mu.Lock()
    defer b.mu.Unlock()
    b.Amount += delta
    b.LastUpdatedAt = time.Now()
}

// GetAmount safely retrieves the balance amount
func (b *Balance) GetAmount() float64 {
    b.mu.RLock()
    defer b.mu.RUnlock()
    return b.Amount
}