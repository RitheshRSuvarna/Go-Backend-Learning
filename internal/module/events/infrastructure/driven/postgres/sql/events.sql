-- name: CreateEvents :one
INSERT INTO events(day_session_id, type, ts, payload_json)
VALUES($1,$2,$3,$4)
RETURNING id, day_session_id, type, ts, payload_json, created_at;

-- name: GetEvents :many
SELECT id, day_session_id, type, ts, payload_json, created_at
FROM events
WHERE day_session_id=$1
ORDER BY created_at DESC;

