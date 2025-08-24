# Go Backend Project

A backend system built with Golang for managing users, transactions, balances, and audit logs.

## Features

- Modular and clean project structure
- Configuration via environment variables
- Structured logging with zerolog
- Graceful shutdown for the HTTP server
- RESTful API endpoints for core operations
- Middleware for authentication, authorization, validation, and error handling
- Dockerized deployment with docker-compose

## Getting Started

1. Clone the repository
2. Install Go (version 1.21 or newer)
3. Create a `.env` file and set your environment variables
4. Run the server:
   ```
   go run ./cmd/main.go
   ```

## Project Structure

- `cmd/` — Application entry point
- `pkg/` — Reusable packages
- `internal/` — Private application code
- `configs/` — Configuration files

## Contributing

- Use conventional commits and feature branches
- Open issues for new features or bugs
- Submit pull requests with clear