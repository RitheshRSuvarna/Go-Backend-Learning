. (Join-Path $PSScriptRoot "_state.ps1")

$state = Read-AskHubState
$api = $state.apiBase
Require-StateValue $state "daySessionId"

$resp = Invoke-RestMethod "$api/day-sessions/$($state.daySessionId)/replan" -Method Post -ContentType "application/json" -Body "{}"

$state.suggestedPlanId = $resp.suggested_plan_version_id
Write-AskHubState $state

Write-Host "OK: suggestedPlanId = $($state.suggestedPlanId)"
Write-Host "Reason codes: $((($resp.reason_codes) -join ', '))"
Write-Host "State saved to $(Get-AskHubStatePath)"

