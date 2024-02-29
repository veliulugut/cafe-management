package reservation

import (
	"context"
	"time"
)

type ReservationService interface {
	CreateReservation(ctx context.Context, c *CreateReservationModel) error
	DeleteReservation(ctx context.Context, id int) error
	UpdateReservation(ctx context.Context, id int, c *UpdateReservationModel) error
	GetByNameReservation(ctx context.Context, name string) (*ReservationModel, error)
	CheckAvailability(ctx context.Context, startTime, endTime time.Time, tableID int) (bool, error)
	GetReservationHistory(ctx context.Context, username string) ([]*ReservationModel, error)
	CancelReservation(ctx context.Context, id int) error
	ListReservationByTable(ctx context.Context, tableID int) (*ReservationModel, error)
}
