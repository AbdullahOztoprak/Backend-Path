package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AbdullahOztoprak/Backend-Path/internal/api"
	"github.com/AbdullahOztoprak/Backend-Path/internal/api/dto"
	"github.com/AbdullahOztoprak/Backend-Path/internal/api/handler"
)

type fakeTransactionUseCase struct{}

func (f fakeTransactionUseCase) Create(_ context.Context, input handler.CreateTransactionInput) (handler.CreateTransactionOutput, error) {
	return handler.CreateTransactionOutput{
		ID:        "tx-1",
		Status:    "created",
		CreatedAt: time.Now(),
	}, nil
}

func (f fakeTransactionUseCase) List(_ context.Context, userID string, page, pageSize int) (handler.ListTransactionsOutput, error) {
	return handler.ListTransactionsOutput{
		Items: []dto.TransactionResponse{
			{ID: 1, FromUserID: 1, ToUserID: 2, Amount: 100.50, Description: "Payment", CreatedAt: time.Now().Format(time.RFC3339)},
		},
		Page:       page,
		PageSize:   pageSize,
		TotalItems: 1,
	}, nil
}

func TestTransactionContractE2E(t *testing.T) {
	token := issueTestToken(t)
	router := api.NewRouter(api.Dependencies{
		HealthHandler:      handler.NewHealthHandler(),
		TransactionHandler: handler.NewTransactionHandler(fakeTransactionUseCase{}),
	})

	transaction := map[string]interface{}{
		"from_user_id": 1,
		"to_user_id":   2,
		"amount":       100.50,
		"description":  "Payment for services",
	}

	body, _ := json.Marshal(transaction)
	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/transactions", bytes.NewReader(body))
	createReq.Header.Set("Authorization", "Bearer "+token)
	createReq.Header.Set("Content-Type", "application/json")
	createResp := httptest.NewRecorder()
	router.ServeHTTP(createResp, createReq)
	assert.Equal(t, http.StatusCreated, createResp.Code)

	listReq := httptest.NewRequest(http.MethodGet, "/api/v1/transactions?page=1&page_size=20", nil)
	listReq.Header.Set("Authorization", "Bearer "+token)
	listResp := httptest.NewRecorder()
	router.ServeHTTP(listResp, listReq)
	assert.Equal(t, http.StatusOK, listResp.Code)

	var listBody map[string]interface{}
	err := json.Unmarshal(listResp.Body.Bytes(), &listBody)
	require.NoError(t, err)
	_, hasItems := listBody["items"]
	assert.True(t, hasItems)
}

func issueTestToken(t *testing.T) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   "user-1",
		"roles": []string{"user"},
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	signed, err := token.SignedString([]byte("dev-secret"))
	require.NoError(t, err)
	return signed
}
