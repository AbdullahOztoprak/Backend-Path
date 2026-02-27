package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/persistence/postgres"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

func TestBalanceRepositoryIntegration(t *testing.T) {
	ctx := context.Background()
	db, cleanup := SetupPostgres(t)
	defer cleanup()

	balanceRepo := postgres.NewBalanceRepository(db)

	t.Run("Create and Get Balance", func(t *testing.T) {
		balance := &repository.Balance{
			UserID: 1,
			Amount: 100.00,
		}

		// Create balance
		err := balanceRepo.Create(ctx, balance)
		require.NoError(t, err)

		// Get balance
		fetchedBalance, err := balanceRepo.GetByUserID(ctx, balance.UserID)
		require.NoError(t, err)
		assert.Equal(t, balance.Amount, fetchedBalance.Amount)
	})

	t.Run("Update Balance", func(t *testing.T) {
		balance := &repository.Balance{
			UserID: 1,
			Amount: 150.00,
		}

		// Update balance
		err := balanceRepo.Update(ctx, balance)
		require.NoError(t, err)

		// Verify update
		fetchedBalance, err := balanceRepo.GetByUserID(ctx, balance.UserID)
		require.NoError(t, err)
		assert.Equal(t, balance.Amount, fetchedBalance.Amount)
	})

	t.Run("Delete Balance", func(t *testing.T) {
		err := balanceRepo.Delete(ctx, 1)
		require.NoError(t, err)

		// Verify deletion
		fetchedBalance, err := balanceRepo.GetByUserID(ctx, 1)
		assert.Error(t, err)
		assert.Nil(t, fetchedBalance)
	})
}