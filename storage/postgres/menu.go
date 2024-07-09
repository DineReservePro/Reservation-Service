package postgres

import (
	pb "reservation-service/generated/reservation_service"
	"strconv"
	"strings"
)

func (r *ReservationRepo) CreateMenuItem(menuItem *pb.CreateMenuItemRequest) (*pb.CreateMenuItemResponse, error) {

	menu := pb.CreateMenuItemResponse{}
	err := r.DB.QueryRow(`
		INSERT INTO 
		Menu(
			restaurant_id,
			name,
			description,
			price
		)
		VALUES(
			$1,
			$2,
			$3,
			$4
		)
		returning
			id,
			restaurant_id,
			name,
			description,
			price`,
		menuItem.RestaurantId, menuItem.Name, menuItem.Description, menuItem.Price,
	).Scan(menu.MenuItem.Id, menu.MenuItem.RestaurantId, menu.MenuItem.Name, menu.MenuItem.Description, menu.MenuItem.Price)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *ReservationRepo) ListMenuItems(listMenu *pb.ListMenuItemsRequest) (*pb.ListMenuItemsResponse, error) {
	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := `SELECT
					id,
					restaurant_id,
					name,
					description,
					price
				FROM
					Menu
					WHERE
						true `
	if listMenu.RestaurantId != "" {
		params["restaurant_id"] = listMenu.RestaurantId
		filter += "AND restaurant_id = :restaurant_id"
	}

	if listMenu.Name != "" {
		params["name"] = listMenu.Name
		filter += "AND name = :name"
	}

	if listMenu.Price > 0 {
		params["price"] = listMenu.Price
		filter += "AND price = :price"
	}

	query += filter

	query, args = ReplaceQueryParams(query, params)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	ListMenu := pb.ListMenuItemsResponse{}
	for rows.Next() {
		menu := pb.MenuItem{}
		err := rows.Scan(menu.Id, menu.RestaurantId, menu.Name, menu.Description, menu.Price)
		if err != nil {
			return nil, err
		}
		ListMenu.MenuItems = append(ListMenu.MenuItems, &menu)
	}
	return &ListMenu, nil
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		ind  int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(ind))
			args = append(args, v)
			ind++
		}
	}

	return namedQuery, args
}

func (r *ReservationRepo) GetMenuItem(id *pb.GetMenuItemRequest) (*pb.GetMenuItemResponse, error) {
	itemMenu := pb.GetMenuItemResponse{}
	err := r.DB.QueryRow(`	SELECT
								id,
								restaurant_id,
								name,
								description,
								price
							FROM
								Menu
							WHERE
								id = $1`,
		id.Id).
		Scan(
			itemMenu.MenuItem.Id,
			itemMenu.MenuItem.RestaurantId,
			itemMenu.MenuItem.Name,
			itemMenu.MenuItem.Description,
			itemMenu.MenuItem.Price,
		)
	if err != nil {
		return nil, err
	}
	return &itemMenu, nil
}

func (r *ReservationRepo) UpdateMenuItem(updateMenu *pb.UpdateMenuItemRequest) (*pb.UpdateMenuItemResponse, error) {
	menu := pb.UpdateMenuItemResponse{}
	err := r.DB.QueryRow(`	
						UPDATE 
						MENU
					SET
						restaurant_id = $1,
						name = $2,
						description = $3,
						price = $4
					WHERE
						id = $5
					returning
						id,
						restaurant_id,
						name,
						description,
						price
					`, updateMenu.RestaurantId, updateMenu.Name, updateMenu.Description, updateMenu.Price, updateMenu.Id).
		Scan(menu.MenuItem.Id, menu.MenuItem.RestaurantId, menu.MenuItem.Name, menu.MenuItem.Description, menu.MenuItem.Price)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *ReservationRepo) DeleteMenuItem(id *pb.DeleteMenuItemRequest) (*pb.DeleteMenuItemResponse, error) {
	_,err := r.DB.Exec(`	DELETE
				FROM
					Menu
				WHERE
					id = $1`,id.Id)
	if err != nil{
		return &pb.DeleteMenuItemResponse{
			Message: "FAILD TO DELETED MENU ITEM",
		},err
	}
	return &pb.DeleteMenuItemResponse{
		Message: "DELETED SUCCESFULLY MENU ITEM",
	},nil
}
