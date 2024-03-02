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

-- name: UpdateTripLocationByDriver :exec
UPDATE "trip" SET "currentTripLocation"=$1 WHERE id=$2 AND "driverId"=$3;

-- name: UpdateTripLocationByUser :exec
UPDATE "trip" SET "currentTripLocationFromUser"=$1 WHERE id=$2 AND "userId"=$3;