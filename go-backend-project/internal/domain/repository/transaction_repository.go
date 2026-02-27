package repository

import (
	"context"

	"go-backend-project/internal/domain/entity"
)

// TransactionRepository defines the interface for transaction data operations.
type TransactionRepository interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
	GetByID(ctx context.Context, id string) (*entity.Transaction, error)
	ListByUserID(ctx context.Context, userID string) ([]entity.Transaction, error)
	Update(ctx context.Context, transaction *entity.Transaction) error
	Delete(ctx context.Context, id string) error
}