# Project Guidelines

## Code Style
- Language: Go. Keep package-local layering consistent: `entity.go`, `input.go`, `service.go`, `pg_repository.go`, `handler.go`, `routes.go`, `init.go` (see `internal/user`, `internal/account`, `internal/transaction`).
- Prefer constructor helpers `newX(...)` and panic on invalid required infra dependencies (example: `newPgRepository` in `internal/account/pg_repository.go`).
- Use context-aware `pgx` DB calls (`QueryRow(ctx, ...)`) and explicit field scans into entities.
- Keep HTTP flow in handlers: bind -> validate -> call service -> map domain errors to status codes (see `internal/account/handler.go`).
- Keep exported surface minimal; expose service interfaces from `init.go` (example: `internal/account/init.go`).

## Architecture
- Entry point: `cmd/server/main.go`.
- App wiring happens in `main`: create Echo, DB, then initialize modules in dependency order: `user` -> `account` (depends on user service) -> `transaction` (depends on account service).
- Each bounded context under `internal/*` owns its HTTP routes, domain/service logic, and Postgres repository implementation.
- Shared infrastructure is split into:
  - DB config/env parsing in `configs/postgres.go`
  - DB connector in `pkg/postgres/main.go`
  - HTTP helper(s) in `internal/platform/httpx/request_id.go`

## Build and Test
- Start local infra: `make up` (uses `deployments/compose.yml`).
- Stop infra: `make down`; cleanup volumes: `make purge`; inspect: `make logs`, `make ps`.
- Run server locally: `go run ./cmd/server`.
- Run all tests: `go test ./...`.
- Run one package tests: `go test ./internal/user -v` (same pattern for other packages).

## Project Conventions
- Route registration is centralized per module in `routes.go` (example: `POST /accounts` in `internal/account/routes.go`).
- Domain errors are sentinel vars and checked with `errors.Is(...)` (examples: `internal/user/error.go`, `internal/account/handler.go`).
- Input structs expose `Validate()` and `ToEntity(...)` patterns before service calls.
- Repositories return domain entities and avoid leaking HTTP concerns.

## Integration Points
- HTTP framework: Echo (`github.com/labstack/echo`) with middleware in server setup (`cmd/server/main.go`).
- Postgres driver/pool: `github.com/jackc/pgx/v5` + `pgxpool`; SQL is inline in repository files.
- UUIDs use `github.com/google/uuid` in entities and service/repo boundaries.
- Environment-driven DB settings: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSLMODE`, pool/lifetime vars in `configs/postgres.go`.

## Security
- No authn/authz layer is implemented yet; do not assume protected routes.
- Treat request validation in `input.Validate()` as mandatory before service/repository calls.
- Keep DB credentials only in environment variables (or compose env), never hardcode secrets.
- Preserve request correlation via `X-Request-ID` usage (`internal/platform/httpx/request_id.go`) for traceability.
