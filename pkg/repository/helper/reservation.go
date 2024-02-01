package helper

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
)

func DbReservationTo(d *ent.Reservation) *dto.Reservation {
	return &dto.Reservation{
		FullName:    d.FullName,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		TableID:     d.TableID,
		PhoneNumber: d.PhoneNumber,
	}
}
