package reservation

import (
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"
	"time"
)

var _ ReservationService = (*Reservation)(nil)

func New(repo entadp.RepositoryInterface) *Reservation {
	return &Reservation{
		repo: repo,
	}
}

type Reservation struct {
	repo entadp.RepositoryInterface
}

func (r *Reservation) CreateReservation(ctx context.Context, c *CreateReservationModel) error {
	reservationDTO := dto.Reservation{
		FullName:    c.FullName,
		TableID:     c.TableID,
		PhoneNumber: c.PhoneNumber,
		StartTime:   c.StartTime,
		EndTime:     c.EndTime,
		Status:      c.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := r.repo.Reservation().CreateReservation(ctx, &reservationDTO); err != nil {
		return fmt.Errorf("reservation / create reservation :%w", err)
	}

	return nil
}

func (r *Reservation) CancelReservation(ctx context.Context, id int) error {
	if err := r.repo.Reservation().CancelReservation(ctx, id); err != nil {
		return fmt.Errorf("reservation / cancel reservation :%w", err)
	}

	return nil
}

func (r *Reservation) CheckAvailability(ctx context.Context, startTime time.Time, endTime time.Time, tableID int) (bool, error) {
	return r.repo.Reservation().CheckAvailability(ctx, startTime, endTime, tableID)
}

func (r *Reservation) DeleteReservation(ctx context.Context, id int) error {
	if err := r.repo.Reservation().DeleteReservation(ctx, id); err != nil {
		return fmt.Errorf("reservation / delete reservation :%w", err)
	}

	return nil
}

func (r *Reservation) GetByNameReservation(ctx context.Context, name string) (*ReservationModel, error) {
	reservationDTO, err := r.repo.Reservation().GetByNameReservation(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("reservation / get reservation by name :%w", err)
	}

	return &ReservationModel{
		FullName:    reservationDTO.FullName,
		TableID:     reservationDTO.TableID,
		StartTime:   reservationDTO.StartTime,
		EndTime:     reservationDTO.EndTime,
		PhoneNumber: reservationDTO.PhoneNumber,
		Status:      reservationDTO.Status,
		CreatedAt:   reservationDTO.CreatedAt,
		UpdatedAt:   reservationDTO.UpdatedAt,
	}, nil
}

func (r *Reservation) GetReservationHistory(ctx context.Context, username string) ([]*ReservationModel, error) {
	reservations, err := r.repo.Reservation().GetReservationHistory(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("reservation / get reservation history :%w", err)
	}

	var reservationModels []*ReservationModel
	for _, res := range reservations {
		reservationModels = append(reservationModels, &ReservationModel{
			FullName:    res.FullName,
			TableID:     res.TableID,
			StartTime:   res.StartTime,
			EndTime:     res.EndTime,
			PhoneNumber: res.PhoneNumber,
			Status:      res.Status,
			CreatedAt:   res.CreatedAt,
			UpdatedAt:   res.UpdatedAt,
		})
	}

	return reservationModels, nil
}

func (r *Reservation) ListReservationByTable(ctx context.Context, tableID int) (*ReservationModel, error) {
	reservationDTO, err := r.repo.Reservation().ListReservationByTable(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("reservation / list reservation by table :%w", err)
	}

	return &ReservationModel{
		FullName:    reservationDTO.FullName,
		TableID:     reservationDTO.TableID,
		StartTime:   reservationDTO.StartTime,
		EndTime:     reservationDTO.EndTime,
		PhoneNumber: reservationDTO.PhoneNumber,
		Status:      reservationDTO.Status,
		CreatedAt:   reservationDTO.CreatedAt,
		UpdatedAt:   reservationDTO.UpdatedAt,
	}, nil
}

func (r *Reservation) UpdateReservation(ctx context.Context, id int, c *UpdateReservationModel) error {
	updateReservationDTO := dto.Reservation{
		FullName:    c.FullName,
		TableID:     c.TableID,
		PhoneNumber: c.PhoneNumber,
		StartTime:   c.StartTime,
		EndTime:     c.EndTime,
		Status:      c.Status,
		UpdatedAt:   time.Now(),
	}

	if err := r.repo.Reservation().UpdateReservation(ctx, id, &updateReservationDTO); err != nil {
		return fmt.Errorf("reservation / update reservation :%w", err)
	}

	return nil
}
