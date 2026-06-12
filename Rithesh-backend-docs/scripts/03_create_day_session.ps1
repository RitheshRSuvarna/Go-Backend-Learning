. (Join-Path $PSScriptRoot "_state.ps1")

$state = Read-AskHubState
$api = $state.apiBase
Require-StateValue $state "tripId"

if (-not $state.ContainsKey("dayDate")) { $state.dayDate = [string]$state.startDate }
if (-not $state.ContainsKey("startTime")) { $state.startTime = "09:00" }
if (-not $state.ContainsKey("startLabel")) { $state.startLabel = "Hotel" }

$body = New-AskHubJsonBody @{
  trip_id = [string]$state.tripId
  date = [string]$state.dayDate
  start_time = [string]$state.startTime
  start_label = [string]$state.startLabel
}

$resp = Invoke-RestMethod "$api/day-sessions" -Method Post -ContentType "application/json" -Body $body

$state.daySessionId = $resp.id
Write-AskHubState $state

Write-Host "OK: daySessionId = $($state.daySessionId)"
Write-Host "State saved to $(Get-AskHubStatePath)"

