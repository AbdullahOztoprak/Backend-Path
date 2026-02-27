package observability

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
    Status string `json:"status"`
}

// HealthCheck handles health check requests
func HealthCheck(c *gin.Context) {
    response := HealthCheckResponse{
        Status: "healthy",
    }
    c.JSON(http.StatusOK, response)
}