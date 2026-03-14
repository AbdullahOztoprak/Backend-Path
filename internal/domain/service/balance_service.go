package service

import (
	"context"
	"errors"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

type BalanceService struct {
	balanceRepo repository.BalanceRepository
}

func NewBalanceService(balanceRepo repository.BalanceRepository) *BalanceService {
	return &BalanceService{balanceRepo: balanceRepo}
}

func (s *BalanceService) GetBalance(ctx context.Context, userID int) (*entity.Balance, error) {
	_ = ctx
	balance, err := s.balanceRepo.GetBalance(userID)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (s *BalanceService) UpdateBalance(ctx context.Context, userID int, amount float64) error {
	if amount < 0 {
		return errors.New("amount cannot be negative")
	}

	_ = ctx
	balance, err := s.balanceRepo.GetBalance(userID)
	if err != nil {
		return err
	}

	balance.Amount += amount
	if err := s.balanceRepo.UpdateBalance(balance); err != nil {
		return err
	}
	return nil
}

func (s *BalanceService) TransferFunds(ctx context.Context, fromUserID, toUserID int, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	_ = ctx
	fromBalance, err := s.balanceRepo.GetBalance(fromUserID)
	if err != nil {
		return err
	}

	if fromBalance.Amount < amount {
		return errors.New("insufficient funds")
	}

	toBalance, err := s.balanceRepo.GetBalance(toUserID)
	if err != nil {
		return err
	}

	fromBalance.Amount -= amount
	toBalance.Amount += amount

	if err := s.balanceRepo.UpdateBalance(fromBalance); err != nil {
		return err
	}
	if err := s.balanceRepo.UpdateBalance(toBalance); err != nil {
		return err
	}
	return nil
}