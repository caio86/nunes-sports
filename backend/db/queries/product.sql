-- name: FindAll :many
SELECT * FROM products;

-- name: Find :many
SELECT * FROM products
LIMIT $1
OFFSET $2;

-- name: Count :one
SELECT count(*) FROM products;

-- name: FindByID :one
SELECT *
  FROM products
  WHERE id = $1;

-- name: Create :one
INSERT INTO products (
  id, name, description, price
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: Delete :exec
DELETE FROM products
  WHERE id = $1;

-- name: Update :one
UPDATE products
  SET name = $2,
  description = $3,
  price = $4
  WHERE id = $1
  RETURNING *;
