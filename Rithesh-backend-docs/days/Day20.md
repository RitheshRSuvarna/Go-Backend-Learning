# Day 20 - Swagger completeness + API examples

## Goal (end of day)
- Swagger reflects all implemented endpoints
- Example request docs exist for critical flows

## Learn (20-45 minutes)
- OpenAPI spec: `https://swagger.io/specification/`
- Swaggo docs: `https://github.com/swaggo/swag`

## Tasks
### 1) Complete swagger annotations
Ensure each public route has:
- summary
- tags
- params
- success/failure responses

### 2) Regenerate docs
```powershell
cd services/api
swag init -g main.go -o ./internal/swaggerdocs
```

### 3) Serve UI and verify
```powershell
# with API running
Start-Process "http://localhost:8080/swagger/index.html"
```

### 4) Add examples folder
Create `Rithesh-backend-docs/examples/` with markdown files:
- `auth-flow.md`
- `planning-flow.md`
- `replan-flow.md`

Each file should include runnable PowerShell snippets.

## Acceptance checks
- Swagger UI loads with all key endpoints
- example docs are runnable without editing more than IDs

## Deliverables
- API is self-discoverable by another developer.
