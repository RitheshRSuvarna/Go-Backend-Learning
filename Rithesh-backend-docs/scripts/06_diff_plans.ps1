. (Join-Path $PSScriptRoot "_state.ps1")

$state = Read-AskHubState
$api = $state.apiBase
Require-StateValue $state "planV1Id"
Require-StateValue $state "suggestedPlanId"

$resp = Invoke-RestMethod "$api/plan-versions/$($state.suggestedPlanId)/diff?baseId=$($state.planV1Id)" -Method Get

$resp | ConvertTo-Json -Depth 20

