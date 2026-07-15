-- name: CreateDaySession :one
INSERT INTO day_sessions (trip_id, date, start_time,start_label)
VALUES($1, $2, $3, $4)
RETURNING id, trip_id, date, start_time, start_label, created_at;

-- name: GetDaySessionByIDAndDate :one
SELECT id, trip_id, date, start_time, start_label, created_at
FROM day_sessions 
WHERE trip_id = $1
AND date = $2
LIMIT 1;

-- name: GetByID :many  
SELECT id, trip_id, date, start_time, start_label, created_at
FROM day_sessions 
WHERE trip_id= $1;