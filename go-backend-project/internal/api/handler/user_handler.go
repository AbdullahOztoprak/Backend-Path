package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your_project/internal/application/usecase"
	"github.com/your_project/internal/domain/entity"
	"github.com/your_project/internal/infrastructure/observability"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
	logger      observability.Logger
}

func NewUserHandler(userUseCase usecase.UserUseCase, logger observability.Logger) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
		logger:      logger,
	}
}

// RegisterUser handles user registration
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Failed to bind user data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.userUseCase.RegisterUser(&user); err != nil {
		h.logger.Error("Failed to register user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// GetUser retrieves a user by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := h.userUseCase.GetUserByID(userID)
	if err != nil {
		h.logger.Error("Failed to get user", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates user information
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Failed to bind user data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.userUseCase.UpdateUser(&user); err != nil {
		h.logger.Error("Failed to update user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user by ID
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if err := h.userUseCase.DeleteUser(userID); err != nil {
		h.logger.Error("Failed to delete user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deletion failed"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}