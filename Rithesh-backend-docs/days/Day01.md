# Day 01 - Add Postgres (Docker) + migrations tool (goose)

## Goal (end of day)
- Postgres runs locally in Docker
- Go API connects using `DATABASE_URL`
- migrations run from API startup

## Learn (45-90 minutes)
- HTTP status codes: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Status`
- Postgres basics: `https://www.postgresql.org/docs/`
- goose migrations: `https://github.com/pressly/goose`

## Tasks
### 1) Add Postgres compose service
Create `infra/docker-compose.yml` (or root `docker-compose.yml`) with Postgres 16.

### 2) Add API env file
Create `services/api/.env`:
- `DATABASE_URL=postgres://askhub:askhub@localhost:5432/askhubtravel?sslmode=disable`
- `PORT=8080`

### 3) Add DB connection package
Create `services/api/internal/db/db.go`:
- reads `DATABASE_URL`
- returns `pgxpool.Pool`

### 4) Add migration runner
Create:
- `services/api/internal/migrate/migrate.go`
- `services/api/internal/migrate/migrations/` (empty for now)

Wire migration runner from `main.go`.

### 5) Run DB + API
```powershell
docker compose up -d
cd services/api
go run .
```

## Acceptance checks
- API starts while DB is running
- migration runner executes (no panic)

## Deliverables
- API connected to Postgres with migration wiring.
