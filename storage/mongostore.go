package storage

import (
    "context"
    "log"
    "rest-book/model"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
    uri			string 
    Client		*mongo.Client
    restaurantCol	*mongo.Collection
    userCol		*mongo.Collection
    productCol          *mongo.Collection
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

func(m *MongoStore) UserCollection(databaseName, collectionName string) {
    m.userCol = m.Client.Database(databaseName).Collection(collectionName)
}

func(m *MongoStore) ProductCollection(databaseName, collectionName string) {
    m.productCol = m.Client.Database(databaseName).Collection(collectionName)
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

func(m *MongoStore) GetAllRestaurantDetails() ([]*model.RestaurantDb, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
    defer cancel()

    filter := bson.M{}
    cursor, err := m.restaurantCol.Find(ctx, filter)
    if err != nil {
        return nil, err
    }

    rests := []*model.RestaurantDb{}
    for cursor.Next(context.Background()) {
        rest := new(model.RestaurantDb)
        cursor.Decode(rest)
        rests = append(rests, rest)
    }
    return rests, nil
}

func(m *MongoStore) DeleteRestaurantDetails(id string) (error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
    defer cancel() 

    primID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    filter := bson.M{"_id": primID}
    // res := m.restaurantCol.FindOneAndDelete(ctx, filter)
    res, err := m.restaurantCol.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }
    log.Println("INFO: Deleted Result:", res.DeletedCount)

    return nil
}

func(m *MongoStore) AddUser(user *model.User) (primitive.ObjectID, error) {
    ctx,cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    res, err := m.userCol.InsertOne(ctx, user)
    if err != nil {
        return primitive.NilObjectID, err
    }
    log.Println("INFO: Inserted into User Collection ->", res.InsertedID)
    return res.InsertedID.(primitive.ObjectID), nil
}

func(m *MongoStore) GetUsers() ([]*model.UserDb, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel() 

    users := []*model.UserDb{}
    filter := bson.M{}

    cursor, err := m.userCol.Find(ctx, filter)
    if err != nil {
        return nil, err
    }

    for cursor.Next(context.Background()) {
        user := &model.UserDb{}

        if err := cursor.Decode(user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func(m *MongoStore) UpdateUser(idStr string, user *model.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel() 
    // id, err := primitive.ObjectIDFromHex(idStr)
    // if err != nil {
    // 	return err
    // }

    res, err := m.userCol.UpdateByID(ctx, idStr, user)
    if err != nil {
        return err
    }
    log.Println("INFO: Update Result:", res.UpsertedID)
    return err
}

func(m *MongoStore) AddProduct(product *model.Product) (primitive.ObjectID, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel() 

    res, err := m.productCol.InsertOne(ctx, product)
    if err != nil {
        return primitive.NilObjectID, err
    }
    log.Println("INFO: Add Product:", res)
    return res.InsertedID.(primitive.ObjectID), nil
}

func(m *MongoStore) GetAllProducts() ([]*model.Product, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel() 
    
    filter := bson.M{}
    cursor, err := m.productCol.Find(ctx, filter)
    if err != nil {
        return nil, err
    }

    products := []*model.Product{}
    for cursor.Next(context.Background()) {
        product := &model.Product{}
        if err := cursor.Decode(product); err != nil {
            return nil, err
        }
        products = append(products, product) 
    }
    
    return products, nil
}
