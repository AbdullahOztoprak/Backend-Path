# Go Backend Project

A scalable backend system built with Golang, featuring user management, transactions, balances, audit logs, and robust API endpoints.

## Features

- Modular project structure
- Environment-based configuration
- Structured logging with zerolog
- Graceful shutdown for HTTP server
- User, transaction, balance, and audit log domain models
- RESTful API endpoints
- Middleware for authentication, authorization, validation, error handling, and monitoring
- Dockerized deployment with docker-compose

## Getting Started

1. **Clone the repository**
2. **Install Go (>=1.21)**
3. **Copy `.env.example` to `.env` and configure environment variables**
4. **Run the server**
   ```
   go run ./cmd/main.go
   ```

## Project Structure

```
cmd/         # Application entry point
pkg/         # Reusable packages
internal/    # Private application code
configs/     # Configuration files
```

## Contributing

- Use conventional commits and feature branches
- Open issues for new features or bugs
- Submit pull requests with clear descriptions

## License