package helper

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
)

func DbMenuTo(d *ent.Menu) *dto.Menu {
	return &dto.Menu{
		MenuID:       d.ID,
		Name:         d.Name,
		Category:     d.Category,
		Price:        d.Price,
		Description:  d.Description,
		MenuImageUrl: d.MenuImageURL,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}
