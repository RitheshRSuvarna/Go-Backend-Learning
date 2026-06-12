# Day 00 - Create backend project skeleton (from zero)

## Goal (end of day)
You can run a tiny Go API locally that returns JSON from:
- `GET /health`

No database yet. Today is about folder structure, `go.mod`, `main.go`, and running a server.

## Before you start
Complete setup first:
- `Rithesh-backend-docs/DAY_MINUS_1_SETUP.md`

## Learn (30-90 minutes)
- Go Tour (basics): `https://tour.golang.org/`
- `net/http` handlers: `https://pkg.go.dev/net/http`
- Project layout inspiration: `https://github.com/golang-standards/project-layout`

## Tasks (PowerShell)
### 1) Create folder structure
From workspace root (same level as `Rithesh-backend-docs`), create:
- `services/api/`
- `services/api/internal/`
- `services/api/internal/model/`
- `services/api/internal/repo/`
- `services/api/internal/db/`
- `services/api/internal/migrate/`
- `services/api/internal/migrate/migrations/`

### 2) Initialize Go module
```powershell
cd services/api
go mod init askhubtravel/api
```

### 3) Create `main.go` with one endpoint
Requirements:
- listen on `8080`
- `GET /health` returns:
  - `rue, "service": "a{ "ok": tpi", "timestamp": "..." }`

### 4) Run server
```powershell
cd services/api
go run .
```

### 5) Test endpoint
```powershell
Invoke-RestMethod http://localhost:8080/health -Method Get | ConvertTo-Json -Depth 5
```

## Acceptance checks
- server starts without error
- `/health` returns valid JSON

## Deliverables
- working `main.go`
- successful `/health` response in terminal.
