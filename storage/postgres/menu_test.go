package postgres

import (
	pb "reservation-service/generated/reservation_service"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateMenuItem(t *testing.T) {
	db, err := Conn()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	menuRepo := NewRRestaurantRepo(db)

	menu := pb.CreateMenuItemRequest{
		RestaurantId: "",
		Name:         "",
		Description:  "",
		Price:        0,
	}
	res, err := menuRepo.CreateMenuItem(&menu)
	if err != nil {
		t.Errorf("failed to created menuItem : %v", err)
		return
	}

	if menu.RestaurantId == res.MenuItem.RestaurantId && menu.Name == res.MenuItem.Name &&
		menu.Description == res.MenuItem.Description && menu.Price == res.MenuItem.Price {
		t.Errorf("created mismatch (-want +got):\n%s", cmp.Diff(&menu, res))
		return
	}
}
