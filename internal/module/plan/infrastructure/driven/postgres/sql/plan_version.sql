-- name: CreatePlanVersion :one
INSERT INTO plan_versions ( day_session_id, version, notes)
VALUES ($1, $2, $3)
RETURNING id, day_session_id, version, notes, created_at;

-- name: GetPlanVersionByID :one
SELECT id, day_session_id, version, notes, created_at
FROM plan_versions
WHERE id = $1;

-- name: ListPlanVersionsByDaySessionID :many
SELECT id, day_session_id, version, notes, created_at
FROM plan_versions
WHERE day_session_id = $1
ORDER BY version ASC;