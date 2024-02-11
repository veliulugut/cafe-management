package dto

import "time"

/*
	field.String("product_name").Unique(),
	field.String("description"),
	field.Float("price"),
	field.Int("quantity"),
	field.String("product_type"),
	field.Time("created_at").Default(time.Now().UTC()),
	field.Time("updated_at").Default(time.Now().UTC()),
*/

type Product struct {
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	ProductType string    `json:"product_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
