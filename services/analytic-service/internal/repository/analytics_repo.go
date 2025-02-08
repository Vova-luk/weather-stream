package repository

import (
	"fmt"

	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/models"
	"github.com/jmoiron/sqlx"
)

type AnalyticsRepository struct {
	db *sqlx.DB
}

func NewAnalyticsRepository(db *sqlx.DB) *AnalyticsRepository {
	return &AnalyticsRepository{
		db: db,
	}
}

func (a *AnalyticsRepository) GetAnalyticsById(locationId int32, interval string) (*models.WeatherAnalytic, error) {
	query := fmt.Sprintf(`SELECT location_id, 
							AVG(temperature) AS avg_temperature, 
							AVG(humidity) AS avg_humidity, 
							AVG(wind_speed) AS avg_wind_speed, 
							AVG(pressure) AS avg_pressure, 
							AVG(precip) AS avg_precip, 
							AVG(cloud) AS avg_cloud
			  			 FROM weather_analysis 
						 WHERE updated_at >= NOW() - INTERVAL '%s' AND location_id = $1
			  			 GROUP BY location_id;`, interval)

	var weatherAnalytic models.WeatherAnalytic
	err := a.db.Get(&weatherAnalytic, query, locationId)
	if err != nil {
		return nil, err
	}

	return &weatherAnalytic, nil

}

func (a *AnalyticsRepository) AddWeatherData(weatherData *models.WeatherData) error {
	query := `INSERT INTO weather_analysis (location_id, temperature, humidity, wind_speed, pressure, precip, cloud)
			VALUES (:location_id, :temperature, :humidity, :wind_speed, :pressure, :precip, :cloud)`

	_, err := a.db.NamedExec(query, weatherData)

	return err
}
