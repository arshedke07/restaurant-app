package routes

import (
	"fmt"

	"github.com/arshedke07/restaurant-app/model"
	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
)

type IAddUserAddressRoute interface {
	AddUserAddress(c *fiber.Ctx) error
}

type AddAddressRoute struct {
	Service     services.IAddressService
	FindService services.ILoginService
	Store       *session.Store
}

func NewAddressRoute(service services.IAddressService, store *session.Store, findService services.ILoginService) IAddUserAddressRoute {
	return AddAddressRoute{
		Service:     service,
		Store:       store,
		FindService: findService,
	}
}

func (route AddAddressRoute) AddUserAddress(c *fiber.Ctx) error {
	// id := c.Params("id")
	sess, _ := route.Store.Get(c)
	id := sess.Get("ID")
	str := fmt.Sprintf("%v", id)
	if c.Method() == "GET" {
		data, err := route.FindService.FindUserById(str)
		if err != nil {
			return err
		}
		return c.Render("addaddress", fiber.Map{
			"Title":    "Register User",
			"UserName": sess.Get("NAME"),
			"Data":     data,
		}, "layout")
	} else if c.Method() == "POST" {
		address := model.Address{
			Add1:    c.FormValue("add1"),
			Add2:    c.FormValue("add2"),
			Add3:    c.FormValue("add3"),
			City:    c.FormValue("city"),
			State:   c.FormValue("state"),
			Pincode: c.FormValue("pincode"),
		}
		_, err := route.Service.AddAdress(&address, str)
		if err != nil {
			return err
		}

		return c.Render("home", fiber.Map{
			"Title":    "The Foodie",
			"UserName": sess.Get("NAME"),
		}, "layout")
	}
	return nil
}
