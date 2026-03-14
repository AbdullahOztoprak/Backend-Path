package usecase

import (
	"context"
	"errors"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

type GetBalanceUseCase struct {
    balanceRepo repository.BalanceRepository
}

func NewGetBalanceUseCase(balanceRepo repository.BalanceRepository) *GetBalanceUseCase {
    return &GetBalanceUseCase{
        balanceRepo: balanceRepo,
    }
}

func (uc *GetBalanceUseCase) Execute(ctx context.Context, userID int) (*entity.Balance, error) {
	_ = ctx

    if userID <= 0 {
        return nil, errors.New("invalid user ID")
    }

    balance, err := uc.balanceRepo.GetBalance(userID)
    if err != nil {
        return nil, err
    }

    return balance, nil
}