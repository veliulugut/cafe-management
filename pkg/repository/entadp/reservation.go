package entadp

import (
	"cafe-management/ent"
	"cafe-management/ent/reservation"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/helper"
	"context"
	"fmt"
	"time"
)

/*
FullName    string    `json:"full_name"`

	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	TableID     int       `json:"table_id"`
	PhoneNumber string    `json:"phone_number"`
*/
var _ ReservationRepository = (*Reservation)(nil)

func NewReservation(dbClient *ent.Client) *Reservation {
	return &Reservation{
		dbClient: dbClient,
	}
}

type Reservation struct {
	dbClient *ent.Client
}

func (r *Reservation) CheckAvailability(ctx context.Context, startTime, endTime time.Time, tableID int) (bool, error) {
	reservations, err := r.dbClient.Reservation.Query().
		Where(
			reservation.TableID(tableID),
			reservation.StartTimeLTE(startTime),
			reservation.EndTimeGTE(endTime),
		).Count(ctx)

	if err != nil {
		return false, fmt.Errorf("reservation / check availability :%w", err)
	}

	return reservations == 0, nil
}

func (r *Reservation) CreateReservation(ctx context.Context, c *dto.Reservation) error {
	err := r.dbClient.Reservation.Create().
		SetFullName(c.FullName).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetTableID(c.TableID).
		SetStartTime(c.StartTime).
		SetEndTime(c.EndTime).
		SetPhoneNumber(c.PhoneNumber).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("reservation / create reservation :%w", err)
	}

	return nil
}

func (r *Reservation) DeleteReservation(ctx context.Context, id int) error {
	err := r.dbClient.Reservation.DeleteOneID(id).Exec(ctx)
	if ent.IsNotFound(err) {
		return fmt.Errorf("reservation / delete reservation :%w", ErrorNotFound)
	}

	return nil
}

func (r *Reservation) UpdateReservation(ctx context.Context, id int, c *dto.Reservation) error {
	err := r.dbClient.Reservation.UpdateOneID(id).SetFullName(c.FullName).SetUpdatedAt(time.Now()).SetTableID(c.TableID).SetPhoneNumber(c.PhoneNumber).Exec(ctx)
	if err != nil {
		return fmt.Errorf("reservation / update reservation :%w", err)
	}

	return nil
}

func (r *Reservation) ListReservation(ctx context.Context, date time.Time) (*dto.Reservation, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Reservation) ListReservationByTable(ctx context.Context, tableID int) (*dto.Reservation, error) {
	var (
		db  *ent.Reservation
		err error
	)

	if db, err = r.dbClient.Reservation.Get(ctx, tableID); err != nil {
		return nil, fmt.Errorf("reservation / list reservation by table :%w", err)
	}

	return helper.DbReservationTo(db), nil
}

func (r *Reservation) ListReservationByUser(ctx context.Context, userID int) (*dto.Reservation, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Reservation) GetByNameReservation(ctx context.Context, name string) (*dto.Reservation, error) {
	var (
		db  *ent.Reservation
		err error
	)

	if db, err = r.dbClient.Reservation.Query().Where(reservation.FullName(name)).First(ctx); err != nil {
		return nil, fmt.Errorf("reservation / get reservation by name :%w", err)
	}

	return helper.DbReservationTo(db), nil
}
