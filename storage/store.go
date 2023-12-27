package storage 

import (
	"rest-book/model"
)
type Store interface {
	GetRestaurantDetails(string)						(*model.Restaurant, error )
	AddRestaurantDetails(*model.Restaurant)				error
	UpdateRestaurantDetails(string, *model.Restaurant)	(error)
	GetAllRestaurantDetails()							([]*model.RestaurantDb, error)
	DeleteRestaurantDetails(string)						(error) 
}
