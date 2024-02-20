-- name: CreateTrip :one
INSERT INTO "trip" (
    "userId",
    "driverId",
    "transactionId",
    "pickUpLocation",
    "destination",
    "currentTripLocation"
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;