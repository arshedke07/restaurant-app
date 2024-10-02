package services

import (
	"database/sql"

	"github.com/arshedke07/restaurant-app/model"
)

type IAddItemService interface {
	AddNewItem(item *model.Items) (*model.Items, error)
	FindAllItemsById(id string) (*[]model.Items, error)
}

type AddItemService struct {
	Config *model.DbConfig
}

func NewItemService(config *model.DbConfig) IAddItemService {
	return AddItemService{
		Config: config,
	}
}

func (service AddItemService) AddNewItem(item *model.Items) (*model.Items, error) {
	insertstatement := "INSERT INTO item(restaurant_id, item_name, price, quantity, description, picture) VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var id string

	row := db.QueryRow(insertstatement, item.RestaurantId, item.ItemName, item.Price, item.Quantity, item.Description, item.Picture)
	scanErr := row.Scan(&id)
	if scanErr != nil {
		return nil, scanErr
	}

	newitem := model.Items{
		Id:           id,
		RestaurantId: item.RestaurantId,
		ItemName:     item.ItemName,
		Price:        item.Price,
		Quantity:     item.Quantity,
		Description:  item.Description,
		Picture:      item.Picture,
	}

	return &newitem, nil
}

func (service AddItemService) FindAllItemsById(id string) (*[]model.Items, error) {
	selectstatement := "SELECT id, restaurant_id, item_name, price, quantity, description, picture FROM item WHERE restaurant_id = $1"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, selectErr := db.Query(selectstatement, id)
	if selectErr != nil {
		return nil, selectErr
	}

	defer rows.Close()

	items := []model.Items{}
	for rows.Next() {
		item := model.Items{}
		scanerr := rows.Scan(&item.Id, &item.RestaurantId, &item.ItemName, &item.Price, &item.Quantity, &item.Description, &item.Picture)
		if scanerr != nil {
			return nil, scanerr
		}
		items = append(items, item)
	}
	return &items, nil
}
