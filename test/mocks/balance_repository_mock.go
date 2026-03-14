package mocks

import (
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
)

type BalanceRepositoryMock struct {
	GetBalanceFunc    func(userID int) (*entity.Balance, error)
	UpdateBalanceFunc func(balance *entity.Balance) error
	CreateBalanceFunc func(balance *entity.Balance) error
	DeleteBalanceFunc func(userID int) error
}

func (m *BalanceRepositoryMock) GetBalance(userID int) (*entity.Balance, error) {
	if m.GetBalanceFunc != nil {
		return m.GetBalanceFunc(userID)
	}
	return nil, nil
}

func (m *BalanceRepositoryMock) UpdateBalance(balance *entity.Balance) error {
	if m.UpdateBalanceFunc != nil {
		return m.UpdateBalanceFunc(balance)
	}
	return nil
}

func (m *BalanceRepositoryMock) CreateBalance(balance *entity.Balance) error {
	if m.CreateBalanceFunc != nil {
		return m.CreateBalanceFunc(balance)
	}
	return nil
}

func (m *BalanceRepositoryMock) DeleteBalance(userID int) error {
	if m.DeleteBalanceFunc != nil {
		return m.DeleteBalanceFunc(userID)
	}
	return nil
}