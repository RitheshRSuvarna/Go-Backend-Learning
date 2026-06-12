# Day 15 - Contract tests (Go to LLM)

## Goal (end of day)
- Contract tests prove:
  - Go sends expected JSON
  - Go validates LLM response shape

## Learn (30-75 minutes)
- Go testing package: `https://pkg.go.dev/testing`
- `httptest` package: `https://pkg.go.dev/net/http/httptest`
- Table-driven tests: `https://go.dev/wiki/TableDrivenTests`

## Prerequisites
- Day14 complete

## Tasks
### 1) Create test file for LLM client/handler flow
Use `httptest.NewServer` to simulate LLM service.

Test cases:
- valid response (happy path)
- missing `stops` field
- stop with missing required field
- invalid time format

### 2) Verify Go request payload
In mock server handler, assert request fields:
- required ids and times are present
- JSON field names are exact

### 3) Add handler-level bad request test
If API receives invalid input (missing day session id, invalid request body), expect `400 bad_request`.

### 4) Run tests
```powershell
cd services/api
go test ./...
```

## Acceptance checks
- `go test ./...` passes
- at least one invalid contract test fails before fix and passes after fix

## Deliverables
- Contract tests committed with clear test names.
