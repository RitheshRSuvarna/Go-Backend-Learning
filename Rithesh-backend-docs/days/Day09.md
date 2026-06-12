# Day 09 - Plan diff endpoint (server-side diff)

## Goal (end of day)
- compare plan versions via `GET /plan-versions/{id}/diff?baseId=...`
- explain added/removed/moved changes clearly

## Learn (20-45 minutes)
- sequence comparison fundamentals

## Tasks
### 1) Implement diff endpoint
Add:
- `GET /plan-versions/{suggestedId}/diff?baseId={baseId}`

Behavior:
- load both plans and stops
- compare sequence by stop title (v1)
- return changes: `added|removed|moved|unchanged`

Response shape:
- `{ base: <planVersion>, suggested: <planVersion>, diff: [...] }`

### 2) Create test data
Use Day08 flow to produce v1 + v2.

### 3) Run diff
```powershell
. .\Rithesh-backend-docs\scripts\05_replan.ps1
. .\Rithesh-backend-docs\scripts\06_diff_plans.ps1
```

## Acceptance checks
- diff output contains at least one meaningful change entry

## Deliverables
- server-side comparison for UI explainability.
