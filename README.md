# Backend-Path

A production-grade Go backend API for fintech-style money transfers. Built with clean architecture, JWT authentication, PostgreSQL transaction hardening, Redis rate limiting, Prometheus metrics, and full CI/CD via GitHub Actions.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.23 |
| Router | gorilla/mux |
| Database | PostgreSQL 15 (pgx/v4) |
| Cache / Rate Limit | Redis (go-redis/v8) |
| Auth | JWT (dgrijalva/jwt-go) |
| Logging | logrus |
| Config | YAML + env |
| Metrics | Prometheus + Grafana |
| Containerization | Docker + Docker Compose |
| CI/CD | GitHub Actions |
| Testing | testify, testcontainers |

---

## Features

- **JWT Authentication** — access token + refresh token
- **Role-Based Access Control** — `admin` and `user` roles
- **Transaction Hardening** — `SERIALIZABLE` isolation, `SELECT FOR UPDATE`, deadlock-safe ordering, idempotency keys
- **Balance Management** — thread-safe balance reads and updates
- **Rate Limiting** — per-IP request throttling via Redis
- **CORS Middleware** — configurable allowed origins
- **Structured Logging** — request ID and correlation ID on every log line
- **Prometheus Metrics** — `/metrics` endpoint with request counters and latency histograms
- **Health Checks** — `/api/v1/health` liveness probe
- **Graceful Shutdown** — context-aware server shutdown
- **Background Workers** — async transaction processing with retry and dead-letter queue
- **Migrations** — versioned SQL migrations via shell script
- **Full Test Suite** — unit, integration (testcontainers), E2E, and load tests (k6)

---

## Project Structure

```
Backend-Path/
├── cmd/
│   └── main.go                        # Entry point, wiring, graceful shutdown
├── configs/
│   ├── config.go                      # Config struct loader
│   ├── config.yaml                    # Default config
│   └── config.production.yaml        # Production overrides
├── internal/
│   ├── api/
│   │   ├── router.go                  # Route registration
│   │   ├── handler/                   # HTTP handlers
│   │   ├── middleware/                # Auth, CORS, logging, rate limiter, RBAC
│   │   └── dto/                       # Request / response types
│   ├── domain/
│   │   ├── entity/                    # Core domain models
│   │   ├── repository/                # Repository interfaces
│   │   └── service/                   # Domain service interfaces
│   ├── application/
│   │   ├── usecase/                   # Business logic (register, login, transfer, ...)
│   │   └── validator/                 # Input validation rules
│   ├── infrastructure/
│   │   ├── persistence/postgres/      # PostgreSQL repository implementations
│   │   ├── persistence/redis/         # Redis token store and rate limiter
│   │   ├── auth/                      # JWT provider, bcrypt hasher, RBAC policy
│   │   ├── observability/             # Logger, metrics, tracing, health
│   │   └── messaging/                 # Event publisher
│   ├── worker/                        # Async transaction worker, retry, DLQ
│   └── db/
│       ├── schema.sql
│       └── migrations/                # Versioned up/down SQL migrations
├── pkg/
│   ├── apperror/                      # Structured application errors
│   ├── idempotency/                   # Idempotency key helpers
│   └── pagination/                    # Cursor/offset pagination helpers
├── test/
│   ├── unit/                          # Unit tests (usecase, service, validator)
│   ├── integration/                   # Repo tests with testcontainers
│   ├── e2e/                           # Full HTTP round-trip tests
│   ├── load/                          # k6 load test script
│   ├── mocks/                         # Generated mocks
│   └── fixtures/                      # Seed data helpers
├── deployments/
│   ├── docker/                        # Dockerfile, Dockerfile.test
│   ├── docker-compose.yml
│   ├── docker-compose.test.yml
│   └── k8s/                           # Kubernetes manifests
├── monitoring/
│   ├── prometheus.yml
│   ├── alerting/rules.yml
│   └── grafana/dashboard.json
├── scripts/
│   ├── migrate.sh
│   ├── seed.sh
│   └── run_tests.sh
├── .github/workflows/                 # CI (build + test) and CD (Docker + K8s)
├── .env.example
├── .gitignore
├── .golangci.yml
├── Makefile
├── go.mod
└── go.sum
```

---

## Quick Start

### Prerequisites

- Go 1.23+
- Docker Desktop

### Option 1 — Docker Compose (recommended)

```bash
git clone https://github.com/AbdullahOztoprak/Backend-Path.git
cd Backend-Path

cp .env.example .env

docker-compose -f deployments/docker-compose.yml up --build
```

Server: `http://localhost:8081`

### Option 2 — Run locally

```bash
git clone https://github.com/AbdullahOztoprak/Backend-Path.git
cd Backend-Path

# Start Postgres and Redis
docker run -d --name pg -e POSTGRES_PASSWORD=secret -p 5432:5432 postgres:15
docker run -d --name redis -p 6379:6379 redis:alpine

cp .env.example .env
# Edit .env with your DB credentials

go mod tidy
make migrate
make run
```

---

## Environment Variables

| Variable | Default | Description |
|---|---|---|
| `PORT` | `8081` | HTTP listen port |
| `DATABASE_URL` | — | Full Postgres connection string |
| `REDIS_URL` | `redis://localhost:6379` | Redis connection string |
| `JWT_SECRET` | — | Access token signing secret |
| `REFRESH_TOKEN_SECRET` | — | Refresh token signing secret |
| `LOG_LEVEL` | `info` | Logging level |
| `RATE_LIMIT` | `100` | Requests per minute per IP |
| `CORS_ORIGINS` | `http://localhost:3000` | Comma-separated allowed origins |
| `IDEMPOTENCY_KEY_EXPIRY` | `24h` | TTL for idempotency records |

---

## API Reference

All protected routes require `Authorization: Bearer <access_token>`.

### Health

| Method | Path | Auth | Description |
|---|---|---|---|
| GET | `/api/v1/health` | No | Liveness probe |

### Auth

| Method | Path | Auth | Description |
|---|---|---|---|
| POST | `/api/v1/auth/login` | No | Login, returns access + refresh tokens |
| POST | `/api/v1/auth/refresh` | No | Rotate refresh token |

#### Login
```bash
curl -X POST http://localhost:8081/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"john_doe","password":"securepassword"}'
```
```json
{
  "access_token": "eyJ...",
  "refresh_token": "eyJ...",
  "expires_in": 3600
}
```

### Users

| Method | Path | Auth | Description |
|---|---|---|---|
| POST | `/api/v1/users` | No | Register a new user |
| GET | `/api/v1/users` | Yes | List all users (admin) |

#### Register
```bash
curl -X POST http://localhost:8081/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"username":"john_doe","email":"john@example.com","password":"securepassword","role":"user"}'
```

### Transactions

| Method | Path | Auth | Description |
|---|---|---|---|
| POST | `/api/v1/transactions` | Yes | Transfer funds |
| GET | `/api/v1/transactions` | Yes | List transactions |

#### Transfer funds
```bash
curl -X POST http://localhost:8081/api/v1/transactions \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -H "Idempotency-Key: uuid-v4-here" \
  -d '{"from_user_id":1,"to_user_id":2,"amount":100.50,"description":"Payment"}'
```

### Balances

| Method | Path | Auth | Description |
|---|---|---|---|
| GET | `/api/v1/balances` | Yes | Get current user balance |

---

## Development

### Make commands

```bash
make build    # Build binary to ./bin/
make run      # Build and run
make test     # Run all tests
make migrate  # Run database migrations
make seed     # Seed test data
make clean    # Remove build artifacts
```

### Running tests

```bash
# All tests
go test ./... -v

# Unit tests only
go test ./test/unit/... -v

# Integration tests (requires Docker)
go test ./test/integration/... -v

# E2E tests
go test ./test/e2e/... -v

# Load test (requires k6)
k6 run test/load/k6_load_test.js
```

### Linting

```bash
golangci-lint run
```

---

## CI/CD

GitHub Actions runs on every push to `main`:

1. `go mod tidy` + `go build`
2. `go test ./... -v`
3. Docker image build
4. Push to Docker registry
5. Deploy to Kubernetes

Workflow files: [.github/workflows/](.github/workflows/)

---

## Monitoring

Prometheus scrapes `/metrics`. Import `monitoring/grafana/dashboard.json` into Grafana for a pre-built dashboard.

```bash
# Start full monitoring stack
docker-compose -f deployments/docker-compose.yml up prometheus grafana
```

- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000`

---

## Troubleshooting

**Port 8081 already in use**
```bash
lsof -ti:8081 | xargs kill -9
```

**Database not ready**
```bash
docker logs <postgres-container>
# Wait for "database system is ready to accept connections", then restart the app
docker-compose restart
```

**Module import errors**
```bash
go mod tidy
```

---

## Contributing

1. Fork the repo
2. Create a branch: `git checkout -b feat/your-feature`
3. Commit using conventional commits: `git commit -m "feat: add X"`
4. Push and open a Pull Request

---

## License

MIT — see [LICENSE](LICENSE)

---

## Author

**Abdullah Öztoprak**
- GitHub: [@AbdullahOztoprak](https://github.com/AbdullahOztoprak)
- Project: [github.com/AbdullahOztoprak/Backend-Path](https://github.com/AbdullahOztoprak/Backend-Path)