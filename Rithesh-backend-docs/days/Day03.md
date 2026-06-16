# Day 03 - Day Sessions (trip/day context)

## Goal (end of day)
- `day_sessions` table exists
- endpoints work:
  - `POST /day-sessions`
  - `GET /day-sessions?tripId=&date=`
  - `GET /day-sessions/{id}`
  - `PUT /day-sessions/{id}`

## Learn (30-75 minutes)
- SQL constraints + indexes: `https://www.postgresql.org/docs/current/ddl-constraints.html`
- Postgres UUID type: `https://www.postgresql.org/docs/current/datatype-uuid.html`

## Tasks
### 1) Migration
Create `0003_create_day_sessions.sql` with:
- `day_sessions(id uuid, trip_id uuid fk trips, date date, start_time text, start_label text, start_lat float8 null, start_lon float8 null, created_at timestamptz)`
- unique `(trip_id, date)`
- index `(trip_id, date)`

### 2) Model
Create `internal/model/day_session.go`:
- `DaySession`
- `CreateDaySessionRequest`
- `UpdateDaySessionRequest`

### 3) Repo
Create `internal/repo/day_session_repo.go`:
- create
- find by `trip_id + date`
- get by id
- update

### 4) Handlers
Add routes and validation in `main.go`.

### 5) Test
```powershell
$trip = Invoke-RestMethod http://localhost:8080/trips -Method Post -ContentType 'application/json' -Body '{"destination":"Dubai","start_date":"2026-10-12","end_date":"2026-10-18","travelers_count":2}'
$ds = Invoke-RestMethod http://localhost:8080/day-sessions -Method Post -ContentType 'application/json' -Body ("{`"trip_id`":`"$($trip.id)`",`"date`":`"2026-10-12`",`"start_time`":`"09:00`",`"start_label`":`"Hotel`"}")
Invoke-RestMethod "http://localhost:8080/day-sessions/$($ds.id)" -Method Get | ConvertTo-Json -Depth 8
```

## Acceptance checks
- created day session is retrievable by id and by trip/date

## Deliverables
- day context entity completed.
