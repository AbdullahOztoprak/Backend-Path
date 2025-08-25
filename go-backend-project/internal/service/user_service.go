package service

import "github.com/AbdullahOztoprak/go-backend-project/internal/models"

type UserService interface {
    Register(user *models.User) error
    Authenticate(username, password string) (*models.User, error)
    GetByID(id int64) (*models.User, error)
    Update(user *models.User) error
    Delete(id int64) error
}
