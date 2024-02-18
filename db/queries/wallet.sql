-- name: CreateWallet :one 
INSERT INTO "wallets" (
    "userId",
    balance
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetWallet :one
SELECT balance FROM "wallets" WHERE "userId"=$1;

-- name: UpdateBalance :one
UPDATE "wallets" SET balance = balance + $1 WHERE "userId" = $2 RETURNING *;