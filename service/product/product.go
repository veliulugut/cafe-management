package product

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"
)

var _ ServiceProduct = (*Product)(nil)

func New(repo entadp.RepositoryInterface) *Product {
	return &Product{
		repo: repo,
	}
}

type Product struct {
	repo entadp.RepositoryInterface
}

func (p *Product) CreateProduct(ctx context.Context, c *CreateProductModel) error {
	productDTO := dto.Product{
		ProductName: c.ProductName,
		Description: c.Description,
		Price:       c.Price,
		Quantity:    c.Quantity,
		ProductType: c.ProductType,
	}

	if err := p.repo.Product().CreateProduct(ctx, &productDTO); err != nil {
		return fmt.Errorf("product srv / create product : %w", err)
	}

	return nil
}

func (p *Product) DeleteProduct(ctx context.Context, id int) error {
	if err := p.repo.Product().DeleteProduct(ctx, id); err != nil {
		return fmt.Errorf("product srv / delete product : %w", err)
	}

	return nil
}

func (p *Product) GetById(ctx context.Context, id int) (*ProductModel, error) {
	var (
		dbProduct *dto.Product
		err       error
	)

	if dbProduct, err = p.repo.Product().GetById(ctx, id); err != nil {
		return nil, fmt.Errorf("product srv / get product : %w", err)
	}

	return &ProductModel{
		ProductName: dbProduct.ProductName,
		Description: dbProduct.Description,
		Price:       dbProduct.Price,
		Quantity:    dbProduct.Quantity,
		ProductType: dbProduct.ProductType,
	}, nil

}

func (p *Product) ListProduct(ctx context.Context) ([]*ProductModel, error) {
	var (
		dbProducts []*ent.Product
		err        error
	)

	if dbProducts, err = p.repo.Product().ListProduct(ctx); err != nil {
		return nil, fmt.Errorf("product srv / list product : %w", err)
	}

	var products []*ProductModel
	for _, dbProduct := range dbProducts {
		products = append(products, &ProductModel{
			ProductName: dbProduct.ProductName,
			Description: dbProduct.Description,
			Price:       dbProduct.Price,
			Quantity:    dbProduct.Quantity,
			ProductType: dbProduct.ProductType,
		})
	}

	return products, nil
}

func (p *Product) UpdateProduct(ctx context.Context, id int, c *UpdateProductModel) error {
	productDTO := dto.Product{
		ProductName: c.ProductName,
		Description: c.Description,
		Price:       c.Price,
		Quantity:    c.Quantity,
		ProductType: c.ProductType,
	}

	if err := p.repo.Product().UpdateProduct(ctx, id, &productDTO); err != nil {
		return fmt.Errorf("product srv / update product : %w", err)
	}

	return nil
}
