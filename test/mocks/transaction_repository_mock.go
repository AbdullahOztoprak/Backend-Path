package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) Create(ctx context.Context, transaction *entity.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) GetByID(ctx context.Context, id string) (*entity.Transaction, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) ListByUserID(ctx context.Context, userID string) ([]entity.Transaction, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) Update(ctx context.Context, transaction *entity.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}