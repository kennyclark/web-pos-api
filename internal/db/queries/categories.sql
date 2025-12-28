-- name: GetAllCategories :many
SELECT * FROM categories
ORDER BY name;

-- name: GetCategoryById :one
SELECT * FROM categories WHERE id = $1;

-- name: CreateCategory :one
INSERT INTO categories (name) VALUES ($1) RETURNING *;

-- name: UpdateCategory :one
UPDATE categories SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;
