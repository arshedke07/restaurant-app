package services

import (
	"database/sql"

	"github.com/arshedke07/restaurant-app/model"
)

type IOrderService interface {
	OrderItemService(string, string, string) (*[]model.OrderItem, error)
	UserOrderService(string, string, string) (*model.Order, error)
}

type OrderService struct {
	DbConfig    *model.DbConfig
	CartService IAddToCartService
}

func NewOrderService(dbconfig *model.DbConfig, cart IAddToCartService) IOrderService {
	return OrderService{
		DbConfig:    dbconfig,
		CartService: cart,
	}
}

func (service OrderService) OrderItemService(id string, orderId string, restaurantId string) (*[]model.OrderItem, error) {
	data, _, err := service.CartService.FindCartItems(id, restaurantId)
	if err != nil {
		return nil, err
	}

	insertstatement := "INSERT INTO order_item(order_id, item, quantity, price, picture) VALUES($1, $2, $3, $4, $5) RETURNING id"
	db, dberr := sql.Open("postgres", service.DbConfig.ConnectionString)
	if dberr != nil {
		return nil, dberr
	}

	defer db.Close()
	var itemId int

	orders := []model.OrderItem{}
	for i := 0; i < len(data); i++ {
		row := db.QueryRow(insertstatement, orderId, data[i].Item, data[i].Quantity, data[i].Price, data[i].Picture)
		scanErr := row.Scan(&itemId)
		if scanErr != nil {
			return nil, scanErr
		}
		order := model.OrderItem{
			Id:       itemId,
			OrderId:  orderId,
			Item:     data[i].Item,
			Quantity: data[i].Quantity,
			Price:    data[i].Price,
			Picture:  data[i].Picture,
		}
		orders = append(orders, order)
	}

	deletestatment := "DELETE FROM order_item_temp"
	_, newErr := db.Exec(deletestatment)
	if newErr != nil {
		return nil, newErr
	}

	return &orders, nil
}

func (service OrderService) UserOrderService(userId string, restaurantId string, address string) (*model.Order, error) {
	_, price, err := service.CartService.FindCartItems(userId, restaurantId)
	if err != nil {
		return nil, err
	}

	insertstatement := "INSERT INTO user_order(user_id, tax, total, grand_total, status, restaurant_id, address) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	db, dbErr := sql.Open("postgres", service.DbConfig.ConnectionString)
	if dbErr != nil {
		return nil, dbErr
	}

	defer db.Close()
	var id string

	newprice := float32(*price)
	tax := float32(5.0 / 100.0 * newprice)
	total := newprice + tax

	row := db.QueryRow(insertstatement, userId, tax, price, total, 0, restaurantId, address)
	scanerr := row.Scan(&id)
	if scanerr != nil {
		return nil, scanerr
	}

	neworder := model.Order{
		Id:           id,
		UserId:       userId,
		Tax:          tax,
		Total:        *price,
		GrandTotal:   total,
		Status:       0,
		RestaurantId: restaurantId,
		Address:      address,
	}

	return &neworder, nil
}
