package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/helper"
	"context"
	"fmt"
	"time"
)

var _ OrderRepository = (*Order)(nil)

func NewOrderRepository(dbClient *ent.Client) *Order {
	return &Order{
		dbClient: dbClient,
	}
}

type Order struct {
	dbClient *ent.Client
}

func (o *Order) CreateOrder(ctx context.Context, c *dto.Order) error {
	_, err := o.dbClient.Order.Create().
		SetOrderID(c.OrderID).
		SetTableID(c.TableID).
		SetUserID(c.UserId).
		SetOrderType(c.OrderType).
		SetStatus(c.Status).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("order / create order :%w", err)
	}

	return nil
}

func (o *Order) DeleteOrder(ctx context.Context, id int) error {
	if err := o.dbClient.Order.DeleteOneID(id).Exec(ctx); err != nil {
		return fmt.Errorf("order / delete order :%w", err)
	}

	return nil
}

func (o *Order) GetById(ctx context.Context, id int) (*dto.Order, error) {
	var (
		order *ent.Order
		err   error
	)

	order, err = o.dbClient.Order.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("order / get order :%w", err)
	}

	return helper.DbOrderTo(order), nil
}

func (o *Order) ListOrder(ctx context.Context) ([]*ent.Order, error) {
	var (
		order []*ent.Order
		err   error
	)

	order, err = o.dbClient.Order.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("order / list order :%w", err)
	}

	return order, nil
}

func (o *Order) UpdateOrder(ctx context.Context, id int, c *dto.Order) error {
	err := o.dbClient.Order.UpdateOneID(id).SetOrderID(c.OrderID).
		SetTableID(c.TableID).
		SetUserID(c.UserId).
		SetOrderType(c.OrderType).
		SetStatus(c.Status).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("order / update order :%w", err)
	}

	return nil
}
