package helper

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
)

func DbProductTo(d *ent.Product) *dto.Product {
	return &dto.Product{
		ProductName: d.ProductName,
		Description: d.Description,
		Price:       d.Price,
		Quantity:    d.Quantity,
		ProductType: d.ProductType,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
