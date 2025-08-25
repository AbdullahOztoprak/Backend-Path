package service

import "github.com/AbdullahOztoprak/go-backend-project/internal/models"

type TransactionService interface {
    Create(tx *models.Transaction) error
    GetByID(id int64) (*models.Transaction, error)
    ListByUser(userID int64) ([]*models.Transaction, error)
    UpdateStatus(id int64, status models.TransactionStatus) error
}