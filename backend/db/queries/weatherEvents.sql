-- name: CreateWeatherEvent :one
INSERT INTO weather_events (
    region,
    wind_speed,
    rainfall_mm,
    severity,
    occurred_at
) VALUES (
             $1, $2, $3, $4, $5
         )
RETURNING *;



-- name: GetWeatherEventByID :one
SELECT *
FROM weather_events
WHERE id = $1;


-- name: ListWeatherEvents :many
SELECT *
FROM weather_events
ORDER BY occurred_at DESC;



-- name: ListWeatherEventsByRegion :many
SELECT *
FROM weather_events
WHERE region = $1
ORDER BY occurred_at DESC;


-- name: ListRecentWeatherEvents :many
SELECT *
FROM weather_events
WHERE occurred_at >= now() - INTERVAL '24 hours'
ORDER BY occurred_at DESC;


-- name: ListSevereWeatherEvents :many
SELECT *
FROM weather_events
WHERE severity = 'high'
ORDER BY occurred_at DESC;


-- name: ListHighWindEvents :many
SELECT *
FROM weather_events
WHERE wind_speed >= $1
ORDER BY occurred_at DESC;


-- name: ListHeavyRainEvents :many
SELECT *
FROM weather_events
WHERE rainfall_mm >= $1
ORDER BY occurred_at DESC;


-- name: GetLatestSevereWeatherByRegion :one
SELECT *
FROM weather_events
WHERE region = $1
  AND severity = 'high'
ORDER BY occurred_at DESC
LIMIT 1;
