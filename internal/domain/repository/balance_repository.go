package repository

import "github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"

// BalanceRepository defines the interface for balance data operations.
type BalanceRepository interface {
    GetBalance(userID int) (*entity.Balance, error)
    UpdateBalance(balance *entity.Balance) error
    CreateBalance(balance *entity.Balance) error
    DeleteBalance(userID int) error
}