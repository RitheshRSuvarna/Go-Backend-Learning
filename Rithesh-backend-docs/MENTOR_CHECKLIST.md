# Mentor Checklist

Use this to review progress each week.

## Code quality

- naming is clear and consistent
- handlers validate input and return consistent errors
- repo layer contains DB logic, not HTTP logic
- no copied code blocks without understanding

## API behavior

- status codes are correct (`400/401/404/500`)
- error JSON shape is stable
- auth session behavior is correct
- migration flow works from clean database

## Testing

- unit tests added for core pure logic
- contract tests exist for LLM integration
- `go test ./...` passes on clean machine

## Operational readiness

- runbook exists and is usable
- logs include request id and useful context
- common failure modes documented

## Learning signals

Rithesh R Suvarna can explain:
- why password hashes are stored (not plain text)
- why token hash is stored instead of raw token
- when to use DB migration vs code change
- difference between deterministic backend rules and LLM behavior
