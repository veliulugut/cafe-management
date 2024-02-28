package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"context"
	"fmt"
)

var _ OrdersTypeRepository = (*OrderType)(nil)

func NewOrdersTypeRepository(dbClient *ent.Client) *OrderType {
	return &OrderType{
		dbClient: dbClient,
	}
}

type OrderType struct {
	dbClient *ent.Client
}

func (o *OrderType) CreateOrderType(ctx context.Context, c *dto.OrderType) error {
	_, err := o.dbClient.OrderType.Create().SetID(c.ID).SetName(c.Name).Save(ctx)
	if err != nil {
		return fmt.Errorf("order type / create order type :%w", err)
	}

	return nil
}

func (o *OrderType) DeleteOrderType(ctx context.Context, id int) error {
	err := o.dbClient.OrderType.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("order type / delete order type :%w", err)
	}

	return nil
}

func (o *OrderType) UpdateOrderType(ctx context.Context, id int, c *dto.OrderType) error {
	err := o.dbClient.OrderType.UpdateOneID(id).SetName(c.Name).Exec(ctx)
	if err != nil {
		return fmt.Errorf("order type / update order type :%w", err)
	}

	return nil
}

func (o *OrderType) ListOrderType(ctx context.Context) ([]*ent.OrderType, error) {
	orderTypes, err := o.dbClient.OrderType.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("order type / list order type :%w", err)
	}

	return orderTypes, nil
}
