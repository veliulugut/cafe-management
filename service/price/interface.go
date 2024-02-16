package price

import "context"

type ServicePrice interface {
	CreatePrice(ctx context.Context, c *CreatePriceModel) error
	DeletePrice(ctx context.Context, id int) error
	UpdatePrice(ctx context.Context, id int, c *UpdatePriceModel) error
	ListPrice(ctx context.Context) ([]*PriceModel, error)
}
