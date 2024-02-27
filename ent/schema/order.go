package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_id"),
		field.Int("table_id"),
		field.Int("user_id"),
		field.Int("order_type"),
		field.String("status"),
		field.Time("created_at").Default(time.Now().UTC()),
		field.Time("updated_at").Default(time.Now().UTC()),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
