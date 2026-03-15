# Architecture

## High-Level System Overview
Backend-Path is a layered Go backend that exposes HTTP endpoints for authentication, transfer operations, and balance retrieval. The design separates transport, application orchestration, domain contracts, and infrastructure adapters.

## Layered Architecture
Core direction:
- transport -> use case -> domain
- infrastructure implements domain-facing contracts

This keeps business rules independent from HTTP framework and persistence details.

## API Layer Responsibilities
Location: `internal/api`
- Route composition (`router.go`).
- Middleware chain registration.
- Request decode/response encode concerns in handlers.
- Delegation to application use case interfaces.

## Application / Use Case Layer Responsibilities
Location: `internal/application/usecase`
- Orchestrates business workflows (register, login, refresh, transfer, list transactions, get balance).
- Coordinates repository/auth/idempotency dependencies.
- Uses `context.Context` from incoming requests.

## Domain Layer Responsibilities
Location: `internal/domain`
- Defines entities (`User`, `Transaction`, `Balance`).
- Defines repository interfaces as contracts.
- Provides business-oriented invariants and model semantics.

## Infrastructure Layer Responsibilities
Location: `internal/infrastructure`
- Implements concrete adapters for auth, persistence, messaging, and observability.
- PostgreSQL and Redis integration code lives here.
- Must conform to domain repository interfaces.

## Worker Layer Responsibilities
Location: `internal/worker`
- Handles asynchronous/retry-oriented processing.
- Includes retry and dead-letter queue style modules.
- Keeps background concerns separate from request-response paths.

## Dependency Direction Rules
- Domain must not depend on API or infrastructure packages.
- Use cases depend on domain contracts, not concrete DB clients.
- API depends on use case contracts.
- Infrastructure depends on domain contracts and external libraries.

## Request Lifecycle
1. Router receives request.
2. Middleware chain executes before handlers.
3. Handler validates payload and calls use case with `r.Context()`.
4. Use case coordinates domain rules and repository operations.
5. Infrastructure adapters read/write external systems.
6. Handler returns HTTP response.

## Persistence and Repository Pattern
Repository interfaces are defined in `internal/domain/repository`. PostgreSQL implementations in `internal/infrastructure/persistence/postgres` satisfy those interfaces. This enforces stable use-case contracts and reduces coupling to storage details.

## Middleware Chain
Requests flow through middleware before handlers, including:
- request ID
- logging
- recovery
- CORS
- auth and RBAC for protected routes
- rate limiting

## Observability Role
Observability utilities under `internal/infrastructure/observability` provide:
- metrics middleware and handler
- tracer setup helpers
- logger setup helpers

## Reliability Concerns
Current reliability mechanisms include:
- middleware safety layers (recovery, rate limiting, auth)
- worker modules for retries/background handling
- idempotency utility package for transfer safety patterns

## Current Hardening Priorities
- Complete runtime dependency wiring in `cmd/main.go`.
- Expand integration test confidence around persistence/auth paths.
- Continue CI/deployment hardening without weakening architectural boundaries.
