package models

import (
	"time"

	analyticPb "github.com/Vova-luk/weather-stream/services/analytic-service/proto"
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
type WeatherAnalytic struct {
	LocationId     int32   `db:"location_id"`
	AvgTemperature float64 `db:"avg_temperature"`
	AvgHumidity    float64 `db:"avg_humidity"`
	AvgWindSpeed   float64 `db:"avg_wind_speed"`
	AvgPressure    float64 `db:"avg_pressure"`
	AvgPrecip      float64 `db:"avg_precip"`
	AvgCloud       float64 `db:"avg_cloud"`
}

func (w *WeatherAnalytic) ToProto() *analyticPb.AnalyticsWeather {
	return &analyticPb.AnalyticsWeather{
		LocationId:     w.LocationId,
		AvgTemperature: w.AvgTemperature,
		AvgHumidity:    w.AvgHumidity,
		AvgWindSpeed:   w.AvgWindSpeed,
		AvgPressure:    w.AvgPressure,
		AvgPrecip:      w.AvgPrecip,
		AvgCloud:       w.AvgCloud,
	}
}
