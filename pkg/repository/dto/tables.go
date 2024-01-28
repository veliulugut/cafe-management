package dto

import "time"

/*
	     field.Int("table_id"),
		field.Int("number_of_guests"),
		field.Int("table_number"),
		field.String("description"),
		field.String("table_type"),
		field.Time("created_at").Default(time.Now().UTC()),
		field.Time("updated_at").Default(time.Now().UTC()),
*/
type Tables struct {
	NumberOfGuests int       `json:"number_of_guests"`
	TableNumber    int       `json:"table_number"`
	Description    string    `json:"description"`
	TableType      string    `json:"table_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
