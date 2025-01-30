package handler

import (
	"context"

	"github.com/Vova-luk/weather-stream/location-service/internal/models"
	"github.com/Vova-luk/weather-stream/location-service/internal/service"
	pb "github.com/Vova-luk/weather-stream/location-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sirupsen/logrus"
)

type LocationHandler struct {
	pb.UnimplementedLocationServiceServer
	locationService *service.LocationService
	log             *logrus.Logger
}

func NewLocationHanadler(locationService *service.LocationService, log *logrus.Logger) *LocationHandler {
	return &LocationHandler{locationService: locationService,
		log: log}
}

func (l *LocationHandler) CreateLocation(ctx context.Context, request *pb.CreateLocationRequest) (*pb.CreateLocationResponse, error) {
	location := &models.Location{
		Name:        request.Name,
		Coordinates: request.Coordinates,
	}
	locationId, err := l.locationService.CreateLocationService(location)
	if err != nil {

		return nil, status.Errorf(codes.Internal, "failed to create location: %v", err.Error())
	}

	return &pb.CreateLocationResponse{
		LocationId: int32(locationId),
	}, nil
}

func (l *LocationHandler) GetLocations(ctx context.Context, _ *pb.Empty) (*pb.GetLocationsResponse, error) {
	locations, err := l.locationService.GetLocationsService()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get locations: %v", err.Error())
	}
	return &pb.GetLocationsResponse{
		Items: locations,
	}, nil
}
