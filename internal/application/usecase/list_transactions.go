package usecase

import (
	"context"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

type ListTransactionsUseCase struct {
	transactionRepo repository.TransactionRepository
}

func NewListTransactionsUseCase(repo repository.TransactionRepository) *ListTransactionsUseCase {
	return &ListTransactionsUseCase{
		transactionRepo: repo,
	}
}

func (uc *ListTransactionsUseCase) Execute(ctx context.Context, userID string) ([]entity.Transaction, error) {
	transactions, err := uc.transactionRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}