package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/helper"
	"context"
	"fmt"
	"time"
)

var _ MenuRepository = (*Menu)(nil)

func NewMenuRepository(d *ent.Client) *Menu {
	return &Menu{
		dbClient: d,
	}
}

type Menu struct {
	dbClient *ent.Client
}

func (m *Menu) CreateMenu(ctx context.Context, c *dto.Menu) error {
	_, err := m.dbClient.Menu.Create().
		SetMenuID(c.MenuID).
		SetName(c.Name).
		SetCategory(c.Category).
		SetPrice(c.Price).
		SetDescription(c.Description).
		SetMenuImageURL(c.MenuImageUrl).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return fmt.Errorf("menu / create menu :%w", err)
	}

	return nil
}

func (m *Menu) DeleteMenu(ctx context.Context, id int) error {
	err := m.dbClient.Menu.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("menu / delete menu :%w", err)
	}

	return nil
}

func (m *Menu) GetById(ctx context.Context, id int) (*dto.Menu, error) {
	var (
		menu *ent.Menu
		err  error
	)

	if menu, err = m.dbClient.Menu.Get(ctx, id); err != nil {
		return nil, fmt.Errorf("menu / get menu :%w", err)
	}

	return helper.DbMenuTo(menu), nil
}

func (m *Menu) ListMenu(ctx context.Context) ([]*ent.Menu, error) {
	var (
		menus []*ent.Menu
		err   error
	)

	menus, err = m.dbClient.Menu.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("menu / list menu :%w", err)
	}

	return menus, nil
}

func (m *Menu) UpdateMenu(ctx context.Context, id int, c *dto.Menu) error {
	err := m.dbClient.Menu.UpdateOneID(id).SetMenuID(c.MenuID).
		SetName(c.Name).
		SetCategory(c.Category).
		SetPrice(c.Price).
		SetDescription(c.Description).
		SetMenuImageURL(c.MenuImageUrl).
		SetUpdatedAt(time.Now()).Exec(ctx)
	if err != nil {
		return fmt.Errorf("menu / update menu :%w", err)
	}

	return nil
}
