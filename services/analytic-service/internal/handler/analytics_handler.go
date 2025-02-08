package handler

import (
	"context"

	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/service"
	analyticsPb "github.com/Vova-luk/weather-stream/services/analytic-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AnalyticsHandler struct {
	analyticsPb.UnimplementedAnalyticServiceServer
	analyticsService *service.AnalyticsService
}

func NewAnalyticsHandler(analyticsService *service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{
		analyticsService: analyticsService,
	}
}

func (a *AnalyticsHandler) GetAnalyticsById(ctx context.Context, request *analyticsPb.GetAnalyticsByIdRequest) (*analyticsPb.GetAnalyticsByIdResponse, error) {
	locationId := request.GetLocationId()
	period := request.GetPeriod()

	weatherAnalytic, err := a.analyticsService.GetAnalyticsById(locationId, period)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error when receiving analytics: %v", err)
	}

	return &analyticsPb.GetAnalyticsByIdResponse{
		Analytics: weatherAnalytic.ToProto(),
	}, nil
}
