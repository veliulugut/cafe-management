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

func TestProduct_Create(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewProductRepository(client)

	product := []dto.Product{
		{
			ProductName: "water",
			Description: "water",
			Price:       100,
			Quantity:    100,
			ProductType: "drink",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ProductName: "orange juice",
			Description: "orange juice",
			Price:       3131,
			Quantity:    3131,
			ProductType: "drink",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	t.Run("CreateProduct", func(t *testing.T) {
		for _, p := range product {
			err := repo.CreateProduct(context.Background(), &p)
			if err != nil {
				t.Error(err)
			}
		}
	})

}

func TestProduct_ListProduct(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewProductRepository(client)

	product := []dto.Product{
		{
			ProductName: "water",
			Description: "water",
			Price:       100,
			Quantity:    100,
			ProductType: "drink",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ProductName: "orange juice",
			Description: "orange juice",
			Price:       3131,
			Quantity:    3131,
			ProductType: "drink",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	t.Run("CreateProduct", func(t *testing.T) {
		for _, p := range product {
			err := repo.CreateProduct(context.Background(), &p)
			if err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("ListProduct", func(t *testing.T) {
		list, err := repo.ListProduct(context.Background())
		if err != nil {
			t.Error(err)
		}

		expectedCount := 2
		if len(list) != expectedCount {
			t.Errorf("expected %d products, got %d", expectedCount, len(list))
		}

		for i, p := range product {
			if p.ProductName != list[i].ProductName || p.Price != list[i].Price || p.ProductType != list[i].ProductType {
				t.Errorf("expected %v, got %v", p, list[i])
			}
		}
	})
}

func TestProduct_UpdateProduct(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewProductRepository(client)

	product := dto.Product{
		ProductName: "water",
		Description: "water",
		Price:       100,
		Quantity:    100,
		ProductType: "drink",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	updateProduct := dto.Product{
		ProductName: "milk",
		Description: "milk",
		Price:       100,
		Quantity:    100,
		ProductType: "drink",
		UpdatedAt:   time.Now(),
	}

	t.Run("CreateProduct", func(t *testing.T) {
		err := repo.CreateProduct(context.Background(), &product)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("UpdateProduct", func(t *testing.T) {
		err := repo.UpdateProduct(context.Background(), 1, &updateProduct)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestProduct_DeleteProduct(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewProductRepository(client)

	product := dto.Product{
		ProductName: "water",
		Description: "water",
		Price:       100,
		Quantity:    100,
		ProductType: "drink",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("CreateProduct", func(t *testing.T) {
		err := repo.CreateProduct(context.Background(), &product)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		if err := repo.DeleteProduct(context.Background(), 1); err != nil {
			t.Error(err)
		}
	})

}

func TestProduct_GetById(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewProductRepository(client)

	product := dto.Product{
		ProductName: "water",
		Description: "water",
		Price:       100,
		Quantity:    100,
		ProductType: "drink",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("CreateProduct", func(t *testing.T) {
		err := repo.CreateProduct(context.Background(), &product)
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
