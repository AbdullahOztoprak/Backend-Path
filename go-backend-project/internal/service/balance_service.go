package service

import "github.com/AbdullahOztoprak/go-backend-project/internal/models"

type BalanceService interface {
    GetByUserID(userID int64) (*models.Balance, error)
    UpdateAmount(userID int64, delta float64) error
    GetHistory(userID int64) ([]*models.Balance, error)
}