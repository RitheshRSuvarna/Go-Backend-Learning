# Day 14 - Go API to LLM integration (HTTP client + contracts)

## Goal (end of day)
- Go API calls the Python LLM service over HTTP
- Response is validated before storing a new plan version

## Learn (45-90 minutes)
- Go HTTP client: `https://pkg.go.dev/net/http`
- JSON encode/decode: `https://pkg.go.dev/encoding/json`
- Go context timeouts (preview for Day16): `https://pkg.go.dev/context`

## Prerequisites
- Day13 complete and LLM service running on `http://localhost:8090`

## Tasks
### 1) Add env var
In `services/api/.env` add:
- `LLM_BASE_URL=http://localhost:8090`

### 2) Create LLM client package
Create `services/api/internal/llmclient/` with:
- request/response structs
- `Plan(...)`
- `Replan(...)`

Validation rules in Go:
- `stops` length >= 1
- each stop has non-empty `title`, `category_label`, `planned_arrival`, `planned_departure`
- time strings follow `HH:MM` basic check

### 3) Add API endpoint
Add endpoint:
- `POST /day-sessions/{id}/llm/plan`

Behavior:
1. load day session
2. call LLM `/plan`
3. validate response
4. store as new `plan_version`
5. return created plan version id

### 4) Add fail-safe error mapping
If LLM response is invalid, return:
- `400 bad_request` if request payload is bad
- `502` or `500` with clear error if LLM response is invalid/unavailable

### 5) Test manually
```powershell
# Ensure LLM is running first
Invoke-RestMethod "http://localhost:8080/day-sessions/<daySessionId>/llm/plan" -Method Post -ContentType 'application/json' -Body '{}' | ConvertTo-Json -Depth 10
```

## Acceptance checks
- Endpoint creates a plan version from LLM response
- Invalid LLM response is rejected with clear JSON error
- API does not panic if LLM is down

## Deliverables
- LLM integration works with strict contract validation.
