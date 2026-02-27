package mocks

import (
	"github.com/stretchr/testify/mock"
	"your_project/internal/domain/entity"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) GetByID(id string) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByUsername(username string) (*entity.User, error) {
	args := m.Called(username)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepositoryMock) List() ([]*entity.User, error) {
	args := m.Called()
	return args.Get(0).([]*entity.User), args.Error(1)
}