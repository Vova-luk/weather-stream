package main

import (
	"net"

	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/config"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/handler"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/logger"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/repository"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/service"
	"github.com/Vova-luk/weather-stream/services/analytic-service/pkg/db"
	analyticsPb "github.com/Vova-luk/weather-stream/services/analytic-service/proto"
	"google.golang.org/grpc"
)

func main() {
	log := logger.InitLogger()

	cfg := config.LoadConfig()

	database, err := db.ConnectPostgre(cfg)
	if err != nil {
		log.Fatalf("Bad Connect to Postgre %s", err.Error())
	}

	analyticsRepository := repository.NewAnalyticsRepository(database)
	analyticsService := service.NewAnalyticsService(analyticsRepository, log)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsService)

	grpcServer := grpc.NewServer()

	analyticsPb.RegisterAnalyticServiceServer(grpcServer, analyticsHandler)

	grpcAddr := cfg.Server.Port

	log.Infof("Starting gRPC server on port %s", grpcAddr)
	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen on port:%s, %v", grpcAddr, err)
	}

	log.Infof("gRPC server started")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC: %s", err)
	}
}
