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
		BikeID    string    `json:"bike_id,omitempty" bson:"bike_id,omitempty"`
		Name      string    `json:"name,omitempty" bson:"name,omitempty"`
		Number    string    `json:"number,omitempty" bson:"number,omitempty"`
		Color     string    `json:"color,omitempty" bson:"color,omitempty"`
		Cost      int       `json:"rental,omitempty" bson:"rental,omitempty"`
		Rate      int       `json:"rate" bson:"rate,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdateAt  time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
		Status    Status    `json:"status" bson:"status,omitempty"`
	}

	CreateBike struct {
		Number string `json:"number,omitempty"  validate:"required"`
		Name   string `json:"name,omitempty"  validate:"required"`
		Color  string `json:"color,omitempty" validate:"required"`
		Cost   int    `json:"cost,omitempty"  validate:"required"`
	}

	UpdateBike struct {
		Number string `json:"number,omitempty"`
		Name   string `json:"name,omitempty"`
		Color  string `json:"color,omitempty"`
		Cost   int    `json:"cost,omitempty" `
		Status Status `json:"status,omitempty"`
	}
)
