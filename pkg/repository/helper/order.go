package helper

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
)

func DbOrderTo(d *ent.Order) *dto.Order {
	return &dto.Order{
		OrderID:   d.ID,
		UserId:    d.UserID,
		OrderType: d.OrderType,
		TableID:   d.TableID,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
