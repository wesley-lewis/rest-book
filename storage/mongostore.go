package storage 

import (
	"rest-book/model"
)

type MongoStore struct {
	uri			string 
}

func NewMongoStore(uri string) Store {
	return &MongoStore{
		uri: uri,
	}
}

func (m *MongoStore) GetRestaurantDetails(id string) (*model.Restaurant, error) {
	return nil, nil
}

func(m *MongoStore) AddRestaurantDetails(*model.Restaurant) error {
	return nil
}
