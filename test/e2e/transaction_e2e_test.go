package e2e

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/AbdullahOztoprak/Backend-Path/internal/api"
    "github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/persistence/postgres"
)

func TestTransactionE2E(t *testing.T) {
    // Setup the database connection and router
    db, err := postgres.NewConnection()
    if err != nil {
        t.Fatalf("could not connect to database: %v", err)
    }
    defer db.Close()

    router := api.SetupRouter(db)

    // Create a test user for authentication
    user := map[string]interface{}{
        "username": "test_user",
        "email":    "test@example.com",
        "password": "password123",
        "role":     "user",
    }

    userBody, _ := json.Marshal(user)
    req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(userBody))
    req.Header.Set("Content-Type", "application/json")

    // Create user
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)
    assert.Equal(t, http.StatusCreated, resp.Code)

    // Authenticate the user
    login := map[string]string{
        "username": "test_user",
        "password": "password123",
    }

    loginBody, _ := json.Marshal(login)
    req, _ = http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(loginBody))
    req.Header.Set("Content-Type", "application/json")

    resp = httptest.NewRecorder()
    router.ServeHTTP(resp, req)
    assert.Equal(t, http.StatusOK, resp.Code)

    var loginResponse map[string]interface{}
    json.Unmarshal(resp.Body.Bytes(), &loginResponse)
    token := loginResponse["token"].(string)

    // Create a transaction
    transaction := map[string]interface{}{
        "from_user_id": 1,
        "to_user_id":   2,
        "amount":       100.50,
        "description":  "Payment for services",
    }

    transactionBody, _ := json.Marshal(transaction)
    req, _ = http.NewRequest("POST", "/api/v1/transactions", bytes.NewBuffer(transactionBody))
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    resp = httptest.NewRecorder()
    router.ServeHTTP(resp, req)
    assert.Equal(t, http.StatusCreated, resp.Code)

    // Verify transaction was created
    // Additional checks can be added here to verify the transaction in the database
}