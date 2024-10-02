package services

import (
	"database/sql"

	"github.com/arshedke07/restaurant-app/model"
)

type IAddressService interface {
	AddAdress(*model.Address, string) (*model.Address, error)
}

type AddressService struct {
	Config *model.DbConfig
}

func NewAddressService(config *model.DbConfig) IAddressService {
	return AddressService{
		Config: config,
	}
}

func (service AddressService) AddAdress(address *model.Address, userId string) (*model.Address, error) {
	insertstatement := "INSERT INTO user_address(user_id, add1, add2, add3, city, state, pincode) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var id int
	row := db.QueryRow(insertstatement, userId, address.Add1, address.Add2, address.Add3, address.City, address.State, address.Pincode)
	inserterr := row.Scan(&id)
	if inserterr != nil {
		return nil, inserterr
	}
	newaddress := model.Address{
		Id:      id,
		UserId:  userId,
		Add1:    address.Add1,
		Add2:    address.Add2,
		Add3:    address.Add3,
		City:    address.City,
		State:   address.State,
		Pincode: address.Pincode,
	}
	return &newaddress, nil
}
