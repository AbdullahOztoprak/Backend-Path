package service

import (
	"errors"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
	"github.com/AbdullahOztoprak/Backend-Path/pkg/apperror"
)

type AuthService struct {
	userRepo       repository.UserRepository
	jwtProvider    auth.JWTProvider
	bcryptHasher   auth.BcryptHasher
	tokenStore     auth.TokenStore
}

func NewAuthService(userRepo repository.UserRepository, jwtProvider auth.JWTProvider, bcryptHasher auth.BcryptHasher, tokenStore auth.TokenStore) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		jwtProvider: jwtProvider,
		bcryptHasher: bcryptHasher,
		tokenStore:  tokenStore,
	}
}

func (s *AuthService) Register(username, email, password string) (*entity.User, error) {
	if username == "" || email == "" || password == "" {
		return nil, apperror.NewBadRequestError("username, email, and password are required")
	}

	hashedPassword, err := s.bcryptHasher.Hash(password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if !s.bcryptHasher.Compare(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := s.jwtProvider.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) RefreshToken(oldToken string) (string, error) {
	claims, err := s.jwtProvider.ValidateToken(oldToken)
	if err != nil {
		return "", err
	}

	if time.Now().After(claims.ExpiresAt) {
		return "", errors.New("token expired")
	}

	newToken, err := s.jwtProvider.GenerateToken(claims.UserID, claims.Role)
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func (s *AuthService) Logout(userID string) error {
	return s.tokenStore.Remove(userID)
}