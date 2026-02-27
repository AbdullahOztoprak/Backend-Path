package entity

import (
	"errors"
)

type Balance struct {
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
}

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidAmount     = errors.New("invalid amount")
)

// NewBalance creates a new Balance instance with validation.
func NewBalance(userID int, amount float64, currency string) (*Balance, error) {
	if amount < 0 {
		return nil, ErrInvalidAmount
	}
	return &Balance{
		UserID:   userID,
		Amount:   amount,
		Currency: currency,
	}, nil
}

// Deposit adds an amount to the balance.
func (b *Balance) Deposit(amount float64) error {
	if amount < 0 {
		return ErrInvalidAmount
	}
	b.Amount += amount
	return nil
}

// Withdraw subtracts an amount from the balance.
func (b *Balance) Withdraw(amount float64) error {
	if amount < 0 {
		return ErrInvalidAmount
	}
	if b.Amount < amount {
		return ErrInsufficientFunds
	}
	b.Amount -= amount
	return nil
}