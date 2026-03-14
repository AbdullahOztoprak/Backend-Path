package usecase

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
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
	userRepo    repository.UserRepository
	jwtProvider *auth.JWTProvider
	hasher      *auth.BcryptHasher
}

func NewLoginUserUseCase(userRepo repository.UserRepository, jwtProvider *auth.JWTProvider, hasher *auth.BcryptHasher) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepo:    userRepo,
		jwtProvider: jwtProvider,
		hasher:      hasher,
	}
}

func (uc *LoginUserUseCase) Execute(ctx context.Context, input LoginUserInput) (LoginUserOutput, error) {
	_ = ctx

	if uc.userRepo == nil || uc.jwtProvider == nil || uc.hasher == nil {
		return LoginUserOutput{}, errors.New("login use case dependencies are not configured")
	}

	user, err := uc.userRepo.GetByUsername(input.Username)
	if err != nil {
		return LoginUserOutput{}, err
	}
	if user == nil {
		return LoginUserOutput{}, errors.New("invalid credentials")
	}

	if err := uc.hasher.Compare(user.Password, input.Password); err != nil {
		return LoginUserOutput{}, errors.New("invalid credentials")
	}

	token, err := uc.jwtProvider.GenerateToken(strconv.FormatInt(user.ID, 10), []string{user.Role})
	if err != nil {
		return LoginUserOutput{}, err
	}

	return LoginUserOutput{
		Token:        token,
		RefreshToken: token,
		ExpiresAt:    time.Now().Add(24 * time.Hour), // Token expiration time
	}, nil
}