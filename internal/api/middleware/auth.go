package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	infraauth "github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
)

type contextKey string

const (
	userIDContextKey contextKey = "user_id"
	rolesContextKey  contextKey = "roles"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		if tokenString == "" {
			http.Error(w, "authorization header is missing", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "dev-secret"
		}

		jwtProvider := infraauth.NewJWTProvider(secret, "backend-path", 24*time.Hour)
		userID, err := jwtProvider.ValidateToken(tokenString)
		if err != nil || userID == "" {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		roles := []string{}
		ctx := context.WithValue(r.Context(), userIDContextKey, userID)
		ctx = context.WithValue(ctx, rolesContextKey, roles)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDContextKey).(string)
	return userID, ok
}

func RolesFromContext(ctx context.Context) []string {
	roles, ok := ctx.Value(rolesContextKey).([]string)
	if !ok {
		return nil
	}
	return roles
}

