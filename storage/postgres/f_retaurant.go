package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "reservation-service/generated/reservation_service"
)

func (r *RRestaurantRepo) CheckReservation(ctx context.Context, in *pb.CheckReservationRequest) (*pb.CheckReservationResponse, error) {
	query := `SELECT * FROM reservations WHERE id = $1 AND reservation_time < now() at time zone 'UTC'`
	_, err := r.DB.Query(query, in.RestaurantId)
	if err != nil {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return &pb.CheckReservationResponse{Available: false}, nil
	}
	return &pb.CheckReservationResponse{Available: true}, nil
}

func (r *RRestaurantRepo) OrderMeals(ctx context.Context, in *pb.OrderMealsRequest) (*pb.OrderMealsResponse, error) {
	r.DB.QueryRow("select id from menu where name = $1", in.Meals)
	for _, meal := range in.Meals {
		row := r.DB.QueryRow("select id from menu where name = $1", meal.Name)
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, fmt.Errorf("meal with name: '%s'  not found", meal)
		}
		err := row.Scan(&meal.Id)
		if err != nil {
			return nil, err
		}
		_, err = r.DB.Exec(`insert into reservationorders (reservation_id, menu_item_id, quantity)
values ($1, $2, $3)`, in.ReservationId, meal.Id, 1)
		return nil, err
	}
	return &pb.OrderMealsResponse{}, nil
}

func (r *RRestaurantRepo) PayReservation(ctx context.Context, in *pb.PayReservationRequest) (*pb.PayReservationResponse, error) {
	return nil, nil
}
