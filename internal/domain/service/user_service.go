package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/AbdullahOztoprak/Backend-Path/internal/application/validator"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
)

type UserService struct {
	userRepo repository.UserRepository
	hasher   *auth.BcryptHasher
	validator *validator.UserValidator
}

func NewUserService(userRepo repository.UserRepository, hasher *auth.BcryptHasher, userValidator *validator.UserValidator) *UserService {
	return &UserService{
		userRepo: userRepo,
		hasher:   hasher,
		validator: userValidator,
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

	_ = ctx
	return s.userRepo.Create(user)
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	_ = ctx
	return s.userRepo.GetByID(strconv.Itoa(id))
}

func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := s.validator.Validate(user); err != nil {
		return err
	}

	_ = ctx
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	_ = ctx
	return s.userRepo.Delete(strconv.Itoa(id))
}

func (s *UserService) AuthenticateUser(ctx context.Context, username, password string) (*entity.User, error) {
	_ = ctx
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if err := s.hasher.Compare(user.Password, password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) ListUsers(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	_ = ctx
	_, _ = limit, offset
	return s.userRepo.List()
}

func (s *UserService) ChangePassword(ctx context.Context, userID int, newPassword string) error {
	_ = ctx

	hashedPassword, err := s.hasher.Hash(newPassword)
	if err != nil {
		return err
	}

	user, err := s.userRepo.GetByID(strconv.Itoa(userID))
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	user.Password = hashedPassword
	return s.userRepo.Update(user)
}

func (s *UserService) SetLastLogin(ctx context.Context, userID int) error {
	_ = ctx
	_, err := s.userRepo.GetByID(strconv.Itoa(userID))
	return err
}