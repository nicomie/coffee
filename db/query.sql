-- name: GetCoffee :one
SELECT * FROM coffees
WHERE id = $1 LIMIT 1;

-- name: ListCoffees :many
SELECT * FROM coffees
ORDER BY name;

-- name: CreateCoffee :one
INSERT INTO coffees (
  name, flavor, acidity, image_src
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteCoffee :exec
DELETE FROM coffees
WHERE id = $1;

-- name: UpdateCoffee :one
UPDATE coffees
set name = $2,
flavor = $3,
acidity = $4,
image_src = $5
WHERE id = $1
RETURNING *;