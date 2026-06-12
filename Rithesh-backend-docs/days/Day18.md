# Day 18 - Logging + request IDs

## Goal (end of day)
- Every request has request id
- Logs include method, path, status, duration, request_id

## Learn (20-45 minutes)
- HTTP headers overview: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers`
- Structured logging basics

## Tasks
### 1) Add request-id middleware
Behavior:
- if `X-Request-Id` exists, keep it
- else generate one
- always include it in response header

### 2) Add request logging middleware
Log one line per request with fields:
- `request_id`
- `method`
- `path`
- `status`
- `duration_ms`

### 3) Verify error path includes request id
Make one bad request and confirm:
- response header has `X-Request-Id`
- log line has same id

## Acceptance checks
- successful request and failing request both include request id
- you can correlate client response to a specific server log line

## Deliverables
- basic observability layer for debugging.
