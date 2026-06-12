# Day 06 - Assistant suggestions (non-LLM rule engine)

## Goal (end of day)
- suggestion data is stored in DB
- endpoints work:
  - `GET /day-sessions/{id}/suggestion`
  - `POST /day-sessions/{id}/suggestion`
  - `PUT /assistant-suggestions/{id}`

## Learn (30-60 minutes)
- REST patterns: `https://restfulapi.net/`
- Postgres enum tradeoffs: `https://www.postgresql.org/docs/current/datatype-enum.html`

## Tasks
### 1) Migration
Create `0005_create_assistant_suggestions.sql`:
- fields: id, day_session_id, message, status, created_at
- status values: `pending|accepted|snoozed`

### 2) Model + repo
Create:
- `internal/model/assistant_suggestion.go`
- `internal/repo/assistant_suggestion_repo.go`

### 3) Add handlers
- `GET /day-sessions/{id}/suggestion`
- `POST /day-sessions/{id}/suggestion`
- `PUT /assistant-suggestions/{id}`

Generation rule v1:
- if active plan has high busy stop, return warning message
- else return balanced-day message

### 4) Test
Create suggestion, then update status to `snoozed`.

## Acceptance checks
- suggestion is persisted and status change is reflected on next fetch

## Deliverables
- first assistant-like feature without LLM.
