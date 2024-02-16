package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OrderType holds the schema definition for the OrderType entity.
type OrderType struct {
	ent.Schema
}

// Fields of the OrderType.
func (OrderType) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
	}
}

// Edges of the OrderType.
func (OrderType) Edges() []ent.Edge {
	return nil
}
