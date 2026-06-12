. (Join-Path $PSScriptRoot "_state.ps1")

$state = Read-AskHubState
$api = $state.apiBase

# Change these if you want.
if (-not $state.ContainsKey("email")) { $state.email = "rithesh.r.suvarna@example.com" }
if (-not $state.ContainsKey("password")) { $state.password = "Password123!" }
if (-not $state.ContainsKey("fullName")) { $state.fullName = "Rithesh R Suvarna" }

$email = [string]$state.email
$password = [string]$state.password
$fullName = [string]$state.fullName

$session = $null

try {
  $body = New-AskHubJsonBody @{ email = $email; password = $password; full_name = $fullName }
  Invoke-WebRequest "$api/auth/signup" -Method Post -ContentType "application/json" -Body $body -SessionVariable session | Out-Null
} catch {
  # If signup fails (likely because user already exists), try login.
  $body = New-AskHubJsonBody @{ email = $email; password = $password }
  Invoke-WebRequest "$api/auth/login" -Method Post -ContentType "application/json" -Body $body -SessionVariable session | Out-Null
}

$state.hasSession = $true
Write-AskHubState $state

Write-Host "OK: authenticated as $email"
Write-Host "State saved to $(Get-AskHubStatePath)"
