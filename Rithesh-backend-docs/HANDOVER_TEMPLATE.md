# Handover Template

## Service map
- API: `services/api` (`:8080`)
- LLM: `services/llm` (`:8090`)
- Postgres: Docker compose

## Environment variables

### API
- `PORT`
- `DATABASE_URL`
- `LLM_BASE_URL`

### LLM
- list service-specific env vars (if any)

## Run commands
- DB start
- API start
- LLM start
- smoke checks

## Known issues
- issue 1:
- issue 2:

## Troubleshooting
- common failure + fix

## Out of scope (intentional)
- production auth hardening
- real email delivery
- real LLM evaluation framework

## Suggested next engineering tasks
- short task list with priority
