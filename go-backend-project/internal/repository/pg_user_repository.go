package repository

import (
    "context"
    "github.com/AbdullahOztoprak/go-backend-project/internal/models"
    "github.com/jackc/pgx/v5"
)

type PGUserRepository struct {
    Conn *pgx.Conn
}

func NewPGUserRepository(conn *pgx.Conn) *PGUserRepository {
    return &PGUserRepository{Conn: conn}
}

func (r *PGUserRepository) Create(user *models.User) error {
    _, err := r.Conn.Exec(context.Background(),
        "INSERT INTO users (username, email, password_hash, role) VALUES ($1, $2, $3, $4)",
        user.Username, user.Email, user.PasswordHash, user.Role)
    return err
}

func (r *PGUserRepository) GetByID(id int64) (*models.User, error) {
    row := r.Conn.QueryRow(context.Background(),
        "SELECT id, username, email, password_hash, role, created_at, updated_at FROM users WHERE id=$1", id)
    var user models.User
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *PGUserRepository) Update(user *models.User) error {
    _, err := r.Conn.Exec(context.Background(),
        "UPDATE users SET username=$1, email=$2, password_hash=$3, role=$4, updated_at=NOW() WHERE id=$5",
        user.Username, user.Email, user.PasswordHash, user.Role, user.ID)
    return err
}

func (r *PGUserRepository) Delete(id int64) error {
    _, err := r.Conn.Exec(context.Background(),
        "DELETE FROM users WHERE id=$1", id)
    return err
}

func (r *PGUserRepository) List() ([]*models.User, error) {
    rows, err := r.Conn.Query(context.Background(),
        "SELECT id, username, email, password_hash, role, created_at, updated_at FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, &user)
    }
    return users, nil
}
