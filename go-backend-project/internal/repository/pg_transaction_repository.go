package repository

import (
    "context"
    "github.com/AbdullahOztoprak/go-backend-project/internal/models"
    "github.com/jackc/pgx/v5"
)

type PGTransactionRepository struct {
    Conn *pgx.Conn
}

func NewPGTransactionRepository(conn *pgx.Conn) *PGTransactionRepository {
    return &PGTransactionRepository{Conn: conn}
}

func (r *PGTransactionRepository) Create(tx *models.Transaction) error {
    _, err := r.Conn.Exec(context.Background(),
        "INSERT INTO transactions (from_user_id, to_user_id, amount, type, status) VALUES ($1, $2, $3, $4, $5)",
        tx.FromUserID, tx.ToUserID, tx.Amount, tx.Type, tx.Status)
    return err
}

func (r *PGTransactionRepository) GetByID(id int64) (*models.Transaction, error) {
    row := r.Conn.QueryRow(context.Background(),
        "SELECT id, from_user_id, to_user_id, amount, type, status, created_at FROM transactions WHERE id=$1", id)
    var tx models.Transaction
    err := row.Scan(&tx.ID, &tx.FromUserID, &tx.ToUserID, &tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &tx, nil
}

func (r *PGTransactionRepository) ListByUser(userID int64) ([]*models.Transaction, error) {
    rows, err := r.Conn.Query(context.Background(),
        "SELECT id, from_user_id, to_user_id, amount, type, status, created_at FROM transactions WHERE from_user_id=$1 OR to_user_id=$1", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var transactions []*models.Transaction
    for rows.Next() {
        var tx models.Transaction
        err := rows.Scan(&tx.ID, &tx.FromUserID, &tx.ToUserID, &tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt)
        if err != nil {
            return nil, err
        }
        transactions = append(transactions, &tx)
    }
    return transactions, nil
}

func (r *PGTransactionRepository) UpdateStatus(id int64, status models.TransactionStatus) error {
    _, err := r.Conn.Exec(context.Background(),
        "UPDATE transactions SET status=$1 WHERE id=$2", status, id)
    return err
}
