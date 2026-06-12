# Day 23 - Replan guardrails (constraints input)

## Goal (end of day)
- replan supports constraints
- constraints are validated and respected

## Learn (20-40 minutes)
- Input validation patterns in Go handlers

## Tasks
### 1) Define request contract
Add JSON body for replan, for example:
- `must_keep_stop_ids: []`
- `avoid_stop_ids: []`

### 2) Add validation
Rules:
- same stop id cannot appear in both lists
- ids must exist in active plan
- max list size guard (avoid abuse)

### 3) Apply guardrails in heuristic
- must-keep stops cannot be removed
- avoid stops should be deprioritized when possible

### 4) Return explicit conflict errors
If constraints are impossible, return `400 bad_request` with clear message.

## Acceptance checks
- impossible constraints return clear 400 error
- valid constraints alter replan output in expected direction

## Deliverables
- replan endpoint is safer and user-guided.
