# Run Log (Day 29)

## Run 1
- Date/time: 2026-05-11
- Commands executed:
  - `cd test/test`
  - `task compose:up`
  - `task migrate:up`
- Result: pass (DB boot + migrations)
- Errors:
  - `atlas migrate apply` initially failed with `password authentication failed for user "test"` (stale docker volume credentials).
- Fixes applied:
  - `task migrate:up` auto-ran `docker compose ... down -v` and restarted postgres, then migrations applied cleanly.

## Run 2
- Date/time: 2026-05-11
- Commands executed:
  - Start DB:
    - `cd test/test`
    - `docker compose -f ./tools/docker/docker-compose.yml up -d postgres`
  - Migrate:
    - `task migrate:up`
  - Start LLM (local):
    - `cd test/test/services/llm`
    - `python3 -m pip install --user -r requirements.txt`
    - `python3 -m uvicorn app:app --port 8090`
  - Start API:
    - `cd test/test`
    - `APP_ENV=dev PORT=8080 DATABASE_URL='postgres://test:test@localhost:5439/test?sslmode=disable' LLM_BASE_URL='http://localhost:8090' go run ./cmd/api`
  - Seed:
    - `curl -X POST http://localhost:8080/demo/reset -H "Content-Type: application/json" -d '{}'`
  - Story checks:
    - `POST /api/auth/login`
    - `GET /api/auth/me`
    - `GET /api/day-sessions/{id}/active-plan`
    - `POST /api/day-sessions/{id}/replan`
    - `GET /api/plan-versions/{suggested}/diff?baseId={base}`
    - `PUT /api/day-sessions/{id}/active-plan/{planVersionId}`
- Result: pass
- Errors:
  - API initially panicked on startup due to duplicate `ServeMux` pattern registrations for `"/api/"`.
  - `POST /demo/reset` originally hung and timed out because it called a service that started a nested transaction.
- Fixes applied:
  - Register `"/api/"` once and delegate by sub-path (`/auth/*` vs planning routes).
  - Removed nested transaction usage from demo reset: it now creates plan versions/stops directly via repos inside the single transaction.

## Run 3
- Date/time: 2026-05-11
- Commands executed:
  - `POST /demo/reset` (repeat)
  - `POST /api/auth/login`
  - `GET /api/auth/me`
  - `GET /api/day-sessions/{id}/active-plan`
  - `POST /api/day-sessions/{id}/replan`
  - `GET /api/plan-versions/{suggested}/diff?baseId={base}`
  - `PUT /api/day-sessions/{id}/active-plan/{planVersionId}`
- Result: pass
- Errors: none
- Fixes applied: none

## Summary
- What still feels risky:
  - LLM production calls require valid `OPENAI_API_KEY` and provider availability.
- What is stable:
  - DB lifecycle + migrations are reproducible.
  - Demo reset + core story endpoints behave deterministically.

# Day29 Run Log (Fill This)

Use `RUN_LOG_TEMPLATE_DAY29.md` as base and fill results from your 3 rehearsals.
