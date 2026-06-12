# Demo Script (Day 30)

## 1) Start Postgres

```bash
cd test/test
task compose:up
```

## 2) Apply migrations

```bash
cd test/test
task migrate:up
```

## 3) Run LLM service (terminal 1)

```bash
cd test/test/services/llm
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

# Optional: use real provider
export OPENAI_API_KEY="..."
export OPENAI_MODEL="gpt-4o-mini"

uvicorn app:app --reload --port 8090
```

## 4) Run API (terminal 2)

```bash
cd test/test
export APP_ENV=dev
export PORT=8080
export DATABASE_URL="postgres://test:test@localhost:5439/test?sslmode=disable"
export LLM_BASE_URL="http://localhost:8090"

go run ./cmd/api
```

## 5) Open Swagger artifacts

- OpenAPI JSON: `http://localhost:8080/api/swagger/swagger.json`
- (Optional) also check `Rithesh-backend-docs/examples/*.md` for runnable requests.

## 6) Seed demo data (dev-only)

```bash
curl -sS -X POST http://localhost:8080/demo/reset -H "Content-Type: application/json" -d '{}' | jq .
```

## 7) Execute story (manual, using examples)

Use:
- `Rithesh-backend-docs/examples/auth-flow.md`
- `Rithesh-backend-docs/examples/planning-flow.md`
- `Rithesh-backend-docs/examples/replan-flow.md`

## Demo talking points

- Deterministic rules (busy risk, drift, replan heuristics) vs LLM-driven plan generation
- DDD/hexagonal boundaries: domain entities + app services + ports/adapters + sqlc repos
- Reliability baseline: request IDs, timeouts, consistent error shape, and error logs
