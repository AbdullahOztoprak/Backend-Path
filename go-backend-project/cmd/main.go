package main

import (
    "os"

    "github.com/AbdullahOztoprak/go-backend-project/internal/db"
    "github.com/AbdullahOztoprak/go-backend-project/internal/api"

    "github.com/joho/godotenv"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
	
    "github.com/AbdullahOztoprak/go-backend-project/internal/repository"
    "github.com/AbdullahOztoprak/go-backend-project/internal/service"
	"context"
    "net/http"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    err := godotenv.Load()
    if err != nil {
        log.Fatal().Err(err).Msg("Error loading .env file")
    } 

    conn, err := db.Connect()
    if err != nil {
        log.Fatal().Err(err).Msg("Database connection failed")
        return
    }
    defer conn.Close(context.Background())

    port := os.Getenv("PORT")
    log.Info().Msgf("Server will start at port: %s", port)

    // Initialize repositories
    userRepo := repository.NewPGUserRepository(conn)
    txRepo := repository.NewPGTransactionRepository(conn)
    balanceRepo := repository.NewPGBalanceRepository(conn)

    // Initialize services
    userService := service.NewUserService(userRepo)
    transactionService := txRepo // PGTransactionRepository implements TransactionService
    balanceService := balanceRepo // PGBalanceRepository implements BalanceService

    router := api.NewRouter(userService, transactionService, balanceService)
    srv := &http.Server{
        Addr:    ":" + port,
        Handler: router,
    }

    // Run server in a goroutine
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal().Err(err).Msg("Server error")
        }
    }()
    log.Info().Msg("Server started")

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit
    log.Info().Msg("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Error().Err(err).Msg("Server forced to shutdown")
    }

    log.Info().Msg("Server exited")
}