package service

import (
    "github.com/AbdullahOztoprak/go-backend-project/internal/models"
    "github.com/AbdullahOztoprak/go-backend-project/internal/repository"
)

type UserServiceImpl struct {
    Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
    return &UserServiceImpl{Repo: repo}
}

func (s *UserServiceImpl) Register(user *models.User) error {
    if err := user.Validate(); err != nil {
        return err
    }
    return s.Repo.Create(user)
}

func (s *UserServiceImpl) Authenticate(username, password string) (*models.User, error) {
    // In a real implementation, you would hash the password and check against stored hash
    // For now, this is a placeholder
    return nil, nil
}

func (s *UserServiceImpl) GetByID(id int64) (*models.User, error) {
    return s.Repo.GetByID(id)
}

func (s *UserServiceImpl) Update(user *models.User) error {
    if err := user.Validate(); err != nil {
        return err
    }
    return s.Repo.Update(user)
}

func (s *UserServiceImpl) Delete(id int64) error {
    return s.Repo.Delete(id)
}
