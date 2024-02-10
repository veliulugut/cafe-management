package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tables holds the schema definition for the Tables entity.
type Tables struct {
	ent.Schema
}

// Fields of the Tables.
func (Tables) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number_of_guests"),
		field.Int("table_number"),
		field.String("description"),
		field.String("table_type"),
		field.Time("created_at").Default(time.Now().UTC()),
		field.Time("updated_at").Default(time.Now().UTC()),
	}
}

// table_id reservation table_id olarak kayıtlı olacak
// Edges of the Tables.
func (Tables) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reservation", Reservation.Type),
	}
}
