package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/AbdullahOztoprak/Backend-Path/internal/application/usecase"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/auth"
	"github.com/AbdullahOztoprak/Backend-Path/internal/api/dto"
)

type AuthHandler struct {
	authService usecase.AuthService
}

func NewAuthHandler(authService usecase.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// RegisterUser handles user registration
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.authService.RegisterUser(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login
func (h *AuthHandler) LoginUser(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := h.authService.LoginUser(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newToken, err := h.authService.RefreshToken(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}