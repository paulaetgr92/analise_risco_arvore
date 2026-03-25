-- name: CreateTree :one
INSERT INTO trees (
    latitude,
    longitude,
    species,
    height,
    diameter,
    age,
    health
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         )
RETURNING *;


-- name: GetTreeByID :one
SELECT *
FROM trees
WHERE id = $1;


-- name: ListTrees :many
SELECT *
FROM trees
ORDER BY created_at DESC;


-- name: ListTreesByBoundingBox :many
SELECT *
FROM trees
WHERE latitude BETWEEN $1 AND $2
  AND longitude BETWEEN $3 AND $4;


-- name: UpdateTree :one
UPDATE trees
SET
    latitude   = $1,
    longitude  = $2,
    species    = $3,
    height     = $4,
    diameter   = $5,
    age        = $6,
    health     = $7
WHERE id = $8
RETURNING *;


-- name: DeleteTree :exec
DELETE FROM trees
WHERE id = $1;


-- name: ListPotentialRiskTrees :many
SELECT *
FROM trees
WHERE
    age > 20
   OR height > 10
   OR health IN ('poor', 'critical');
