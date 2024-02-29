package ordertype

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"
)

var _ ServiceOrderType = (*OrderType)(nil)

func New(repo entadp.RepositoryInterface) *OrderType {
	return &OrderType{
		repo: repo,
	}
}

type OrderType struct {
	repo entadp.RepositoryInterface
}

func (o *OrderType) CreateOrderType(ctx context.Context, c *CreateOrderTypeModel) error {
	orderTypeDTO := dto.OrderType{
		ID:   c.ID,
		Name: c.Name,
	}

	if err := o.repo.OrderType().CreateOrderType(ctx, &orderTypeDTO); err != nil {
		return fmt.Errorf("order type / create order type :%w", err)
	}

	return nil
}

func (o *OrderType) DeleteOrderType(ctx context.Context, id int) error {
	if err := o.repo.OrderType().DeleteOrderType(ctx, id); err != nil {
		return fmt.Errorf("order type / delete order type :%w", err)
	}
	return nil
}

func (o *OrderType) ListOrderType(ctx context.Context) ([]*ListOrderType, error) {
	var (
		orderTypes []*ent.OrderType
		err        error
	)

	if orderTypes, err = o.repo.OrderType().ListOrderType(ctx); err != nil {
		return nil, fmt.Errorf("order type / list order type :%w", err)
	}

	var orderTypeModels []*ListOrderType
	for _, orderType := range orderTypes {
		orderTypeModels = append(orderTypeModels, &ListOrderType{
			ID:   orderType.ID,
			Name: orderType.Name,
		})
	}

	return orderTypeModels, nil
}

func (o *OrderType) UpdateOrderType(ctx context.Context, id int, c *UpdateOrderType) error {
	orderTypeUpdate := dto.OrderType{
		ID:   c.ID,
		Name: c.Name,
	}

	if err := o.repo.OrderType().UpdateOrderType(ctx, id, &orderTypeUpdate); err != nil {
		return fmt.Errorf("order type / update order type :%w", err)
	}

	return nil
}
