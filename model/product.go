package model

import "time"

type Product struct {
    Name                string `json:"product"`
    DateOfManufacturing time.Time  `json:"date_of_manufacturing"`
}
