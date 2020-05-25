package types

import "time"

type (
	Rental struct {
		RentalID   string    `json:"rental_id,omitempty" bson:"rental_id"`
		CustomerID string    `json:"customer_id,omitempty" bson:"customer_id"`
		BikeID     string    `json:"bike_id,omitempty" bson:"bike_id"`
		CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at"`
		From       time.Time `json:"from,omitempty" bson:"from"`
		To         time.Time `json:"to,omitempty" bson:"to"`
		TotalCost  float64   `json:"total_cost,omitempty" bson:"total_cost"`
	}
)
