# Day 16 - Reliability pass (timeouts + consistent errors)

## Goal (end of day)
- Outbound calls have timeouts
- Error responses are consistent across endpoints

## Learn (30-60 minutes)
- Go contexts and cancellation: `https://pkg.go.dev/context`
- Timeout patterns in Go HTTP clients

## Tasks
### 1) Standardize error response shape
Use one shape everywhere:
```json
{
  "error": {
    "code": "bad_request",
    "message": "..."
  }
}
```

Allowed codes:
- `bad_request`
- `unauthorized`
- `not_found`
- `internal_error`

### 2) Add timeout wrappers
- DB heavy calls: context timeout (for example 2-3s)
- LLM calls: context timeout (for example 3-5s)

Do not leave long-running requests unbounded.

### 3) Ensure timeout errors map cleanly
When context deadline happens, return deterministic JSON error.

### 4) Validate with intentional failure
Temporarily stop LLM and call an LLM endpoint.

## Acceptance checks
- Intentional failure returns clean JSON error
- API process does not crash
- error shape is identical across multiple endpoints

## Deliverables
- Reliability baseline in place for all core routes.
