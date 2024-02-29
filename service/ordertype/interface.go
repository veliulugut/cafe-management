package ordertype

import (
	"context"
)

type ServiceOrderType interface {
	CreateOrderType(ctx context.Context, c *CreateOrderTypeModel) error
	DeleteOrderType(ctx context.Context, id int) error
	UpdateOrderType(ctx context.Context, id int, c *UpdateOrderType) error
	ListOrderType(ctx context.Context) ([]*ListOrderType, error)
}
