package handler

import (
	"context"
	"net/http"

	"github.com/AbdullahOztoprak/Backend-Path/internal/api/middleware"
)

type BalanceOutput struct {
	UserID           string  `json:"user_id"`
	AvailableBalance float64 `json:"available_balance"`
	Currency         string  `json:"currency"`
}

type BalanceUseCase interface {
	Get(ctx context.Context, userID string) (BalanceOutput, error)
}

type BalanceHandler struct {
	balanceUseCase BalanceUseCase
}

func NewBalanceHandler(balanceUseCase BalanceUseCase) *BalanceHandler {
	return &BalanceHandler{balanceUseCase: balanceUseCase}
}

func (h *BalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	if h.balanceUseCase == nil {
		writeError(w, http.StatusServiceUnavailable, "balance service not configured")
		return
	}

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok || userID == "" {
		writeError(w, http.StatusUnauthorized, "missing authenticated user")
		return
	}

	balance, err := h.balanceUseCase.Get(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, balance)
}
