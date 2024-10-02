package model

type Order struct {
	Id           string
	UserId       string
	Tax          float32
	Total        int
	GrandTotal   float32
	Status       int
	RestaurantId string
	Address      string
}
