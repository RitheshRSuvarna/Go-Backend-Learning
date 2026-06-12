## Planning flow (PowerShell)

### Option A) Fast demo seed (dev-only)

Requires: `APP_ENV=dev` or `DEMO_RESET_ENABLED=true`

```powershell
Invoke-RestMethod "http://localhost:8080/demo/reset" -Method Post -ContentType "application/json" -Body "{}" | ConvertTo-Json -Depth 20
```

### Option B) Manual flow

```powershell
$base = "http://localhost:8080/api"

# 1) Create trip
$trip = Invoke-RestMethod "$base/trips" `
  -Method Post `
  -ContentType "application/json" `
  -Body (@{
    destination     = "Paris"
    start_date      = "2026-05-11"
    end_date        = "2026-05-13"
    travelers_count = 2
  } | ConvertTo-Json)

# 2) Create day session
$ds = Invoke-RestMethod "$base/day-sessions" `
  -Method Post `
  -ContentType "application/json" `
  -Body (@{
    trip_id     = $trip.id
    date        = "2026-05-11"
    start_time  = "09:00"
    start_label = "Hotel"
  } | ConvertTo-Json)

# 3) Create plan version
$plan = Invoke-RestMethod "$base/day-sessions/$($ds.id)/plan-versions" `
  -Method Post `
  -ContentType "application/json" `
  -Body (@{
    notes = "v1"
    stops = @(
      @{
        position = 1; title = "Louvre Museum"; category_label = "Sightseeing"
        planned_arrival = "10:00"; planned_departure = "12:00"
        travel_minutes = 30; stay_minutes = 120
      },
      @{
        position = 2; title = "Lunch"; category_label = "Dining"
        planned_arrival = "12:30"; planned_departure = "13:30"
        travel_minutes = 15; stay_minutes = 60
      }
    )
  } | ConvertTo-Json -Depth 20)

# 4) Get active plan
Invoke-RestMethod "$base/day-sessions/$($ds.id)/active-plan" -Method Get | ConvertTo-Json -Depth 20
```

## Planning flow example

```powershell
. .\Rithesh-backend-docs\scripts\01_auth_signup_or_login.ps1
. .\Rithesh-backend-docs\scripts\02_create_trip.ps1
. .\Rithesh-backend-docs\scripts\03_create_day_session.ps1
. .\Rithesh-backend-docs\scripts\04_create_plan_v1.ps1
```

Then inspect active plan:

```powershell
$state = Get-Content .\Rithesh-backend-docs\scripts\.state.json -Raw | ConvertFrom-Json
Invoke-RestMethod "http://localhost:8080/day-sessions/$($state.daySessionId)/active-plan" -Method Get | ConvertTo-Json -Depth 20
```

Expected:
- day session exists
- plan v1 exists
- active plan has stops
