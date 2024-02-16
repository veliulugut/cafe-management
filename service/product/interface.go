package product

import "context"

type ServiceProduct interface {
	CreateProduct(ctx context.Context, c *CreateProductModel) error
	DeleteProduct(ctx context.Context, id int) error
	UpdateProduct(ctx context.Context, id int, c *UpdateProductModel) error
	ListProduct(ctx context.Context) ([]*ProductModel, error)
	GetById(ctx context.Context, id int) (*ProductModel, error)
}
