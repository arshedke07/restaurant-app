package model

type OrderItem struct {
	Id       int
	OrderId  string
	Item     string
	Quantity string
	Price    int
	Picture  string
}

type TempOrder struct {
	UserId       string
	Item         string
	Quantity     string
	Price        int
	Picture      string
	RestaurantId string
}
