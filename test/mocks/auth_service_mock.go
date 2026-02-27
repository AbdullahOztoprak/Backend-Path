package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) Register(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *AuthServiceMock) Login(ctx context.Context, username, password string) (string, error) {
	args := m.Called(ctx, username, password)
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) RefreshToken(ctx context.Context, token string) (string, error) {
	args := m.Called(ctx, token)
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) ValidateToken(ctx context.Context, token string) (*entity.User, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *AuthServiceMock) GetUserFromToken(ctx context.Context, token string) (*entity.User, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(*entity.User), args.Error(1)
}