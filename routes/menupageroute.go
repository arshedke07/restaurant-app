package routes

import (
	"fmt"
	"strconv"

	"github.com/arshedke07/restaurant-app/model"
	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
)

type IMenuPageRoute interface {
	MenuPageRoute(c *fiber.Ctx) error
	MyCartRoute(c *fiber.Ctx) error
}

type MenuPageRoute struct {
	Store          *session.Store
	Service        services.IAddItemService
	DbService      services.IAddToCartService
	ProfileService services.IProfileService
}

func NewMenuRoute(service services.IAddItemService, store *session.Store, dbservice services.IAddToCartService, profileservice services.IProfileService) IMenuPageRoute {
	return MenuPageRoute{
		Store:          store,
		Service:        service,
		DbService:      dbservice,
		ProfileService: profileservice,
	}
}

func (route MenuPageRoute) MenuPageRoute(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := c.Params("id")
	data, err := route.Service.FindAllItemsById(id)
	if err != nil {
		return err
	}
	if c.Method() == "GET" {
		return c.Render("ordermenupage", fiber.Map{
			"Title":        "Order Online",
			"UserName":     sess.Get("NAME"),
			"Data":         data,
			"RestaurantId": id,
		})
	} else if c.Method() == "POST" {
		userId := sess.Get("ID")
		str := fmt.Sprintf("%v", userId)

		price, _ := strconv.Atoi(c.FormValue("price"))
		order := model.TempOrder{
			UserId:   str,
			Item:     c.FormValue("itemname"),
			Quantity: c.FormValue("quantity"),
			Price:    price,
			Picture:  c.FormValue("picture"),
		}

		Id := c.FormValue("restaurantid")

		_, err := route.DbService.AddItemToTemp(&order, Id)
		if err != nil {
			return err
		}
		return c.Render("ordermenupage", fiber.Map{
			"Title":        "Order Online",
			"UserName":     sess.Get("NAME"),
			"Data":         data,
			"RestaurantId": id,
		})
	}
	return nil
}

func (route MenuPageRoute) MyCartRoute(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := sess.Get("ID")
	str := fmt.Sprintf("%v", id)
	restaurantId := c.Params("id")

	data, price, err := route.DbService.FindCartItems(str, restaurantId)
	if err != nil {
		return err
	}

	address, err := route.ProfileService.AddressService(str)
	if err != nil {
		return err
	}

	// userAddress := address.Add1 + " " + address.Add2 + " " + address.Add3 + " " + address.City + " " + address.State + " " + address.Pincode

	newprice := float32(*price)
	tax := float32(5.0 / 100.0 * newprice)
	total := newprice + tax

	sess.Set("Price", price)
	sess.Set("Tax", tax)
	sess.Set("Total", total)

	return c.Render("mycart", fiber.Map{
		"Title":        "My Cart",
		"UserName":     sess.Get("NAME"),
		"Data":         data,
		"OrderTotal":   price,
		"Tax":          tax,
		"GrandTotal":   total,
		"RestaurantId": restaurantId,
		"Address":      address,
	}, "layout")
}
