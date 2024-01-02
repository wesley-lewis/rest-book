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
        log.Println("ERROR:", err)
        return nil, err
    }

    return rest, nil
}

func(m *MongoStore) AddRestaurantDetails(rest *model.Restaurant) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
    defer cancel()

    _, err := m.restaurantCol.InsertOne(ctx, rest)
    if err != nil {
        log.Println("ERROR:", err)
    }
    return err
}

func(m *MongoStore) UpdateRestaurantDetails(id string, rest *model.Restaurant) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
    defer cancel()

    result, err := m.restaurantCol.UpdateByID(ctx, id, rest)
    if err != nil {
        log.Println("INFO: Update Result:",result)
        log.Println("ERROR:", err)
    }
    return err
}

func(m *MongoStore) GetAllRestaurantDetails() ([]*model.RestaurantDb, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
    defer cancel()

    filter := bson.M{}
    cursor, err := m.restaurantCol.Find(ctx, filter)
    if err != nil {
        log.Println("ERROR:", err)
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
        log.Println("ERROR:", err)
        return err
    }
    filter := bson.M{"_id": primID}
    // res := m.restaurantCol.FindOneAndDelete(ctx, filter)
    res, err := m.restaurantCol.DeleteOne(ctx, filter)
    if err != nil {
        log.Println("ERROR:", err)
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
        log.Println("ERROR:", err)
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
        log.Println("ERROR:", err)
        return nil, err
    }

    for cursor.Next(context.Background()) {
        user := &model.UserDb{}

        if err := cursor.Decode(user); err != nil {
            log.Println("ERROR:", err)
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func(m *MongoStore) GetUserByEmail(email string) (*model.UserDb, error) {
    filter := bson.M{"email": email}
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10) 
    defer cancel()

    cur, err := m.userCol.Find(ctx, filter)
    if err != nil {
        log.Println("ERROR:", err)
        return nil, err
    }
    user := &model.UserDb{}
    if cur.Next(context.Background()) {
        if err := cur.Decode(user);err != nil {
            log.Println("ERROR:", err)
            return nil, err
        }
    }

    return user, nil
}

func(m *MongoStore) UpdateUser(idStr string, user *model.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel() 

    res, err := m.userCol.UpdateByID(ctx, idStr, user)
    if err != nil {
        log.Println("ERROR:", err)
        return err
    }
    log.Println("INFO: Update Result:", res.UpsertedID)
    return err
}

func(m *MongoStore) DeleteUser(id string) (error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objId}
    res, err := m.userCol.DeleteOne(ctx, filter)
    if err != nil {
        log.Println("ERROR:", err)
        log.Println("INFO: Deleted user result:", res.DeletedCount)
    }
    return err
}

func(m *MongoStore) AddItem(product *model.Item) (primitive.ObjectID, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel() 

    res, err := m.productCol.InsertOne(ctx, product)
    if err != nil {
        log.Println("ERROR:", err)
        return primitive.NilObjectID, err
    }
    log.Println("INFO: Add Item:", res)
    return res.InsertedID.(primitive.ObjectID), nil
}

func(m *MongoStore) GetAllItems() ([]*model.ItemDb, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel() 
    
    filter := bson.M{}
    cursor, err := m.productCol.Find(ctx, filter)
    if err != nil {
        log.Println("ERROR:", err)
        return nil, err
    }

    products := []*model.ItemDb{}
    for cursor.Next(context.Background()) {
        product := &model.ItemDb{}
        if err := cursor.Decode(product); err != nil {
            log.Println("ERROR:", err)
            return nil, err
        }
        products = append(products, product) 
    }
    
    return products, nil
}

func(m *MongoStore) DeleteItem(idStr string) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
    defer cancel()

    id, _ := primitive.ObjectIDFromHex(idStr)
    filter := bson.M{"_id": id}
    res,err := m.productCol.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }
    log.Println("INFO: Deleted User:", id, " Count:", res.DeletedCount)
    return nil
}
