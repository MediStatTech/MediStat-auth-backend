-- name: GetPersonalByID :one
SELECT 
    p.*,
    sr.code as role_code,
    sr.name as role_name
FROM personal p
JOIN staff_roles sr ON p.role_id = sr.role_id
WHERE p.staff_id = $1;

-- name: GetPersonalByEmail :one
SELECT 
    p.*,
    sr.code as role_code,
    sr.name as role_name
FROM personal p
JOIN staff_roles sr ON p.role_id = sr.role_id
WHERE p.email = $1;

-- name: ListPersonal :many
SELECT 
    p.*,
    sr.code as role_code,
    sr.name as role_name
FROM personal p
JOIN staff_roles sr ON p.role_id = sr.role_id
WHERE 
    (p.dismissed_at IS NULL OR p.dismissed_at > CURRENT_DATE)
ORDER BY p.staff_id
LIMIT $1 OFFSET $2;

-- name: ListPersonalByRole :many
SELECT 
    p.*,
    sr.code as role_code,
    sr.name as role_name
FROM personal p
JOIN staff_roles sr ON p.role_id = sr.role_id
WHERE 
    sr.code = $1
    AND (p.dismissed_at IS NULL OR p.dismissed_at > CURRENT_DATE)
ORDER BY p.staff_id;

-- name: ListPersonalByDepartment :many
SELECT 
    p.*,
    sr.code as role_code,
    sr.name as role_name
FROM personal p
JOIN staff_roles sr ON p.role_id = sr.role_id
WHERE 
    p.department = $1
    AND (p.dismissed_at IS NULL OR p.dismissed_at > CURRENT_DATE)
ORDER BY p.staff_id;

-- name: CreatePersonal :one
INSERT INTO personal (
    full_name, 
    role_id, 
    department, 
    phone, 
    email, 
    hired_at
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdatePersonal :one
UPDATE personal
SET 
    full_name = COALESCE($2, full_name),
    role_id = COALESCE($3, role_id),
    department = COALESCE($4, department),
    phone = COALESCE($5, phone),
    email = COALESCE($6, email),
    updated_at = NOW()
WHERE staff_id = $1
RETURNING *;

-- name: DismissPersonal :one
UPDATE personal
SET 
    dismissed_at = CURRENT_DATE,
    updated_at = NOW()
WHERE staff_id = $1
RETURNING *;

-- name: RestorePersonal :one
UPDATE personal
SET 
    dismissed_at = NULL,
    updated_at = NOW()
WHERE staff_id = $1
RETURNING *;

-- name: DeletePersonal :exec
DELETE FROM personal
WHERE staff_id = $1;

-- name: CountActivePersonal :one
SELECT COUNT(*) as total
FROM personal
WHERE dismissed_at IS NULL OR dismissed_at > CURRENT_DATE;

-- name: CountPersonalByRole :one
SELECT COUNT(*) as total
FROM personal p
JOIN staff_roles sr ON p.role_id = sr.role_id
WHERE 
    sr.code = $1
    AND (p.dismissed_at IS NULL OR p.dismissed_at > CURRENT_DATE);

-- name: SearchPersonal :many
SELECT 
    p.*,
    sr.code as role_code,
    sr.name as role_name
FROM personal p
JOIN staff_roles sr ON p.role_id = sr.role_id
WHERE 
    (p.full_name ILIKE '%' || $1 || '%' 
     OR p.email ILIKE '%' || $1 || '%'
     OR p.phone ILIKE '%' || $1 || '%')
    AND (p.dismissed_at IS NULL OR p.dismissed_at > CURRENT_DATE)
ORDER BY p.staff_id
LIMIT $2 OFFSET $3;