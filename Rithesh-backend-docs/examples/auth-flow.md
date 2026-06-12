## Auth flow (PowerShell)

### 1) Signup

```powershell
$base = "http://localhost:8080/api"

Invoke-RestMethod "$base/auth/signup" `
  -Method Post `
  -ContentType "application/json" `
  -Body (@{
    email     = "demo@example.com"
    full_name = "Demo User"
    password  = "password123"
  } | ConvertTo-Json)
```

### 2) Login (saves cookie session)

```powershell
$base = "http://localhost:8080/api"
$session = New-Object Microsoft.PowerShell.Commands.WebRequestSession

Invoke-RestMethod "$base/auth/login" `
  -Method Post `
  -ContentType "application/json" `
  -WebSession $session `
  -Body (@{
    email    = "demo@example.com"
    password = "password123"
  } | ConvertTo-Json)
```

### 3) Me

```powershell
$base = "http://localhost:8080/api"

Invoke-RestMethod "$base/auth/me" `
  -Method Get `
  -WebSession $session
```

### 4) Logout

```powershell
$base = "http://localhost:8080/api"

Invoke-RestMethod "$base/auth/logout" `
  -Method Post `
  -WebSession $session
```

## Auth flow example

```powershell
$email = "demo$(Get-Random)@example.com"
$sess = New-Object Microsoft.PowerShell.Commands.WebRequestSession

Invoke-RestMethod http://localhost:8080/auth/signup -Method Post -ContentType 'application/json' -Body ("{`"email`":`"$email`",`"full_name`":`"Demo User`",`"password`":`"Password123!`"}") -WebSession $sess | ConvertTo-Json -Depth 5
Invoke-RestMethod http://localhost:8080/auth/me -Method Get -WebSession $sess | ConvertTo-Json -Depth 5
Invoke-RestMethod http://localhost:8080/auth/logout -Method Post -ContentType 'application/json' -Body '{}' -WebSession $sess | ConvertTo-Json -Depth 5
```

Expected:
- signup/login returns user
- `/auth/me` works before logout
- `/auth/me` returns 401 after logout
