package handler

import (
	"context"

	"github.com/Vova-luk/weather-stream/services/location-service/internal/service"
	locationAnalyticsPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location_analytics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AnalyticsHandler struct {
	analyticsService *service.AnalyticsService
}

func NewAnalyticsHandler(analyticsService *service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{
		analyticsService: analyticsService,
	}
}

func (a *AnalyticsHandler) GetAnalyticsById(ctx context.Context, request *locationAnalyticsPb.GetAnalyticsByIdRequest) (*locationAnalyticsPb.GetAnalyticsByIdResponse, error) {
	locationId := request.GetLocationId()
	period := request.GetPeriod()

	analyticsData, err := a.analyticsService.GetAnalyticsById(locationId, period)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error receiving analytics by ID: %v", err)
	}

	return &locationAnalyticsPb.GetAnalyticsByIdResponse{
		Analytics: analyticsData,
	}, nil
}
