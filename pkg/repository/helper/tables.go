package helper

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
)

func DBTableTo(t *ent.Tables) *dto.Tables {
	return &dto.Tables{
		NumberOfGuests: t.NumberOfGuests,
		TableNumber:    t.TableNumber,
		Description:    t.Description,
		TableType:      t.TableType,
	}
}
