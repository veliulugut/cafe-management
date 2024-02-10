package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"context"
	"fmt"
	"time"
)

var _ PriceRepository = (*Price)(nil)

type Price struct {
	dbClient *ent.Client
}

func NewPriceRepository(dbClient *ent.Client) *Price {
	return &Price{
		dbClient: dbClient,
	}
}

func (p *Price) CreatePrice(ctx context.Context, c *dto.Price) error {
	_, err := p.dbClient.Price.Create().
		SetPriceName(c.PriceName).
		SetPrice(c.Price).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("price / create price :%w", err)
	}

	return nil
}

func (p *Price) DeletePrice(ctx context.Context, id int) error {
	if err := p.dbClient.Price.DeleteOneID(id).Exec(ctx); err != nil {
		return fmt.Errorf("price / delete price :%w", err)
	}

	return nil

}

func (p *Price) ListPrice(ctx context.Context) ([]*ent.Price, error) {
	var (
		prices []*ent.Price
		err    error
	)

	prices, err = p.dbClient.Price.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("price / list price: %w", err)
	}

	return prices, nil
}

func (p *Price) UpdatePrice(ctx context.Context, id int, c *dto.Price) error {
	err := p.dbClient.Price.UpdateOneID(id).
		SetPriceName(c.PriceName).
		SetPrice(c.Price).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("update price:%w", ErrorNotFound)
		}
	}

	return nil
}
