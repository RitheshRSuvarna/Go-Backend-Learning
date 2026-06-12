# Day 11 - Auth: signup/login/logout/me (cookie sessions)

## Goal (end of day)
- account creation and login work
- session cookie auth works for `/auth/me`
- logout clears active session

## Learn (30-75 minutes)
- OWASP password storage: `https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html`
- OWASP auth basics: `https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html`
- Cookie basics: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies`

## Tasks
### 1) Migration: users + sessions + password_resets
Create `services/api/internal/migrate/migrations/0009_create_auth_tables.sql`.

Tables:
- `users`
- `sessions`
- `password_resets`

### 2) Password hashing helper
Add dependency:
```powershell
cd services/api
go get golang.org/x/crypto/bcrypt@latest
go mod tidy
```

Create `services/api/internal/auth/password.go`:
- `HashPassword(password) -> string`
- `CheckPassword(hash, password) -> bool`

### 3) Auth repository methods
Add in `auth_repo.go`:
- create user
- get user by email
- create session (store token hash)
- resolve session to user
- delete session

### 4) HTTP handlers + cookies
Implement:
- `POST /auth/signup`
- `POST /auth/login`
- `POST /auth/logout`
- `GET /auth/me`

Session approach:
- generate random token
- store only `SHA256(token)` in DB
- cookie name `askhub_session`, `HttpOnly`, `SameSite=Lax`

### 5) Test with session cookie
```powershell
$email = "u$(Get-Random)@example.com"
$sess = New-Object Microsoft.PowerShell.Commands.WebRequestSession
Invoke-RestMethod http://localhost:8080/auth/signup -Method Post -ContentType 'application/json' -Body ("{`"email`":`"$email`",`"full_name`":`"User`",`"password`":`"password123`"}") -WebSession $sess | ConvertTo-Json -Depth 5
Invoke-RestMethod http://localhost:8080/auth/me -Method Get -WebSession $sess | ConvertTo-Json -Depth 5
Invoke-RestMethod http://localhost:8080/auth/logout -Method Post -ContentType 'application/json' -Body '{}' -WebSession $sess | ConvertTo-Json -Depth 3
```

## Acceptance checks
- `/auth/me` returns user before logout
- `/auth/me` returns 401 after logout

## Deliverables
- complete cookie-session auth baseline.
