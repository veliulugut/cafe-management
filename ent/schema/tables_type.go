package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tables_type holds the schema definition for the Tables_type entity.
type Tables_type struct {
	ent.Schema
}

// Fields of the Tables_type.
func (Tables_type) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tables_id"),
		field.String("name"),
	}
}

// Edges of the Tables_type.
func (Tables_type) Edges() []ent.Edge {
	return nil
}
