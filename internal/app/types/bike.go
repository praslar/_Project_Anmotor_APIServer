package types

import "time"

const (
	Rented Status = iota
	Avaiable
	Unavaible
)

type (
	Status int
	Bike   struct {
		BikeID    string    `json:"bike_id,omitempty" bson:"bike_id"`
		Number    string    `json:"number,omitempty" bson:"number"`
		Color     string    `json:"color,omitempty" bson:"color"`
		Cost      int       `json:"rental,omitempty" bson:"rental"`
		Rate      int       `json:"rate" bson:"rate"`
		CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
		UpdateAt  time.Time `json:"updated_at,omitempty" bson:"updated_at"`
		Status    Status    `json:"status" bson:"status"`
	}
)
