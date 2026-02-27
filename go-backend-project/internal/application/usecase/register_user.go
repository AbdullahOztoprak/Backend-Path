package usecase

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"your_project/internal/domain/entity"
	"your_project/internal/domain/repository"
	"your_project/internal/application/validator"
)

type RegisterUserInput struct {
	Username     string
	Email        string
	Password     string
	Role         string
}

type RegisterUserOutput struct {
	UserID string
}

type RegisterUserUseCase struct {
	userRepo repository.UserRepository
	validator *validator.UserValidator
}

func NewRegisterUserUseCase(userRepo repository.UserRepository, validator *validator.UserValidator) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo: userRepo,
		validator: validator,
	}
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, input RegisterUserInput) (RegisterUserOutput, error) {
	if err := uc.validator.Validate(input.Username, input.Email, input.Password); err != nil {
		return RegisterUserOutput{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterUserOutput{}, errors.New("failed to hash password")
	}

	user := entity.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	userID, err := uc.userRepo.Create(ctx, user)
	if err != nil {
		return RegisterUserOutput{}, err
	}

	return RegisterUserOutput{UserID: userID}, nil
}