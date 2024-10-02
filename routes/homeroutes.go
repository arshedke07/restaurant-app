package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type IHomeRoute interface {
	Home(c *fiber.Ctx) error
}

type HomeRoute struct {
	Store *session.Store
}

func NewHomeRoute(store *session.Store) IHomeRoute {
	return HomeRoute{
		Store: store,
	}
}

func (route HomeRoute) Home(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)

	return c.Render("home", fiber.Map{
		"Title":    "The Foodie",
		"UserName": sess.Get("NAME"),
	}, "layout")
}
