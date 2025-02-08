package service

import (
	"context"

	analyticsPb "github.com/Vova-luk/weather-stream/services/analytic-service/proto"
	locationAnalyticsPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location_analytics"
	"github.com/sirupsen/logrus"
)

type AnalyticsService struct {
	analyticsClient analyticsPb.AnalyticServiceClient
	log             *logrus.Logger
}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{}
}

func (a *AnalyticsService) GetAnalyticsById(locationId int32, period string) (*locationAnalyticsPb.AnalyticsWeather, error) {
	request := &analyticsPb.GetAnalyticsByIdRequest{
		LocationId: locationId,
		Period:     period,
	}
	response, err := a.analyticsClient.GetAnalyticsById(context.Background(), request)
	if err != nil {
		a.log.Warnf("error getting analytics by ID %d: %v", locationId, err)
		return nil, err
	}

	analyticsData := response.GetAnalytics()

	analyticsResponse := &locationAnalyticsPb.AnalyticsWeather{
		LocationId:     analyticsData.LocationId,
		AvgTemperature: analyticsData.AvgTemperature,
		AvgHumidity:    analyticsData.AvgHumidity,
		AvgWindSpeed:   analyticsData.AvgWindSpeed,
		AvgPressure:    analyticsData.AvgPressure,
		AvgPrecip:      analyticsData.AvgPrecip,
		AvgCloud:       analyticsData.AvgCloud,
	}

	return analyticsResponse, nil
}
