package routes

import (
	"fmt"

	"github.com/arshedke07/restaurant-app/model"
	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
)

type IAddRestaurantRoute interface {
	AddRestaurant(c *fiber.Ctx) error
	MyRestaurant(c *fiber.Ctx) error
}

type AddRestaurantRoute struct {
	Service services.IRestaurantService
	Store   *session.Store
}

func NewRestaurantRoute(service services.IRestaurantService, store *session.Store) IAddRestaurantRoute {
	return AddRestaurantRoute{
		Service: service,
		Store:   store,
	}
}

func (route AddRestaurantRoute) AddRestaurant(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)

	id := sess.Get("ID")
	str := fmt.Sprintf("%v", id)

	if c.Method() == "GET" {
		return c.Render("addrestaurant", fiber.Map{
			"Title":    "Register Restaurant",
			"UserName": sess.Get("NAME"),
		}, "layout")
	} else if c.Method() == "POST" {
		restaurant := model.Restaurant{
			UserId:      str,
			Name:        c.FormValue("name"),
			Add1:        c.FormValue("add1"),
			Add2:        c.FormValue("add2"),
			City:        c.FormValue("city"),
			State:       c.FormValue("state"),
			Pincode:     c.FormValue("pincode"),
			Picture:     c.FormValue("picture"),
			Description: c.FormValue("description"),
			Cuisine:     c.FormValue("cuisine"),
		}
		user, err := route.Service.AddNewRestaurant(&restaurant)
		if err != nil {
			return err
		}

		return c.Render("myrestaurantpage", fiber.Map{
			"Title":       "My Restaurant",
			"UserName":    sess.Get("NAME"),
			"Name":        user.Name,
			"Description": user.Description,
			"Address":     user.Add1 + "," + user.Add2 + "," + user.City + "," + user.State + "," + user.Pincode,
			"Picture":     user.Picture,
			"Cuisine":     user.Cuisine,
		}, "layout")
	}
	return nil
}

func (route AddRestaurantRoute) MyRestaurant(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)

	id := sess.Get("ID")
	str := fmt.Sprintf("%v", id)

	user, err := route.Service.MyRestaurantService(str)
	if err != nil {
		return err
	}

	return c.Render("myrestaurantpage", fiber.Map{
		"Title":       "My Restaurant",
		"UserName":    sess.Get("NAME"),
		"Name":        user.Name,
		"Description": user.Description,
		"Address":     user.Add1 + "," + user.Add2 + "," + user.City + "," + user.State + "," + user.Pincode,
		"Picture":     user.Picture,
		"Id":          user.Id,
		"Cuisine":     user.Cuisine,
	}, "layout")
}
