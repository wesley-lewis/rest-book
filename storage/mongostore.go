package storage

import (
	"context"
	"log"
	"rest-book/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	uri			string 
	Client		*mongo.Client
}

func (m *MongoStore) Connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	opt := options.Client().ApplyURI(m.uri)
	client, err := mongo.Connect(ctx, opt)
	return client, err 
}

func NewMongoStore(uri string) Store {
	m :=  &MongoStore{
		uri: uri,
	}
	client, err := m.Connect()
	if err != nil {
		log.Fatal(err)
	}
	m.Client = client
	return m
}

func (m *MongoStore) GetRestaurantDetails(id string) (*model.Restaurant, error) {
	return nil, nil
}

func(m *MongoStore) AddRestaurantDetails(*model.Restaurant) error {
	return nil
}
