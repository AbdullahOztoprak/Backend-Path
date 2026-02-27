package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-backend-project/internal/domain/entity"
	"go-backend-project/internal/domain/repository"
	"go-backend-project/internal/application/validator"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindByUsername(username string) (*entity.User, error) {
	args := m.Called(username)
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestRegisterUser_Success(t *testing.T) {
	userRepo := new(UserRepositoryMock)
	userValidator := validator.NewUserValidator()
	usecase := NewRegisterUser(userRepo, userValidator)

	user := &entity.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "securepassword",
	}

	userRepo.On("Create", user).Return(nil)

	err := usecase.Execute(user)

	assert.NoError(t, err)
	userRepo.AssertExpectations(t)
}

func TestRegisterUser_UsernameTaken(t *testing.T) {
	userRepo := new(UserRepositoryMock)
	userValidator := validator.NewUserValidator()
	usecase := NewRegisterUser(userRepo, userValidator)

	existingUser := &entity.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "securepassword",
	}
	userRepo.On("FindByUsername", existingUser.Username).Return(existingUser, nil)

	err := usecase.Execute(existingUser)

	assert.Error(t, err)
	assert.Equal(t, "username already taken", err.Error())
	userRepo.AssertExpectations(t)
}

func TestRegisterUser_InvalidEmail(t *testing.T) {
	userRepo := new(UserRepositoryMock)
	userValidator := validator.NewUserValidator()
	usecase := NewRegisterUser(userRepo, userValidator)

	user := &entity.User{
		Username: "testuser",
		Email:    "invalid-email",
		Password: "securepassword",
	}

	err := usecase.Execute(user)

	assert.Error(t, err)
	assert.Equal(t, "invalid email format", err.Error())
}