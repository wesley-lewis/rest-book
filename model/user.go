package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint8  `json:"age"`
	Email     string `json:"email"`
	Phone     string `json:"mobile_no"`
}

type UserDb struct {
	Id				primitive.ObjectID		`bson:"_id" json:"id"`
	LastName  string `json:"last_name"`
	Age       uint8  `json:"age"`
	Email     string `json:"email"`
	Phone     string `json:"mobile_no"`
}
