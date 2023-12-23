package storage 

import (
	"rest-book/model"
)
type Store interface {
	GetRestaurantDetails(string) *model.Restaurant
}
