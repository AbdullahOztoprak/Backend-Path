package postgres

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *entity.User) error {
	query := `INSERT INTO users (username, email, password_hash, role) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, user.Username, user.Email, user.Password, user.Role).Scan(&user.ID)
}

func (r *UserRepositoryImpl) GetByID(id string) (*entity.User, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	query := `SELECT id, username, email, password_hash, role FROM users WHERE id = $1`
	user := &entity.User{}
	err = r.db.QueryRow(query, idInt).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetByUsername(username string) (*entity.User, error) {
	query := `SELECT id, username, email, password_hash, role FROM users WHERE username = $1`
	user := &entity.User{}
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Update(user *entity.User) error {
	query := `UPDATE users SET username = $1, email = $2, password_hash = $3, role = $4 WHERE id = $5`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Role, user.ID)
	return err
}

func (r *UserRepositoryImpl) Delete(id string) error {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	query := `DELETE FROM users WHERE id = $1`
	_, err = r.db.Exec(query, idInt)
	return err
}

func (r *UserRepositoryImpl) List() ([]*entity.User, error) {
	query := `SELECT id, username, email, password_hash, role FROM users ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// NewUserRepo keeps backward compatibility with older test call sites.
func NewUserRepo(db *sql.DB) repository.UserRepository {
	return NewUserRepository(db)
}