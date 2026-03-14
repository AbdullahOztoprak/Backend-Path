package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AbdullahOztoprak/Backend-Path/internal/api"
	"github.com/AbdullahOztoprak/Backend-Path/internal/api/handler"
)

func TestHealthCheck(t *testing.T) {
	router := api.NewRouter(api.Dependencies{
		HealthHandler: handler.NewHealthHandler(),
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status code 200 OK")
}