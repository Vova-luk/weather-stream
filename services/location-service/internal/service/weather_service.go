package service

import (
	"context"

	locationWeatherPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location_weather"
	weatherPb "github.com/Vova-luk/weather-stream/services/weather-service/proto"
	"github.com/sirupsen/logrus"
)

type WeatherService struct {
	weatherClient weatherPb.WeatherServiceClient
	log           *logrus.Logger
}

func NewWeatherService(weatherClient weatherPb.WeatherServiceClient, log *logrus.Logger) *WeatherService {
	return &WeatherService{
		weatherClient: weatherClient,
		log:           log,
	}
}

func (w *WeatherService) GetLocationById(locationId int32) (*locationWeatherPb.WeatherData, error) {
	request := weatherPb.GetWeatherByIdRequest{LocationId: locationId}

	responce, err := w.weatherClient.GetWeatherById(context.Background(), &request)
	if err != nil {
		w.log.Warnf("error receiving response from weather-service: %v", err)
		return nil, err
	}
	weatherData := responce.GetWeather()

	locationResponce := &locationWeatherPb.WeatherData{
		LocationId:  weatherData.LocationId,
		Temperature: weatherData.Temperature,
		Humidity:    weatherData.Humidity,
		WindSpeed:   weatherData.WindSpeed,
		Pressure:    weatherData.Pressure,
		Precip:      weatherData.Precip,
		Cloud:       weatherData.Cloud,
		UpdateAt:    weatherData.UpdateAt,
	}

	return locationResponce, nil
}
