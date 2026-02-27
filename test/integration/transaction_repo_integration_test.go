package integration_test

import (
	"context"
	"testing"

	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/persistence/postgres"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	transactionRepo repository.TransactionRepository
	ctx             context.Context
)

func setup() {
	// Initialize the database connection and transaction repository
	db, err := postgres.NewConnection()
	require.NoError(t, err)

	transactionRepo = postgres.NewTransactionRepo(db)
	ctx = context.Background()
}

func TestCreateTransaction(t *testing.T) {
	setup()
	defer teardown()

	transaction := &entity.Transaction{
		FromUserID: 1,
		ToUserID:   2,
		Amount:     100.50,
		Description: "Payment for services",
	}

	createdTransaction, err := transactionRepo.Create(ctx, transaction)
	require.NoError(t, err)
	assert.NotNil(t, createdTransaction.ID)
	assert.Equal(t, transaction.FromUserID, createdTransaction.FromUserID)
	assert.Equal(t, transaction.ToUserID, createdTransaction.ToUserID)
	assert.Equal(t, transaction.Amount, createdTransaction.Amount)
	assert.Equal(t, transaction.Description, createdTransaction.Description)
}

func TestGetTransactionByID(t *testing.T) {
	setup()
	defer teardown()

	transaction := &entity.Transaction{
		FromUserID: 1,
		ToUserID:   2,
		Amount:     100.50,
		Description: "Payment for services",
	}

	createdTransaction, err := transactionRepo.Create(ctx, transaction)
	require.NoError(t, err)

	fetchedTransaction, err := transactionRepo.GetByID(ctx, createdTransaction.ID)
	require.NoError(t, err)
	assert.Equal(t, createdTransaction.ID, fetchedTransaction.ID)
	assert.Equal(t, createdTransaction.FromUserID, fetchedTransaction.FromUserID)
	assert.Equal(t, createdTransaction.ToUserID, fetchedTransaction.ToUserID)
	assert.Equal(t, createdTransaction.Amount, fetchedTransaction.Amount)
	assert.Equal(t, createdTransaction.Description, fetchedTransaction.Description)
}

func TestListTransactions(t *testing.T) {
	setup()
	defer teardown()

	transaction1 := &entity.Transaction{
		FromUserID: 1,
		ToUserID:   2,
		Amount:     100.50,
		Description: "Payment for services",
	}

	transaction2 := &entity.Transaction{
		FromUserID: 2,
		ToUserID:   3,
		Amount:     200.75,
		Description: "Payment for goods",
	}

	_, err := transactionRepo.Create(ctx, transaction1)
	require.NoError(t, err)
	_, err = transactionRepo.Create(ctx, transaction2)
	require.NoError(t, err)

	transactions, err := transactionRepo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, transactions, 2)
}

func teardown() {
	// Clean up resources, if necessary
}