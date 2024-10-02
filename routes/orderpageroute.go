package routes

import (
	"errors"

	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
)

type IOrderRoute interface {
	OrderPageRoute(c *fiber.Ctx) error
}

type OrderRoute struct {
	Service services.IRestaurantService
	Store   *session.Store
}

func NewOrderRoute(service services.IRestaurantService, store *session.Store) IOrderRoute {
	return OrderRoute{
		Service: service,
		Store:   store,
	}
}

func (route OrderRoute) OrderPageRoute(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := sess.Get("ID")

	data, err := route.Service.FindAllRestaurants()
	if err != nil {
		return err
	}

	if id != nil {
		return c.Render("orderpage", fiber.Map{
			"Title":    "The Foodie",
			"UserName": sess.Get("NAME"),
			"Data":     data,
		}, "layout")
	}
	return errors.New("you are Not logged in")
}
