package services

import (
	"database/sql"

	"github.com/arshedke07/restaurant-app/model"
)

type IAddToCartService interface {
	AddItemToTemp(*model.TempOrder, string) (*model.TempOrder, error)
	FindCartItems(id string, restaurantId string) ([]model.TempOrder, *int, error)
}

type AddToCartService struct {
	Dbconfig *model.DbConfig
}

func NewAddToCartService(dbconfig *model.DbConfig) IAddToCartService {
	return AddToCartService{
		Dbconfig: dbconfig,
	}
}

func (service AddToCartService) AddItemToTemp(order *model.TempOrder, id string) (*model.TempOrder, error) {
	insertstatement := "INSERT INTO order_item_temp(user_id, item, quantity, price, picture, restaurant_id) VALUES($1, $2, $3, $4, $5, $6)"
	db, err := sql.Open("postgres", service.Dbconfig.ConnectionString)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	_, scanErr := db.Exec(insertstatement, order.UserId, order.Item, order.Quantity, order.Price, order.Picture, id)
	if scanErr != nil {
		return nil, scanErr
	}

	newOrder := model.TempOrder{
		UserId:       order.UserId,
		Item:         order.Item,
		Price:        order.Price,
		Quantity:     order.Quantity,
		RestaurantId: id,
	}

	return &newOrder, nil
}

func (service AddToCartService) FindCartItems(id string, restaurantId string) ([]model.TempOrder, *int, error) {
	selectstatement := "SELECT user_id, item, quantity, price, picture, restaurant_id FROM order_item_temp WHERE(user_id = $1 AND restaurant_id = $2)"
	db, err := sql.Open("postgres", service.Dbconfig.ConnectionString)
	if err != nil {
		return nil, nil, err
	}

	rows, _ := db.Query(selectstatement, id, restaurantId)
	carts := []model.TempOrder{}
	var total int

	for rows.Next() {
		cart := model.TempOrder{}
		scanErr := rows.Scan(&cart.UserId, &cart.Item, &cart.Quantity, &cart.Price, &cart.Picture, &cart.RestaurantId)
		if scanErr != nil {
			return nil, nil, scanErr
		}
		total = total + cart.Price
		carts = append(carts, cart)
	}
	return carts, &total, nil
}
