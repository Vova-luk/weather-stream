package main

import (
	"net"

	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/config"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/handler"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/logger"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/repository"
	"github.com/Vova-luk/weather-stream/services/analytic-service/internal/service"
	"github.com/Vova-luk/weather-stream/services/analytic-service/pkg/db"
	"github.com/Vova-luk/weather-stream/services/analytic-service/pkg/kafka"
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

	brokers := cfg.Kafka.Brokers
	group_id := cfg.Kafka.GroupId

	consumerGroup, err := kafka.NewConsumerGroup(brokers, group_id)
	if err != nil {
		log.Fatalf("error when creating consumer group %v", err)
	}

	analyticsRepository := repository.NewAnalyticsRepository(database)
	analyticsService := service.NewAnalyticsService(analyticsRepository, consumerGroup, log)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsService)

	go func() {
		analyticsService.StartKafkaConsumer(cfg.Kafka.Topic)
	}()

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
