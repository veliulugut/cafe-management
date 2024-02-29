package order

import "context"

type OrderService interface {
	CreateOrder(ctx context.Context, c *CreateOrderModel) error
	DeleteOrder(ctx context.Context, id int) error
	UpdateOrder(ctx context.Context, id int, c *UpdateOrderModel) error
	GetById(ctx context.Context, id int) (*OrderModel, error)
	ListOrder(ctx context.Context) ([]*OrderModel, error)
}
