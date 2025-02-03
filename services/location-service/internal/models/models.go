package models

import (
	locationPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location"
)

type Location struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Coordinates string `db:"coordinates" json:"coordinates"`
}

func ToProto(locations []*Location) []*locationPb.Location {
	locationsProto := make([]*locationPb.Location, len(locations))
	for idx, location := range locations {
		locationsProto[idx] = &locationPb.Location{
			LocationId:  int32(location.ID),
			Name:        location.Name,
			Coordinates: location.Coordinates,
		}
	}
	return locationsProto
}
