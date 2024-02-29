package order

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"
)

var _ OrderService = (*Order)(nil)

func New(repo entadp.RepositoryInterface) *Order {
	return &Order{
		repo: repo,
	}
}

type Order struct {
	repo entadp.RepositoryInterface
}

func (o *Order) CreateOrder(ctx context.Context, c *CreateOrderModel) error {
	orderDTO := dto.Order{
		OrderID:   c.OrderID,
		TableID:   c.TableID,
		UserId:    c.UserId,
		OrderType: c.OrderType,
		Status:    c.Status,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}

	if err := o.repo.Order().CreateOrder(ctx, &orderDTO); err != nil {
		return fmt.Errorf("order / create order :%w", err)
	}

	return nil
}

func (o *Order) DeleteOrder(ctx context.Context, id int) error {
	if err := o.repo.Order().DeleteOrder(ctx, id); err != nil {
		return fmt.Errorf("order / delete order :%w", err)
	}
	return nil
}

func (o *Order) GetById(ctx context.Context, id int) (*OrderModel, error) {
	var (
		dbOrder *dto.Order
		err     error
	)

	if dbOrder, err = o.repo.Order().GetById(ctx, id); err != nil {
		return nil, fmt.Errorf("order / get order :%w", err)
	}

	return &OrderModel{
		OrderID:   dbOrder.OrderID,
		TableID:   dbOrder.TableID,
		UserId:    dbOrder.UserId,
		OrderType: dbOrder.OrderType,
		Status:    dbOrder.Status,
		CreatedAt: dbOrder.CreatedAt,
		UpdatedAt: dbOrder.UpdatedAt,
	}, nil
}

func (o *Order) ListOrder(ctx context.Context) ([]*OrderModel, error) {
	var (
		dbOrders []*ent.Order
		err      error
	)

	if dbOrders, err = o.repo.Order().ListOrder(ctx); err != nil {
		return nil, fmt.Errorf("order / list order :%w", err)
	}

	var orderModels []*OrderModel
	for _, dbOrder := range dbOrders {
		orderModels = append(orderModels, &OrderModel{
			OrderID:   dbOrder.OrderID,
			TableID:   dbOrder.TableID,
			UserId:    dbOrder.UserID,
			OrderType: dbOrder.OrderType,
			Status:    dbOrder.Status,
			CreatedAt: dbOrder.CreatedAt,
			UpdatedAt: dbOrder.UpdatedAt,
		})
	}

	return orderModels, nil
}

func (o *Order) UpdateOrder(ctx context.Context, id int, c *UpdateOrderModel) error {
	updateOrderDTO := dto.Order{
		OrderID:   c.OrderID,
		TableID:   c.TableID,
		UserId:    c.UserId,
		OrderType: c.OrderType,
		Status:    c.Status,
		UpdatedAt: c.UpdatedAt,
	}

	if err := o.repo.Order().UpdateOrder(ctx, id, &updateOrderDTO); err != nil {
		return fmt.Errorf("order / update order :%w", err)
	}

	return nil
}
