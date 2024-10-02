package routes

import (
	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type IPendingOrderRoute interface {
	GetPendingOrders(c *fiber.Ctx) error
	GetPendingItems(c *fiber.Ctx) error
	DispatchOrder(c *fiber.Ctx) error
}

type PendingOrderRoute struct {
	Store   *session.Store
	Service services.IDeliveryService
}

func NewPendingOrderRoute(store *session.Store, service services.IDeliveryService) IPendingOrderRoute {
	return PendingOrderRoute{
		Store:   store,
		Service: service,
	}
}

func (route PendingOrderRoute) GetPendingOrders(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := c.Params("id")

	data, err := route.Service.PendingOrders(id)
	if err != nil {
		return err
	}

	return c.Render("pendingorders", fiber.Map{
		"Title":    "Pending Orders",
		"UserName": sess.Get("NAME"),
		"Data":     data,
	})
}

func (route PendingOrderRoute) GetPendingItems(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := c.Params("id")

	data, err := route.Service.PendingOrderItems(id)
	if err != nil {
		return err
	}

	return c.Render("orderdetails", fiber.Map{
		"Title":    "Pending Orders",
		"UserName": sess.Get("NAME"),
		"Data":     data,
	})
}

func (route PendingOrderRoute) DispatchOrder(c *fiber.Ctx) error {
	// sess, _ := route.Store.Get(c)
	id := c.Params("id")

	err := route.Service.DispatchOrderService(id)
	if err != nil {
		return err
	}

	return nil
}
