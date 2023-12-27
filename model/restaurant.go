package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Restaurant struct {
	Email		string				`json:"email"`
	Name		string				`json:"name_of_restaurant"`
	Address		*Address			`json:"address"`	
	Timings		*Timing				`json:"timing"`
}

type RestaurantDb struct {
	Id			primitive.ObjectID	`bson:"_id" json:"id"`
	Email		string				`json:"email"`
	Name		string				`json:"name_of_restaurant"`
	Address		*Address			`json:"address"`	
	Timings		*Timing				`json:"timing"`
}

func( r *Restaurant) String() string {
	return fmt.Sprintf("Name: %s | Address: %s %s %s", r.Name, r.Address.Building, r.Address.Street, r.Address.Landmark)
}
