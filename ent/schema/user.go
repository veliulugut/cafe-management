package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("first_name"),
		field.String("last_name"),
		field.String("password"),
		field.String("user_name").Unique(),
		field.String("email").Unique(),
		field.String("avatar"),
		field.String("phone"),
		field.Time("created_at").Default(time.Now().UTC()),
		field.Time("updated_at").Default(time.Now().UTC()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
