package dto

import "time"

/*
field.Int("user_id"),
field.String("first_name"),
field.String("last_name"),
field.String("password"),
field.String("user_name"),
field.String("avatar"),
field.String("phone"),
field.Time("created_at").Default(time.Now().UTC()),
field.Time("updated_at").Default(time.Now().UTC()),
*/
type User struct {
	UserID          int       `json:"user_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Password        string    `json:"password"`
	UserName        string    `json:"user_name"`
	Email           string    `json:"email"`
	Avatar          string    `json:"avatar"`
	Phone           string    `json:"phone"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
	ConfirmPassword string    `json:"confirm_password"`
}
