package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/api/dto"
	"github.com/AbdullahOztoprak/Backend-Path/internal/api/middleware"
)

type CreateTransactionInput struct {
	FromUserID  int
	ToUserID    int
	Amount      float64
	Description string
}

type CreateTransactionOutput struct {
	ID        string
	Status    string
	CreatedAt time.Time
}

type ListTransactionsOutput struct {
	Items      []dto.TransactionResponse `json:"items"`
	Page       int                       `json:"page"`
	PageSize   int                       `json:"page_size"`
	TotalItems int                       `json:"total_items"`
}

type TransactionUseCase interface {
	Create(ctx context.Context, input CreateTransactionInput) (CreateTransactionOutput, error)
	List(ctx context.Context, userID string, page, pageSize int) (ListTransactionsOutput, error)
}

type TransactionHandler struct {
	transactionUseCase TransactionUseCase
}

func NewTransactionHandler(transactionUseCase TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{transactionUseCase: transactionUseCase}
}

// TransferFunds creates a transaction.
func (h *TransactionHandler) TransferFunds(w http.ResponseWriter, r *http.Request) {
	if h.transactionUseCase == nil {
		writeError(w, http.StatusServiceUnavailable, "transaction service not configured")
		return
	}

	var req dto.FundTransferRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	out, err := h.transactionUseCase.Create(r.Context(), CreateTransactionInput{
		FromUserID:  req.FromUserID,
		ToUserID:    req.ToUserID,
		Amount:      req.Amount,
		Description: req.Description,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"id":         out.ID,
		"status":     out.Status,
		"created_at": out.CreatedAt,
	})
}

// ListTransactions returns transactions for authenticated user.
func (h *TransactionHandler) ListTransactions(w http.ResponseWriter, r *http.Request) {
	if h.transactionUseCase == nil {
		writeError(w, http.StatusServiceUnavailable, "transaction service not configured")
		return
	}

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok || userID == "" {
		writeError(w, http.StatusUnauthorized, "missing authenticated user")
		return
	}

	page, pageSize := parsePagination(r)
	out, err := h.transactionUseCase.List(r.Context(), userID, page, pageSize)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, out)
}

func parsePagination(r *http.Request) (int, int) {
	page := 1
	pageSize := 20

	if raw := r.URL.Query().Get("page"); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil && v > 0 {
			page = v
		}
	}
	if raw := r.URL.Query().Get("page_size"); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil && v > 0 {
			pageSize = v
		}
	}

	return page, pageSize
}
