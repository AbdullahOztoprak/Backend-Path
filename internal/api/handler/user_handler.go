package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/AbdullahOztoprak/Backend-Path/internal/api/dto"
)

type CreateUserInput struct {
	Username string
	Email    string
	Password string
	Role     string
}

type CreateUserOutput struct {
	ID        string
	Username  string
	Email     string
	Role      string
	CreatedAt time.Time
}

type UserUseCase interface {
	Create(ctx context.Context, input CreateUserInput) (CreateUserOutput, error)
}

type UserHandler struct {
	userUseCase UserUseCase
}

func NewUserHandler(userUseCase UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

// RegisterUser handles user registration.
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if h.userUseCase == nil {
		writeError(w, http.StatusServiceUnavailable, "user service not configured")
		return
	}

	var req dto.UserRegistrationRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid input")
		return
	}

	role := req.Role
	if role == "" {
		role = "user"
	}

	out, err := h.userUseCase.Create(r.Context(), CreateUserInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     role,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"id":         out.ID,
		"username":   out.Username,
		"email":      out.Email,
		"role":       out.Role,
		"created_at": out.CreatedAt,
	})
}
