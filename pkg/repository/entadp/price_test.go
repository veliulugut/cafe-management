package entadp

import (
	"cafe-management/ent"
	"cafe-management/ent/enttest"
	"cafe-management/pkg/repository/dto"
	"context"
	"errors"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestPrice_CreatePrice(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewPriceRepository(client)

	price := dto.Price{
		PriceName: "water",
		Price:     100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("CreatePrice", func(t *testing.T) {
		err := repo.CreatePrice(context.Background(), &price)
		if err != nil {
			t.Error(err)
		}
	})

}

func TestPrice_DeletePrice(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewPriceRepository(client)

	price := dto.Price{
		PriceName: "water",
		Price:     100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("CreatePrice", func(t *testing.T) {
		err := repo.CreatePrice(context.Background(), &price)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeletePrice", func(t *testing.T) {
		if err := repo.DeletePrice(context.Background(), 1); err != nil {
			t.Error(err)
		}
	})
}

func TestPrice_UpdatePrice(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewPriceRepository(client)

	price := dto.Price{
		PriceName: "water",
		Price:     100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("CreatePrice", func(t *testing.T) {
		err := repo.CreatePrice(context.Background(), &price)
		if err != nil {
			t.Error(err)
		}
	})

	test := []struct {
		description string
		price_id    int
		expectedErr error
	}{
		{
			description: "success",
			price_id:    1,
			expectedErr: nil,
		},
		{
			description: "fail",
			price_id:    2,
			expectedErr: ErrorNotFound,
		},
	}

	for _, tt := range test {
		t.Run(tt.description, func(t *testing.T) {
			err := repo.UpdatePrice(context.Background(), tt.price_id, &price)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}

}

func TestPrice_ListPrice(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewPriceRepository(client)

	prices := []dto.Price{
		{
			PriceName: "water",
			Price:     100,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			PriceName: "orange juice",
			Price:     3131,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("CreatePrice", func(t *testing.T) {
		for _, p := range prices {
			err := repo.CreatePrice(context.Background(), &p)
			if err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("ListPrice", func(t *testing.T) {
		listedPrices, err := repo.ListPrice(context.Background())
		if err != nil {
			t.Error(err)
		}

		expectedCount := 2
		if len(listedPrices) != expectedCount {
			t.Errorf("expected %d prices, got %d", expectedCount, len(listedPrices))
		}

		for i, expected := range prices {
			if expected.PriceName != listedPrices[i].PriceName || expected.Price != listedPrices[i].Price {
				t.Errorf("expected %v, got %v", expected, listedPrices[i])
			}
		}
	})

}
