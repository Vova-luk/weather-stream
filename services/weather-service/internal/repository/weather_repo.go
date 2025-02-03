package repository

import (
	"github.com/Vova-luk/weather-stream/services/weather-service/internal/models"
	"github.com/jmoiron/sqlx"
)

type WeatherRepository struct {
	db *sqlx.DB
}

func NewWeatherRepository(db *sqlx.DB) *WeatherRepository {
	return &WeatherRepository{
		db: db,
	}
}

func (w *WeatherRepository) GetWeatherById(locationId int32) (*models.WeatherData, error) {
	var location *models.WeatherData
	query := `SELECT * FROM weathers WHERE location_id=&1`

	err := w.db.Select(&location, query, locationId)
	if err != nil {
		return &models.WeatherData{}, err
	}

	return location, nil
}

func (w *WeatherRepository) AddForcefullyWeather(weatherLocation *models.WeatherData) error {
	query := `INSERT INTO weathers(location_id, temperature, humidity, wind_speed, pressure, precip, cloud)
		VALUES (:location_id, :temperature, :humidity, :wind_speed, :pressure, :precip, :cloud)`

	_, err := w.db.NamedExec(query, weatherLocation)

	return err
}
