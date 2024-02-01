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
		EndTime:     time.Now(),
	}

	t.Run("CreateReservation", func(t *testing.T) {
		err := repo.CreateReservation(context.Background(), &reservation)
		if err != nil {
			t.Error(err)
		}
	})

	test := []struct {
		description string
		datetime    time.Time
		tableID     int
		expected    bool
	}{
		{
			description: "available",
			datetime:    time.Now().Add(2 * time.Hour),
			tableID:     1,
			expected:    true, //available
		},
		{
			description: "not available",
			datetime:    time.Now().Add(2 * time.Hour),
			tableID:     1,
			expected:    false,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("CheckAvailability %s", tt.description), func(t *testing.T) {
			t.Run(fmt.Sprintf("DeleteReservation %s", tt.description), func(t *testing.T) {
				err := repo.DeleteReservation(context.Background(), 1)
				if err != nil {
					t.Error(err)
				}
			})

			t.Run(fmt.Sprintf("CreateReservation %s", tt.description), func(t *testing.T) {
				err := repo.CreateReservation(context.Background(), &reservation)
				if err != nil {
					t.Error(err)
				}
			})

			result, err := repo.CheckAvailability(context.Background(), tt.datetime, tt.tableID)
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
