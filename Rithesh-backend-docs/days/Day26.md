# Day 26 - Hardening: edge cases + validation

## Goal (end of day)
- bad payloads are rejected consistently
- stop/time validation is explicit and testable

## Learn (20-40 minutes)
- HTTP 400 semantics
- Defensive input validation patterns

## Tasks
### 1) Validate stop schedule fields
Check:
- `planned_arrival` format `HH:MM`
- `planned_departure` format `HH:MM`
- arrival <= departure

### 2) Validate stop ordering and uniqueness
- position values are unique
- positions form expected sequence

### 3) Validate durations
- `travel_minutes >= 0`
- `stay_minutes >= 0`

### 4) Add targeted tests
At least one test per validation family.

### 5) Keep error messages clear
Client should understand what to fix from one response.

## Acceptance checks
- malformed payloads return `400 bad_request`
- valid payload still works for happy path

## Deliverables
- API input layer is significantly safer.
