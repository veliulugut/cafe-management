package entadp_test

import (
	"cafe-management/ent"
	"cafe-management/ent/enttest"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestMenu_CreateMenu(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := entadp.NewMenuRepository(client)

	menu := dto.Menu{
		MenuID:       1,
		Name:         "water",
		Category:     "drink",
		Price:        100.0,
		Description:  "water",
		MenuImageUrl: "image",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("CreateMenu", func(t *testing.T) {
		err := repo.CreateMenu(context.Background(), &menu)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestMenu_DeleteMenu(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := entadp.NewMenuRepository(client)

	menu := dto.Menu{
		MenuID:       1,
		Name:         "water",
		Category:     "drink",
		Price:        100.0,
		Description:  "water",
		MenuImageUrl: "image",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("CreateMenu", func(t *testing.T) {
		err := repo.CreateMenu(context.Background(), &menu)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteMenu", func(t *testing.T) {
		if err := repo.DeleteMenu(context.Background(), 1); err != nil {
			t.Error(err)
		}
	})
}

func TestMenu_UpdateMenu(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := entadp.NewMenuRepository(client)

	menu := dto.Menu{
		MenuID:       1,
		Name:         "water",
		Category:     "drink",
		Price:        100.0,
		Description:  "water",
		MenuImageUrl: "image",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	updateMenu := dto.Menu{
		MenuID:       1,
		Name:         "orange juice",
		Category:     "drink",
		Price:        100.0,
		Description:  "water",
		MenuImageUrl: "image",
		UpdatedAt:    time.Now(),
	}

	t.Run("CreateMenu", func(t *testing.T) {
		err := repo.CreateMenu(context.Background(), &menu)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("UpdateMenu", func(t *testing.T) {
		if err := repo.UpdateMenu(context.Background(), 1, &updateMenu); err != nil {
			t.Error(err)
		}
	})
}

func TestMenu_GetById(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := entadp.NewMenuRepository(client)

	menu := dto.Menu{
		MenuID:       1,
		Name:         "water",
		Category:     "drink",
		Price:        100.0,
		Description:  "water",
		MenuImageUrl: "image",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("CreateMenu", func(t *testing.T) {
		err := repo.CreateMenu(context.Background(), &menu)
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

func TestMenu_ListMenu(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := entadp.NewMenuRepository(client)

	menu := dto.Menu{
		MenuID:       1,
		Name:         "water",
		Category:     "drink",
		Price:        100.0,
		Description:  "water",
		MenuImageUrl: "image",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("CreateMenu", func(t *testing.T) {
		err := repo.CreateMenu(context.Background(), &menu)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("ListMenu", func(t *testing.T) {
		_, err := repo.ListMenu(context.Background())
		if err != nil {
			t.Error(err)
		}
	})
}
