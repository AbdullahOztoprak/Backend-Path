package usecase

import (
	"errors"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/service"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
)

type RefreshTokenUseCase struct {
	authService service.AuthService
	tokenStore   auth.TokenStore
}

func NewRefreshTokenUseCase(authService service.AuthService, tokenStore auth.TokenStore) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		authService: authService,
		tokenStore:  tokenStore,
	}
}

func (uc *RefreshTokenUseCase) Execute(oldToken string) (string, error) {
	if oldToken == "" {
		return "", errors.New("refresh token is required")
	}

	claims, err := uc.authService.ValidateRefreshToken(oldToken)
	if err != nil {
		return "", err
	}

	if err := uc.tokenStore.ValidateToken(claims.ID); err != nil {
		return "", err
	}

	newToken, err := uc.authService.GenerateToken(claims.UserID, claims.Roles)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	if err := uc.tokenStore.StoreToken(claims.UserID, newToken, expirationTime); err != nil {
		return "", err
	}

	return newToken, nil
}