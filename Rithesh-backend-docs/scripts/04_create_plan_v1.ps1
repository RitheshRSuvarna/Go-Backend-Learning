. (Join-Path $PSScriptRoot "_state.ps1")

$state = Read-AskHubState
$api = $state.apiBase
Require-StateValue $state "daySessionId"

$payload = @{
  notes = "v1"
  stops = @(
    @{
      position = 1
      title = "Dubai Mall"
      category_label = "Shopping"
      planned_arrival = "09:30"
      planned_departure = "11:00"
      travel_minutes = 30
      stay_minutes = 90
    },
    @{
      position = 2
      title = "Burj Khalifa"
      category_label = "Sightseeing"
      planned_arrival = "11:15"
      planned_departure = "12:45"
      travel_minutes = 15
      stay_minutes = 90
    }
  )
}

$body = New-AskHubJsonBody $payload
$resp = Invoke-RestMethod "$api/day-sessions/$($state.daySessionId)/plan-versions" -Method Post -ContentType "application/json" -Body $body

$state.planV1Id = $resp.id
Write-AskHubState $state

Write-Host "OK: planV1Id = $($state.planV1Id)"
Write-Host "State saved to $(Get-AskHubStatePath)"

