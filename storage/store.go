package storage 

import (
	"rest-book/model"
)
type Store interface {
	New() *Store
	GetRestaurantDetails(string) *model.Restaurant
}


