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
*** Begin Patch
*** Add File: /Users/abdullah/Github2Backendpath/Backend-Path/README.md
#+ Backend-Path

Backend-Path is a production-oriented Go backend that demonstrates secure user management, balance handling, and transactional transfer workflows using a layered architecture, robust middleware stack, and observability tooling.

## Why This Project Matters
Backend-Path focuses on engineering concerns that matter in production systems: secure authentication, transactional integrity for transfers, idempotent operations, resilient background processing, and actionable observability. It's intended as a portfolio-quality reference for backend engineers and hiring teams.

## Current Status
- Transport refactor complete: canonical `gorilla/mux` + `net/http` stack.
- Transport, middleware, and JSON helpers standardized across handlers.
- E2E tests updated to match the canonical API contract.
- DI wiring in `cmd/main.go` is scaffolded; production dependency bootstrapping is being finalized.

## Implemented
- Layered architecture (transport → use cases → domain → infra).
- Authentication (JWT), RBAC middleware, structured request logging, and request tracing.
- Transactional transfer flows with idempotency support and DB migrations.
- Prometheus metrics and health/readiness endpoints.
- Background worker scaffolding for retries and dead-letter handling.
- Test fixtures and E2E tests under `test/`.

## In Progress
- Complete DI wiring to inject repository adapters and services into `cmd/main.go`.
- Harden CI: linting, caching, and pinned action versions.
- Add full integration tests that run against ephemeral DB/redis in CI.

## Planned
- Production deployment manifests and k8s validation (Helm/manifest tests).
- Performance tuning, benchmarking, and load testing for transfer throughput.
- Enhanced observability dashboards and SLO-driven alerts.

## Key Features
- Secure JWT-based authentication and role-based access control.
- Transactional transfers with idempotency and retry-safety.
- Consistent JSON API and centralized error model.
- Request-level correlation (`request_id`) and structured logs.
- Rate limiting, recovery, and observability (metrics + health checks).

## Architecture Overview
Backend-Path follows a clean, layered architecture: transport (router + middleware) handles HTTP concerns and validation; handlers map requests to application use cases; use cases contain orchestration and business rules; domain models express core invariants; infrastructure implements persistence, messaging, and observability. Background workers process deferred and retryable tasks outside user transactions.

![Architecture Diagram](docs/assets/architecture-diagram.png)

## Request Flow
1. The router receives an HTTP request on `/api/v1/*` and assigns a `request_id`.
2. Middleware applies tracing, structured logging, authentication, RBAC, and rate limiting.
3. The handler decodes input, validates it, and calls an application use case.
4. The use case coordinates domain rules and repository operations (within DB transactions when required).
5. Results are returned via consistent JSON responses; long-running work is pushed to the worker system when appropriate.

![Request Flow](docs/assets/request-flow-diagram.png)

## Tech Stack
- Language: Go (modules)
- HTTP router: `github.com/gorilla/mux` + `net/http`
- Database: PostgreSQL (pgx/sql drivers)
- Cache/queue: Redis (optional)
- Auth: JWT
- Observability: Prometheus + structured logging
- Containers: Docker, `docker-compose` for local stacks
- CI: GitHub Actions

## Project Structure
- `cmd/` — application entrypoint(s) (`cmd/main.go`)
- `internal/api/` — router, handlers, middleware
- `internal/application/` — use cases and validators
- `internal/domain/` — entities and business rules
- `internal/infrastructure/` — persistence, auth adapters, messaging, observability
- `internal/db/` — migrations and schema
- `internal/worker/` — background jobs and retry handlers
- `pkg/` — shared utilities (`apperror`, `idempotency`, `pagination`)
- `test/` — unit, integration, and e2e tests and fixtures

## Quick Start
Prerequisites: Go 1.20+, PostgreSQL, optional Redis.

1. Copy and edit configuration:

```bash
cp configs/config.yaml.example configs/config.yaml
# or use .env for local env vars
```

2. Run migrations:

```bash
./scripts/migrate.sh up
```

3. Build and run:

```bash
go build -o bin/backend ./cmd
./bin/backend --config configs/config.yaml
```

Or with Docker Compose for a local stack:

```bash
docker-compose -f deployments/docker-compose.yml up --build
```

## Local Development
- Use `configs/config.yaml` or environment variables for secrets and DB URLs.
- Focus on small, testable changes: handlers should delegate to use cases; use cases encapsulate business logic.
- Useful files: `cmd/main.go`, `internal/api/router.go`, `internal/application/usecase/*`.

## Testing
- Unit tests:

```bash
go test ./... -v
```

- Integration/E2E:

```bash
./scripts/run_tests.sh
```

- Benchmarks:

```bash
go test ./... -bench=. -benchmem
```

Notes: `test/` includes fixtures and example E2E flows. Integration tests expect a running DB and optional Redis (use the provided `docker-compose` for reproducible local runs).

## Observability
- Metrics endpoint: `/metrics` for Prometheus scraping.
- Health: `/healthz`; readiness: `/readyz`.
- Logs: structured JSON including `request_id` and trace fields.
- Dashboards and alerting rules are under `monitoring/` and `monitoring/alerting`.

## API Notes
- Base path: `/api/v1`
- Auth: Bearer JWT. Middleware injects user context into requests.
- Idempotency: Transfer endpoints accept `Idempotency-Key` headers.
- Error format: `{ "code": "...", "message": "...", "request_id": "..." }`.

Example transfer request:

```bash
curl -X POST http://localhost:8080/api/v1/transfers \
  -H "Authorization: Bearer $TOKEN" \
  -H "Idempotency-Key: <uuid>" \
  -H "Content-Type: application/json" \
  -d '{"from_user_id":"...","to_user_id":"...","amount":100.00}'
```

## Demo
- Start the local stack (`docker-compose`) and run example fixtures from `test/fixtures` to exercise auth, transfer, and balance flows.
- Review `test/e2e/` for sample flows that demonstrate typical usage.

## Design Decisions
- Transport: `gorilla/mux` + `net/http` chosen for explicit control and minimal runtime abstraction.
- Layering: transport handlers delegate to use cases to ensure business logic is testable and framework-agnostic.
- Transactions & Idempotency: transfers are executed inside DB transactions with idempotency keys to avoid double-processing.
- Workers: retryable and deferred tasks are handled by background workers to keep API latency bounded.

## Performance and Benchmarking
- Focus areas: transfer throughput, DB transaction latency, worker concurrency.
- Run targeted benchmarks:

```bash
go test ./internal/... -bench BenchmarkTransfer -benchmem
```

- For load testing, see `test/load/k6_load_test.js` and run against a staging environment.

## Roadmap
- Complete DI wiring and production-ready `cmd` bootstrapping.
- Harden CI/CD with security scanning and dependency pinning.
- Add deployment manifests and automated k8s validation.
- Implement SLOs and on-call playbooks for critical flows.

## Contributing
- Open issues and PRs are welcome. Keep changes small and focused.
- Run linters and tests locally before submitting PRs.
- Refer to `CONTRIBUTING.md` (if present) for branch and commit guidance.

## Release and Versioning
- Semantic Versioning (MAJOR.MINOR.PATCH).
- Tag releases and include changelogs describing migration steps and breaking changes.

## License
See the `LICENSE` file at the repository root for license terms. Contact the repository owner if no license is present.

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