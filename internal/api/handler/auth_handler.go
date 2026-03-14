package handler

import (
	"context"
	"net/http"

	"github.com/AbdullahOztoprak/Backend-Path/internal/api/dto"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type AuthUseCase interface {
	Login(ctx context.Context, username, password string) (TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (TokenPair, error)
}

type AuthHandler struct {
	authUseCase AuthUseCase
}

func NewAuthHandler(authUseCase AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

// LoginUser handles user login.
func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	if h.authUseCase == nil {
		writeError(w, http.StatusServiceUnavailable, "auth service not configured")
		return
	}

	var req dto.UserLoginRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid input")
		return
	}

	tokens, err := h.authUseCase.Login(r.Context(), req.Username, req.Password)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, tokens)
}

// RefreshToken rotates and returns a new token pair.
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if h.authUseCase == nil {
		writeError(w, http.StatusServiceUnavailable, "auth service not configured")
		return
	}

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid input")
		return
	}
	if req.RefreshToken == "" {
		writeError(w, http.StatusBadRequest, "refresh_token is required")
		return
	}

	tokens, err := h.authUseCase.Refresh(r.Context(), req.RefreshToken)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, tokens)
}
