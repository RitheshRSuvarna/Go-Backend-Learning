# Day 19 - Demo seed endpoint (dev-only)

## Goal (end of day)
- One endpoint resets and seeds demo data quickly

## Learn (20-45 minutes)
- SQL transaction basics: `https://www.postgresql.org/docs/current/tutorial-transactions.html`

## Tasks
### 1) Add endpoint
Add:
- `POST /demo/reset`

Behavior:
- clear relevant demo data (safe for local only)
- create trip
- create day session
- create plan v1
- create suggested plan v2
- create events + suggestion
- return key IDs in response

### 2) Guard endpoint for development only
Require explicit flag/env check (example `APP_ENV=dev`).

### 3) Use transaction for seed flow
If one step fails, rollback to keep state clean.

### 4) Test
```powershell
Invoke-RestMethod http://localhost:8080/demo/reset -Method Post -ContentType 'application/json' -Body '{}' | ConvertTo-Json -Depth 20
```

## Acceptance checks
- endpoint returns ids for trip, day session, plan versions
- repeated calls produce clean usable state
- endpoint is blocked when app is not in dev mode

## Deliverables
- one-command demo data reset.
