package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"your_project/internal/domain/entity"
)

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name     string
		user     entity.User
		expected bool
	}{
		{
			name: "valid user",
			user: entity.User{
				Username: "valid_user",
				Email:    "valid@example.com",
				Password: "SecurePassword123!",
				Role:     "user",
			},
			expected: true,
		},
		{
			name: "invalid username",
			user: entity.User{
				Username: "",
				Email:    "valid@example.com",
				Password: "SecurePassword123!",
				Role:     "user",
			},
			expected: false,
		},
		{
			name: "invalid email",
			user: entity.User{
				Username: "valid_user",
				Email:    "invalid-email",
				Password: "SecurePassword123!",
				Role:     "user",
			},
			expected: false,
		},
		{
			name: "invalid password",
			user: entity.User{
				Username: "valid_user",
				Email:    "valid@example.com",
				Password: "short",
				Role:     "user",
			},
			expected: false,
		},
		{
			name: "invalid role",
			user: entity.User{
				Username: "valid_user",
				Email:    "valid@example.com",
				Password: "SecurePassword123!",
				Role:     "invalid_role",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateUser(tt.user)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidateUserCreation(t *testing.T) {
	tests := []struct {
		name     string
		user     entity.User
		expected bool
	}{
		{
			name: "valid user creation",
			user: entity.User{
				Username: "new_user",
				Email:    "new@example.com",
				Password: "StrongPassword123!",
				Role:     "user",
			},
			expected: true,
		},
		{
			name: "duplicate username",
			user: entity.User{
				Username: "existing_user",
				Email:    "new@example.com",
				Password: "StrongPassword123!",
				Role:     "user",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateUserCreation(tt.user)
			require.Equal(t, tt.expected, result)
		})
	}
}