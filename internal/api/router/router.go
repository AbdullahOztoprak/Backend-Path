package router

import (
    "github.com/gorilla/mux"

    "github.com/AbdullahOztoprak/Backend-Path/internal/api"
)

// NewRouter is a thin compatibility wrapper that forwards to api.NewRouter.
// This file provides a stable import path `internal/api/router` for CI or
// external references that expect the former package location.
func NewRouter(deps api.Dependencies) *mux.Router {
    return api.NewRouter(deps)
}
