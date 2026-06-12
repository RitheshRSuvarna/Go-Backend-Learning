# Day 10 - Swagger (generate OpenAPI + serve UI)

## Goal (end of day)
- Swagger UI loads at `/swagger/index.html`
- docs reflect implemented endpoints

## Learn (20-45 minutes)
- OpenAPI basics: `https://swagger.io/specification/`
- Swaggo: `https://github.com/swaggo/swag`

## Tasks
### 1) Install `swag`
```powershell
cd services/api
go install github.com/swaggo/swag/cmd/swag@latest
```

### 2) Add swagger dependencies
```powershell
cd services/api
go get github.com/swaggo/http-swagger@latest
go get github.com/swaggo/swag@latest
go mod tidy
```

### 3) Add top-level annotations
In `main.go` add API metadata above `main()`:
- title
- version
- description
- basePath

### 4) Generate docs
```powershell
cd services/api
swag init -g main.go -o ./internal/swaggerdocs
```

### 5) Serve swagger route
Add:
- `GET /swagger/*`

### 6) Verify
Open:
- `http://localhost:8080/swagger/index.html`

## Acceptance checks
- Swagger lists at least `/trips`, `/day-sessions`, `/auth/*`
- changing one annotation and regenerating docs updates UI

## Deliverables
- working Swagger UI connected to current API.
