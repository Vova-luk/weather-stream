package handler

import (
	"context"

	"github.com/Vova-luk/weather-stream/services/location-service/internal/models"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/service"
	locationPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sirupsen/logrus"
)

type LocationHandler struct {
	locationPb.UnimplementedLocationServiceServer
	locationService *service.LocationService
	log             *logrus.Logger
}

func NewLocationHanadler(locationService *service.LocationService, log *logrus.Logger) *LocationHandler {
	return &LocationHandler{locationService: locationService,
		log: log}
}

func (l *LocationHandler) CreateLocation(ctx context.Context, request *locationPb.CreateLocationRequest) (*locationPb.CreateLocationResponse, error) {
	location := &models.Location{
		Name:        request.Name,
		Coordinates: request.Coordinates,
	}
	locationId, err := l.locationService.CreateLocationService(location)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create location: %v", err.Error())
	}

	return &locationPb.CreateLocationResponse{
		LocationId: locationId,
	}, nil
}

func (l *LocationHandler) GetLocations(ctx context.Context, _ *locationPb.Empty) (*locationPb.GetLocationsResponse, error) {
	locations, err := l.locationService.GetLocationsService()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get locations: %v", err.Error())
	}
	return &locationPb.GetLocationsResponse{
		Locations: locations,
	}, nil
}
