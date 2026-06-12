# Day 24 - Performance pass (DB + endpoints)

## Goal (end of day)
- hot endpoints stay fast on larger stop counts
- indexes support real query patterns

## Learn (30-60 minutes)
- Postgres indexes: `https://www.postgresql.org/docs/current/indexes.html`
- Explain plans: `https://www.postgresql.org/docs/current/using-explain.html`

## Tasks
### 1) Identify hot paths
Focus on:
- `GET /day-sessions/{id}/active-plan`
- `POST /day-sessions/{id}/replan`
- event and plan version lookups

### 2) Review repo queries
Check for:
- N+1 patterns
- repeated lookups that can be joined
- missing order/index usage

### 3) Add justified indexes
Create migration only for indexes you can justify from query patterns.

### 4) Quick baseline measurement
Run each hot endpoint multiple times and note average latency.

### 5) Re-check after changes
Compare before/after numbers.

## Acceptance checks
- active-plan endpoint is stable with around 50 stops
- no new correctness regressions
- migration applies cleanly on fresh DB

## Deliverables
- documented perf improvements with before/after notes.
