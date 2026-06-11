-- name: CreateTrip :one
INSERT INTO trips (destination, start_date, end_date, travelers_count)
VALUES($1,$2,$3,$4)
RETURNING id, destination,start_date, end_date, travelers_count, created_at;

-- name: ListTrips :many
SELECT id, destination, start_date, end_date, travelers_count, created_at
FROM trips
ORDER BY created_at DESC;