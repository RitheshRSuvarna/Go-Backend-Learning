# Day 04 - Plan versions + stops (store itinerary in DB)

## Goal (end of day)
- tables exist:
  - `plan_versions`
  - `plan_stops`
- endpoints work:
  - `POST /day-sessions/{id}/plan-versions`
  - `GET /day-sessions/{id}/active-plan`

## Learn (45-90 minutes)
- SQL transactions: `https://www.postgresql.org/docs/current/tutorial-transactions.html`
- Data modeling basics: `https://www.postgresql.org/docs/current/ddl.html`

## Tasks
### 1) Migration
Create `0004_create_plan_versions_and_stops.sql`.

`plan_versions`:
- id, day_session_id, version, notes, created_at
- unique `(day_session_id, version)`

`plan_stops`:
- id, plan_version_id, position, title, category_label, image_url
- planned_arrival, planned_departure
- travel_minutes, stay_minutes, busy_risk_label, created_at
- unique `(plan_version_id, position)`

### 2) Models
Create `internal/model/plan.go`:
- `PlanVersion`
- `PlanStop`
- `CreatePlanVersionRequest`

### 3) Repo
Create `internal/repo/plan_repo.go`:
- create plan version transactionally
- get active/latest plan with ordered stops

### 4) Handlers
Add routes in `main.go`.

### 5) Test
```powershell
. .\Rithesh-backend-docs\scripts\01_auth_signup_or_login.ps1
. .\Rithesh-backend-docs\scripts\02_create_trip.ps1
. .\Rithesh-backend-docs\scripts\03_create_day_session.ps1
. .\Rithesh-backend-docs\scripts\04_create_plan_v1.ps1
```

## Acceptance checks
- active plan returns ordered stops from DB

## Deliverables
- itinerary persistence ready.
