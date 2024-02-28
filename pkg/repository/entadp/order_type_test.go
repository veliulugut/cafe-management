package entadp

import (
	"cafe-management/ent"
	"cafe-management/ent/enttest"
	"cafe-management/pkg/repository/dto"
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateOrderType(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewOrdersTypeRepository(client)

	orderType := dto.OrderType{
		ID:   1,
		Name: "Food",
	}

	t.Run("CreateOrderType", func(t *testing.T) {
		err := repo.CreateOrderType(context.Background(), &orderType)
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_ListOrderType(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewOrdersTypeRepository(client)

	orderType := dto.OrderType{
		ID:   1,
		Name: "Food",
	}

	t.Run("CreateOrderType", func(t *testing.T) {
		err := repo.CreateOrderType(context.Background(), &orderType)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("ListOrderType", func(t *testing.T) {
		_, err := repo.ListOrderType(context.Background())
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_UpdateOrderType(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewOrdersTypeRepository(client)

	orderType := dto.OrderType{
		ID:   1,
		Name: "Food",
	}

	orderUpdate := dto.OrderType{
		ID:   1,
		Name: "drink",
	}

	t.Run("CreateOrderType", func(t *testing.T) {
		err := repo.CreateOrderType(context.Background(), &orderType)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("UpdateOrderType", func(t *testing.T) {
		err := repo.UpdateOrderType(context.Background(), 1, &orderUpdate)
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_DeleteOrderType(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewOrdersTypeRepository(client)

	orderType := dto.OrderType{
		ID:   1,
		Name: "Food",
	}

	t.Run("CreateOrderType", func(t *testing.T) {
		err := repo.CreateOrderType(context.Background(), &orderType)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteOrderType", func(t *testing.T) {
		err := repo.DeleteOrderType(context.Background(), 1)
		if err != nil {
			t.Error(err)
		}
	})
}
