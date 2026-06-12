. (Join-Path $PSScriptRoot "_state.ps1")

$state = Read-AskHubState
$api = $state.apiBase
Require-StateValue $state "daySessionId"
Require-StateValue $state "suggestedPlanId"

Invoke-RestMethod "$api/day-sessions/$($state.daySessionId)/active-plan/$($state.suggestedPlanId)" -Method Put | Out-Null

$active = Invoke-RestMethod "$api/day-sessions/$($state.daySessionId)/active-plan" -Method Get
$active | ConvertTo-Json -Depth 20

