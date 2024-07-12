package main

import (
	"net"
	"reservation-service/config"
	pb "reservation-service/generated/reservation_service"
	"reservation-service/service"
	"reservation-service/storage/postgres"
	"reservation-service/storage/redis"

	"google.golang.org/grpc"
)

func main() {
	config.InitLogger()
	logger := config.Logger
	logger.Info("Starting the application...")

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := redis.ConnectR()
	cfg := config.Load()
	listener, err := net.Listen("tcp", cfg.GRPC_PORT)
	if err != nil {
		panic(err)
	}
	s := service.NewRRestaurantService(*postgres.NewRRestaurantRepo(db, r))
	server := grpc.NewServer()
	pb.RegisterReservationServiceServer(server, s)
	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}
