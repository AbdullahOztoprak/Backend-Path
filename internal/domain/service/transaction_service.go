package service

import (
	"context"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	if transaction.Amount <= 0 {
		return entity.ErrInvalidTransactionAmount
	}

	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	return s.repo.Create(ctx, transaction)
}

func (s *TransactionService) GetTransactionByID(ctx context.Context, id string) (*entity.Transaction, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TransactionService) ListTransactions(ctx context.Context, userID string, limit, offset int) ([]entity.Transaction, error) {
	_, _ = limit, offset
	transactions, err := s.repo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	if transaction.Amount <= 0 {
		return entity.ErrInvalidTransactionAmount
	}

	transaction.UpdatedAt = time.Now()
	return s.repo.Update(ctx, transaction)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}