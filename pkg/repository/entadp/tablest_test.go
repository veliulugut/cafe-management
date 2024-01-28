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

func TestTable_CreateTable(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	tableRepo := NewTablesRepository(client)

	tables := []dto.Tables{
		{
			NumberOfGuests: 3,
			TableNumber:    1,
			Description:    "test",
			TableType:      "vip",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	for _, u := range tables {
		err := tableRepo.CreateTable(context.Background(), &u)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}
}

func TestTable_DeleteTable(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	tableRepo := NewTablesRepository(client)

	tables := []dto.Tables{
		{
			NumberOfGuests: 3,
			TableNumber:    1,
			Description:    "test",
			TableType:      "vip",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	for _, u := range tables {
		err := tableRepo.CreateTable(context.Background(), &u)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}

}
