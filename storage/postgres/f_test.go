package postgres

import (
	"context"
	"log"
	pb "reservation-service/generated/reservation_service"
	"testing"
)

//func TestCheckReservation(t *testing.T) {
//	var db, err = Conn()
//	if err != nil {
//		fmt.Println(err)
//		log.Fatalln(err)
//	}
//	fmt.Println(err)
//	defer db.Close()
//
//	reserv := NewRRestaurantRepo(db)
//
//	checkResv := pb.CheckReservationRequest{RestaurantId: "839b5373-3e7f-4dce-8b4d-284b3343e72c"}
//	ch, err := reserv.CheckReservation(context.Background(), &checkResv)
//	if err != nil {
//		fmt.Println(err)
//	}
//	assert.True(t, !ch.Available)
//}

func TestOrderMales(t *testing.T) {
	db, err := Conn()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	reserv := NewRRestaurantRepo(db)

	var id string
	q := db.QueryRow("INSERT INTO menu(restaurant_id, name, price) VALUES ($1, $2, $3) RETURNING id",
		"839b5373-3e7f-4dce-8b4d-284b3343e72c", "hot-dog", 55).Scan(&id)
	if q.Error() != "" {
		log.Fatal(q.Error())
	}

	mr := pb.OrderMealsRequest{
		ReservationId: id,
		Meals: []*pb.MenuItem{
			{
				Name: "hot-dog",
			},
		},
	}

	_, err = reserv.OrderMeals(context.Background(), &mr)
	if err != nil {
		log.Fatal(err)
	}
}
