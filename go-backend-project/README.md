# Quick Start

1. Install Go and Docker Desktop.
2. Clone the repo:
	```
	git clone <your-repo-url>
	cd go-backend-project
	```
3. Start Docker Desktop.
4. Run PostgreSQL:
	```
	docker run --name postgres-db -e POSTGRES_PASSWORD=abdullah -p 5432:5432 -d postgres:17
	```
5. Create the database and tables:
	```
	docker exec -it postgres-db psql -U postgres
	CREATE DATABASE go_backend_db;
	\q
	docker cp internal/db/schema.sql postgres-db:/schema.sql
	docker exec -it postgres-db psql -U postgres -d go_backend_db -f /schema.sql
	```
6. Copy `.env.example` to `.env` and fill in:
	```
	PORT=8081
	DATABASE_URL=postgres://postgres:abdullah@localhost:5432/go_backend_db?sslmode=disable
	```
7. Run the server:
	```
	go run ./cmd/main.go
	```
# Go Backend Project

A simple backend API built with Go for managing users, transactions, and balances.

## What This Project Does

- **User Management**: Register and manage user accounts
- **Transactions**: Handle money transfers between users
- **Balance Tracking**: Keep track of user balances safely
- **Concurrent Processing**: Process multiple transactions at the same time
- **Database**: Store data in PostgreSQL

## Features

- Clean and simple code structure
- Environment-based configuration (`.env` file)
- Logging with structured logs
- Safe server shutdown
- Thread-safe operations for balances
- Worker pools for processing transactions
- Docker support for easy setup

## Quick Start

### 1. Requirements
- Go 1.21 or newer
- Docker Desktop
- Git

### 2. Setup
```bash
# Clone the project
git clone <your-repo-url>
cd go-backend-project

# Start PostgreSQL database
docker run --name postgres-db -e POSTGRES_PASSWORD=abdullah -p 5432:5432 -d postgres:17

# Create database and tables
docker exec -it postgres-db psql -U postgres
CREATE DATABASE go_backend_db;
\q

# Load database schema
docker cp internal/db/schema.sql postgres-db:/schema.sql
docker exec -it postgres-db psql -U postgres -d go_backend_db -f /schema.sql

# Create .env file
PORT=8080
DATABASE_URL=postgres://postgres:abdullah@localhost:5432/go_backend_db

# Run the application
go run ./cmd/main.go
```

### 3. Test
The server will start on `http://localhost:8080`

## Project Structure

```
├── cmd/                    # Main application
├── internal/
│   ├── models/            # Data structures (User, Transaction, Balance)
│   ├── repository/        # Database operations
│   ├── service/           # Business logic
│   ├── worker/            # Background processing
│   └── db/               # Database connection and schema
├── pkg/                   # Shared packages
└── configs/              # Configuration files
```

## Key Components

- **Models**: Define data structures with validation
- **Repositories**: Handle database operations
- **Services**: Contain business logic
- **Workers**: Process transactions concurrently
- **Balance**: Thread-safe balance updates

## Contributing

1. Create a new branch: `git checkout -b feature/your-feature`
2. Make your changes
3. Commit: `git commit -m "feat: add your feature"`
4. Push: `git push origin feature/your-feature`
5. Create a Pull Request

## License

MIT
