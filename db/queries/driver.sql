-- name: CreateDriver :one
INSERT INTO "drivers" (
    email,
    password
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetDriver :one
SELECT * FROM "drivers" WHERE email=$1 LIMIT 1;

-- name: GetDriverById :one 
SELECT * FROM "drivers" WHERE id=$1 LIMIT 1;

-- name: UpdateDriverLastLocation :exec
UPDATE "drivers" SET "lastLocation"=$1 WHERE id=$2;
