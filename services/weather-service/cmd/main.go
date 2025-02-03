package main

import (
	"net"

	"github.com/Vova-luk/weather-stream/services/weather-service/internal/config"
	"github.com/Vova-luk/weather-stream/services/weather-service/internal/handler"
	"github.com/Vova-luk/weather-stream/services/weather-service/internal/logger"
	"github.com/Vova-luk/weather-stream/services/weather-service/internal/repository"
	"github.com/Vova-luk/weather-stream/services/weather-service/internal/service"
	"github.com/Vova-luk/weather-stream/services/weather-service/pkg/db"
	"github.com/Vova-luk/weather-stream/services/weather-service/pkg/kafka"
	weatherPb "github.com/Vova-luk/weather-stream/services/weather-service/proto"
	"google.golang.org/grpc"
)

func main() {
	log := logger.InitLogger()

	cfg := config.LoadConfig()

	db, err := db.ConnectPostgre(cfg)
	if err != nil {
		log.Fatalf("Bad Connect to Postgre %s", err.Error())
	}

	consumerGroup, err := kafka.NewConsumer(cfg.Kafka.Brokers, cfg.Kafka.GroupId)
	if err != nil {
		log.Fatalf("error creating consumer %v", err)
	}

	weatherRepo := repository.NewWeatherRepository(db)
	weatherService := service.NewWeatherService(weatherRepo, consumerGroup, log, cfg)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	weatherService.StartKafkaConsumer(cfg.Kafka.LocationTopic)

	grpcServer := grpc.NewServer()
	weatherPb.RegisterWeatherServiceServer(grpcServer, weatherHandler)

	grpcAddr := ":" + cfg.Server.Port

	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen on port %v", err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC: %s", err)
	}
}
