package services

import (
	"database/sql"

	"github.com/arshedke07/restaurant-app/model"
)

type IDeliveryService interface {
	PendingOrders(string) ([]model.Order, error)
	PendingOrderItems(string) (*[]model.OrderItem, error)
	DispatchOrderService(string) error
}

type DeliveryService struct {
	Config *model.DbConfig
}

func NewDeliveryService(config *model.DbConfig) IDeliveryService {
	return DeliveryService{
		Config: config,
	}
}

func (service DeliveryService) PendingOrders(id string) ([]model.Order, error) {
	selectstatement := "SELECT id, user_id, tax, total, grand_total, status, restaurant_id, address FROM user_order WHERE (status=$1 AND restaurant_id=$2)"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row, rowErr := db.Query(selectstatement, 0, id)
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

	return orders, nil
}

func (service DeliveryService) PendingOrderItems(id string) (*[]model.OrderItem, error) {
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

func (service DeliveryService) DispatchOrderService(id string) error {
	updatestatement := "UPDATE user_order SET status = $1 WHERE id = $2"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	_, updateErr := db.Exec(updatestatement, 1, id)
	if updateErr != nil {
		return updateErr
	}
	return nil
}
