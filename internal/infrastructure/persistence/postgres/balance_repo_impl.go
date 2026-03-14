package postgres

import (
	"database/sql"
	"errors"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

type BalanceRepoImpl struct {
	db *sql.DB
}

func NewBalanceRepo(db *sql.DB) repository.BalanceRepository {
	return &BalanceRepoImpl{db: db}
}

func (r *BalanceRepoImpl) GetBalance(userID int) (*entity.Balance, error) {
	var balance entity.Balance
	query := "SELECT user_id, amount FROM balances WHERE user_id = $1"
	err := r.db.QueryRow(query, userID).Scan(&balance.UserID, &balance.Amount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &balance, nil
}

func (r *BalanceRepoImpl) UpdateBalance(balance *entity.Balance) error {
	query := "UPDATE balances SET amount = $1, currency = $2 WHERE user_id = $3"
	result, err := r.db.Exec(query, balance.Amount, balance.Currency, balance.UserID)
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

func (r *BalanceRepoImpl) CreateBalance(balance *entity.Balance) error {
	query := "INSERT INTO balances (user_id, amount, currency) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, balance.UserID, balance.Amount, balance.Currency)
	return err
}

func (r *BalanceRepoImpl) DeleteBalance(userID int) error {
	query := "DELETE FROM balances WHERE user_id = $1"
	_, err := r.db.Exec(query, userID)
	return err
}

// NewBalanceRepository keeps backward compatibility with older test call sites.
func NewBalanceRepository(db *sql.DB) repository.BalanceRepository {
	return NewBalanceRepo(db)
}