# Day 02 - Trips table + create/list endpoints (first real feature)

## Goal (end of day)
- `trips` table exists via migration
- endpoints work:
  - `POST /trips`
  - `GET /trips`

## Learn (45-90 minutes)
- Postgres DDL: `https://www.postgresql.org/docs/current/ddl.html`
- HTTP POST: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST`

## Tasks
### 1) Create migrations
Create:
- `0001_create_trips.sql`
- `0002_add_trip_prefs.sql`

`0001_create_trips.sql` should include:
- `CREATE EXTENSION IF NOT EXISTS pgcrypto;`
- `trips` columns:
  - `id uuid primary key default gen_random_uuid()`
  - `destination text not null`
  - `start_date date not null`
  - `end_date date not null`
  - `travelers_count integer not null check (travelers_count >= 1)`
  - `pace text`
  - `budget_level integer`
  - `interests_json jsonb`
  - `dining_json jsonb`
  - `created_at timestamptz not null default now()`
- index `idx_trips_created_at` on `(created_at DESC)`

`0002_add_trip_prefs.sql` can use `ADD COLUMN IF NOT EXISTS` for pref columns.

### 2) Create model + repo
Create:
- `internal/model/trip.go`
- `internal/repo/trip_repo.go`

### 3) Add handlers
In `main.go` add:
- `POST /trips`
- `GET /trips`

Validate:
- destination required
- start/end date required
- travelers count >= 1

### 4) Test
```powershell
Invoke-RestMethod http://localhost:8080/trips -Method Post -ContentType 'application/json' -Body '{"destination":"Dubai","start_date":"2026-10-12","end_date":"2026-10-18","travelers_count":2}' | ConvertTo-Json -Depth 8
Invoke-RestMethod http://localhost:8080/trips -Method Get | ConvertTo-Json -Depth 8
```

## Acceptance checks
- newly created trip appears in list output

## Deliverables
- first DB-backed feature complete.
