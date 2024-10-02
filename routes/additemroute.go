package routes

import (
	"fmt"

	"github.com/arshedke07/restaurant-app/model"
	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type IAddItemRoute interface {
	AddItem(c *fiber.Ctx) error
	MyMenuPage(c *fiber.Ctx) error
}

type AddItemRoute struct {
	Service           services.IAddItemService
	Store             *session.Store
	RestaurantService services.IRestaurantService
}

func NewItemRoute(service services.IAddItemService, store *session.Store, rservice services.IRestaurantService) IAddItemRoute {
	return AddItemRoute{
		Service:           service,
		Store:             store,
		RestaurantService: rservice,
	}
}

func (route AddItemRoute) AddItem(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := sess.Get("ID")
	str := fmt.Sprintf("%v", id)
	data, err := route.RestaurantService.MyRestaurantService(str)
	if err != nil {
		return err
	}

	if c.Method() == "GET" {
		return c.Render("addmenuitem", fiber.Map{
			"Title":    "Add Item",
			"UserName": sess.Get("NAME"),
		}, "layout")
	} else if c.Method() == "POST" {
		item := model.Items{
			RestaurantId: data.Id,
			ItemName:     c.FormValue("item_name"),
			Price:        c.FormValue("price"),
			Quantity:     c.FormValue("quantity"),
			Description:  c.FormValue("description"),
			Picture:      c.FormValue("picture"),
		}

		_, err := route.Service.AddNewItem(&item)
		if err != nil {
			return err
		}

		data, err := route.Service.FindAllItemsById(data.Id)
		if err != nil {
			return err
		}

		return c.Render("mymenupage", fiber.Map{
			"Title":    "My Restaurant",
			"UserName": sess.Get("NAME"),
			"Data":     data,
		}, "layout")
	}
	return nil
}

func (route AddItemRoute) MyMenuPage(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	id := sess.Get("ID")
	str := fmt.Sprintf("%v", id)
	data, err := route.RestaurantService.MyRestaurantService(str)
	if err != nil {
		return err
	}

	user, findErr := route.Service.FindAllItemsById(data.Id)
	if findErr != nil {
		return findErr
	}
	return c.Render("mymenupage", fiber.Map{
		"Title":    "My Restaurant",
		"UserName": sess.Get("NAME"),
		"Data":     user,
	}, "layout")
}
