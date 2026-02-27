package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/AbdullahOztoprak/Backend-Path/internal/api/middleware"
    "github.com/AbdullahOztoprak/Backend-Path/internal/api/router"
    "github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/observability"
    "github.com/AbdullahOztoprak/Backend-Path/internal/infrastructure/persistence/postgres"
    "github.com/AbdullahOztoprak/Backend-Path/configs"
)

func main() {
    // Load configuration
    config, err := configs.LoadConfig()
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    // Initialize observability
    observability.InitLogger(config.LogLevel)
    observability.InitMetrics()

    // Connect to the database
    db, err := postgres.Connect(config.DatabaseURL)
    if err != nil {
        log.Fatalf("could not connect to database: %v", err)
    }
    defer db.Close()

    // Set up the router
    r := mux.NewRouter()
    r.Use(middleware.LoggingMiddleware)
    r.Use(middleware.CORSMiddleware)
    r.Use(middleware.RequestIDMiddleware)
    r.Use(middleware.RateLimiterMiddleware)

    // Initialize routes
    router.SetupRoutes(r)

    // Start the server
    log.Printf("Starting server on port %s", config.Port)
    if err := http.ListenAndServe(":"+config.Port, r); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}