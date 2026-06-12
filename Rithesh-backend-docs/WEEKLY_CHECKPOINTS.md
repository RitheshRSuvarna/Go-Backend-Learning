# Weekly Checkpoints

Do not move to the next week unless current checkpoint passes.

## Week 1 (Day00-Day07)

Must pass:
- `GET /health`
- `POST/GET /trips`
- `POST/GET /day-sessions`
- `POST /day-sessions/{id}/plan-versions`
- `GET /day-sessions/{id}/active-plan`
- `POST/GET /day-sessions/{id}/events`

Validation:

```powershell
. .\Rithesh-backend-docs\scripts\01_auth_signup_or_login.ps1
. .\Rithesh-backend-docs\scripts\02_create_trip.ps1
. .\Rithesh-backend-docs\scripts\03_create_day_session.ps1
. .\Rithesh-backend-docs\scripts\04_create_plan_v1.ps1
```

## Week 2 (Day08-Day15)

Must pass:
- `POST /day-sessions/{id}/replan`
- `GET /plan-versions/{id}/diff?baseId=...`
- auth flow (`signup/login/me/logout`)
- forgot/reset password
- basic swagger docs
- LLM service local call from API

Validation:

```powershell
. .\Rithesh-backend-docs\scripts\05_replan.ps1
. .\Rithesh-backend-docs\scripts\06_diff_plans.ps1
go test ./...
```

## Week 3 (Day16-Day23)

Must pass:
- consistent error JSON
- request id in logs and response header
- replan constraints respected
- key logic moved out of `main.go` where planned

Validation:
- Run one failing request and confirm clean error format.
- Capture one request id and find matching log line.

## Week 4 (Day24-Day30)

Must pass:
- critical endpoints still work after perf/hardening
- replace-stop flow creates new plan version
- runbook and handover docs are complete
- full rehearsal passes 3 times

Validation:
- `go test ./...`
- end-to-end story from Day29
- demo script from Day30 runs without improvisation
