package services

import (
	"database/sql"

	"github.com/arshedke07/restaurant-app/model"
)

type IUserService interface {
	AddNewUser(user *model.AppUser) (*model.AppUser, error)
	// UpdateUser(user *model.User) (*model.User, error)
	// Delete(id int) error
}

type UserService struct {
	Config *model.DbConfig
}

func NewUserService(config *model.DbConfig) IUserService {
	return UserService{
		Config: config,
	}
}

func (service UserService) AddNewUser(user *model.AppUser) (*model.AppUser, error) {
	insertstatement := "INSERT INTO app_user(firstname, lastname, emailid, gender, active, mobile, password, usertype, createddate, updateddate) VALUES($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var id int
	row := db.QueryRow(insertstatement, user.FirstName, user.LastName, user.EmailId, user.Gender, user.Status, user.Mobile, user.Password, user.UserType)
	inserterr := row.Scan(&id)
	if inserterr != nil {
		return nil, inserterr
	}

	newuser := model.AppUser{
		Id:          id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		EmailId:     user.EmailId,
		Gender:      user.Gender,
		Status:      user.Status,
		Mobile:      user.Mobile,
		Password:    user.Password,
		UserType:    user.UserType,
		CreatedDate: user.CreatedDate,
		UpdatedDate: user.UpdatedDate,
	}
	return &newuser, nil
}

// func (service UserService) UpdateUser(user *model.User) (*model.User, error) {

// }

// func (service UserService) Delete(id int) error {

// }
