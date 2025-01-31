package service

import (
	"github.com/Vova-luk/weather-stream/services/location-service/internal/models"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/repository"
	pb "github.com/Vova-luk/weather-stream/services/location-service/proto"
	"github.com/sirupsen/logrus"
)

type LocationService struct {
	locationRepository *repository.LocationRepository
	log                *logrus.Logger
}

func NewLocationService(locationRepository *repository.LocationRepository, log *logrus.Logger) *LocationService {
	return &LocationService{locationRepository: locationRepository,
		log: log}
}

func (l *LocationService) CreateLocationService(location *models.Location) (int, error) {
	locationId, err := l.locationRepository.CreateLocation(location)
	if err != nil {
		l.log.Errorf("error when creating location: %v", err.Error())
		return 0, err
	}
	l.log.Infof("Location saved successfully, ID : %d", locationId)
	return locationId, nil
}

func (l *LocationService) GetLocationsService() ([]*pb.Item, error) {
	locations, err := l.locationRepository.GetLocations()
	if err != nil {
		l.log.Errorf("error when getting locations: %v", err.Error())
		return nil, err
	}
	locationsItems := make([]*pb.Item, len(locations))
	for ind, location := range locations {
		locationsItems[ind] = &pb.Item{
			LocationId:  int32(location.ID),
			Name:        location.Name,
			Coordinates: location.Coordinates,
		}
	}
	l.log.Info("Locations successfully received")
	return locationsItems, nil
}
