package middleware

import (
	"net/http"
	"github.com/google/uuid"
)

// RequestIDMiddleware generates a unique request ID for each incoming request
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r)
	})
}