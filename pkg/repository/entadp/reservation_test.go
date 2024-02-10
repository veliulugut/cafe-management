package entadp

import (
	"cafe-management/ent"
	"cafe-management/ent/enttest"
	"cafe-management/pkg/repository/dto"
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestReservation_CreateReservation(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewReservation(client)

	reservation := dto.Reservation{
		FullName:    "ufo361",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		TableID:     1,
		PhoneNumber: "331 3131 3311",
		Status:      "pending",
	}

	t.Run("CreateReservation", func(t *testing.T) {
		err := repo.CreateReservation(context.Background(), &reservation)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestReservation_DeleteReservation(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewReservation(client)

	reservation := dto.Reservation{
		FullName:    "ufo361",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		TableID:     1,
		PhoneNumber: "331 3131 3311",
		Status:      "pending",
	}

	err := repo.CreateReservation(context.Background(), &reservation)
	if err != nil {
		t.Error(err)
	}

	test := []struct {
		id          int
		description string
		expectedErr error
	}{
		{
			id:          1,
			description: "success",
			expectedErr: nil,
		},
		{
			id:          2,
			description: "fail",
			expectedErr: ErrorNotFound,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("DeleteReservation %s", tt.description), func(t *testing.T) {
			err := repo.DeleteReservation(context.Background(), tt.id)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func TestReservation_CheckAvailability(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewReservation(client)

	reservation := dto.Reservation{
		FullName:    "ufo361",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		TableID:     1,
		PhoneNumber: "31 31 31",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(3 * time.Hour),
		Status:      "pending",
	}

	t.Run("CreateReservation", func(t *testing.T) {
		err := repo.CreateReservation(context.Background(), &reservation)
		if err != nil {
			t.Error(err)
		}
	})

	test := []struct {
		description string
		startTime   time.Time
		endTime     time.Time
		tableID     int
		expected    bool
	}{
		{
			description: "available",
			startTime:   time.Now().Add(4 * time.Hour),
			endTime:     time.Now().Add(5 * time.Hour),
			tableID:     1,
			expected:    true, //available
		},
		{
			description: "not available",
			startTime:   time.Now().Add(1 * time.Hour),
			endTime:     time.Now().Add(2 * time.Hour),
			tableID:     1,
			expected:    false,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("CheckAvailability %s", tt.description), func(t *testing.T) {
			result, err := repo.CheckAvailability(context.Background(), tt.startTime, tt.endTime, tt.tableID)
			if err != nil {
				t.Errorf("unexpected error :%v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}

}

func TestReservation_GetReservationHistory(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewReservation(client)

	reservations := []dto.Reservation{
		{
			FullName:    "ufo361",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			TableID:     1,
			PhoneNumber: "331 3131 3311",
			Status:      "pending",
		},
		{
			FullName:    "ufo361",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			TableID:     2,
			PhoneNumber: "323 3131 3213",
			Status:      "confirmed",
		},
	}

	t.Run("CreateReservation", func(t *testing.T) {
		for _, r := range reservations {
			err := repo.CreateReservation(context.Background(), &r)
			if err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("GetReservationHistory", func(t *testing.T) {
		historyReservations, err := repo.GetReservationHistory(context.Background(), "ufo361")
		if err != nil {
			t.Error(err)
		}

		if len(historyReservations) != 2 {
			t.Errorf("expected 2 reservation, got %d", len(reservations))
		}

		if historyReservations[0].FullName != reservations[0].FullName {
			t.Errorf("expected full name %s, got %s", reservations[0].FullName, historyReservations[0].FullName)
		}
	})

}

func TestReservation_CancelReservation(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewReservation(client)

	reservation := dto.Reservation{
		FullName:    "ufo361",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		TableID:     1,
		PhoneNumber: "331 3131 3311",
		Status:      "pending",
	}

	err := repo.CreateReservation(context.Background(), &reservation)
	if err != nil {
		t.Error(err)
	}

	test := []struct {
		id          int
		description string
		expectedErr error
	}{
		{
			id:          1,
			description: "success",
			expectedErr: nil,
		},
		{
			id:          2,
			description: "fail (not found)",
			expectedErr: ErrorNotFound,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("CancelReservation %s", tt.description), func(t *testing.T) {
			err := repo.CancelReservation(context.Background(), tt.id)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				return
			}

			if tt.expectedErr == nil {
				// Check if the reservation is actually cancelled
				dbRes, err := repo.dbClient.Reservation.Get(context.Background(), tt.id)
				if err != nil {
					t.Errorf("failed to retrieve reservation: %v", err)
					return
				}

				if dbRes.Status != "cancelled" {
					t.Errorf("expected status to be 'cancelled', got %s", dbRes.Status)
				}
			}
		})
	}
}
