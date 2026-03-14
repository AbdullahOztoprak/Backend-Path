package service

import (
	"errors"
	"strconv"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
)

type AuthService struct {
	userRepo      repository.UserRepository
	jwtProvider   *auth.JWTProvider
	bcryptHasher  *auth.BcryptHasher
}

func NewAuthService(userRepo repository.UserRepository, jwtProvider *auth.JWTProvider, bcryptHasher *auth.BcryptHasher) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		jwtProvider:  jwtProvider,
		bcryptHasher: bcryptHasher,
	}
}

func (s *AuthService) Register(username, email, password string) (*entity.User, error) {
	if username == "" || email == "" || password == "" {
		return nil, errors.New("username, email, and password are required")
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
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := s.bcryptHasher.Compare(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := s.jwtProvider.GenerateToken(strconv.FormatInt(user.ID, 10), []string{user.Role})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) RefreshToken(oldToken string) (string, error) {
	userID, err := s.jwtProvider.ValidateToken(oldToken)
	if err != nil {
		return "", err
	}

	newToken, err := s.jwtProvider.GenerateToken(userID, []string{"user"})
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func (s *AuthService) Logout(userID string) error {
	if userID == "" {
		return errors.New("user id is required")
	}
	return nil
}