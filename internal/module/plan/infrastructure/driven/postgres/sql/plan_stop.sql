-- name: CreatePlanStop :one
INSERT INTO plan_stops (
    plan_version_id,
    position,
    title,
    category_label,
    image_url,
    planned_arrival,
    planned_departure,
    travel_minutes,
    stay_minutes,
    busy_risk_label
)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING
    id,
    plan_version_id,
    position,
    title,
    category_label,
    image_url,
    planned_arrival,
    planned_departure,
    travel_minutes,
    stay_minutes,
    busy_risk_label,
    created_at;

-- name: ListPlanStopsByPlanVersionID :many
SELECT
    id,
    plan_version_id,
    position,
    title,
    category_label,
    image_url,
    planned_arrival,
    planned_departure,
    travel_minutes,
    stay_minutes,
    busy_risk_label,
    created_at
FROM plan_stops
WHERE plan_version_id = $1
ORDER BY position ASC;