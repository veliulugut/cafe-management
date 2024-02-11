package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"context"
	"fmt"
	"time"
)

var _ ProductRepository = (*Product)(nil)

type Product struct {
	dbClient *ent.Client
}

func NewProductRepository(dbClient *ent.Client) *Product {
	return &Product{
		dbClient: dbClient,
	}
}

func (p *Product) CreateProduct(ctx context.Context, c *dto.Product) error {
	_, err := p.dbClient.Product.Create().
		SetProductName(c.ProductName).
		SetDescription(c.Description).
		SetPrice(c.Price).
		SetQuantity(c.Quantity).
		SetProductType(c.ProductType).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("product / create product :%w", err)
	}

	return nil
}

func (p *Product) DeleteProduct(ctx context.Context, id int) error {
	if err := p.dbClient.Product.DeleteOneID(id).Exec(ctx); err != nil {
		return fmt.Errorf("product / delete product :%w", err)
	}

	return nil

}

func (p *Product) ListProduct(ctx context.Context) ([]*ent.Product, error) {
	var (
		products []*ent.Product
		err      error
	)

	products, err = p.dbClient.Product.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("product / list product :%w", err)
	}

	return products, nil
}

func (p *Product) UpdateProduct(ctx context.Context, id int, c *dto.Product) error {
	err := p.dbClient.Product.UpdateOneID(id).SetProductName(c.ProductName).
		SetDescription(c.Description).
		SetPrice(c.Price).
		SetQuantity(c.Quantity).
		SetProductType(c.ProductType).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("product / update product :%w", err)
	}

	return nil
}
