# Day 07 - Events (arrived/left/skip) + drift status

## Goal (end of day)
- events are stored in DB
- endpoints work:
  - `POST /day-sessions/{id}/events`
  - `GET /day-sessions/{id}/events`
- `GET /day-sessions/{id}` includes drift status

## Learn (30-60 minutes)
- RFC3339 timestamps: `https://www.rfc-editor.org/rfc/rfc3339`
- Go time: `https://pkg.go.dev/time`

## Tasks
### 1) Migration
Create `0007_create_events.sql` with:
- id, day_session_id, type, ts, payload_json, created_at

### 2) Model + repo
Create:
- `internal/model/event.go`
- `internal/repo/event_repo.go`

### 3) Add event endpoints
- `POST /day-sessions/{id}/events`
- `GET /day-sessions/{id}/events`

### 4) Drift module
Create `internal/drift/drift.go` to compute:
- `drift_minutes`
- `drift_status` (`on_time|late|early`)
- `replan_recommended` (`true` when drift >= 20)

### 5) Attach drift on day-session read
Return:
- `{ day_session: {...}, status: {...} }`

## Acceptance checks
- posting `arrived` event changes drift output

## Deliverables
- live status foundation for replanning.
