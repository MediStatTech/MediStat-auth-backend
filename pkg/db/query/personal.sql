-- name: GetPersonalByID :one
SELECT * FROM personal
WHERE personal_id = $1;

-- name: GetAllPersonal :many
SELECT * FROM personal
ORDER BY created_at DESC;

-- name: CreatePersonal :one
INSERT INTO personal (
    personal_id,
    first_name,
    last_name,
    email,
    phone,
    password_hash,
    status,
    departure,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: UpdatePersonal :one
UPDATE personal
SET
    first_name = $2,
    last_name = $3,
    email = $4,
    phone = $5,
    status = $6,
    departure = $7,
    updated_at = $8
WHERE personal_id = $1
RETURNING *;

-- name: GetPersonalByEmail :one
SELECT * FROM personal
WHERE email = $1;

-- name: UpdatePersonalStatus :one
UPDATE personal
SET
    status = $2,
    updated_at = $3
WHERE personal_id = $1
RETURNING *;

-- name: ExistsPersonalByEmail :one
SELECT EXISTS(
    SELECT 1
    FROM personal
    WHERE email = $1
) AS exists;