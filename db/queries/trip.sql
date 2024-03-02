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

-- name: UpdateTripLocationByDriver :one
UPDATE "trip" SET "currentTripLocation"=$1 WHERE id=$2 AND "driverId"=$3 RETURNING *;

-- name: UpdateTripLocationByUser :one
UPDATE "trip" SET "currentTripLocationFromUser"=$1 WHERE id=$2 AND "userId"=$3 RETURNING *;