-- name: GetAllCustomers :many
SELECT * FROM customers
ORDER BY name;

-- name: GetCustomerById :one
SELECT * FROM customers WHERE id = $1;

-- name: CreateCustomer :one
INSERT INTO customers (
  name,
  address,
  contact_number
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers SET
  name = $2,
  address = $3,
  contact_number = $4,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers WHERE id = $1;
