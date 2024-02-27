package entadp

import (
	"cafe-management/ent"
	"cafe-management/ent/enttest"
	"cafe-management/pkg/repository/dto"
	"context"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*
		OrderID   int       `json:"order_id"`
		TableID   int       `json:"table_id"`
		UserId    int       `json:"user_id"`
		OrderType int       `json:"order_type"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
*/
func TestOrder_CreateOrder(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewOrderRepository(client)

	orderDTO := dto.Order{
		OrderID:   1,
		TableID:   1,
		UserId:    1,
		OrderType: 1,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("CreateOrder", func(t *testing.T) {
		err := repo.CreateOrder(context.Background(), &orderDTO)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestOrder_DeleteOrder(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewOrderRepository(client)

	orderDTO := dto.Order{
		OrderID:   1,
		TableID:   1,
		UserId:    1,
		OrderType: 1,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("CreateOrder", func(t *testing.T) {
		err := repo.CreateOrder(context.Background(), &orderDTO)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteOrder", func(t *testing.T) {
		err := repo.DeleteOrder(context.Background(), 1)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestOrder_GetById(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewOrderRepository(client)

	orderDTO := dto.Order{
		OrderID:   1,
		TableID:   1,
		UserId:    1,
		OrderType: 1,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("CreateOrder", func(t *testing.T) {
		err := repo.CreateOrder(context.Background(), &orderDTO)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("GetById", func(t *testing.T) {
		_, err := repo.GetById(context.Background(), 1)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestOrder_ListOrder(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewOrderRepository(client)

	orderDTO := dto.Order{
		OrderID:   1,
		TableID:   1,
		UserId:    1,
		OrderType: 1,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("CreateOrder", func(t *testing.T) {
		err := repo.CreateOrder(context.Background(), &orderDTO)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("ListOrder", func(t *testing.T) {
		_, err := repo.ListOrder(context.Background())
		if err != nil {
			t.Error(err)
		}
	})
}

func TestOrder_UpdateOrder(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewOrderRepository(client)

	orderDTO := dto.Order{
		OrderID:   1,
		TableID:   1,
		UserId:    1,
		OrderType: 1,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	updateOrder := dto.Order{
		OrderID:   1,
		TableID:   1,
		UserId:    1,
		OrderType: 1,
		Status:    "ready",
		UpdatedAt: time.Now(),
	}

	t.Run("CreateOrder", func(t *testing.T) {
		err := repo.CreateOrder(context.Background(), &orderDTO)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("UpdateOrder", func(t *testing.T) {
		err := repo.UpdateOrder(context.Background(), 1, &updateOrder)
		if err != nil {
			t.Error(err)
		}
	})
}
