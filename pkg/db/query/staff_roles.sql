-- name: GetStaffRole :one
SELECT * FROM staff_roles
WHERE role_id = $1;

-- name: GetStaffRoleByCode :one
SELECT * FROM staff_roles
WHERE code = $1;

-- name: ListStaffRoles :many
SELECT * FROM staff_roles
ORDER BY role_id;

-- name: CreateStaffRole :one
INSERT INTO staff_roles (code, name)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateStaffRole :one
UPDATE staff_roles
SET name = $2
WHERE role_id = $1
RETURNING *;

-- name: DeleteStaffRole :exec
DELETE FROM staff_roles
WHERE role_id = $1;