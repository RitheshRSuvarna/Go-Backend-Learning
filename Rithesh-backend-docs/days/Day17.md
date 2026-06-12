# Day 17 - Tests: BusyScore + Drift + Auth basics

## Goal (end of day)
- Unit tests exist for core deterministic logic

## Learn (30-60 minutes)
- Go tests: `https://pkg.go.dev/testing`
- Subtests and table tests

## Tasks
### 1) BusyScore tests
In `internal/busy` add table tests for labels:
- low
- medium
- high

Cover at least:
- peak hour
- off-peak
- unknown category fallback

### 2) Drift tests
In `internal/drift` cover:
- on-time case
- moderate delay
- heavy delay triggering replan recommendation

### 3) Auth helper tests
In `internal/auth` test:
- `HashPassword` produces non-plain output
- `CheckPassword` true for correct password
- `CheckPassword` false for wrong password

### 4) Run tests
```powershell
cd services/api
go test ./...
```

## Acceptance checks
- all tests pass
- test names make behavior obvious

## Deliverables
- baseline unit test suite for core rules.
