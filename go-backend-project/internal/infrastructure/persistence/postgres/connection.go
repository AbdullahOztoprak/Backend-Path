package postgres

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/lib/pq"
)

type PostgresConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DbName   string
}

func NewPostgresConnection() (*sql.DB, error) {
    config := PostgresConfig{
        Host:     getEnv("DB_HOST", "localhost"),
        Port:     getEnvAsInt("DB_PORT", 5432),
        User:     getEnv("DB_USER", "postgres"),
        Password: getEnv("DB_PASSWORD", "your_secure_password"),
        DbName:   getEnv("DB_NAME", "go_backend_db"),
    }

    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.Host, config.Port, config.User, config.Password, config.DbName)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Set connection pool settings
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(time.Minute * 5)

    // Test the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Successfully connected to the PostgreSQL database")
    return db, nil
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

func getEnvAsInt(key string, fallback int) int {
    if value, exists := os.LookupEnv(key); exists {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return fallback
}