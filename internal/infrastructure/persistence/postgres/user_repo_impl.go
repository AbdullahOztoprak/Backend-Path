package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (username, email, password_hash, role) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.PasswordHash, user.Role).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	query := `SELECT id, username, email, password_hash, role FROM users WHERE id = $1`
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `SELECT id, username, email, password_hash, role FROM users WHERE username = $1`
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET username = $1, email = $2, password_hash = $3, role = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.PasswordHash, user.Role, user.ID)
	return err
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepositoryImpl) ListUsers(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	query := `SELECT id, username, email, password_hash, role FROM users ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, username, email, password_hash, role FROM users WHERE email = $1`
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetUserWithTimeout(ctx context.Context, id int, timeout time.Duration) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return r.GetUserByID(ctx, id)
}