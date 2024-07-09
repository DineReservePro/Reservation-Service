package postgres

import (
	pb "reservation-service/generated/reservation_service"
	"testing"
)

func TestRestaurant(t *testing.T) {
	db, err := Conn()
	if err != nil {
		t.Errorf("Faieled database connection")
	}
	restaurantRepo := NewRRestaurantRepo(db)
	restaurant := pb.CreateRestaurantRequest{
		Name:        "S",
		Address:     "Chilonzor",
		PhoneNumber: "991234567",
		Description: "jwlejf",
	}
	restaurantRepo.CreateRestaurant(&restaurant)
}
