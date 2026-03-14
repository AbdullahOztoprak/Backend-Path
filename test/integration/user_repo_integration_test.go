//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/persistence/postgres"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var userRepo repository.UserRepository

func setup() {
	db, err := postgres.NewConnection()
	if err != nil {
		panic(err)
	}
	userRepo = postgres.NewUserRepo(db)
}

func TestCreateUser(t *testing.T) {
	setup()
	defer teardown()

	user := &repository.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "securepassword",
	}

	createdUser, err := userRepo.Create(context.Background(), user)
	require.NoError(t, err)
	assert.NotNil(t, createdUser.ID)
	assert.Equal(t, user.Username, createdUser.Username)
}

func TestGetUserByID(t *testing.T) {
	setup()
	defer teardown()

	user := &repository.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "securepassword",
	}

	createdUser, err := userRepo.Create(context.Background(), user)
	require.NoError(t, err)

	fetchedUser, err := userRepo.GetByID(context.Background(), createdUser.ID)
	require.NoError(t, err)
	assert.Equal(t, createdUser.ID, fetchedUser.ID)
	assert.Equal(t, createdUser.Username, fetchedUser.Username)
}

func TestUpdateUser(t *testing.T) {
	setup()
	defer teardown()

	user := &repository.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "securepassword",
	}

	createdUser, err := userRepo.Create(context.Background(), user)
	require.NoError(t, err)

	createdUser.Email = "updated@example.com"
	updatedUser, err := userRepo.Update(context.Background(), createdUser)
	require.NoError(t, err)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}

func TestDeleteUser(t *testing.T) {
	setup()
	defer teardown()

	user := &repository.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "securepassword",
	}

	createdUser, err := userRepo.Create(context.Background(), user)
	require.NoError(t, err)

	err = userRepo.Delete(context.Background(), createdUser.ID)
	require.NoError(t, err)

	fetchedUser, err := userRepo.GetByID(context.Background(), createdUser.ID)
	assert.Error(t, err)
	assert.Nil(t, fetchedUser)
}

func teardown() {
	// Clean up resources, if necessary
}