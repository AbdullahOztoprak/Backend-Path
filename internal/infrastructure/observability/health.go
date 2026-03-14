package observability

import (
    "encoding/json"
    "net/http"
)

type HealthCheckResponse struct {
    Status string `json:"status"`
}

// HealthCheck handles health check requests
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
    response := HealthCheckResponse{
        Status: "healthy",
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(response)
}