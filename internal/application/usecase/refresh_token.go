package usecase

import (
	"errors"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
)

type RefreshAuthService interface {
	ValidateRefreshToken(token string) (string, error)
}

type RefreshTokenStore interface {
	SaveToken(userID string, token string, expiration time.Duration) error
}

type RefreshTokenUseCase struct {
	authService RefreshAuthService
	tokenStore  RefreshTokenStore
	jwtProvider *auth.JWTProvider
}

func NewRefreshTokenUseCase(authService RefreshAuthService, tokenStore RefreshTokenStore, jwtProvider *auth.JWTProvider) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		authService: authService,
		tokenStore:  tokenStore,
		jwtProvider: jwtProvider,
	}
}

func (uc *RefreshTokenUseCase) Execute(oldToken string) (string, error) {
	if oldToken == "" {
		return "", errors.New("refresh token is required")
	}

	if uc.authService == nil || uc.tokenStore == nil || uc.jwtProvider == nil {
		return "", errors.New("refresh token dependencies are not configured")
	}

	userID, err := uc.authService.ValidateRefreshToken(oldToken)
	if err != nil {
		return "", err
	}

	newToken, err := uc.jwtProvider.GenerateToken(userID, []string{"user"})
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	if err := uc.tokenStore.SaveToken(userID, newToken, time.Until(expirationTime)); err != nil {
		return "", err
	}

	return newToken, nil
}