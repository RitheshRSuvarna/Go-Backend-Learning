# Day 12 - Forgot/reset password (dev-token flow)

## Goal (end of day)
- forgot-password request works safely
- reset-password updates stored hash

## Learn (20-60 minutes)
- OWASP account enumeration guidance: `https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html`

## Notes
`password_resets` table already exists from Day11 migration.

## Tasks
### 1) Add repo methods
In `auth_repo.go` add:
- `CreatePasswordReset(user_id, token_hash, expires_at)`
- `UsePasswordReset(token_hash) -> user_id` (marks used)
- `UpdateUserPassword(user_id, new_hash)`

### 2) Add endpoints
- `POST /auth/forgot`
  - always returns `{ok:true}`
  - if email exists, create reset token
  - in dev mode return `dev_reset_token`
- `POST /auth/reset`
  - validates token exists, not expired, not used
  - updates password hash

### 3) Manual flow test
```powershell
$sess = New-Object Microsoft.PowerShell.Commands.WebRequestSession
$email = "reset$(Get-Random)@example.com"
Invoke-RestMethod http://localhost:8080/auth/signup -Method Post -ContentType 'application/json' -Body ("{`"email`":`"$email`",`"full_name`":`"User`",`"password`":`"password123`"}") -WebSession $sess | Out-Null
$resp = Invoke-RestMethod http://localhost:8080/auth/forgot -Method Post -ContentType 'application/json' -Body ("{`"email`":`"$email`"}")
$token = $resp.dev_reset_token
Invoke-RestMethod http://localhost:8080/auth/reset -Method Post -ContentType 'application/json' -Body ("{`"token`":`"$token`",`"password`":`"newpassword123`"}") | ConvertTo-Json -Depth 3
$sess2 = New-Object Microsoft.PowerShell.Commands.WebRequestSession
Invoke-RestMethod http://localhost:8080/auth/login -Method Post -ContentType 'application/json' -Body ("{`"email`":`"$email`",`"password`":`"newpassword123`"}") -WebSession $sess2 | ConvertTo-Json -Depth 5
Invoke-RestMethod http://localhost:8080/auth/me -Method Get -WebSession $sess2 | ConvertTo-Json -Depth 5
```

## Acceptance checks
- forgot endpoint returns `ok` even for unknown email
- reset token can be used once
- login works with new password

## Deliverables
- full dev reset-password loop.
