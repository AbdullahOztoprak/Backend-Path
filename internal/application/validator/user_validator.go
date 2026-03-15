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
	Password     string
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
		Password:     user.Password,
		Role:         user.Role,
	})
}

func (v *UserValidator) ValidateCredentials(username, email, password string) error {
	return ValidateUser(User{
		Username:     username,
		Email:        email,
		Password:     password,
		Role:         "user",
	})
}

var (
	ErrInvalidUsername     = errors.New("invalid username")
	ErrInvalidEmail        = errors.New("invalid email")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInvalidRole         = errors.New("invalid role")
)

func ValidateUser(user User) error {
	if err := validateUsername(user.Username); err != nil {
		return err
	}
	if err := validateEmail(user.Email); err != nil {
		return err
	}
	if err := validatePassword(user.Password); err != nil {
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

func validatePassword(password string) error {
	if len(password) == 0 {
		return ErrInvalidPassword
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