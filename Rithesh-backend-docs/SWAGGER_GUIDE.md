## Swagger / OpenAPI guide (Go API, Windows + PowerShell)

Use this after Day10 and keep docs updated as routes change.

## 1) Install `swag`

```powershell
cd services/api
go install github.com/swaggo/swag/cmd/swag@latest
```

Verify:

```powershell
swag --version
```

If command is missing, add Go bin to PATH:
- `%USERPROFILE%\go\bin`

## 2) Add dependencies (one-time)

```powershell
cd services/api
go get github.com/swaggo/http-swagger@latest
go get github.com/swaggo/swag@latest
go mod tidy
```

## 3) Add base annotations in `main.go`

Add API metadata above `func main()`:
- title
- version
- description
- host/basePath (if used)

Then annotate handlers with:
- summary
- tags
- params
- success/failure responses

## 4) Generate docs

```powershell
cd services/api
swag init -g main.go -o ./internal/swaggerdocs
```

## 5) Serve Swagger UI route

Add route in API:
- `GET /swagger/*`

Open:
- `http://localhost:8080/swagger/index.html`

## 6) Daily workflow

Whenever endpoint contract changes:
1. update annotations
2. run `swag init`
3. verify in UI

## 7) Common issues

### `swag` not found
- PATH missing Go bin path

### Empty or outdated docs
- regenerate after code changes
- ensure annotations are on exported types/handlers as needed

### UI route 404
- verify swagger route wiring in router

## 8) Quality checklist

- every public endpoint is documented
- request/response examples exist for major flows
- error schema is documented consistently
