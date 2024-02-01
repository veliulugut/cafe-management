package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Reservation holds the schema definition for the Reservation entity.
type Reservation struct {
	ent.Schema
}

// table_id = table_type
// Fields of the Reservation.
func (Reservation) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").Optional(),
		field.Time("created_at").Default(time.Now().UTC()),
		field.Time("updated_at").Default(time.Now().UTC()),
		field.Int("table_id"),
		field.String("phone_number").Optional(),
		field.Time("start_time").Default(time.Now().UTC()),
		field.Time("end_time").Default(time.Now().UTC()),
	}
}

// Edges of the Reservation.
func (Reservation) Edges() []ent.Edge {
	return nil
}
