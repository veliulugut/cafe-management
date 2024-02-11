package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

/*
	product_type = food,drink etc.
*/

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("product_name").Unique(),
		field.String("description"),
		field.Float("price"),
		field.Int("quantity"),
		field.String("product_type"),
		field.Time("created_at").Default(time.Now().UTC()),
		field.Time("updated_at").Default(time.Now().UTC()),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
