package service

import (
	"context"
	"errors"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/pkg/apperror"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	if err := transaction.Validate(); err != nil {
		return apperror.NewValidationError("Invalid transaction data", err)
	}

	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	return s.repo.Create(ctx, transaction)
}

func (s *TransactionService) GetTransactionByID(ctx context.Context, id string) (*entity.Transaction, error) {
	transaction, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, apperror.NewNotFoundError("Transaction not found")
		}
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) ListTransactions(ctx context.Context, userID string, limit, offset int) ([]entity.Transaction, error) {
	transactions, err := s.repo.FindByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	if err := transaction.Validate(); err != nil {
		return apperror.NewValidationError("Invalid transaction data", err)
	}

	transaction.UpdatedAt = time.Now()
	return s.repo.Update(ctx, transaction)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}