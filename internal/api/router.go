package api

import (
    "github.com/gorilla/mux"
    "net/http"
    "github.com/AbdullahOztoprak/Backend-Path/internal/api/middleware"
    "github.com/AbdullahOztoprak/Backend-Path/internal/api/handler"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // Health check
    router.HandleFunc("/api/v1/health", handler.HealthHandler).Methods(http.MethodGet)

    // User routes
    userRouter := router.PathPrefix("/api/v1/users").Subrouter()
    userRouter.HandleFunc("", handler.CreateUserHandler).Methods(http.MethodPost)
    userRouter.HandleFunc("", handler.ListUsersHandler).Methods(http.MethodGet)
    userRouter.Use(middleware.AuthMiddleware)

    // Authentication routes
    authRouter := router.PathPrefix("/api/v1/auth").Subrouter()
    authRouter.HandleFunc("/login", handler.LoginHandler).Methods(http.MethodPost)
    authRouter.HandleFunc("/refresh", handler.RefreshTokenHandler).Methods(http.MethodPost)

    // Transaction routes
    transactionRouter := router.PathPrefix("/api/v1/transactions").Subrouter()
    transactionRouter.HandleFunc("", handler.TransferFundsHandler).Methods(http.MethodPost)
    transactionRouter.HandleFunc("", handler.ListTransactionsHandler).Methods(http.MethodGet)
    transactionRouter.Use(middleware.AuthMiddleware)

    // Balance routes
    balanceRouter := router.PathPrefix("/api/v1/balances").Subrouter()
    balanceRouter.HandleFunc("", handler.GetBalanceHandler).Methods(http.MethodGet)
    balanceRouter.Use(middleware.AuthMiddleware)

    // Apply CORS middleware
    router.Use(middleware.CORSMiddleware)

    return router
}