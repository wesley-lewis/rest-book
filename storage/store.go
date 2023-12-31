package storage

import (
	"rest-book/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Store interface {
	GetRestaurantDetails(string)					(*model.Restaurant, error )
	AddRestaurantDetails(*model.Restaurant)				error
	UpdateRestaurantDetails(string, *model.Restaurant)	(error)
	GetAllRestaurantDetails()							([]*model.RestaurantDb, error)
	DeleteRestaurantDetails(string)				(error) 
	AddUser(*model.User)									(primitive.ObjectID, error)
	GetUsers()														([]*model.UserDb, error)
}
