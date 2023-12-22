package model 

type Restaurant struct {
	Name		string		`json:"name_of_restaurant"`
	Address		*Address	`json:"address"`	
Timings		*Timing		`json:"timing"`
}
