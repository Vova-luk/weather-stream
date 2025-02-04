package handler

import (
	"context"

	"github.com/Vova-luk/weather-stream/services/location-service/internal/service"
	locationWeatherPb "github.com/Vova-luk/weather-stream/services/location-service/proto/weather"
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

func (w *WeatherHandler) GetLocationById(ctx context.Context, request *locationWeatherPb.GetLocationByIdRequest) (*locationWeatherPb.GetLocationByIdResponce, error) {
	locationId := request.GetLocationId()

	responce, err := w.weatherService.GetLocationById(locationId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting current weather: %v", err)
	}
	return &locationWeatherPb.GetLocationByIdResponce{Weather: responce}, nil
}
