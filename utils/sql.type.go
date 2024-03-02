package utils

import "github.com/jackc/pgx/v5/pgtype"

type SqlPoint struct {
	Longitude float64
	Latitude  float64
}

func SqlTypeInt32(val int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: val,
		Valid: true,
	}
}

func SqlTypeInt64(val int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: val,
		Valid: true,
	}
}

func SqlTypeText(val string) pgtype.Text {
	return pgtype.Text{
		String: val,
		Valid:  true,
	}
}

func SqlTypePoint(val SqlPoint) pgtype.Point {
	return pgtype.Point{
		P: pgtype.Vec2{
			X: val.Longitude,
			Y: val.Latitude,
		},
		Valid: true,
	}
}
