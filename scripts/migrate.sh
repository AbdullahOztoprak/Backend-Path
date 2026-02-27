#!/bin/bash

set -e

# Database migration script for the Go backend project

# Load environment variables from .env file
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

# Database connection parameters
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-your_secure_password}
DB_NAME=${DB_NAME:-go_backend_db}

# Run migrations
echo "Running database migrations..."
migrate -path internal/db/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

echo "Migrations completed successfully."