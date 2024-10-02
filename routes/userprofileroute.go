package routes

import (
	"fmt"

	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type IProfileRoutes interface {
	GetProfile(c *fiber.Ctx) error
	GetAddress(c *fiber.Ctx) error
}

type ProfileRoute struct {
	Store   *session.Store
	Service services.IProfileService
}

func NewProfileRoute(store *session.Store, service services.IProfileService) IProfileRoutes {
	return ProfileRoute{
		Store:   store,
		Service: service,
	}
}

func (route ProfileRoute) GetProfile(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)

	return c.Render("userprofile", fiber.Map{
		"Title":    "My Profile",
		"UserName": sess.Get("NAME"),
		"Name":     sess.Get("NAME"),
	}, "layout")
}

func (route ProfileRoute) GetAddress(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := sess.Get("ID")
	str := fmt.Sprintf("%v", id)

	data, err := route.Service.AddressService(str)
	if err != nil {
		return err
	}

	return c.Render("myaddresses", fiber.Map{
		"Title":    "My Addresses",
		"UserName": sess.Get("NAME"),
		"Data":     data,
	}, "layout")
}
