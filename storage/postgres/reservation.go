package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	pb "reservation-service/generated/reservation_service"

	_ "github.com/lib/pq"
)

type RReservationRepo struct {
	pb.UnimplementedReservationServiceServer
	DB *sql.DB
}

func NewRReservationRepo(db *sql.DB) *RReservationRepo {
	return &RReservationRepo{DB: db}
}

func (r *RReservationRepo) CreateReservation(ctx context.Context, req *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	query := `
		INSERT INTO reservations (
			user_id, 
			restaurant_id, 
			reservation_time, 
			status
		)
		VALUES (
			$1, 
			$2, 
			$3, 
			$4
		)
		RETURNING 
			id, 
			user_id, 
			restaurant_id, 
			reservation_time, 
			status;
	`
	reservation := &pb.Reservation{}
	err := r.DB.QueryRowContext(ctx, query, req.UserId, req.RestaurantId, req.ReservationTime, req.Status).Scan(
		&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.ReservationTime, &reservation.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to create reservation: %v", err)
	}
	return &pb.CreateReservationResponse{Reservation: reservation}, nil
}

func (r *RReservationRepo) ListReservations(ctx context.Context, req *pb.ListReservationsRequest) (*pb.ListReservationsResponse, error) {
	query := `
		SELECT 
			id, 
			user_id, 
			restaurant_id, 
			reservation_time, 
			status 
		FROM 
			reservations 
		WHERE deleted_at = 0;`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list reservations: %v", err)
	}
	defer rows.Close()

	var reservations []*pb.Reservation
	for rows.Next() {
		reservation := &pb.Reservation{}
		if err := rows.Scan(&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.ReservationTime, &reservation.Status); err != nil {
			return nil, fmt.Errorf("failed to scan reservations: %v", err)
		}
		reservations = append(reservations, reservation)
	}
	return &pb.ListReservationsResponse{Reservations: reservations}, nil
}

func (r *RReservationRepo) GetReservation(ctx context.Context, req *pb.GetReservationRequest) (*pb.GetReservationResponse, error) {
	query := `
		SELECT 
			id, 
			user_id, 
			restaurant_id, 
			reservation_time, 
			status 
		FROM 
			reservations 
		WHERE id = $1 AND deleted_at = 0;
	`
	reservation := &pb.Reservation{}
	err := r.DB.QueryRowContext(ctx, query, req.Id).Scan(
		&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.ReservationTime, &reservation.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("reservation not found")
		}
		return nil, fmt.Errorf("failed to get reservation: %v", err)
	}
	return &pb.GetReservationResponse{Reservation: reservation}, nil
}

func (r *RReservationRepo) UpdateReservation(ctx context.Context, req *pb.UpdateReservationRequest) (*pb.UpdateReservationResponse, error) {
	query := `
		UPDATE 
			reservations 
		SET 
			user_id = $2, 
			restaurant_id = $3, 
			reservation_time = $4, 
			status = $5, 
			updated_at = CURRENT_TIMESTAMP
		WHERE 
			id = $1 AND deleted_at = 0
		RETURNING 
			id, 
			user_id, 
			restaurant_id, 
			reservation_time, 
			status;
	`
	reservation := &pb.Reservation{}
	err := r.DB.QueryRowContext(ctx, query, req.Id, req.UserId, req.RestaurantId, req.ReservationTime, req.Status).Scan(
		&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.ReservationTime, &reservation.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to update reservation: %v", err)
	}
	return &pb.UpdateReservationResponse{Reservation: reservation}, nil
}

func (r *RReservationRepo) DeleteReservation(ctx context.Context, req *pb.DeleteReservationRequest) (*pb.DeleteReservationResponse, error) {
	query := `
		UPDATE 
			reservations 
		SET 
			deleted_at = EXTRACT(EPOCH FROM CURRENT_TIMESTAMP)
		WHERE 
			id = $1 AND deleted_at = 0;
	`
	
	_, err := r.DB.ExecContext(ctx, query, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("reservation not found")
		}
		return nil, fmt.Errorf("failed to delete reservation: %v", err)
	}
	return &pb.DeleteReservationResponse{Message: "Reservation deleted successfully"}, nil
}
