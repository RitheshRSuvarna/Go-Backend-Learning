## Replan flow (PowerShell)

This flow assumes you already have a `day_session_id`. The easiest way is to run the dev seed:

```powershell
$seed = Invoke-RestMethod "http://localhost:8080/demo/reset" -Method Post -ContentType "application/json" -Body "{}"
$daySessionId = $seed.day_session_id
$base = "http://localhost:8080/api"
```

### 1) Trigger deterministic replan (server-side)

```powershell
Invoke-RestMethod "$base/day-sessions/$daySessionId/replan" -Method Post -ContentType "application/json" -Body "{}" | ConvertTo-Json -Depth 20
```

### 2) Create a plan via LLM service (uses LLM microservice)

```powershell
Invoke-RestMethod "$base/day-sessions/$daySessionId/llm/plan" -Method Post -ContentType "application/json" -Body "{}" | ConvertTo-Json -Depth 20
```

### 3) Diff two plan versions

```powershell
# Use ids from /active-plan or from the response of /replan or /llm/plan
$suggested = "<suggested_plan_version_id>"
$basePlan = "<base_plan_version_id>"

Invoke-RestMethod "$base/plan-versions/$suggested/diff?baseId=$basePlan" -Method Get | ConvertTo-Json -Depth 20
```

## Replan flow example

```powershell
. .\Rithesh-backend-docs\scripts\05_replan.ps1
. .\Rithesh-backend-docs\scripts\06_diff_plans.ps1
. .\Rithesh-backend-docs\scripts\07_accept_plan.ps1
```

Expected:
- replan returns `suggested_plan_version_id`
- diff endpoint shows at least one change
- accepted plan becomes active
