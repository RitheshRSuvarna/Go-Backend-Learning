. (Join-Path $PSScriptRoot "_state.ps1")

$state = Read-AskHubState
$api = $state.apiBase

if (-not $state.ContainsKey("destination")) { $state.destination = "Dubai" }
if (-not $state.ContainsKey("startDate")) { $state.startDate = "2026-10-12" }
if (-not $state.ContainsKey("endDate")) { $state.endDate = "2026-10-18" }
if (-not $state.ContainsKey("travelersCount")) { $state.travelersCount = 2 }

$body = New-AskHubJsonBody @{
  destination = [string]$state.destination
  start_date = [string]$state.startDate
  end_date = [string]$state.endDate
  travelers_count = [int]$state.travelersCount
}

$resp = Invoke-RestMethod "$api/trips" -Method Post -ContentType "application/json" -Body $body

$state.tripId = $resp.id
Write-AskHubState $state

Write-Host "OK: tripId = $($state.tripId)"
Write-Host "State saved to $(Get-AskHubStatePath)"

