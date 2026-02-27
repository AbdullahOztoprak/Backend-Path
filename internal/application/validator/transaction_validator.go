package validator

import (
	"errors"
	"regexp"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
)

var (
	errInvalidTransactionAmount = errors.New("transaction amount must be greater than zero")
	errInvalidTransactionDescription = errors.New("transaction description cannot be empty")
	errInvalidUserID = errors.New("user ID must be a valid positive integer")
)

// TransactionValidator validates transaction data.
type TransactionValidator struct{}

// NewTransactionValidator creates a new instance of TransactionValidator.
func NewTransactionValidator() *TransactionValidator {
	return &TransactionValidator{}
}

// Validate validates the transaction entity.
func (v *TransactionValidator) Validate(transaction *entity.Transaction) error {
	if transaction.Amount <= 0 {
		return errInvalidTransactionAmount
	}
	if transaction.Description == "" {
		return errInvalidTransactionDescription
	}
	if !isValidUserID(transaction.FromUserID) || !isValidUserID(transaction.ToUserID) {
		return errInvalidUserID
	}
	return nil
}

// isValidUserID checks if the user ID is a valid positive integer.
func isValidUserID(userID int) bool {
	return userID > 0
}