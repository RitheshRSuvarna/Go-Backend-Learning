## Reusable PowerShell scripts

These scripts remove guesswork for IDs such as:
- `tripId`
- `daySessionId`
- `planVersionId`

They store state in:
- `Rithesh-backend-docs/scripts/.state.json`

## Workspace assumption

Run commands from your workspace root (the folder containing `Rithesh-backend-docs`).

Expected layout by the time you use these scripts:

```text
backend-learning/
  Rithesh-backend-docs/
  services/api/
```

These scripts are mainly useful after the API endpoints from early days exist.

## Prerequisites

From workspace root:

```powershell
docker compose up -d
```

In another terminal:

```powershell
cd services/api
go run .
```

Optional preflight:

```powershell
. .\Rithesh-backend-docs\scripts\00_preflight.ps1
```

## Recommended order

1. Auth (signup or login)

```powershell
. .\Rithesh-backend-docs\scripts\01_auth_signup_or_login.ps1
```

2. Create trip

```powershell
. .\Rithesh-backend-docs\scripts\02_create_trip.ps1
```

3. Create day session

```powershell
. .\Rithesh-backend-docs\scripts\03_create_day_session.ps1
```

4. Create plan v1

```powershell
. .\Rithesh-backend-docs\scripts\04_create_plan_v1.ps1
```

5. Replan to create suggested plan

```powershell
. .\Rithesh-backend-docs\scripts\05_replan.ps1
```

6. Diff suggested vs base

```powershell
. .\Rithesh-backend-docs\scripts\06_diff_plans.ps1
```

7. Accept suggested plan as active

```powershell
. .\Rithesh-backend-docs\scripts\07_accept_plan.ps1
```

## Reset state (optional)

If you want a clean run:

```powershell
. .\Rithesh-backend-docs\scripts\99_reset_state.ps1
```

Then rerun scripts from step 1.
