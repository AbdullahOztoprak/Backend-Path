package validator

import (
	"errors"
	"regexp"
	"strings"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
)

type User struct {
	Username     string
	Email        string
	PasswordHash string
	Role         string
}

type UserValidator struct{}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (v *UserValidator) Validate(user *entity.User) error {
	return ValidateUser(User{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.Password,
		Role:         user.Role,
	})
}

func (v *UserValidator) ValidateCredentials(username, email, password string) error {
	return ValidateUser(User{
		Username:     username,
		Email:        email,
		PasswordHash: password,
		Role:         "user",
	})
}

var (
	ErrInvalidUsername     = errors.New("invalid username")
	ErrInvalidEmail        = errors.New("invalid email")
	ErrInvalidPasswordHash = errors.New("invalid password hash")
	ErrInvalidRole         = errors.New("invalid role")
)

func ValidateUser(user User) error {
	if err := validateUsername(user.Username); err != nil {
		return err
	}
	if err := validateEmail(user.Email); err != nil {
		return err
	}
	if err := validatePasswordHash(user.PasswordHash); err != nil {
		return err
	}
	if err := validateRole(user.Role); err != nil {
		return err
	}
	return nil
}

func validateUsername(username string) error {
	if len(username) < 3 || len(username) > 20 {
		return ErrInvalidUsername
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(username) {
		return ErrInvalidUsername
	}
	return nil
}

func validateEmail(email string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

func validatePasswordHash(passwordHash string) error {
	if len(passwordHash) == 0 {
		return ErrInvalidPasswordHash
	}
	return nil
}

func validateRole(role string) error {
	validRoles := []string{"user", "admin"}
	for _, r := range validRoles {
		if strings.EqualFold(role, r) {
			return nil
		}
	}
	return ErrInvalidRole
}