package dto

import "time"

type Order struct {
	OrderID   int       `json:"order_id"`
	TableID   int       `json:"table_id"`
	UserId    int       `json:"user_id"`
	OrderType int       `json:"order_type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
