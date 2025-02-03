package handler

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Vova-luk/weather-stream/services/weather-service/internal/service"
	weatherPb "github.com/Vova-luk/weather-stream/services/weather-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WeatherHandler struct {
	weatherPb.UnimplementedWeatherServiceServer
	weatherService *service.WeatherService
}

func NewWeatherHandler(weatherService *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}

func (w *WeatherHandler) GetWeatherById(ctx context.Context, request *weatherPb.GetWeatherByIdRequest) (*weatherPb.GetWeatherByIdResponse, error) {
	location, err := w.weatherService.GetWeatherById(request.LocationId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "Weather data not found for LocationID %d", request.LocationId)
		}
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &weatherPb.GetWeatherByIdResponse{
		Weather: location.ToProto(),
	}, nil

}
