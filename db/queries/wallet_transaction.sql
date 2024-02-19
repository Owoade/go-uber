-- name: CreateWalletTransaction :one
INSERT INTO "wallet_transactions" (
    "walletId",
    "amount",
    "type"
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;