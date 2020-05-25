package types

import "time"

type (
	Customer struct {
		CustomerID string    `json:"customer_id,omitempty" bson:"customer_id"`
		Name       string    `json:"name,omitempty" bson:"name"`
		Age        int       `json:"age,omitempty" bson:"age"`
		Address    string    `json:"address,omitempty" bson:"address"`
		CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at"`
		UpdateAt   time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	}
)
