package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "reservation-service/generated/reservation_service"

	_ "github.com/lib/pq"
)

type RRestaurantRepo struct {
	pb.UnimplementedReservationServiceServer
	DB *sql.DB
}

func NewRRestaurantRepo(db *sql.DB) *RRestaurantRepo {
	return &RRestaurantRepo{DB: db}
}

func (r *RRestaurantRepo) CreateRestaurant(ctx context.Context, req *pb.CreateRestaurantRequest) (*pb.CreateRestaurantResponse, error) {
	query := `
		INSERT INTO Restaurants (
			name, 
			address, 
			phone_number, 
			description
		)
		VALUES (
			$1, 
			$2, 
			$3, 
			$4
		)
		RETURNING 
			id, 
			name, 
			address, 
			phone_number, 
			description;
	`
	restaurant := &pb.Restaurant{}

	err := r.DB.QueryRowContext(ctx, query, req.Name, req.Address, req.PhoneNumber, req.Description).Scan(
		&restaurant.Id, &restaurant.Name, &restaurant.Address, &restaurant.PhoneNumber, &restaurant.Description)

	if err != nil {
		return nil, fmt.Errorf("failed to create restaurant: %v", err)
	}
	return &pb.CreateRestaurantResponse{Restaurant: restaurant}, nil
}

func (r *RRestaurantRepo) ListRestaurants(req *pb.ListRestaurantsRequest) (*pb.ListRestaurantsResponse, error) {
	var (
		params = make(map[string]interface{})
		args []interface{}
		filter string
	)

	query := `
		SELECT 
			id, 
			name, 
			address, 
			phone_number, 
			description 
		FROM 
			Restaurants 
		WHERE 
			deleted_at = 0 
	`

	if req.Name != ""{
		params["name"] = req.Name
		filter += " AND name = :name "
	}
	if req.Address != ""{
		params["address"] = req.Address
		filter += " AND address = :address "
	}
	query += filter

	query,args = ReplaceQueryParams(query,params)
	rows, err := r.DB.Query(query,args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list restaurants: %v", err)
	}
	defer rows.Close()

	var restaurants []*pb.Restaurant
	for rows.Next() {
		restaurant := &pb.Restaurant{}
		if err := rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Address, &restaurant.PhoneNumber, &restaurant.Description); err != nil {
			return nil, fmt.Errorf("failed to scan restaurant: %v", err)
		}
		restaurants = append(restaurants, restaurant)
	}
	return &pb.ListRestaurantsResponse{Restaurants: restaurants}, nil
}

func (r *RRestaurantRepo) GetRestaurant(ctx context.Context, req *pb.GetRestaurantRequest) (*pb.GetRestaurantResponse, error) {
	query := `
		SELECT 
			id, 
			name, 
			address, 
			phone_number, 
			description 
		FROM 
			Restaurants 
		WHERE 
			id = $1 AND deleted_at = 0;
	`
	restaurant := &pb.Restaurant{}
	err := r.DB.QueryRow(query, req.Id).Scan(
		&restaurant.Id, &restaurant.Name, &restaurant.Address, &restaurant.PhoneNumber, &restaurant.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("restaurant not found")
		}
		return nil, fmt.Errorf("failed to get restaurant: %v", err)
	}
	return &pb.GetRestaurantResponse{Restaurant: restaurant}, nil
}

func (r *RRestaurantRepo) UpdateRestaurant(ctx context.Context, req *pb.UpdateRestaurantRequest) (*pb.UpdateRestaurantResponse, error) {
	query := `
		UPDATE 
			Restaurants 
		SET 
			name = $2, 
			address = $3, 
			phone_number = $4, 
			description = $5, 
			updated_at = CURRENT_TIMESTAMP
		WHERE 
			id = $1 AND deleted_at = 0
		RETURNING 
			id, 
			name, 
			address, 
			phone_number, 
			description;
	`
	restaurant := &pb.Restaurant{}
	err := r.DB.QueryRow(query, req.Id, req.Name, req.Address, req.PhoneNumber, req.Description).Scan(
		&restaurant.Id, &restaurant.Name, &restaurant.Address, &restaurant.PhoneNumber, &restaurant.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to update restaurant: %v", err)
	}
	return &pb.UpdateRestaurantResponse{Restaurant: restaurant}, nil
}

func (r *RRestaurantRepo) DeleteRestaurant(ctx context.Context, req *pb.DeleteRestaurantRequest) (*pb.DeleteRestaurantResponse, error) {
	query := `
		UPDATE 
			Restaurants 
		SET 
			deleted_at = EXTRACT(EPOCH FROM CURRENT_TIMESTAMP)
		WHERE 
			id = $1 AND deleted_at = 0
	`
	
	_, err := r.DB.Exec(query, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("restaurant not found")
		}
		return nil, fmt.Errorf("failed to delete restaurant: %v", err)
	}
	return &pb.DeleteRestaurantResponse{Message: "Restaurant deleted successfully"}, nil

}
