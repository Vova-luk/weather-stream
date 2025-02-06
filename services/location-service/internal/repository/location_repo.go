package repository

import (
	"github.com/Vova-luk/weather-stream/services/location-service/internal/models"
	"github.com/jmoiron/sqlx"
)

type LocationRepository struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

func (l *LocationRepository) CreateLocation(location *models.Location) (int32, error) {
	var locationId int32
	query := `INSERT INTO locations (name, coordinates) 
			  VALUES ($1, $2) 
			  RETURNING id`
	err := l.db.Get(&locationId, query, location.Name, location.Coordinates)
	if err != nil {
		return 0, err
	}
	return locationId, nil
}

func (l *LocationRepository) GetLocations() ([]*models.Location, error) {
	var locations []*models.Location

	query := `SELECT * FROM locations`

	err := l.db.Select(&locations, query)
	if err != nil {
		return nil, err
	}
	return locations, nil
}
