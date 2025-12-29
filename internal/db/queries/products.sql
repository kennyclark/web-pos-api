-- name: GetAllProducts :many
SELECT * FROM products
ORDER BY name;

-- name: GetProductById :one
SELECT * FROM products WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (
  name,
  description,
  cost_price_cents,
  selling_price_cents,
  quantity_on_hand,
  low_stock_threshold,
  category_id,
  is_active
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET
  name = $2,
  description = $3,
  cost_price_cents = $4,
  selling_price_cents = $5,
  quantity_on_hand = $6,
  low_stock_threshold = $7,
  category_id = $8,
  is_active = $9,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;
