package storage 

import (
	"rest-book/model"
)
type Store interface {
	GetRestaurantDetails(string)						(*model.Restaurant, error )
	AddRestaurantDetails(*model.Restaurant)				error
	UpdateRestaurantDetails(string, *model.Restaurant)	(error)
	GetAllRestaurantDetails()							([]*model.Restaurant, error)
	DeleteRestaurantDetails(string)						(error) 
}
