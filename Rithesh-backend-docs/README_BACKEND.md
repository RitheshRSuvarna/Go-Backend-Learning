# Backend Runbook (Day 28)

This repo contains a Go API (DDD/hexagonal, sqlc, Atlas migrations) plus a Python LLM microservice.

## Services

- **Go API**: `test/test` (default port `8080`)
- **Python LLM service**: `test/test/services/llm` (default port `8090`)
- **Postgres**: Docker Compose (exposes `5439` on host)

## Prerequisites

- Docker Desktop
- Go (1.22+ recommended)
- Python 3.11+
- Taskfile (`task`) installed

## Environment variables

Create `test/test/.env`:

```env
PORT=8080
DATABASE_URL=postgres://test:test@localhost:5439/test?sslmode=disable
LLM_BASE_URL=http://localhost:8090

# Enable dev-only endpoints/debug (demo reset, replan_debug)
APP_ENV=dev
```

Optional:

```env
# If you prefer a dedicated flag instead of APP_ENV
DEMO_RESET_ENABLED=true
REPLAN_DEBUG=true
```

## Run locally (recommended)

### 1) Start Postgres + run migrations + start API (hot reload)

```bash
cd test/test
task dev
```

This will:
- generate code (sqlc/graphql/swagger)
- start Postgres via compose
- apply migrations
- run the API via Air

### 2) Start the LLM service

```bash
cd test/test/services/llm
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
uvicorn app:app --reload --port 8090
```

#### Optional: use real LLM provider (OpenAI)

If you set `OPENAI_API_KEY`, the LLM service will call OpenAI using structured outputs (JSON schema) and return validated responses.

```bash
export OPENAI_API_KEY="..."
export OPENAI_MODEL="gpt-4o-mini"
```

## Smoke checks

```bash
curl -sS http://localhost:8080/health
curl -sS http://localhost:8090/docs | head
```

## Demo seed/reset (dev-only)

Requires `APP_ENV=dev` (or `DEMO_RESET_ENABLED=true`).

```bash
curl -sS -X POST http://localhost:8080/demo/reset -H "Content-Type: application/json" -d '{}' | jq .
```

The response includes IDs for `trip`, `day_session`, and `plan_versions`.

## Example flows

See:
- `Rithesh-backend-docs/examples/auth-flow.md`
- `Rithesh-backend-docs/examples/planning-flow.md`
- `Rithesh-backend-docs/examples/replan-flow.md`

## Troubleshooting

- **`go test ./...` fails under `test/test`**: this repo uses `go.work` with multiple modules; run module tests explicitly:

```bash
cd test/test
go test ./internal/modules/planning/...
go test ./internal/modules/auth/...
```

- **Migrations fail with syntax around `IF NOT EXISTS` in constraints**: Atlas parsing can be stricter than Postgres. Prefer `DROP CONSTRAINT IF EXISTS` in down migrations, and plain `ADD CONSTRAINT` in up migrations.

- **LLM endpoint returns 502 or timeouts**: ensure the LLM service is running on `8090` and `LLM_BASE_URL` matches. The API also uses request timeouts; a down LLM should return a clean JSON error with `request_id`.

