//go:build legacy

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/service"
	"github.com/AbdullahOztoprak/Backend-Path/test/mocks"
)

func TestRegisterUser(t *testing.T) {
	authService := new(mocks.AuthServiceMock)
	userRepo := new(mocks.UserRepositoryMock)
	authService.On("Register", mock.Anything).Return(nil)

	// Call the method
	err := authService.Register(&service.User{Username: "testuser", Password: "password"})

	// Assertions
	assert.NoError(t, err)
	authService.AssertExpectations(t)
}

func TestLoginUser(t *testing.T) {
	authService := new(mocks.AuthServiceMock)
	userRepo := new(mocks.UserRepositoryMock)
	authService.On("Login", mock.Anything).Return("token", nil)

	// Call the method
	token, err := authService.Login(&service.User{Username: "testuser", Password: "password"})

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	authService.AssertExpectations(t)
}

func TestRefreshToken(t *testing.T) {
	authService := new(mocks.AuthServiceMock)
	authService.On("RefreshToken", mock.Anything).Return("new_token", nil)

	// Call the method
	newToken, err := authService.RefreshToken("old_token")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "new_token", newToken)
	authService.AssertExpectations(t)
}

func TestInvalidLogin(t *testing.T) {
	authService := new(mocks.AuthServiceMock)
	authService.On("Login", mock.Anything).Return("", service.ErrInvalidCredentials)

	// Call the method
	token, err := authService.Login(&service.User{Username: "wronguser", Password: "wrongpassword"})

	// Assertions
	assert.Error(t, err)
	assert.Empty(t, token)
	authService.AssertExpectations(t)
}