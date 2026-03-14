package api

import (
    "net/http"

    "github.com/gorilla/mux"

    "github.com/AbdullahOztoprak/Backend-Path/internal/api/handler"
    "github.com/AbdullahOztoprak/Backend-Path/internal/api/middleware"
    "github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/observability"
)

type Dependencies struct {
    HealthHandler      *handler.HealthHandler
    AuthHandler        *handler.AuthHandler
    UserHandler        *handler.UserHandler
    TransactionHandler *handler.TransactionHandler
    BalanceHandler     *handler.BalanceHandler
}

func NewRouter(deps Dependencies) *mux.Router {
    r := mux.NewRouter()
    r.Use(middleware.RequestIDMiddleware)
    r.Use(middleware.LoggingMiddleware)
    r.Use(middleware.RecoveryMiddleware)
    r.Use(middleware.CORSMiddleware)
    r.Use(observability.MetricsMiddleware)
    r.Use(middleware.RateLimiterMiddleware)

    healthHandler := deps.HealthHandler
    if healthHandler == nil {
        healthHandler = handler.NewHealthHandler()
    }

    authHandler := deps.AuthHandler
    if authHandler == nil {
        authHandler = handler.NewAuthHandler(nil)
    }

    userHandler := deps.UserHandler
    if userHandler == nil {
        userHandler = handler.NewUserHandler(nil)
    }

    transactionHandler := deps.TransactionHandler
    if transactionHandler == nil {
        transactionHandler = handler.NewTransactionHandler(nil)
    }

    balanceHandler := deps.BalanceHandler
    if balanceHandler == nil {
        balanceHandler = handler.NewBalanceHandler(nil)
    }

    v1 := r.PathPrefix("/api/v1").Subrouter()
    v1.HandleFunc("/health", healthHandler.HealthCheck).Methods(http.MethodGet)
    v1.HandleFunc("/auth/login", authHandler.LoginUser).Methods(http.MethodPost)
    v1.HandleFunc("/auth/refresh", authHandler.RefreshToken).Methods(http.MethodPost)
    v1.HandleFunc("/users", userHandler.RegisterUser).Methods(http.MethodPost)

    protected := v1.NewRoute().Subrouter()
    protected.Use(middleware.AuthMiddleware)
    protected.HandleFunc("/transactions", transactionHandler.TransferFunds).Methods(http.MethodPost)
    protected.HandleFunc("/transactions", transactionHandler.ListTransactions).Methods(http.MethodGet)
    protected.HandleFunc("/balances", balanceHandler.GetBalance).Methods(http.MethodGet)

    r.Handle("/metrics", observability.MetricsHandler()).Methods(http.MethodGet)

    return r
}