package service

import "math"

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

func (s *UserService) RequestForARide(currentLocation Point, destination Point, userId int32) {

}
