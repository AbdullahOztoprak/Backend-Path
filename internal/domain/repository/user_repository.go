package repository

import "github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"

// UserRepository defines the interface for user data operations.
type UserRepository interface {
    Create(user *entity.User) error
    GetByID(id string) (*entity.User, error)
    GetByUsername(username string) (*entity.User, error)
    Update(user *entity.User) error
    Delete(id string) error
    List() ([]*entity.User, error)
}