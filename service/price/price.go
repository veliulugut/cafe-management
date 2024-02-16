package price

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"
)

var _ ServicePrice = (*Price)(nil)

func New(repo entadp.RepositoryInterface) *Price {
	return &Price{
		repo: repo,
	}
}

type Price struct {
	repo entadp.RepositoryInterface
}

func (p *Price) CreatePrice(ctx context.Context, c *CreatePriceModel) error {
	priceDTO := dto.Price{
		Price:     c.Price,
		PriceName: c.PriceName,
	}

	if err := p.repo.Price().CreatePrice(ctx, &priceDTO); err != nil {
		return fmt.Errorf("price srv / create price : %w", err)
	}

	return nil
}

func (p *Price) DeletePrice(ctx context.Context, id int) error {
	if err := p.repo.Price().DeletePrice(ctx, id); err != nil {
		return fmt.Errorf("price srv / delete price : %w", err)
	}

	return nil
}

func (p *Price) ListPrice(ctx context.Context) ([]*PriceModel, error) {
	var (
		prices []*ent.Price
		err    error
	)

	if prices, err = p.repo.Price().ListPrice(ctx); err != nil {
		return nil, fmt.Errorf("price srv / list price : %w", err)
	}

	var priceModels []*PriceModel
	for _, price := range prices {
		priceModels = append(priceModels, &PriceModel{
			Price:     price.Price,
			PriceName: price.PriceName,
		})
	}

	return priceModels, nil
}

func (p *Price) UpdatePrice(ctx context.Context, id int, c *UpdatePriceModel) error {
	updateDTO := dto.Price{
		Price:     c.Price,
		PriceName: c.PriceName,
	}

	if err := p.repo.Price().UpdatePrice(ctx, id, &updateDTO); err != nil {
		return fmt.Errorf("price srv / update price : %w", err)
	}

	return nil
}
