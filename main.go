package main

import (
	"time"

	"github.com/arshedke07/restaurant-app/model"
	"github.com/arshedke07/restaurant-app/routes"
	"github.com/arshedke07/restaurant-app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views:                 engine,
		DisableStartupMessage: false,
	})

	dbconfig := model.DbConfig{
		ConnectionString: "host=localhost port=5432 user=postgres password=1234 dbname=restaurantapp sslmode=disable",
	}

	postgresStore := postgres.New(postgres.Config{
		ConnectionURI: dbconfig.ConnectionString,
		Table:         "fiber_storage",
		Reset:         false,
		GCInterval:    1 * time.Hour,
	})

	store := session.New(session.Config{
		Storage: postgresStore,
	})

	//Setup services

	homeroute := routes.NewHomeRoute(store)

	loginService := services.NewLoginService(&dbconfig)
	loginRoute := routes.NewLoginRoute(loginService, store)

	userService := services.NewUserService(&dbconfig)
	addUserRoute := routes.NewUserRoute(userService, store)

	restaurantService := services.NewRestaurantService(&dbconfig)
	addRestaurant := routes.NewRestaurantRoute(restaurantService, store)

	addressService := services.NewAddressService(&dbconfig)
	addaddress := routes.NewAddressRoute(addressService, store, loginService)

	itemservice := services.NewItemService(&dbconfig)
	additem := routes.NewItemRoute(itemservice, store, restaurantService)

	cartservice := services.NewAddToCartService(&dbconfig)

	profileservice := services.NewProfileService(&dbconfig)

	orderroute := routes.NewOrderRoute(restaurantService, store)
	menuroute := routes.NewMenuRoute(itemservice, store, cartservice, profileservice)

	orderservice := services.NewOrderService(&dbconfig, cartservice)
	placeorder := routes.NewPlaceOrderRoute(orderservice, store)

	deliveryservice := services.NewDeliveryService(&dbconfig)
	deliveryroute := routes.NewPendingOrderRoute(store, deliveryservice)

	historyservice := services.NewOrderHistory(&dbconfig)
	historyroute := routes.NewOrderHistoryRoute(store, historyservice)

	profileroute := routes.NewProfileRoute(store, profileservice)

	app.Get("/", homeroute.Home)
	app.Get("/logout", loginRoute.LogoutUser)
	app.Get("/login", loginRoute.LoginUser)
	app.Post("/login", loginRoute.LoginUser)
	app.Get("/adduser", addUserRoute.AddUser)
	app.Post("/adduser", addUserRoute.AddUser)
	app.Get("/addrestaurant", addRestaurant.AddRestaurant)
	app.Post("/addrestaurant", addRestaurant.AddRestaurant)
	app.Get("/myrestaurantpage", addRestaurant.MyRestaurant)
	app.Get("/addaddress", addaddress.AddUserAddress)
	app.Post("/addaddress", addaddress.AddUserAddress)
	app.Get("/addmenuitem", additem.AddItem)
	app.Post("/addmenuitem", additem.AddItem)
	app.Get("/mymenupage", additem.MyMenuPage)
	app.Get("/orderpage", orderroute.OrderPageRoute)
	app.Get("/ordermenupage/:id", menuroute.MenuPageRoute)
	app.Post("/ordermenupage/:id", menuroute.MenuPageRoute)
	app.Get("/mycart/:id", menuroute.MyCartRoute)
	app.Post("/mycart/:id", placeorder.PlaceOrder)
	app.Get("/pendingorders/:id", deliveryroute.GetPendingOrders)
	app.Get("/orderdetails/:id", deliveryroute.GetPendingItems)
	app.Post("/orderdetails/:id", deliveryroute.DispatchOrder)
	app.Get("/myorders/:id", historyroute.MyOrderHistory)
	app.Get("/myorderdetails/:id", historyroute.MyOrderItemHistory)
	app.Get("/userprofile", profileroute.GetProfile)
	app.Get("/myaddresses", profileroute.GetAddress)

	app.Listen(":3000")
}
