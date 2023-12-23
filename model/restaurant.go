package model 

import (
	"fmt"
)
type Restaurant struct {
	Name		string		`json:"name_of_restaurant"`
	Address		*Address	`json:"address"`	
	Timings		*Timing		`json:"timing"`
}

func( r *Restaurant) String() string {
	return fmt.Sprintf("Name: %s | Address: %s %s %s", r.Name, r.Address.Building, r.Address.Street, r.Address.Landmark)
}
