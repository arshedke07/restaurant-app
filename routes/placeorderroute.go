package routes

import (
	"fmt"

	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
)

type IPlaceOrderRoute interface {
	PlaceOrder(c *fiber.Ctx) error
}

type PlaceOrderRoute struct {
	Service services.IOrderService
	Store   *session.Store
}

func NewPlaceOrderRoute(service services.IOrderService, store *session.Store) IPlaceOrderRoute {
	return PlaceOrderRoute{
		Service: service,
		Store:   store,
	}
}

func (route PlaceOrderRoute) PlaceOrder(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	userId := sess.Get("ID")
	str := fmt.Sprintf("%v", userId)
	id := c.Params("id")

	address := c.FormValue("address")
	fmt.Println(address)

	user, err := route.Service.UserOrderService(str, id, address)
	if err != nil {
		return err
	}

	_, err2 := route.Service.OrderItemService(str, user.Id, id)
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	price := user.Total
	tax := user.Tax
	total := user.GrandTotal

	return c.Render("paymentpage", fiber.Map{
		"Title":    "Place Order",
		"UserName": sess.Get("NAME"),
		"Price":    price,
		"Tax":      tax,
		"Total":    total,
		"Address":  address,
	}, "layout")
}
