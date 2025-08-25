package db

import (
    "context"
    "os"
    "github.com/jackc/pgx/v5"
    "github.com/rs/zerolog/log"
)

func Connect() (*pgx.Conn, error) {
    dbURL := os.Getenv("DATABASE_URL")
    conn, err := pgx.Connect(context.Background(), dbURL)
    if err != nil {
        log.Error().Err(err).Msg("Failed to connect to database")
        return nil, err
    }
    log.Info().Msg("Connected to database")
    return conn, nil
}