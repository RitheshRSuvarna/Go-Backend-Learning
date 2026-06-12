# Day 08 - Replan endpoint (non-LLM) + plan acceptance

## Goal (end of day)
- `POST /day-sessions/{id}/replan` creates suggested plan version
- active plan can be switched explicitly

## Learn (30-60 minutes)
- Postgres foreign keys: `https://www.postgresql.org/docs/current/ddl-constraints.html#DDL-CONSTRAINTS-FK`

## Tasks
### 1) Add active plan selection support
Migration `0006_add_day_session_active_plan.sql`:
- add `active_plan_version_id uuid null references plan_versions(id)`

Update plan repo:
- `GetActivePlan(daySessionId)` returns referenced active plan when set
- fallback to latest version
- creating new version sets active by default

Add endpoint:
- `PUT /day-sessions/{id}/active-plan/{planVersionId}`

### 2) Implement replan endpoint
Add:
- `POST /day-sessions/{id}/replan`

v1 deterministic heuristic:
- load active plan
- if at least 2 stops and next stop is high busy risk, swap first two stops
- create new plan version
- return `{ suggested_plan_version_id, reason_codes, created_at }`

### 3) Add reason-codes migration
Add `0008_add_plan_version_reasons.sql` with `reason_codes_json jsonb`.

### 4) Test
```powershell
. .\Rithesh-backend-docs\scripts\04_create_plan_v1.ps1
. .\Rithesh-backend-docs\scripts\05_replan.ps1
. .\Rithesh-backend-docs\scripts\07_accept_plan.ps1
```

## Acceptance checks
- suggested plan version is created
- accepting plan changes active plan

## Deliverables
- deterministic replan cycle complete.
