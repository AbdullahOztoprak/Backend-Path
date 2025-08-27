package service_test

import (
	"testing"
	"github.com/AbdullahOztoprak/go-backend-project/internal/models"
	"github.com/AbdullahOztoprak/go-backend-project/internal/service"
)

// Dummy repo for testing

type dummyUserRepo struct{}

func (d *dummyUserRepo) Create(user *models.User) error                 { return nil }
func (d *dummyUserRepo) GetByID(id int64) (*models.User, error)         { return nil, nil }
func (d *dummyUserRepo) Update(user *models.User) error                 { return nil }
func (d *dummyUserRepo) Delete(id int64) error                         { return nil }
func (d *dummyUserRepo) List() ([]*models.User, error)                 { return []*models.User{}, nil }

func TestRegisterValidation(t *testing.T) {
	repo := &dummyUserRepo{}
	us := service.NewUserService(repo)
	user := &models.User{
		Username:     "ab",
		Email:        "bademail",
		PasswordHash: "123",
		Role:         "",
	}
	err := us.Register(user)
	if err == nil {
		t.Error("expected validation error, got nil")
	}
}
