package usecase

import (
	"context"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/AbdullahOztoprak/Backend-Path/internal/application/validator"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
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
	userRepo  repository.UserRepository
	validator *validator.UserValidator
}

func NewRegisterUserUseCase(userRepo repository.UserRepository, validator *validator.UserValidator) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo: userRepo,
		validator: validator,
	}
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, input RegisterUserInput) (RegisterUserOutput, error) {
	_ = ctx

	if err := uc.validator.ValidateCredentials(input.Username, input.Email, input.Password); err != nil {
		return RegisterUserOutput{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterUserOutput{}, err
	}

	user := &entity.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := uc.userRepo.Create(user); err != nil {
		return RegisterUserOutput{}, err
	}

	return RegisterUserOutput{UserID: strconv.FormatInt(user.ID, 10)}, nil
}