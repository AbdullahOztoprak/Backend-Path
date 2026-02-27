package middleware

import (
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/time/rate"
	"golang.org/x/net/context"
)

type RateLimiter struct {
	limiter  *rate.Limiter
	store    *redis.Client
}

func NewRateLimiter(store *redis.Client, r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(r, b),
		store:   store,
	}
}

func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ip := r.RemoteAddr

		// Check if the rate limit is exceeded
		if !rl.limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		// Store the request in Redis for tracking
		rl.store.Incr(ctx, ip)
		rl.store.Expire(ctx, ip, time.Minute)

		next.ServeHTTP(w, r)
	})
}