package configs

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Port              string `yaml:"port"`
    DatabaseURL       string `yaml:"database_url"`
    RedisURL          string `yaml:"redis_url"`
    JWTSecret         string `yaml:"jwt_secret"`
    LogLevel          string `yaml:"log_level"`
    RateLimit         int    `yaml:"rate_limit"`
    Environment       string `yaml:"environment"`
}

func LoadConfig(filePath string) (*Config, error) {
    var config Config

    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using default environment variables")
    }

    // Open the YAML configuration file
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Decode the YAML configuration
    decoder := yaml.NewDecoder(file)
    if err := decoder.Decode(&config); err != nil {
        return nil, err
    }

    // Override with environment variables if set
    if port := os.Getenv("PORT"); port != "" {
        config.Port = port
    }
    if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
        config.DatabaseURL = dbURL
    }
    if redisURL := os.Getenv("REDIS_URL"); redisURL != "" {
        config.RedisURL = redisURL
    }
    if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
        config.JWTSecret = jwtSecret
    }
    if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
        config.LogLevel = logLevel
    }
    if rateLimit := os.Getenv("RATE_LIMIT"); rateLimit != "" {
        parsedRateLimit, err := strconv.Atoi(rateLimit)
        if err != nil {
            return nil, err
        }
        config.RateLimit = parsedRateLimit
    }
    if environment := os.Getenv("ENVIRONMENT"); environment != "" {
        config.Environment = environment
    }

    return &config, nil
}