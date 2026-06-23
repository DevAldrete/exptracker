-- Queries for Expense Tracker

--- USERS ---

-- name: GetUserById :one
SELECT 
  id, 
  name, 
  email, 
  created_at, 
  updated_at, 
  deleted_at 
FROM users 
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, 
  name, 
  email, 
  created_at, 
  updated_at, 
  deleted_at 
FROM users 
WHERE email = $1 
AND deleted_at IS NULL
LIMIT 1;

-- name: GetUsers :many
SELECT id, 
  name, 
  email, 
  created_at, 
  updated_at, 
  deleted_at 
FROM users 
WHERE deleted_at IS NULL
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: GetUsersByRole :many
SELECT 
  users.id AS user_id, 
  users.name AS user_name, 
  users.email AS user_email, 
  roles.id AS role_id, 
  roles.name AS role_name
  FROM users
  JOIN users_roles
  ON users.id = users_roles.user_id
  JOIN roles 
  ON roles.id = users_roles.role_id
  WHERE roles.id = $1 
  AND users_roles.is_active = true 
  AND users.deleted_at IS NULL
  ORDER BY roles.name DESC;

-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  password
) VALUES ( $1, $2, $3 )
RETURNING *;

-- name: UpdateUserById :one
UPDATE users
  SET name = $2, email = $3, password = $4
  WHERE id = $1
RETURNING *;

-- name: DeleteUserById :exec
UPDATE users
  SET deleted_at = NOW()
  WHERE id = $1;

-- name: DeleteUserByEmail :exec
UPDATE users
  SET deleted_at = NOW()
  WHERE email = $1;

--- EXPENSES ---

-- name: GetExpenseById :one
SELECT * 
FROM expenses 
WHERE id = $1 
LIMIT 1;

-- name: GetExpenses :many
SELECT * 
FROM expenses 
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: GetExpensesByUserId :many
SELECT 
  expenses.id AS expense_id, 
  expenses.name AS expense_name, 
  expenses.cents, 
  users.id AS user_id, 
  users.name AS user_name, 
  users.email AS user_email
FROM expenses
  LEFT JOIN users
  ON users.id = expenses.user_id
WHERE expenses.user_id = $1
AND users.deleted_at IS NULL;

-- name: GetExpensesByUserEmail :many
SELECT 
  expenses.id AS expense_id, 
  expenses.name AS expense_name, 
  expenses.cents, 
  users.id AS user_id, 
  users.name AS user_name, 
  users.email AS user_email
FROM expenses
  LEFT JOIN users
  ON users.id = expenses.user_id
WHERE users.email = $1
AND users.deleted_at IS NULL;

-- name: CreateExpense :one
INSERT INTO expenses (
  name,
  description,
  cents,
  user_id
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: UpdateExpenseById :one
UPDATE expenses
  SET name = $2, description = $3, cents = $4 
  WHERE id = $1
RETURNING *;

-- name: DeleteExpenseById :exec
UPDATE expenses
  SET deleted_at = NOW()
  WHERE id = $1;

--- Roles ---

-- GetRoleById :one
SELECT *
FROM roles
WHERE id = $1
LIMIT 1;

-- GetRoles :many
SELECT *
FROM roles
ORDER BY id
LIMIT $1
OFFSET $2;

-- CreateRole :one
INSERT INTO roles (
  name,
  description
) VALUES ( $1, $2 );

-- UpdateRoleById :one
UPDATE roles
  SET name = $2, description = $3
  WHERE id = $1;

-- DeleteRoleById :exec
DELETE FROM roles
  WHERE id = $1;

--- Permissions ---

-- GetPermissionById :one
SELECT *
FROM permissions
WHERE id = $1
LIMIT 1;

-- GetPermissions :many
SELECT *
FROM permissions
ORDER BY id
LIMIT $1
OFFSET $2;

-- CreateRole :one
INSERT INTO permissions (
  name
) VALUES ( $1, $2 );

-- UpdateRoleById :one
UPDATE permissions
  SET name = $2
  WHERE id = $1;

-- DeleteRoleById :exec
DELETE FROM permissions
  WHERE id = $1;
