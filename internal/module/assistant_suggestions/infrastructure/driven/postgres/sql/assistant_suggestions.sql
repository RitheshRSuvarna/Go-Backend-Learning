-- name: CreateAssistantSuggestions :one
INSERT INTO assistant_suggestions(day_session_id, message, status)
VALUES($1, $2, $3)
RETURNING id, day_session_id, message, status, created_at;

-- name: GetAssistantSuggestions :many
SELECT id, day_session_id, message, status, created_at
FROM assistant_suggestions
WHERE day_session_id = $1;

-- name: EditAssistantSuggestion :one
UPDATE assistant_suggestions
SET
    message = $2,
    status = $3
WHERE id = $1
RETURNING id, day_session_id, message, status, created_at;