package model 

type Address struct {
	Building		string		`json:"building"`
	Street			string		`json:"street"`
	Landmark		string		`json:"landmark"`
	Town			string		`json:"town"`
	City			string		`json:"city"`
	State			string		`json:"state"`
	Timing			*Timing		`json:"timing"`
}
