package user

import "time"

type CreateUserModel struct {
	Password        string `json:"password"`
	Email           string `json:"email"`
	UserName        string `json:"user_name"`
	ConfirmPassword string `json:"confirm_password"`
}

type UpdateUserModel struct {
	UserName  string    `json:"user_name"`
	Avatar    string    `json:"avatar"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserModel struct {
	UserID    int       `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	UserName  string    `json:"user_name"`
	Avatar    string    `json:"avatar"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
