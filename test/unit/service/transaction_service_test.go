package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-backend-project/internal/domain/entity"
	"go-backend-project/internal/domain/repository"
	"go-backend-project/internal/domain/service"
	"go-backend-project/test/mocks"
)

func TestCreateTransaction(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	transactionService := service.NewTransactionService(mockTransactionRepo)

	transaction := &entity.Transaction{
		FromUserID: 1,
		ToUserID:   2,
		Amount:     100.50,
		Description: "Payment for services",
	}

	mockTransactionRepo.On("Create", transaction).Return(nil)

	err := transactionService.Create(transaction)

	assert.NoError(t, err)
	mockTransactionRepo.AssertExpectations(t)
}

func TestCreateTransaction_InsufficientFunds(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	transactionService := service.NewTransactionService(mockTransactionRepo)

	transaction := &entity.Transaction{
		FromUserID: 1,
		ToUserID:   2,
		Amount:     1000.00, // Assuming insufficient funds
		Description: "Payment for services",
	}

	mockTransactionRepo.On("Create", transaction).Return(repository.ErrInsufficientFunds)

	err := transactionService.Create(transaction)

	assert.Error(t, err)
	assert.Equal(t, repository.ErrInsufficientFunds, err)
	mockTransactionRepo.AssertExpectations(t)
}

func TestGetTransactionByID(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	transactionService := service.NewTransactionService(mockTransactionRepo)

	transaction := &entity.Transaction{
		ID:          1,
		FromUserID: 1,
		ToUserID:   2,
		Amount:     100.50,
		Description: "Payment for services",
	}

	mockTransactionRepo.On("GetByID", transaction.ID).Return(transaction, nil)

	result, err := transactionService.GetByID(transaction.ID)

	assert.NoError(t, err)
	assert.Equal(t, transaction, result)
	mockTransactionRepo.AssertExpectations(t)
}

func TestGetTransactionByID_NotFound(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	transactionService := service.NewTransactionService(mockTransactionRepo)

	mockTransactionRepo.On("GetByID", 999).Return(nil, repository.ErrNotFound)

	result, err := transactionService.GetByID(999)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockTransactionRepo.AssertExpectations(t)
}