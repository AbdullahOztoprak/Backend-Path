package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:8081/api/v1"

func TestUserRegistrationAndLogin(t *testing.T) {
	// Step 1: User Registration
	user := map[string]interface{}{
		"username":     "test_user",
		"email":        "test@example.com",
		"password_hash": "securepassword",
		"role":         "user",
	}

	userJSON, _ := json.Marshal(user)
	resp, err := http.Post(baseURL+"/users", "application/json", bytes.NewBuffer(userJSON))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Step 2: User Login
	login := map[string]string{
		"username": "test_user",
		"password": "securepassword",
	}

	loginJSON, _ := json.Marshal(login)
	resp, err = http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(loginJSON))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Step 3: Validate JWT Token
	var loginResponse map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&loginResponse)
	token, ok := loginResponse["token"].(string)
	assert.True(t, ok)

	// Step 4: Access Protected Route
	req, err := http.NewRequest("GET", baseURL+"/users", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUnauthorizedAccess(t *testing.T) {
	// Attempt to access protected route without token
	resp, err := http.Get(baseURL + "/users")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestInvalidLogin(t *testing.T) {
	// Attempt to login with invalid credentials
	login := map[string]string{
		"username": "test_user",
		"password": "wrongpassword",
	}

	loginJSON, _ := json.Marshal(login)
	resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(loginJSON))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestCleanup(t *testing.T) {
	// Cleanup test user from database if necessary
	// This can be done by calling a delete endpoint or directly manipulating the database
}