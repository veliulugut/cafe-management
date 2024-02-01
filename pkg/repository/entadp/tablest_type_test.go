package entadp

import (
	"cafe-management/ent"
	"cafe-management/ent/enttest"
	"cafe-management/pkg/repository/dto"
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestTableType_CreateTableType(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewTableTypeRepository(client)

	tables := dto.TablesType{
		TableID:   1,
		TableName: "vip",
	}

	err := repo.CreateTableType(context.Background(), &tables)
	if err != nil {
		t.Error(err)
	}
}

func TestTableType_DeleteTableType(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewTableTypeRepository(client)

	tables := dto.TablesType{
		TableID:   1,
		TableName: "default",
	}

	t.Run("CreateTableType", func(t *testing.T) {
		err := repo.CreateTableType(context.Background(), &tables)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteTableType", func(t *testing.T) {
		err := repo.DeleteTableType(context.Background(), 1)
		if err != nil {
			t.Error(err)
		}
	})
}
