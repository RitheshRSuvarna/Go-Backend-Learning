# Debug Playbook

Use this when something fails.

## 1) API does not start

Check:

```powershell
cd services/api
go run .
```

Common causes:
- missing env var
- DB is not running
- migration syntax error

Fix path:
1. Read first error line, not only last line.
2. If DB error, run `docker compose ps`.
3. If migration error, inspect latest `.sql` file.

## 2) Database connection refused

Check:

```powershell
docker compose ps
docker compose logs postgres --tail 50
```

Then verify `DATABASE_URL` in `services/api/.env`.

## 3) Migration fails on startup

Checklist:
- migration file names are ordered correctly
- each migration is idempotent when needed (`IF NOT EXISTS`)
- SQL syntax matches Postgres

If needed, reset local DB only:

```powershell
docker compose down -v
docker compose up -d
```

## 4) `401 unauthorized` unexpectedly

Likely causes:
- cookie not stored in PowerShell session
- login/signup failed earlier

Use `-WebSession` consistently:

```powershell
$sess = New-Object Microsoft.PowerShell.Commands.WebRequestSession
Invoke-RestMethod http://localhost:8080/auth/login -Method Post -ContentType "application/json" -Body '{"email":"x@example.com","password":"Password123!"}' -WebSession $sess
Invoke-RestMethod http://localhost:8080/auth/me -Method Get -WebSession $sess
```

## 5) `400 bad_request` from JSON endpoint

Check:
- field names match API exactly (snake_case)
- required fields are present
- time/date formats are correct

Tip:
- start from a known working payload, then edit one field at a time.

## 6) Replan/diff not working

Run scripts in order:

```powershell
. .\Rithesh-backend-docs\scripts\01_auth_signup_or_login.ps1
. .\Rithesh-backend-docs\scripts\02_create_trip.ps1
. .\Rithesh-backend-docs\scripts\03_create_day_session.ps1
. .\Rithesh-backend-docs\scripts\04_create_plan_v1.ps1
. .\Rithesh-backend-docs\scripts\05_replan.ps1
. .\Rithesh-backend-docs\scripts\06_diff_plans.ps1
```

## 7) How to ask for help effectively

Always include:
- command you ran
- full error text
- which day/task you are on
- one screenshot or pasted JSON if relevant

Bad help request:
- "It does not work."

Good help request:
- "Day14 Task2: `POST /day-sessions/{id}/llm/plan` returns 500. Error: `invalid llm response: missing stops`."
