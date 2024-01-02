package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
    Name		    string			`json:"name"`
    DateOfManufacturing	    string			`json:"date_of_manufacturing"`
}

type ItemDb struct {
    Id			    primitive.ObjectID		`json:"_id" bson:"_id"`
    Name		    string			`json:"name"`
    DateOfManufacturing	    string			`json:"date_of_manufacturing"`
}
