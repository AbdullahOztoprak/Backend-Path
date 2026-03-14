package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware that logs incoming requests and responses.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request
		log.Printf("incoming request method=%s url=%s remote_addr=%s", r.Method, r.URL.String(), r.RemoteAddr)

		// Create a response writer to capture the response status
		rec := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rec, r)

		// Log the response
		log.Printf("response sent status=%d duration=%s", rec.statusCode, time.Since(start))
	})
}

// responseWriter is a custom http.ResponseWriter that captures the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and calls the original WriteHeader.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}