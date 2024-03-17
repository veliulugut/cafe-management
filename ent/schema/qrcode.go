package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// QrCode holds the schema definition for the QrCode entity.
type QrCode struct {
	ent.Schema
}

// Fields of the QrCode.
func (QrCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").Unique(),
		field.Bytes("image"),
	}
}

// Edges of the QrCode.
func (QrCode) Edges() []ent.Edge {
	return nil
}
