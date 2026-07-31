-- name: CreateDaySession :one
INSERT INTO day_sessions (trip_id, date, start_time,start_label)
VALUES($1, $2, $3, $4)
RETURNING id, trip_id, date, start_time, start_label, active_plan_version_id, created_at;

-- name: GetDaySession :one
SELECT id, trip_id, date, start_time, start_label, active_plan_version_id, created_at
FROM day_sessions 
WHERE id = $1;

-- name: ListDaySession :many  
SELECT id, trip_id, date, start_time, start_label, active_plan_version_id, created_at
FROM day_sessions 
WHERE trip_id = $1;

-- name: UpdateActivePlan :exec
UPDATE day_sessions
SET active_plan_version_id = $2
WHERE id = $1;
