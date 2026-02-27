package service

import (
	"context"
	"errors"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/pkg/apperror"
)

type BalanceService struct {
	balanceRepo repository.BalanceRepository
}

func NewBalanceService(balanceRepo repository.BalanceRepository) *BalanceService {
	return &BalanceService{balanceRepo: balanceRepo}
}

func (s *BalanceService) GetBalance(ctx context.Context, userID int) (*entity.Balance, error) {
	balance, err := s.balanceRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, apperror.NewNotFoundError("balance not found")
	}
	return balance, nil
}

func (s *BalanceService) UpdateBalance(ctx context.Context, userID int, amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}

	balance, err := s.balanceRepo.FindByUserID(ctx, userID)
	if err != nil {
		return apperror.NewNotFoundError("balance not found")
	}

	balance.Amount += amount
	if err := s.balanceRepo.Update(ctx, balance); err != nil {
		return err
	}
	return nil
}

func (s *BalanceService) TransferFunds(ctx context.Context, fromUserID, toUserID int, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	fromBalance, err := s.balanceRepo.FindByUserID(ctx, fromUserID)
	if err != nil {
		return apperror.NewNotFoundError("from user balance not found")
	}

	if fromBalance.Amount < amount {
		return errors.New("insufficient funds")
	}

	toBalance, err := s.balanceRepo.FindByUserID(ctx, toUserID)
	if err != nil {
		return apperror.NewNotFoundError("to user balance not found")
	}

	fromBalance.Amount -= amount
	toBalance.Amount += amount

	if err := s.balanceRepo.Update(ctx, fromBalance); err != nil {
		return err
	}
	if err := s.balanceRepo.Update(ctx, toBalance); err != nil {
		return err
	}
	return nil
}