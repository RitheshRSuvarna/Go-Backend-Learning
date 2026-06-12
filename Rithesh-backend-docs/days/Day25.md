# Day 25 - Observability upgrade (debuggable errors)

## Goal (end of day)
- all error responses include request id
- logs are sufficient for quick root-cause checks

## Learn (15-30 minutes)
- Structured log design basics

## Tasks
### 1) Include request id in error JSON
Suggested shape:
```json
{
  "error": {
    "code": "bad_request",
    "message": "...",
    "request_id": "..."
  }
}
```

### 2) Ensure request id header always present
Even on `400/401/404/500` responses.

### 3) Improve log context
For errors, include:
- request id
- endpoint
- status
- short cause summary

### 4) Validate end-to-end
Make one intentional bad request and map response id to log line.

## Acceptance checks
- every error response has `request_id`
- every log line needed for debug is present

## Deliverables
- production-style debugging baseline.
