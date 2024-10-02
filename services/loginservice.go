package services

import (
	"database/sql"
	"errors"

	"github.com/arshedke07/restaurant-app/model"
)

type ILoginService interface {
	Login(emailid string, password string) (*model.AppUser, error)
	FindUserById(string) (*model.AppUser, error)
}

type LoginService struct {
	Config *model.DbConfig
}

func NewLoginService(config *model.DbConfig) ILoginService {
	return LoginService{
		Config: config,
	}
}

func (service LoginService) Login(emailid string, password string) (*model.AppUser, error) {
	selectstatement := "SELECT id, firstname, lastname, password, usertype FROM app_user WHERE emailid=$1 and password=$2"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}

	defer db.Close()
	var firstname string
	var lastname string
	var id int
	var usertype string
	var pwd string

	pass := db.QueryRow(selectstatement, emailid, password)
	scanerr := pass.Scan(&id, &firstname, &lastname, &pwd, &usertype)
	if scanerr != nil {
		return nil, errors.New("username or password is incorrect")
	}

	user := model.AppUser{
		Id:        id,
		FirstName: firstname,
		LastName:  lastname,
		UserType:  usertype,
	}
	return &user, nil
}
func (service LoginService) FindUserById(id string) (*model.AppUser, error) {
	selectstatement := "SELECT firstname, lastname, emailid, gender, active, mobile, usertype, createddate, updateddate FROM app_user WHERE id = $1"
	db, err := sql.Open("postgres", service.Config.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(selectstatement, id)
	user := model.AppUser{}
	scanErr := row.Scan(&user.FirstName, &user.LastName, &user.EmailId, &user.Gender, &user.Status, &user.Mobile, &user.UserType, &user.CreatedDate, &user.UpdatedDate)
	if scanErr != nil {
		return nil, scanErr
	}
	return &user, nil
}
