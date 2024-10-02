package services

import (
	"database/sql"

	"github.com/arshedke07/restaurant-app/model"
)

type IRestaurantService interface {
	AddNewRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error)
	MyRestaurantService(id string) (*model.Restaurant, error)
	FindAllRestaurants() (*[]model.Restaurant, error)
}

type RestaurantService struct {
	Config *model.DbConfig
}

func NewRestaurantService(config *model.DbConfig) IRestaurantService {
	return RestaurantService{
		Config: config,
	}
}

func (service RestaurantService) AddNewRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error) {

	insertstatement := "INSERT INTO restaurant(user_id, name, add1, add2, city, state, pincode, picture, description, cuisine) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var id string
	row := db.QueryRow(insertstatement, restaurant.UserId, restaurant.Name, restaurant.Add1, restaurant.Add2, restaurant.City, restaurant.State, restaurant.Pincode, restaurant.Picture, restaurant.Description, restaurant.Cuisine)
	inserterr := row.Scan(&id)
	if inserterr != nil {
		return nil, inserterr
	}
	newrestaurant := model.Restaurant{
		Id:          id,
		UserId:      restaurant.UserId,
		Name:        restaurant.Name,
		Add1:        restaurant.Add1,
		Add2:        restaurant.Add2,
		City:        restaurant.City,
		State:       restaurant.State,
		Pincode:     restaurant.Pincode,
		Picture:     restaurant.Picture,
		Description: restaurant.Description,
		Cuisine:     restaurant.Cuisine,
	}
	return &newrestaurant, nil
}

func (service RestaurantService) MyRestaurantService(id string) (*model.Restaurant, error) {
	selectstatement := "SELECT id, name, add1, add2, city, state, pincode, picture, user_id, description, cuisine FROM restaurant WHERE user_id = $1"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(selectstatement, id)
	user := model.Restaurant{}
	scanErr := row.Scan(&user.Id, &user.Name, &user.Add1, &user.Add2, &user.City, &user.State, &user.Pincode, &user.Picture, &user.UserId, &user.Description, &user.Cuisine)
	if scanErr != nil {
		return nil, scanErr
	}
	return &user, nil
}

func (service RestaurantService) FindAllRestaurants() (*[]model.Restaurant, error) {
	selectstatement := "SELECT id, name, add1, add2, city, state, pincode, picture, user_id, description, cuisine FROM restaurant"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, selectErr := db.Query(selectstatement)
	if selectErr != nil {
		return nil, selectErr
	}

	defer rows.Close()

	restaurants := []model.Restaurant{}
	for rows.Next() {
		restaurant := model.Restaurant{}
		scanerr := rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Add1, &restaurant.Add2, &restaurant.City, &restaurant.State, &restaurant.Pincode, &restaurant.Picture, &restaurant.UserId, &restaurant.Description, &restaurant.Cuisine)
		if scanerr != nil {
			return nil, scanerr
		}
		restaurants = append(restaurants, restaurant)
	}
	return &restaurants, nil
}
