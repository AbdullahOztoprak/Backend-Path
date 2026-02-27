package usecase

import (
    "context"
    "errors"
    "github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
    "github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
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
    if userID <= 0 {
        return nil, errors.New("invalid user ID")
    }

    balance, err := uc.balanceRepo.GetBalanceByUserID(ctx, userID)
    if err != nil {
        return nil, err
    }

    return balance, nil
}