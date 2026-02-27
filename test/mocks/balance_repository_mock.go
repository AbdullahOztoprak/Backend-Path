package mocks

import (
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
)

type BalanceRepositoryMock struct {
	FindBalanceFunc func(userID int) (*entity.Balance, error)
	UpdateBalanceFunc func(balance *entity.Balance) error
}

func (m *BalanceRepositoryMock) FindBalance(userID int) (*entity.Balance, error) {
	if m.FindBalanceFunc != nil {
		return m.FindBalanceFunc(userID)
	}
	return nil, nil
}

func (m *BalanceRepositoryMock) UpdateBalance(balance *entity.Balance) error {
	if m.UpdateBalanceFunc != nil {
		return m.UpdateBalanceFunc(balance)
	}
	return nil
}