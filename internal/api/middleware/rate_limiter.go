package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

var defaultLimiter = rate.NewLimiter(20, 50)

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !defaultLimiter.Allow() {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}