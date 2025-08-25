package repository

import (
    "context"
    "github.com/AbdullahOztoprak/go-backend-project/internal/models"
    "github.com/jackc/pgx/v5"
)

type PGBalanceRepository struct {
    Conn *pgx.Conn
}

func NewPGBalanceRepository(conn *pgx.Conn) *PGBalanceRepository {
    return &PGBalanceRepository{Conn: conn}
}

func (r *PGBalanceRepository) GetByUserID(userID int64) (*models.Balance, error) {
    row := r.Conn.QueryRow(context.Background(),
        "SELECT user_id, amount, last_updated_at FROM balances WHERE user_id=$1", userID)
    var balance models.Balance
    err := row.Scan(&balance.UserID, &balance.Amount, &balance.LastUpdatedAt)
    if err != nil {
        return nil, err
    }
    return &balance, nil
}

func (r *PGBalanceRepository) UpdateAmount(userID int64, delta float64) error {
    _, err := r.Conn.Exec(context.Background(),
        "UPDATE balances SET amount = amount + $1, last_updated_at = NOW() WHERE user_id = $2",
        delta, userID)
    return err
}

func (r *PGBalanceRepository) GetHistory(userID int64) ([]*models.Balance, error) {
    // For now, return current balance as history
    // In a real implementation, you might have a separate balance_history table
    balance, err := r.GetByUserID(userID)
    if err != nil {
        return nil, err
    }
    return []*models.Balance{balance}, nil
}
