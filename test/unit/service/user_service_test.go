package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/service"
	"github.com/AbdullahOztoprak/Backend-Path/test/mocks"
)

func TestRegisterUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userService := service.NewUserService(mockUserRepo)

	user := &entity.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "securepassword",
	}

	mockUserRepo.On("Create", user).Return(nil)

	err := userService.Register(user)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userService := service.NewUserService(mockUserRepo)

	user := &entity.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}

	mockUserRepo.On("FindByID", user.ID).Return(user, nil)

	result, err := userService.GetUserByID(user.ID)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockUserRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userService := service.NewUserService(mockUserRepo)

	user := &entity.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}

	mockUserRepo.On("Update", user).Return(nil)

	err := userService.Update(user)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userService := service.NewUserService(mockUserRepo)

	userID := 1

	mockUserRepo.On("Delete", userID).Return(nil)

	err := userService.Delete(userID)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}