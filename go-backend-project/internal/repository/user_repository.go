package repository

import "github.com/AbdullahOztoprak/go-backend-project/internal/models"

type UserRepository interface {
    Create(user *models.User) error
    GetByID(id int64) (*models.User, error)
    Update(user *models.User) error
    Delete(id int64) error
    List() ([]*models.User, error)
}