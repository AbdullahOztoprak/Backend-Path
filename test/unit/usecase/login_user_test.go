//go:build legacy

package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/service"
	"github.com/AbdullahOztoprak/Backend-Path/test/mocks"
)

func TestLoginUser(t *testing.T) {
	mockAuthService := new(mocks.AuthServiceMock)
	mockUserRepo := new(mocks.UserRepositoryMock)

	loginUser := NewLoginUser(mockAuthService, mockUserRepo)

	tests := []struct {
		name          string
		input         entity.User
		expectedToken string
		expectedError error
	}{
		{
			name: "Successful login",
			input: entity.User{
				Username: "test_user",
				Password: "securepassword",
			},
			expectedToken: "valid_token",
			expectedError: nil,
		},
		{
			name: "Invalid credentials",
			input: entity.User{
				Username: "test_user",
				Password: "wrongpassword",
			},
			expectedToken: "",
			expectedError: service.ErrInvalidCredentials,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo.On("FindByUsername", tt.input.Username).Return(&entity.User{
				Username: "test_user",
				Password: "hashed_securepassword",
			}, nil)

			if tt.expectedError == nil {
				mockAuthService.On("ValidatePassword", tt.input.Password, "hashed_securepassword").Return(true)
				mockAuthService.On("GenerateToken", mock.Anything).Return(tt.expectedToken, nil)
			} else {
				mockAuthService.On("ValidatePassword", tt.input.Password, "hashed_securepassword").Return(false)
			}

			token, err := loginUser.Execute(tt.input)

			assert.Equal(t, tt.expectedToken, token)
			assert.Equal(t, tt.expectedError, err)

			mockUserRepo.AssertExpectations(t)
			mockAuthService.AssertExpectations(t)
		})
	}
}