package main

import (
	"github.com/Vova-luk/weather-stream/services/location-service/internal/config"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/handler"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/logger"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/repository"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/service"
	"github.com/Vova-luk/weather-stream/services/location-service/pkg/db"
	"github.com/Vova-luk/weather-stream/services/location-service/pkg/kafka"

	"google.golang.org/grpc"

	analyticsPb "github.com/Vova-luk/weather-stream/services/analytic-service/proto"
	locationPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location"
	locationAnalyticsPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location_analytics"
	locationWeatherPb "github.com/Vova-luk/weather-stream/services/location-service/proto/location_weather"
	weatherPb "github.com/Vova-luk/weather-stream/services/weather-service/proto"

	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	var err error

	log := logger.InitLogger()

	cfg := config.LoadConfig()

	database, err := db.ConnectPostgre(cfg)
	if err != nil {
		log.Fatalf("Bad Connect to Postgre %s", err.Error())
	}

	kafkaBrokers := cfg.Kafka.Brokers

	producer, err := kafka.NewProducer(kafkaBrokers, cfg.Kafka.LocationTopic, log)
	if err != nil {
		log.Fatalf("Error creating producer %s", err.Error())
	}

	locationRepo := repository.NewLocationRepository(database)
	locationService := service.NewLocationService(locationRepo, producer, log, cfg)
	localHandler := handler.NewLocationHanadler(locationService, log)

	grpcAddrWeather := "weather-service:" + cfg.Server.WeatherServicePort

	weatherConn, err := grpc.Dial(grpcAddrWeather, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to weather-service %s", err.Error())
	}
	log.Infof("Connected to weather-service on the port: %s", cfg.Server.WeatherServicePort)

	weatherClient := weatherPb.NewWeatherServiceClient(weatherConn)

	weatherService := service.NewWeatherService(weatherClient, log)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	grpcAddrAnalytic := "analytic-service:" + cfg.Server.AnalyticServicePort

	analyticsConn, err := grpc.Dial(grpcAddrAnalytic, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to analytic-service %s", err.Error())
	}
	log.Infof("Connected to analytic-service on the port: %s", cfg.Server.AnalyticServicePort)

	analyticClient := analyticsPb.NewAnalyticServiceClient(analyticsConn)

	analyticsService := service.NewAnalyticsService(analyticClient, log)
	analyticHandler := handler.NewAnalyticsHandler(analyticsService)

	grpcServer := grpc.NewServer()
	locationPb.RegisterLocationServiceServer(grpcServer, localHandler)
	locationWeatherPb.RegisterWeatherServiceServer(grpcServer, weatherHandler)
	locationAnalyticsPb.RegisterAnalyticsServiceServer(grpcServer, analyticHandler)

	grpcAddr := ":" + cfg.Server.Port
	gatewayAddr := ":" + cfg.Server.GatewayPort

	go func() {
		listen, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("Failed to listen on port %s", err.Error())
		}

		log.Printf("gRPC server started on port: %s", grpcAddr)
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("Failed to serve gRPC: %s", err)
		}

	}()

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := locationPb.RegisterLocationServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to start REST Gateway from location: %v", err)
	}

	if err := locationWeatherPb.RegisterWeatherServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to start REST Gateway from weather: %v", err)
	}

	if err := locationAnalyticsPb.RegisterAnalyticsServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to start REST Gateway from analytic: %v", err)
	}

	log.Infof("REST Gateway started on %s", gatewayAddr)
	if err := http.ListenAndServe(gatewayAddr, mux); err != nil {
		log.Fatalf("Failed to serve REST Gateway: %v", err)
	}

}
