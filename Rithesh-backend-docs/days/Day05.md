# Day 05 - Busy risk heuristic (compute on read)

## Goal (end of day)
- BusyScore function exists
- `GET /day-sessions/{id}/active-plan` returns `busy_risk_label` per stop

## Learn (30-60 minutes)
- Go time basics: `https://pkg.go.dev/time`
- String handling: `https://pkg.go.dev/strings`

## Tasks
### 1) Create BusyScore module
Create `internal/busy/busy.go`:
- `Label(categoryLabel string, plannedArrivalHHMM string) string`

Return labels: `low|med|high`.

Rules v1:
- Shopping: high from `12:00-17:00`
- Sightseeing: med from `11:00-15:00`
- Dining: high from `12:00-14:00` and `19:00-21:00`
- default: low

### 2) Attach labels on read
In active-plan handler, compute label for each stop before responding.

### 3) Test
```powershell
$state = Get-Content .\Rithesh-backend-docs\scripts\.state.json -Raw | ConvertFrom-Json
Invoke-RestMethod "http://localhost:8080/day-sessions/$($state.daySessionId)/active-plan" -Method Get | ConvertTo-Json -Depth 10
```

## Acceptance checks
- at least one stop returns expected `busy_risk_label` for known input time

## Deliverables
- deterministic busy-risk labels in API output.
