package main

import (
    "context"
    "errors"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "log"

    "github.com/AbdullahOztoprak/Backend-Path/internal/api"
    "github.com/AbdullahOztoprak/Backend-Path/internal/api/handler"
)

func main() {
    deps := api.Dependencies{
        HealthHandler:      handler.NewHealthHandler(),
        AuthHandler:        handler.NewAuthHandler(nil),
        UserHandler:        handler.NewUserHandler(nil),
        TransactionHandler: handler.NewTransactionHandler(nil),
        BalanceHandler:     handler.NewBalanceHandler(nil),
    }

    r := api.NewRouter(deps)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }

    server := &http.Server{
        Addr:              ":" + port,
        Handler:           r,
        ReadHeaderTimeout: 5 * time.Second,
    }

    go func() {
        log.Printf("Starting server on port %s", port)
        if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Fatalf("could not start server: %v", err)
        }
    }()

    shutdownSignal := make(chan os.Signal, 1)
    signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)
    <-shutdownSignal

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Printf("graceful shutdown failed: %v", err)
        return
    }

    log.Println("server stopped gracefully")
}