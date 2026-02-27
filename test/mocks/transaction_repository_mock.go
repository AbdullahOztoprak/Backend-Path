package mocks

import (
	"github.com/stretchr/testify/mock"
	"your_project_path/internal/domain/entity"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) GetByID(id int) (*entity.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) ListByUserID(userID int) ([]*entity.Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]*entity.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) Update(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}