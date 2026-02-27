package service

import (
	"context"
	"errors"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
	"github.com/AbdullahOztoprak/Backend-Path/internal/application/validator"
)

type UserService struct {
	userRepo repository.UserRepository
	hasher    auth.BcryptHasher
	validator validator.UserValidator
}

func NewUserService(userRepo repository.UserRepository, hasher auth.BcryptHasher, validator validator.UserValidator) *UserService {
	return &UserService{
		userRepo: userRepo,
		hasher:    hasher,
		validator: validator,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, user *entity.User) error {
	if err := s.validator.Validate(user); err != nil {
		return err
	}

	hashedPassword, err := s.hasher.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.userRepo.Create(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := s.validator.Validate(user); err != nil {
		return err
	}

	return s.userRepo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.userRepo.Delete(ctx, id)
}

func (s *UserService) AuthenticateUser(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if err := s.hasher.Compare(user.Password, password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) ListUsers(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	return s.userRepo.List(ctx, limit, offset)
}

func (s *UserService) ChangePassword(ctx context.Context, userID int, newPassword string) error {
	hashedPassword, err := s.hasher.Hash(newPassword)
	if err != nil {
		return err
	}

	return s.userRepo.UpdatePassword(ctx, userID, hashedPassword)
}

func (s *UserService) SetLastLogin(ctx context.Context, userID int) error {
	return s.userRepo.UpdateLastLogin(ctx, userID, time.Now())
}