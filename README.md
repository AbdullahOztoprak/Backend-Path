# Go Backend API Project

A modern, secure, and scalable backend API built with Go, featuring enhanced transaction hardening, robust authentication and authorization, clean architecture, observability, and automated testing.

## 📸 API Tests (Postman)

### User Registration
![User Registration](go-backend-project/docs/screenshots/register.png)

### User Login
![User Login](go-backend-project/docs/screenshots/login.png)

### Money Transfer
![Money Transfer](go-backend-project/docs/screenshots/transaction.png)

## 🎯 Project Overview

This project provides:
- **User Registration & Authentication** with bcrypt password hashing and JWT tokens
- **Role-Based Access Control (RBAC)** for fine-grained permissions
- **Transaction Management** with enhanced security features
- **Balance Tracking** with thread-safe operations
- **RESTful API** with comprehensive error handling
- **PostgreSQL Database** integration with migrations
- **Redis** for caching and rate limiting
- **Docker Support** for easy deployment
- **Observability** with structured logging, metrics, and tracing
- **Comprehensive Testing** with unit, integration, and end-to-end tests

## 🚀 Technologies Used

- **Go 1.23+** - Modern programming language
- **PostgreSQL** - Reliable database
- **Redis** - In-memory data structure store
- **Docker & Docker Compose** - Containerization
- **bcrypt** - Secure password hashing
- **pgx** - PostgreSQL driver
- **zerolog** - Structured logging
- **godotenv** - Environment configuration
- **Prometheus** - Monitoring and metrics
- **Grafana** - Visualization of metrics

## 📋 Features

✅ **Security**: Password hashing, JWT authentication, RBAC  
✅ **Clean Architecture**: Repository pattern, service layer, dependency injection  
✅ **Error Handling**: Consistent JSON error responses  
✅ **Logging**: Structured logging with request middleware  
✅ **Observability**: Metrics and tracing for monitoring  
✅ **Testing**: Unit, integration, and end-to-end tests  
✅ **Docker Ready**: Multi-stage Dockerfile and docker-compose setup  

## 🛠️ Quick Start

### Prerequisites
- **Go 1.23+** installed
- **Docker Desktop** running
- **Git** for cloning

### Option 1: Docker Compose (Recommended)

**For Linux/Mac:**
```bash
# Clone the repository
git clone https://github.com/AbdullahOztoprak/Backend-Path.git
cd Backend-Path/go-backend-project

# Copy environment file
cp .env.example .env

# Start everything with Docker
docker-compose up --build

# Server will be available at http://localhost:8081
```

**For Windows (PowerShell):**
```powershell
# Clone the repository
git clone https://github.com/AbdullahOztoprak/Backend-Path.git
cd Backend-Path/go-backend-project

# Copy environment file
Copy-Item .env.example .env

# Start everything with Docker
docker-compose up --build

# Server will be available at http://localhost:8081
```

### Option 2: Manual Setup

**For Linux/Mac:**
```bash
# Clone the repository
git clone https://github.com/AbdullahOztoprak/Backend-Path.git
cd Backend-Path/go-backend-project

# Start PostgreSQL database
docker run --name postgres-db -e POSTGRES_PASSWORD=your_password -p 5432:5432 -d postgres:15

# Create database and load schema
docker exec -it postgres-db psql -U postgres -c "CREATE DATABASE go_backend_db;"
docker cp internal/db/schema.sql postgres-db:/schema.sql
docker exec -it postgres-db psql -U postgres -d go_backend_db -f /schema.sql

# Copy and configure environment
cp .env.example .env

# Install dependencies and run
go mod tidy
go run cmd/main.go
```

**For Windows (PowerShell):**
```powershell
# Clone the repository
git clone https://github.com/AbdullahOztoprak/Backend-Path.git
cd Backend-Path/go-backend-project

# Start PostgreSQL database
docker run --name postgres-db -e POSTGRES_PASSWORD=your_password -p 5432:5432 -d postgres:15

# Create database and load schema
docker exec -it postgres-db psql -U postgres -c "CREATE DATABASE go_backend_db;"
docker cp internal/db/schema.sql postgres-db:/schema.sql
docker exec -it postgres-db psql -U postgres -d go_backend_db -f /schema.sql

# Copy and configure environment
Copy-Item .env.example .env

# Install dependencies and run
go mod tidy
go run cmd/main.go
```

## 📚 API Documentation

### Authentication
**POST** `/api/v1/login`
```json
{
  "username": "john_doe",
  "password": "securepassword"
}
```

### User Registration
**POST** `/api/v1/users`
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password_hash": "securepassword",
  "role": "user"
}
```

### Get All Users
**GET** `/api/v1/users`

### Transactions
**POST** `/api/v1/transactions`
```json
{
  "from_user_id": 1,
  "to_user_id": 2,
  "amount": 100.50,
  "description": "Payment for services"
}
```

### Get Balances
**GET** `/api/v1/balances`

## 🧪 Testing the API

### Using PowerShell (Windows)
```powershell
# Test if server is running
Invoke-RestMethod -Uri "http://localhost:8081/api/v1/users" -Method GET

# Create a new user
$body = @{
    username = "test_user"
    email = "test@example.com"
    password_hash = "password123"
    role = "user"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8081/api/v1/users" -Method POST -Body $body -ContentType "application/json"

# Login (if login endpoint exists)
$loginBody = @{
    username = "test_user"
    password = "password123"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8081/api/v1/login" -Method POST -Body $loginBody -ContentType "application/json"
```

### Using curl (Linux/Mac)
```bash
# Test if server is running
curl http://localhost:8081/api/v1/users

# Create a new user
curl -X POST http://localhost:8081/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"username":"test_user","email":"test@example.com","password_hash":"password123","role":"user"}'

# Login
curl -X POST http://localhost:8081/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test_user","password":"password123"}'
```

## 🏗️ Project Structure

```
go-backend-project/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── api/
│   │   └── ...              # API handlers, middleware, and DTOs
│   ├── domain/
│   │   └── ...              # Domain entities, repositories, and services
│   ├── application/
│   │   └── ...              # Use cases and validators
│   ├── infrastructure/
│   │   └── ...              # Persistence, auth, observability, and messaging
│   ├── worker/
│   │   └── ...              # Background processing
│   └── db/
│       └── ...              # Database schema and migrations
├── pkg/                     # Shared utilities
├── test/                    # Testing files
├── configs/                 # Configuration files
├── deployments/             # Deployment configurations
├── monitoring/              # Monitoring configurations
├── scripts/                 # Utility scripts
├── .env.example             # Environment configuration template
├── .github/                 # GitHub workflows
├── Makefile                 # Build and run commands
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── .golangci.yml            # Linting configuration
└── README.md                # Project documentation
```

## 🔧 Development

### Adding New Features
1. Create a new branch: `git checkout -b feat/your-feature-name`
2. Implement your changes following the existing patterns
3. Add tests for your new functionality
4. Update documentation if needed
5. Commit with conventional commit messages: `git commit -m "feat: add new feature"`
6. Push and create a Pull Request

### Code Organization
- **Models**: Define data structures with validation rules
- **Repositories**: Handle all database operations
- **Services**: Contain business logic and validation
- **API**: HTTP handlers and request/response handling
- **Workers**: Background processing and concurrent operations

## 🚀 Deployment

### Production Deployment
1. Set production environment variables
2. Use `docker-compose.prod.yml` for production setup
3. Configure reverse proxy (nginx) if needed
4. Set up SSL certificates
5. Configure monitoring and logging

### Environment Variables
```bash
PORT=8081
DATABASE_URL=postgres://user:password@host:port/dbname?sslmode=require
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_secure_password
DB_NAME=go_backend_db
JWT_SECRET=your_jwt_secret_for_production
REDIS_URL=redis://localhost:6379
```

## 🔧 Troubleshooting

### Common Issues

**Windows Users:**
- Use `Copy-Item .env.example .env` instead of `cp .env.example .env`
- Make sure Docker Desktop is running and set to Linux containers
- If you get "bind: Only one usage of each socket address" error, stop other services using port 8081

**Port Issues:**
```powershell
# Check what's using port 8081
netstat -ano | findstr :8081

# Kill process using the port (replace PID with actual process ID)
taskkill /PID <PID> /F
```

**Docker Issues:**
- Make sure Docker Desktop is running
- Try `docker-compose down` then `docker-compose up --build`
- Clear Docker cache: `docker system prune -a`

**Database Connection Issues:**
- Ensure PostgreSQL container is running: `docker ps`
- Check database logs: `docker logs postgres-db`
- If you see "the database system is starting up" errors, the application will automatically retry thanks to healthcheck and restart policies in docker-compose.yml
- For immediate restart of just the app: `docker-compose restart go_backend_app`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feat/amazing-feature`)
3. Commit your Changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the Branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Abdullah Öztoprak**
- GitHub: [@AbdullahOztoprak](https://github.com/AbdullahOztoprak)
- Project Link: [https://github.com/AbdullahOztoprak/Backend-Path](https://github.com/AbdullahOztoprak/Backend-Path)

## 🙏 Acknowledgments

- Go community for excellent documentation
- PostgreSQL team for reliable database
- Redis for caching and rate limiting
- Docker for simplifying deployment
- Prometheus and Grafana for observability