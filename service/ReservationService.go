package service

import (
	"context"
	pb "reservation-service/generated/reservation_service"
	"reservation-service/storage/postgres"
)

type ReservationService struct{
	pb.UnimplementedReservationServiceServer
	Reservation postgres.ReservationRepo
}

func NewRRestaurantService(reservation postgres.ReservationRepo)*ReservationService{
	return &ReservationService{Reservation: reservation}
}

func (r *ReservationService) CreateRestaurant(ctx context.Context,restaurant *pb.CreateRestaurantRequest)(*pb.CreateRestaurantResponse,error){
	res,err := r.Reservation.CreateRestaurant(restaurant)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) ListRestaurants(ctx context.Context, listRestaurant *pb.ListRestaurantsRequest)(*pb.ListRestaurantsResponse,error){
	res,err := r.Reservation.ListRestaurants(listRestaurant)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) GetRestaurant(ctx context.Context, id *pb.GetRestaurantRequest)(*pb.GetRestaurantResponse,error){
	res,err := r.Reservation.GetRestaurant(id)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) UpdateRestaurant(ctx context.Context,updateRestaurant *pb.UpdateRestaurantRequest)(*pb.UpdateRestaurantResponse,error){
	res,err := r.Reservation.UpdateRestaurant(updateRestaurant)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) DeleteRestaurant(ctx context.Context, id *pb.DeleteRestaurantRequest)(*pb.DeleteRestaurantResponse,error){
	res,err := r.Reservation.DeleteRestaurant(id)
	if err != nil{
		return nil,err
	}
	return res,nil
}






func (r *ReservationService) CreateReservation(ctx context.Context,reservation *pb.CreateReservationRequest)(*pb.CreateReservationResponse,error){
	res,err := r.Reservation.CreateReservation(reservation)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) ListReservations(ctx context.Context,listReservation *pb.ListReservationsRequest)(*pb.ListReservationsResponse,error){
	res,err :=r.Reservation.ListReservations(listReservation)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) GetReservation(ctx context.Context,Reservation *pb.GetReservationRequest)(*pb.GetReservationResponse,error){
	res,err := r.Reservation.GetReservation(Reservation)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) UpdateReservation(ctx context.Context, updateReservation *pb.UpdateReservationRequest)(*pb.UpdateReservationResponse,error){
	res,err := r.Reservation.UpdateReservation(updateReservation)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) DeleteReservation(ctx context.Context, id *pb.DeleteReservationRequest)(*pb.DeleteReservationResponse,error){
	res,err := r.Reservation.DeleteReservation(id)
	if err != nil{
		return nil,err
	}
	return res,nil
}







func (r * ReservationService) CreateMenuItem(ctx context.Context, menu *pb.CreateMenuItemRequest)(*pb.CreateMenuItemResponse,error){
	res,err := r.Reservation.CreateMenuItem(menu)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) ListMenuItems(ctx context.Context, listMenu *pb.ListMenuItemsRequest)(*pb.ListMenuItemsResponse,error){
	res,err := r.Reservation.ListMenuItems(listMenu)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService)  GetMenuItem(ctx context.Context,id *pb.GetMenuItemRequest)(*pb.GetMenuItemResponse,error){
	res,err := r.Reservation.GetMenuItem(id)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (r *ReservationService) UpdateMenuItem(ctx context.Context, menu *pb.UpdateMenuItemRequest)(*pb.UpdateMenuItemResponse,error){
	res,err := r.Reservation.UpdateMenuItem(menu)
	if err != nil{
		 return nil,err
	}
	return res,nil
}

func	(r *ReservationService) DeleteMenuItem(ctx context.Context, id *pb.DeleteMenuItemRequest)(*pb.DeleteMenuItemResponse,error){
	res,err := r.Reservation.DeleteMenuItem(id)
	if err != nil{
		return nil,err
	}
	return res,nil
}