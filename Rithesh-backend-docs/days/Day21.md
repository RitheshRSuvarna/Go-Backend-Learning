# Day 21 - Refactor for maintainability (handlers/services/repos)

## Goal (end of day)
- handlers are thinner
- business logic moved to `internal/service/`
- no behavior regression

## Learn (20-45 minutes)
- Go package organization: `https://go.dev/doc/modules/layout`
- Refactoring without behavior change

## Tasks
### 1) Choose one business flow
Start with replan flow (recommended):
- move heuristic logic out of `main.go`
- keep HTTP parsing/validation in handler

### 2) Create service package
Create `services/api/internal/service/` with clear methods (for example `ReplanService`).

### 3) Inject dependencies
Service should depend on repo interfaces or concrete repos (simple is fine for now).

### 4) Keep behavior identical
Before and after refactor, run same commands and compare output shape.

### 5) Regression check
```powershell
cd services/api
go test ./...
```
Run script chain:
```powershell
. .\Rithesh-backend-docs\scripts\01_auth_signup_or_login.ps1
. .\Rithesh-backend-docs\scripts\02_create_trip.ps1
. .\Rithesh-backend-docs\scripts\03_create_day_session.ps1
. .\Rithesh-backend-docs\scripts\04_create_plan_v1.ps1
. .\Rithesh-backend-docs\scripts\05_replan.ps1
```

## Acceptance checks
- API responses unchanged for replan endpoint
- code in `main.go` is shorter and easier to scan

## Deliverables
- first maintainability refactor completed safely.
