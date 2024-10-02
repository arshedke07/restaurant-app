package services

import (
	"database/sql"

	"github.com/arshedke07/restaurant-app/model"
)

type IOrderHistoryService interface {
	OrderHistory(id string) (*[]model.Order, error)
	OrderItemsHistory(id string) (*[]model.OrderItem, error)
}

type OrderHistoryService struct {
	Config *model.DbConfig
}

func NewOrderHistory(config *model.DbConfig) IOrderHistoryService {
	return OrderHistoryService{
		Config: config,
	}
}

func (service OrderHistoryService) OrderHistory(id string) (*[]model.Order, error) {
	selectstatement := "SELECT id, user_id, tax, total, grand_total, status, restaurant_id, address FROM user_order WHERE (status=$1 AND restaurant_id=$2)"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row, rowErr := db.Query(selectstatement, 1, id)
	if rowErr != nil {
		return nil, rowErr
	}
	defer row.Close()

	orders := []model.Order{}
	for row.Next() {
		order := model.Order{}
		scanErr := row.Scan(&order.Id, &order.UserId, &order.Tax, &order.Total, &order.GrandTotal, &order.Status, &order.RestaurantId, &order.Address)
		if scanErr != nil {
			return nil, scanErr
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func (service OrderHistoryService) OrderItemsHistory(id string) (*[]model.OrderItem, error) {
	selectstatement := "SELECT id, order_id, item, quantity, price, picture FROM order_item WHERE order_id=$1"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row, rowErr := db.Query(selectstatement, id)
	if rowErr != nil {
		return nil, rowErr
	}

	items := []model.OrderItem{}
	for row.Next() {
		item := model.OrderItem{}
		scanErr := row.Scan(&item.Id, &item.OrderId, &item.Item, &item.Quantity, &item.Price, &item.Picture)
		if scanErr != nil {
			return nil, scanErr
		}
		items = append(items, item)
	}
	return &items, nil
}
