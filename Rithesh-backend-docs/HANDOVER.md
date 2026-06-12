# Handover Notes (Day 30)

## Service map

- **Go API**: `test/test` (default `:8080`)
- **Python LLM**: `test/test/services/llm` (default `:8090`)
- **Postgres**: `test/test/tools/docker/docker-compose.yml` (host port `5439`)

## Environment variables

### API
- `PORT` (default `8080`)
- `DATABASE_URL` (default `postgres://test:test@localhost:5439/test?sslmode=disable`)
- `LLM_BASE_URL` (default `http://localhost:8090`)
- `APP_ENV=dev` (enables dev-only endpoints like `POST /demo/reset` and includes `replan_debug`)

### LLM
- Optional:
  - `OPENAI_API_KEY` (enables real provider calls)
  - `OPENAI_MODEL` (default `gpt-4o-mini`)

## Run commands

See `README_BACKEND.md` and `Rithesh-backend-docs/DEMO_SCRIPT.md`.

## Known issues / gotchas

- `go test ./...` from `test/test` may fail due to `go.work` multi-module layout. Prefer:
  - `go test ./internal/modules/planning/...`
  - `go test ./internal/modules/auth/...`

- If migrations fail with auth errors, you likely have an old Docker volume with different credentials:
  - `docker compose -f test/test/tools/docker/docker-compose.yml down -v`

## Troubleshooting

- Use `X-Request-Id` to correlate client errors with server logs.
- Timeouts:
  - DB-heavy handlers: ~3s context timeout
  - LLM/replan handlers: ~5s context timeout

## Out of scope (intentional)

- Production auth hardening (CSRF protection, secure cookies, rate limiting)
- Real email delivery provider
- LLM prompt quality framework + evaluation dataset
- Production infrastructure (deployments, observability stack)

## Suggested next engineering tasks

- Add provider configuration / key management for the LLM service (already supports OpenAI via `OPENAI_API_KEY`).
- Add CI to run `task build` and unit tests on every PR.
- Add proper Swagger UI hosting (optional) if desired.
