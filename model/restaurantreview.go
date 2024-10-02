package model

type RestaurantReview struct {
	Id           string
	RestaurantId string
	Review       string
	Rating       float32
}
