package routes

import (
	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type ILoginRoute interface {
	LoginUser(c *fiber.Ctx) error
	LogoutUser(c *fiber.Ctx) error
}

type LoginRoute struct {
	Service services.ILoginService
	Store   *session.Store
}

func NewLoginRoute(service services.ILoginService, store *session.Store) ILoginRoute {
	return LoginRoute{
		Service: service,
		Store:   store,
	}
}

func (route LoginRoute) LoginUser(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		sess, _ := route.Store.Get(c)
		if sess.Get("ID") == nil {
			return c.Render("loginpage", fiber.Map{
				"Title": "Login Page",
			}, "layout")
		} else {
			return c.Render("home", fiber.Map{
				"Title":    "The Foodie",
				"UserName": sess.Get("NAME"),
			}, "layout")
		}

	} else if c.Method() == "POST" {
		sess, err := route.Store.Get(c)
		if err != nil {
			return err
		}
		emailid := c.FormValue("emailid")
		password := c.FormValue("password")
		user, loginErr := route.Service.Login(emailid, password)
		if loginErr != nil {
			return loginErr
		}
		sess.Set("ID", user.Id)
		sess.Set("NAME", user.FirstName+" "+user.LastName)
		sess.Set("EMAILID", emailid)
		sess.Set("USERTYPE", user.UserType)
		sess.Save()
		return c.Render("home", fiber.Map{
			"Title":    "The Foodie",
			"UserName": user.FirstName + " " + user.LastName,
		}, "layout")
	}
	return nil
}

func (route LoginRoute) LogoutUser(c *fiber.Ctx) error {
	sess, _ := route.Store.Get(c)
	err := sess.Destroy()
	if err != nil {
		return err
	}
	saveErr := sess.Save()
	if saveErr != nil {
		return saveErr
	}
	return c.Render("home", fiber.Map{
		"Title": "The Foodie",
	}, "layout")
}
