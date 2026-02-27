package usecase

import (
	"context"
	"errors"
	"time"

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

func (uc *ListTransactionsUseCase) Execute(ctx context.Context, userID int, startDate, endDate time.Time) ([]entity.Transaction, error) {
	if startDate.After(endDate) {
		return nil, errors.New("start date cannot be after end date")
	}

	transactions, err := uc.transactionRepo.ListByUserIDAndDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}