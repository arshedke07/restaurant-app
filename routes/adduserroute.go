package routes

import (
	"github.com/arshedke07/restaurant-app/model"
	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
)

type IAddUserRoute interface {
	AddUser(c *fiber.Ctx) error
}

type AddUserRoute struct {
	Service services.IUserService
	Store   *session.Store
}

func NewUserRoute(service services.IUserService, store *session.Store) IAddUserRoute {
	return AddUserRoute{
		Service: service,
		Store:   store,
	}
}

func (route AddUserRoute) AddUser(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("adduser", fiber.Map{
			"Title": "Register User",
		}, "layout")
	} else if c.Method() == "POST" {
		sess, err := route.Store.Get(c)
		if err != nil {
			return err
		}
		user := model.AppUser{
			FirstName: c.FormValue("firstname"),
			LastName:  c.FormValue("lastname"),
			EmailId:   c.FormValue("emailid"),
			Gender:    c.FormValue("gender"),
			Status:    "a",
			Mobile:    c.FormValue("mobile"),
			UserType:  c.FormValue("usertype"),
			Password:  c.FormValue("password"),
		}

		data, err := route.Service.AddNewUser(&user)
		if err != nil {
			return err
		}

		sess.Set("ID", data.Id)
		sess.Set("NAME", data.FirstName+" "+data.LastName)
		sess.Set("EMAILID", data.EmailId)
		sess.Set("USERTYPE", data.UserType)
		sess.Save()

		return c.Render("addaddress", fiber.Map{
			"Title": "Register User",
			"Data":  data,
		}, "layout")
	}
	return nil
}
