-- name: CreateDaySession :one
INSERT INTO day_sessions (trip_id, date, start_time,start_label)
VALUES($1, $2, $3, $4)
RETURNING id, trip_id, date, start_time, start_label, created_at;

-- name: GetDaySessionByIDAndDate :one
SELECT id, trip_id, date, start_time, start_label, created_at
FROM day_sessions
ORDER BY created_at DESC;

-- name: GetByID :one
SELECT id, trip_id, date, start_time, start_label, created_at
FROM day_sessions
ORDER BY created_at DESC;