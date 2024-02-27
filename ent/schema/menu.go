package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.Int("menu_id"),
		field.String("name"),
		field.String("category"),
		field.Time("created_at").Default(time.Now().UTC()),
		field.Time("updated_at").Default(time.Now().UTC()),
		field.String("description"),
		field.Float("price"),
		field.String("menu_image_url"),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return nil
}
