package models

import "time"

type WeatherData struct {
	LocationId  int       `db:"location_id"`
	Temperature float64   `db:"temperature"`
	Humidity    int       `db:"humidity"`
	WindSpeed   float64   `db:"wind_speed"`
	Pressure    int       `db:"pressure"`
	Condition   string    `db:"condition"`
	UpdateAt    time.Time `db:"updated_at"`
}
