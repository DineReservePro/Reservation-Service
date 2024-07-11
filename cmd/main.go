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
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := redis.ConnectR()

	config := config.Load()

	listener, err := net.Listen("tcp", config.GRPC_PORT)
	if err != nil{
		panic(err)
	}
	s := service.NewRRestaurantService(*postgres.NewRRestaurantRepo(db, r))
	server := grpc.NewServer()
	pb.RegisterReservationServiceServer(server,s)
	if err = server.Serve(listener);err != nil{
		panic(err)
	}
}
