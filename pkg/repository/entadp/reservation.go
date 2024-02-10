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

func (r *Reservation) CreateReservation(ctx context.Context, c *dto.Reservation) error {
	err := r.dbClient.Reservation.Create().
		SetFullName(c.FullName).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetTableID(c.TableID).
		SetStartTime(c.StartTime).
		SetEndTime(c.EndTime).
		SetPhoneNumber(c.PhoneNumber).
		SetStatus(c.Status).
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
	err := r.dbClient.Reservation.UpdateOneID(id).
		SetFullName(c.FullName).
		SetUpdatedAt(time.Now()).
		SetTableID(c.TableID).
		SetPhoneNumber(c.PhoneNumber).
		SetStatus(c.Status).
		Exec(ctx)
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

	if db, err = r.dbClient.Reservation.Query().Where(reservation.FullNameContains(name)).First(ctx); err != nil {
		return nil, fmt.Errorf("reservation / get reservation by name :%w", err)
	}

	return helper.DbReservationTo(db), nil
}

func (r *Reservation) CancelReservation(ctx context.Context, id int) error {
	tx, err := r.dbClient.Tx(ctx)
	if err != nil {
		return fmt.Errorf("reservation / cancel reservation: starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback() // Geri alma işlemi
		} else {
			tx.Commit() // Onaylama işlemi
		}
	}()

	// İptal edilecek rezervasyonu db'den alma
	dbRes, err := tx.Reservation.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrorNotFound
		}
		return fmt.Errorf("reservation / cancel reservation: fetching reservation: %w", err)
	}

	// Rezervasyon durumunu "iptal edildi" olarak güncelleme
	if err := tx.Reservation.UpdateOne(dbRes).SetStatus("cancelled").Exec(ctx); err != nil {
		return fmt.Errorf("reservation / cancel reservation: updating reservation: %w", err)
	}

	return nil
}

// retrieves the booking history of a specific user
func (r *Reservation) GetReservationHistory(ctx context.Context, username string) ([]*ent.Reservation, error) {
	reservations, err := r.dbClient.Reservation.Query().
		Where(reservation.FullNameEQ(username)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reservation history: %w", err)
	}

	return reservations, nil

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
