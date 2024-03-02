package service

import (
	"context"

	"github.com/Owoade/go-uber/sql"
	"github.com/Owoade/go-uber/utils"
)

func (s *DriverService) UpdateTripLocation(ctx context.Context, tripId int32, driverId int32, location utils.SqlPoint) (err error) {

	payload := sql.UpdateTripLocationByDriverParams{
		CurrentTripLocation: utils.SqlTypePoint(location),
		ID:                  tripId,
		DriverId:            utils.SqlTypeInt32(driverId),
	}

	err = s.repo.UpdateTripLocationByDriver(ctx, payload)

	if err != nil {
		return err
	}

	return *new(error)

}
