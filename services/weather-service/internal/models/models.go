package models

import (
	"time"

	weatherPb "github.com/Vova-luk/weather-stream/services/weather-service/proto"
)

type WeatherData struct {
	LocationId  int32     `db:"location_id"`
	Temperature float64   `db:"temperature"`
	Humidity    int32     `db:"humidity"`
	WindSpeed   float64   `db:"wind_speed"`
	Pressure    float64   `db:"pressure"`
	Precip      float64   `db:"precip"`
	Cloud       int32     `db:"cloud"`
	UpdateAt    time.Time `db:"updated_at"`
}

func (w *WeatherData) ToProto() *weatherPb.WeatherData {
	return &weatherPb.WeatherData{
		LocationId:  w.LocationId,
		Temperature: w.Temperature,
		Humidity:    w.Humidity,
		WindSpeed:   w.WindSpeed,
		Pressure:    w.Pressure,
		Precip:      w.Precip,
		Cloud:       w.Cloud,
		UpdateAt:    w.UpdateAt.GoString()}
}
