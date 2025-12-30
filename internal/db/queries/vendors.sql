-- name: GetAllVendors :many
SELECT * FROM vendors
ORDER BY name;

-- name: GetVendorById :one
SELECT * FROM vendors WHERE id = $1;

-- name: CreateVendor :one
INSERT INTO vendors (
  name,
  address,
  contact_person,
  contact_number
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateVendor :one
UPDATE vendors SET
  name = $2,
  address = $3,
  contact_person = $4,
  contact_number = $5,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteVendor :exec
DELETE FROM vendors WHERE id = $1;
