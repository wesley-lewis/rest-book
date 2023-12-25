package storage

import (
	"context"
	"log"
	"rest-book/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	uri				string 
	Client			*mongo.Client
	restaurantCol	*mongo.Collection
}

func (m *MongoStore) Connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	opt := options.Client().ApplyURI(m.uri)
	client, err := mongo.Connect(ctx, opt)
	return client, err 
}

func(m *MongoStore) RestaurantCollection(databaseName, collectionName string) {
	m.restaurantCol = m.Client.Database(databaseName).Collection(collectionName)
}

func NewMongoStore(uri string) *MongoStore {
	m :=  &MongoStore{
		uri: uri,
	}
	client, err := m.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("INFO: Connected to MongoDB")
	m.Client = client
	return m
}

func (m *MongoStore) GetRestaurantDetails(email string) (*model.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	rest := &model.Restaurant{}
	filter := bson.M{"email": email}
	err := m.restaurantCol.FindOne(ctx, filter).Decode(rest)  
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return rest, nil
}

func(m *MongoStore) AddRestaurantDetails(rest *model.Restaurant) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	_, err := m.restaurantCol.InsertOne(ctx, rest)
	return err
}

func(m *MongoStore) UpdateRestaurantDetails(id string, rest *model.Restaurant) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	result, err := m.restaurantCol.UpdateByID(ctx, id, rest)
	log.Println("INFO: Update Result:",result)
	return err
}

func(m *MongoStore) GetAllRestaurantDetails() ([]*model.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	filter := bson.M{}
	cursor, err := m.restaurantCol.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	rests := []*model.Restaurant{}
	for cursor.Next(context.Background()) {
		rest := new(model.Restaurant)
		cursor.Decode(rest)
		rests = append(rests, rest)
	}
	return rests, nil
}
