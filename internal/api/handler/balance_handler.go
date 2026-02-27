package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AbdullahOztoprak/Backend-Path/internal/application/usecase"
	"github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"
	"github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/observability"
)

type BalanceHandler struct {
	balanceUseCase usecase.GetBalanceUseCase
	logger         observability.Logger
}

func NewBalanceHandler(balanceUseCase usecase.GetBalanceUseCase, logger observability.Logger) *BalanceHandler {
	return &BalanceHandler{
		balanceUseCase: balanceUseCase,
		logger:         logger,
	}
}

func (h *BalanceHandler) GetBalance(c *gin.Context) {
	userID := c.Param("user_id")

	balance, err := h.balanceUseCase.Execute(userID)
	if err != nil {
		h.logger.Error("Failed to get balance", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func (h *BalanceHandler) UpdateBalance(c *gin.Context) {
	var balance entity.Balance
	if err := c.ShouldBindJSON(&balance); err != nil {
		h.logger.Error("Invalid input", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.balanceUseCase.Update(&balance); err != nil {
		h.logger.Error("Failed to update balance", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Balance updated successfully"})
}