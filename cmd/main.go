package main

import (
	"net"
	"reservation-service/config"
	pb "reservation-service/generated/reservation_service"
	"reservation-service/service"
	"reservation-service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.Conn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	config := config.Load()

	listener, err := net.Listen("tcp", ":"+config.GRPC_PORT)
	if err != nil{
		panic(err)
	}
	s := service.NewRRestaurantService(*postgres.NewRRestaurantRepo(db))
	server := grpc.NewServer()
	pb.RegisterReservationServiceServer(server,s)
	if err = server.Serve(listener);err != nil{
		panic(err)
	}
}
