package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AbdullahOztoprak/Backend-Path/internal/api"
	"github.com/AbdullahOztoprak/Backend-Path/internal/api/handler"
)

type fakeAuthUseCase struct{}

func (f fakeAuthUseCase) Login(_ context.Context, username, password string) (handler.TokenPair, error) {
	if username != "test_user" || password != "securepassword" {
		return handler.TokenPair{}, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   "user-1",
		"roles": []string{"user"},
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	signed, err := token.SignedString([]byte("dev-secret"))
	if err != nil {
		return handler.TokenPair{}, err
	}

	return handler.TokenPair{
		AccessToken:  signed,
		RefreshToken: "refresh-token",
		ExpiresIn:    3600,
	}, nil
}

func (f fakeAuthUseCase) Refresh(_ context.Context, _ string) (handler.TokenPair, error) {
	return handler.TokenPair{AccessToken: "new-token", RefreshToken: "new-refresh", ExpiresIn: 3600}, nil
}

type fakeUserUseCase struct{}

func (f fakeUserUseCase) Create(_ context.Context, input handler.CreateUserInput) (handler.CreateUserOutput, error) {
	return handler.CreateUserOutput{
		ID:        "1",
		Username:  input.Username,
		Email:     input.Email,
		Role:      input.Role,
		CreatedAt: time.Now(),
	}, nil
}

func TestUserRegistrationAndLogin(t *testing.T) {
	router := api.NewRouter(api.Dependencies{
		HealthHandler: handler.NewHealthHandler(),
		AuthHandler:   handler.NewAuthHandler(fakeAuthUseCase{}),
		UserHandler:   handler.NewUserHandler(fakeUserUseCase{}),
	})

	// Step 1: User Registration
	user := map[string]interface{}{
		"username": "test_user",
		"email":    "test@example.com",
		"password": "securepassword",
		"role":     "user",
	}

	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)

	// Step 2: User Login
	login := map[string]string{
		"username": "test_user",
		"password": "securepassword",
	}

	loginJSON, _ := json.Marshal(login)
	req = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	var loginResponse map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &loginResponse)
	require.NoError(t, err)
	_, hasAccessToken := loginResponse["access_token"]
	_, hasRefreshToken := loginResponse["refresh_token"]
	assert.True(t, hasAccessToken)
	assert.True(t, hasRefreshToken)
}

func TestUnauthorizedAccess(t *testing.T) {
	router := api.NewRouter(api.Dependencies{
		HealthHandler: handler.NewHealthHandler(),
	})

	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/transactions", nil)
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestInvalidLogin(t *testing.T) {
	router := api.NewRouter(api.Dependencies{
		HealthHandler: handler.NewHealthHandler(),
		AuthHandler:   handler.NewAuthHandler(fakeAuthUseCase{}),
	})

	login := map[string]string{
		"username": "test_user",
		"password": "wrongpassword",
	}

	loginJSON, _ := json.Marshal(login)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}
