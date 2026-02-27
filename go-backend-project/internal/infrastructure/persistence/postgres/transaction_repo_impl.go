package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

type TransactionRepositoryImpl struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) repository.TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) Create(ctx context.Context, transaction *entity.Transaction) error {
	query := `INSERT INTO transactions (from_user_id, to_user_id, amount, description, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	
	err := r.db.QueryRowContext(ctx, query, transaction.FromUserID, transaction.ToUserID, transaction.Amount, transaction.Description, time.Now(), time.Now()).Scan(&transaction.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryImpl) GetByID(ctx context.Context, id int64) (*entity.Transaction, error) {
	query := `SELECT id, from_user_id, to_user_id, amount, description, created_at, updated_at 
			  FROM transactions WHERE id = $1`
	
	var transaction entity.Transaction
	err := r.db.QueryRowContext(ctx, query, id).Scan(&transaction.ID, &transaction.FromUserID, &transaction.ToUserID, &transaction.Amount, &transaction.Description, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepositoryImpl) ListByUserID(ctx context.Context, userID int64) ([]entity.Transaction, error) {
	query := `SELECT id, from_user_id, to_user_id, amount, description, created_at, updated_at 
			  FROM transactions WHERE from_user_id = $1 OR to_user_id = $1 ORDER BY created_at DESC`
	
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.Transaction
	for rows.Next() {
		var transaction entity.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.FromUserID, &transaction.ToUserID, &transaction.Amount, &transaction.Description, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *TransactionRepositoryImpl) Update(ctx context.Context, transaction *entity.Transaction) error {
	query := `UPDATE transactions SET amount = $1, description = $2, updated_at = $3 WHERE id = $4`
	
	_, err := r.db.ExecContext(ctx, query, transaction.Amount, transaction.Description, time.Now(), transaction.ID)
	return err
}

func (r *TransactionRepositoryImpl) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM transactions WHERE id = $1`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}