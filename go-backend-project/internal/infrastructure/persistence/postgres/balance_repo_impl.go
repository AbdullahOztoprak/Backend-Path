package postgres

import (
	"context"
	"database/sql"
	"errors"
	"go-backend-project/internal/domain/entity"
	"go-backend-project/internal/domain/repository"
)

type BalanceRepoImpl struct {
	db *sql.DB
}

func NewBalanceRepo(db *sql.DB) repository.BalanceRepository {
	return &BalanceRepoImpl{db: db}
}

func (r *BalanceRepoImpl) GetBalance(ctx context.Context, userID int) (*entity.Balance, error) {
	var balance entity.Balance
	query := "SELECT user_id, amount FROM balances WHERE user_id = $1"
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&balance.UserID, &balance.Amount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No balance found for the user
		}
		return nil, err // Return any other error
	}
	return &balance, nil
}

func (r *BalanceRepoImpl) UpdateBalance(ctx context.Context, userID int, amount float64) error {
	query := "UPDATE balances SET amount = amount + $1 WHERE user_id = $2"
	result, err := r.db.ExecContext(ctx, query, amount, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no balance updated, user may not exist")
	}
	return nil
}

func (r *BalanceRepoImpl) CreateBalance(ctx context.Context, userID int, amount float64) error {
	query := "INSERT INTO balances (user_id, amount) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, userID, amount)
	return err
}