package routes

import (
	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type IOrderHistoryRoute interface {
	MyOrderHistory(c *fiber.Ctx) error
	MyOrderItemHistory(c *fiber.Ctx) error
}

type OrderHistoryRoute struct {
	Store   *session.Store
	Service services.IOrderHistoryService
}

func NewOrderHistoryRoute(store *session.Store, service services.IOrderHistoryService) IOrderHistoryRoute {
	return OrderHistoryRoute{
		Store:   store,
		Service: service,
	}
}

func (route OrderHistoryRoute) MyOrderHistory(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := c.Params("id")

	data, err := route.Service.OrderHistory(id)
	if err != nil {
		return err
	}

	return c.Render("myorders", fiber.Map{
		"Title":    "Delivered Orders",
		"UserName": sess.Get("NAME"),
		"Data":     data,
	})
}

func (route OrderHistoryRoute) MyOrderItemHistory(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := c.Params("id")

	data, err := route.Service.OrderItemsHistory(id)
	if err != nil {
		return err
	}

	return c.Render("myorderdetails", fiber.Map{
		"Title":    "Delivered Orders",
		"UserName": sess.Get("NAME"),
		"Data":     data,
	})
}
