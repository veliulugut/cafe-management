package order

import "time"

type CreateOrderModel struct {
	OrderID   int       `json:"order_id"`
	TableID   int       `json:"table_id"`
	UserId    int       `json:"user_id"`
	OrderType int       `json:"order_type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateOrderModel struct {
	OrderID   int       `json:"order_id"`
	TableID   int       `json:"table_id"`
	UserId    int       `json:"user_id"`
	OrderType int       `json:"order_type"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderModel struct {
	OrderID   int       `json:"order_id"`
	TableID   int       `json:"table_id"`
	UserId    int       `json:"user_id"`
	OrderType int       `json:"order_type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
