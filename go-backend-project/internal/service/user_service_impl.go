package service

import (
    "github.com/AbdullahOztoprak/go-backend-project/internal/models"
    "github.com/AbdullahOztoprak/go-backend-project/internal/repository"
    "golang.org/x/crypto/bcrypt"
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
    hashed, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.PasswordHash = string(hashed)
    return s.Repo.Create(user)
}

func (s *UserServiceImpl) Authenticate(username, password string) (*models.User, error) {
    // Find user by username (or email)
    users, err := s.Repo.List()
    if err != nil {
        return nil, err
    }
    var found *models.User
    for _, u := range users {
        if u.Username == username || u.Email == username {
            found = u
            break
        }
    }
    if found == nil {
        return nil, nil
    }
    // Compare password
    err = bcrypt.CompareHashAndPassword([]byte(found.PasswordHash), []byte(password))
    if err != nil {
        return nil, nil
    }
    return found, nil
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

func (s *UserServiceImpl) List() ([]*models.User, error) {
    return s.Repo.List()
}
