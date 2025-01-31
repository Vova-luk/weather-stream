package main

import (
	"github.com/Vova-luk/weather-stream/services/location-service/internal/config"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/handler"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/logger"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/repository"
	"github.com/Vova-luk/weather-stream/services/location-service/internal/service"
	"github.com/Vova-luk/weather-stream/services/location-service/pkg/db"

	"google.golang.org/grpc"

	pb "github.com/Vova-luk/weather-stream/services/location-service/proto"

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

	locationRepo := repository.NewLocationRepository(database)
	locationService := service.NewLocationService(locationRepo, log)
	localHandler := handler.NewLocationHanadler(locationService, log)

	grpcServer := grpc.NewServer()
	pb.RegisterLocationServiceServer(grpcServer, localHandler)

	grpcAddr := ":" + cfg.Server.Port
	gatewayAddr := ":" + cfg.Server.GatewayPort

	go func() {
		listen, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("Failed to listen on port %s", err.Error())
		}
		log.Info("gRPC server started")
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("Failed to serve gRPC: %s", err)
		}
	}()

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterLocationServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to start REST Gateway: %v", err)
	}
	log.Infof("REST Gateway started on %s", gatewayAddr)
	if err := http.ListenAndServe(gatewayAddr, mux); err != nil {
		log.Fatalf("Failed to serve REST Gateway: %v", err)
	}

}
