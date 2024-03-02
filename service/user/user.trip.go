package service

import (
	"context"
	"fmt"
	"math"

	"github.com/Owoade/go-uber/sql"
	"github.com/jackc/pgx/v5/pgtype"
)

type Point struct {
	Longitude float64
	Latitude  float64
}

func toRadian(deg float64) float64 {
	return (deg * math.Pi) / 180
}

func calculateRideDistance(currentLoaction Point, destination Point) float64 {

	originLatitude := currentLoaction.Latitude
	originLongitude := currentLoaction.Longitude

	destinationLatitude := destination.Latitude
	destinationLongitude := destination.Longitude

	const (
		PI = math.Pi
		R  = 6371000
	)

	orgLatRadian := toRadian(originLatitude)
	destLatRadian := toRadian(destinationLatitude)

	deltaPhi := destinationLatitude - originLatitude
	deltaLambda := destinationLongitude - originLongitude

	a := math.Pow(math.Sin(deltaPhi/2), 2) + math.Sin(orgLatRadian)*math.Cos(destLatRadian)*math.Pow(math.Sin(deltaLambda/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	meters := R * c

	return meters

}

func getRidePrice(distance float64) int64 {
	return int64(3 * distance)
}

func (s *UserService) RequestForARide(currentLocation Point, destination Point, userId int32) int64 {

	rideDistance := calculateRideDistance(currentLocation, destination)

	ridePrice := getRidePrice(rideDistance)

	return ridePrice

}

func (s *UserService) InitiateTrip(ctx context.Context, ridePrice int64, userId int32, pickupLocation Point, destination Point) (sql.Trip, error) {

	UserId := pgtype.Int4{
		Int32: userId,
		Valid: true,
	}

	wallet, err := s.repo.GetWallet(ctx, UserId)

	if err != nil {
		return *new(sql.Trip), err
	}

	if wallet.Balance.Int64 < ridePrice {
		return *new(sql.Trip), fmt.Errorf("Insufficient balance")
	}

	// create a transaction

	tripTransaction, err := s.repo.CreateWalletTransaction(ctx, sql.CreateWalletTransactionParams{
		WalletId: pgtype.Int4{
			Int32: wallet.ID,
		},
		Amount: pgtype.Int8{
			Int64: ridePrice,
		},
		Type: pgtype.Text{
			String: "debit",
		},
	})

	if err != nil {
		return *new(sql.Trip), fmt.Errorf(err.Error())
	}

	s.repo.UpdateBalance(ctx, sql.UpdateBalanceParams{
		Balance: pgtype.Int8{
			Int64: -ridePrice,
		},
		UserId: pgtype.Int4{
			Int32: userId,
		},
	})

	// create trip

	newTrip, err := s.repo.CreateTrip(ctx, sql.CreateTripParams{
		TransactionId: pgtype.Int4{
			Int32: tripTransaction.ID,
			Valid: true,
		},
		UserId: UserId,
		PickUpLocation: pgtype.Point{
			P: pgtype.Vec2{
				X: pickupLocation.Longitude,
				Y: pickupLocation.Latitude,
			},
			Valid: true,
		},
		Destination: pgtype.Point{
			P: pgtype.Vec2{
				X: destination.Longitude,
				Y: destination.Latitude,
			},
			Valid: true,
		},
	})

	if err != nil {
		return *new(sql.Trip), fmt.Errorf(err.Error())
	}

	return newTrip, nil

}

func (s *UserService) UpdateTripLocation()
