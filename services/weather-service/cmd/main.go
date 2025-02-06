package main

import (
	"net"

	locationPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location"
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

	kafkaBrokers := cfg.Kafka.Brokers

	consumerGroup, err := kafka.NewConsumer(kafkaBrokers, cfg.Kafka.GroupId)
	if err != nil {
		log.Fatalf("error creating consumer %v", err)
	}

	producer, err := kafka.NewProduces(kafkaBrokers, cfg.Kafka.WeatherTopic, log)
	if err != nil {
		log.Fatalf("error creating producer %v", err)
	}

	grpcAddrLocation := "location-service" + cfg.Server.LocationServicePort

	locationConn, err := grpc.Dial(grpcAddrLocation, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to location-service %s", err.Error())
	}
	log.Printf("Connected to location-service on the port: %d", 50051)

	locationClient := locationPb.NewLocationServiceClient(locationConn)

	weatherRepo := repository.NewWeatherRepository(db)
	weatherService := service.NewWeatherService(weatherRepo, producer, consumerGroup, locationClient, log, cfg)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	go func() {
		weatherService.StartKafkaConsumer(cfg.Kafka.LocationTopic)
	}()

	grpcServer := grpc.NewServer()
	weatherPb.RegisterWeatherServiceServer(grpcServer, weatherHandler)

	grpcAddr := ":" + cfg.Server.Port

	log.Infof("🚀 Starting gRPC server on port %s", grpcAddr)
	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen on port:%s, %v", grpcAddr, err)
	}

	log.Info("gRPC server started")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC: %s", err)
	}
}
