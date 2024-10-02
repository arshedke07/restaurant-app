package services

import (
	"database/sql"

	"github.com/arshedke07/restaurant-app/model"
)

type IProfileService interface {
	AddressService(id string) (*[]model.Address, error)
}

type ProfileService struct {
	Config *model.DbConfig
}

func NewProfileService(config *model.DbConfig) IProfileService {
	return ProfileService{
		Config: config,
	}
}

func (service ProfileService) AddressService(id string) (*[]model.Address, error) {
	selectstatement := "SELECT id, user_id, add1, add2, add3, city, state, pincode FROM user_address WHERE user_id=$1"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, rowErr := db.Query(selectstatement, id)
	if rowErr != nil {
		return nil, rowErr
	}

	users := []model.Address{}
	for rows.Next() {
		user := model.Address{}
		scanErr := rows.Scan(&user.Id, &user.UserId, &user.Add1, &user.Add2, &user.Add3, &user.City, &user.State, &user.Pincode)
		if scanErr != nil {
			return nil, scanErr
		}

		users = append(users, user)
	}

	return &users, nil
}
