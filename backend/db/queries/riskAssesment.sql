-- name: CreateRiskAssessment :one
INSERT INTO risk_assessments (
    tree_id,
    score,
    level,
    calculated_at
) VALUES (
             $1, $2, $3, now()
         )
RETURNING *;


-- name: GetRiskByID :one
SELECT *
FROM risk_assessments
WHERE id = $1;


-- name: ListRiskByTree :many
SELECT *
FROM risk_assessments
WHERE tree_id = $1
ORDER BY calculated_at DESC;

-- name: GetLatestRiskByTree :one
SELECT *
FROM risk_assessments
WHERE tree_id = $1
ORDER BY calculated_at DESC
LIMIT 1;


-- name: ListHighRiskTrees :many
SELECT t.*, r.score, r.level
FROM trees t
         JOIN risk_assessments r ON r.tree_id = t.id
WHERE r.level = 'high'
  AND r.calculated_at = (
    SELECT MAX(r2.calculated_at)
    FROM risk_assessments r2
    WHERE r2.tree_id = t.id
);



-- name: ListMediumOrHighRiskTrees :many
SELECT t.*, r.score, r.level
FROM trees t
         JOIN risk_assessments r ON r.tree_id = t.id
WHERE r.level IN ('medium', 'high')
  AND r.calculated_at = (
    SELECT MAX(r2.calculated_at)
    FROM risk_assessments r2
    WHERE r2.tree_id = t.id
);
