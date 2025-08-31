# Go Backend API Project

Modern, secure and scalable backend API built with Go, featuring user management, transactions, and secure authentication.

## 📸 API Testleri (Postman)

### Kullanıcı Oluşturma
![Kullanıcı Oluşturma](docs/screenshots/register.png)

### Kullanıcı Girişi
![Kullanıcı Girişi](docs/screenshots/login.png)

### Para Gönderme
![Para Gönderme](docs/screenshots/transaction.png)

## 🎯 Project Overview

This is a professional Go backend project that provides:
- **User Registration & Authentication** with bcrypt password hashing
- **Transaction Management** between users
- **Balance Tracking** with thread-safe operations
- **RESTful API** with proper error handling
- **PostgreSQL Database** integration
- **Docker Support** for easy deployment
- **Comprehensive Testing** with unit tests

## 🚀 Technologies Used

- **Go 1.23+** - Modern programming language
- **PostgreSQL** - Reliable database
- **Docker & Docker Compose** - Containerization
- **bcrypt** - Secure password hashing
- **pgx** - PostgreSQL driver
- **zerolog** - Structured logging
- **godotenv** - Environment configuration

## 📋 Features

✅ **Security**: Password hashing, input validation, secure authentication  
✅ **Clean Architecture**: Repository pattern, service layer, dependency injection  
✅ **Error Handling**: Consistent JSON error responses  
✅ **Logging**: Structured logging with request middleware  
✅ **Testing**: Unit tests for business logic  
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
docker run --name postgres-db -e POSTGRES_PASSWORD=abdullah -p 5432:5432 -d postgres:15

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
docker run --name postgres-db -e POSTGRES_PASSWORD=abdullah -p 5432:5432 -d postgres:15

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

## 🏗️ Project Structure

```
go-backend-project/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── api/
│   │   └── router.go        # HTTP handlers and routing
│   ├── db/
│   │   ├── connection.go    # Database connection
│   │   └── schema.sql       # Database schema
│   ├── models/
│   │   ├── user.go          # User model with validation
│   │   ├── transaction.go   # Transaction model
│   │   └── balance.go       # Balance model
│   ├── repository/
│   │   ├── user_repository.go           # User data interface
│   │   ├── pg_user_repository.go       # PostgreSQL implementation
│   │   ├── transaction_repository.go   # Transaction data interface
│   │   └── balance_repository.go       # Balance data interface
│   ├── service/
│   │   ├── user_service.go          # User business logic interface
│   │   ├── user_service_impl.go     # User service implementation
│   │   ├── transaction_service.go   # Transaction business logic
│   │   └── balance_service.go       # Balance business logic
│   └── worker/
│       ├── transaction_worker.go    # Background transaction processing
│       └── batch_processor.go      # Batch operations
├── pkg/                    # Shared utilities
├── configs/               # Configuration files
├── Dockerfile            # Multi-stage Docker build
├── docker-compose.yml    # Complete deployment setup
├── .env.example         # Environment configuration template
└── README.md           # Project documentation
```

## 🧪 Running Tests

**For Linux/Mac:**
```bash
# Run all tests
go test ./internal/service/...

# Run tests with coverage
go test -cover ./internal/service/...

# Run specific test
go test -run TestRegisterValidation ./internal/service/...
```

**For Windows (PowerShell):**
```powershell
# Run all tests
go test ./internal/service/...

# Run tests with coverage
go test -cover ./internal/service/...

# Run specific test
go test -run TestRegisterValidation ./internal/service/...
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
- Check database logs: `docker logs go_backend_postgres`

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
- Docker for simplifying deployment
