package service

import (
	"context"
	"fmt"
	"math"

	"github.com/Owoade/go-uber/sql"
	"github.com/Owoade/go-uber/utils"
)

func toRadian(deg float64) float64 {
	return (deg * math.Pi) / 180
}

func calculateRideDistance(currentLoaction utils.SqlPoint, destination utils.SqlPoint) float64 {

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

func (s *UserService) RequestForARide(currentLocation utils.SqlPoint, destination utils.SqlPoint, userId int32) int64 {

	rideDistance := calculateRideDistance(currentLocation, destination)

	ridePrice := getRidePrice(rideDistance)

	return ridePrice

}

func (s *UserService) InitiateTrip(ctx context.Context, ridePrice int64, userId int32, pickupLocation utils.SqlPoint, destination utils.SqlPoint) (sql.Trip, error) {

	UserId := utils.SqlTypeInt32(userId)

	wallet, err := s.repo.GetWallet(ctx, UserId)

	if err != nil {
		return *new(sql.Trip), err
	}

	if wallet.Balance.Int64 < ridePrice {
		return *new(sql.Trip), fmt.Errorf("Insufficient balance")
	}

	// create a transaction

	tripTransaction, err := s.repo.CreateWalletTransaction(ctx, sql.CreateWalletTransactionParams{
		WalletId: utils.SqlTypeInt32(wallet.ID),
		Amount:   utils.SqlTypeInt64(ridePrice),
		Type:     utils.SqlTypeText("debit"),
	})

	if err != nil {
		return *new(sql.Trip), fmt.Errorf(err.Error())
	}

	s.repo.UpdateBalance(ctx, sql.UpdateBalanceParams{
		Balance: utils.SqlTypeInt64(-ridePrice),
		UserId:  utils.SqlTypeInt32(userId),
	})

	// create trip

	newTrip, err := s.repo.CreateTrip(ctx, sql.CreateTripParams{
		TransactionId:  utils.SqlTypeInt32(tripTransaction.ID),
		UserId:         utils.SqlTypeInt32(userId),
		PickUpLocation: utils.SqlTypePoint(pickupLocation),
		Destination:    utils.SqlTypePoint(destination),
	})

	if err != nil {
		return *new(sql.Trip), fmt.Errorf(err.Error())
	}

	return newTrip, nil

}

func (s *UserService) UpdateTripLocation(ctx context.Context, tripId int32, userId int32, location utils.SqlPoint) (err error) {

	payload := sql.UpdateTripLocationByUserParams{
		CurrentTripLocationFromUser: utils.SqlTypePoint(location),
		ID:                          tripId,
		UserId:                      utils.SqlTypeInt32(userId),
	}

	err = s.repo.UpdateTripLocationByUser(ctx, payload)

	if err != nil {
		return err
	}

	return *new(error)

}
