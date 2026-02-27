package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/service"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
	"github.com/AbdullahOztoprak/Backend-Path/pkg/apperror"
)

type LoginUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserOutput struct {
	Token         string    `json:"token"`
	RefreshToken  string    `json:"refresh_token"`
	ExpiresAt     time.Time `json:"expires_at"`
}

type LoginUserUseCase struct {
	authService service.AuthService
	userService service.UserService
}

func NewLoginUserUseCase(authService service.AuthService, userService service.UserService) *LoginUserUseCase {
	return &LoginUserUseCase{
		authService: authService,
		userService: userService,
	}
}

func (uc *LoginUserUseCase) Execute(ctx context.Context, input LoginUserInput) (LoginUserOutput, error) {
	user, err := uc.userService.FindByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return LoginUserOutput{}, apperror.ErrInvalidCredentials
		}
		return LoginUserOutput{}, err
	}

	if err := uc.authService.ComparePassword(user.PasswordHash, input.Password); err != nil {
		return LoginUserOutput{}, apperror.ErrInvalidCredentials
	}

	token, refreshToken, err := uc.authService.GenerateTokens(user)
	if err != nil {
		return LoginUserOutput{}, err
	}

	return LoginUserOutput{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(24 * time.Hour), // Token expiration time
	}, nil
}