package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

var rolePermissions = map[Role][]string{
	Admin: {"/api/v1/users", "/api/v1/transactions", "/api/v1/balances"},
	User:  {"/api/v1/transactions", "/api/v1/balances"},
}

func RBAC(requiredRole Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_jwt_secret"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userRole := Role(claims.Subject) // Assuming the role is stored in the subject field
		if !isAuthorized(userRole, requiredRole, c.Request.URL.Path) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func isAuthorized(userRole Role, requiredRole Role, path string) bool {
	allowedPaths, exists := rolePermissions[userRole]
	if !exists {
		return false
	}

	for _, allowedPath := range allowedPaths {
		if allowedPath == path {
			return true
		}
	}
	return false
}