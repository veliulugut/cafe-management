package dto

import "time"

/*
field.String("price"),
fild.String("name"),
field.Time("created_at").Default(time.Now().UTC()),
field.Time("updated_at").Default(time.Now().UTC()),
*/
type Price struct {
	PriceName string    `json:"price_name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
