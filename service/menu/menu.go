package menu

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"
	"time"
)

var _ ServiceMenu = (*Menu)(nil)

func New(repo entadp.RepositoryInterface) *Menu {
	return &Menu{
		repo: repo,
	}
}

type Menu struct {
	repo entadp.RepositoryInterface
}

func (m *Menu) CreateMenu(ctx context.Context, c *CreateMenuModel) error {
	menuDTO := dto.Menu{
		MenuID:       c.MenuID,
		Name:         c.Name,
		Category:     c.Category,
		Price:        c.Price,
		Description:  c.Description,
		CreatedAt:    c.CreatedAt,
		MenuImageUrl: c.MenuImageUrl,
		UpdatedAt:    c.UpdatedAt,
	}

	if err := m.repo.Menu().CreateMenu(ctx, &menuDTO); err != nil {
		return fmt.Errorf("menu srv / create menu : %w", err)
	}

	return nil
}

func (m *Menu) DeleteMenu(ctx context.Context, id int) error {
	if err := m.repo.Menu().DeleteMenu(ctx, id); err != nil {
		return fmt.Errorf("menu srv / delete menu : %w", err)
	}

	return nil
}

func (m *Menu) GetById(ctx context.Context, id int) (*MenuModel, error) {
	var (
		menu *dto.Menu
		err  error
	)

	if menu, err = m.repo.Menu().GetById(ctx, id); err != nil {
		return nil, fmt.Errorf("menu srv / get menu : %w", err)
	}

	return &MenuModel{
		MenuID:       menu.MenuID,
		Name:         menu.Name,
		Category:     menu.Category,
		Price:        menu.Price,
		Description:  menu.Description,
		CreatedAt:    time.Now(),
		MenuImageUrl: menu.MenuImageUrl,
		UpdatedAt:    time.Now(),
	}, nil
}

func (m *Menu) ListMenu(ctx context.Context) ([]*MenuModel, error) {
	var (
		menus []*ent.Menu
		err   error
	)

	if menus, err = m.repo.Menu().ListMenu(ctx); err != nil {
		return nil, fmt.Errorf("menu srv / list menu : %w", err)
	}

	var menuModels []*MenuModel
	for _, menu := range menus {
		menuModels = append(menuModels, &MenuModel{
			MenuID:       menu.MenuID,
			Name:         menu.Name,
			Category:     menu.Category,
			Price:        menu.Price,
			Description:  menu.Description,
			CreatedAt:    menu.CreatedAt,
			MenuImageUrl: menu.MenuImageURL,
			UpdatedAt:    menu.UpdatedAt,
		})
	}

	return menuModels, nil
}

func (m *Menu) UpdateMenu(ctx context.Context, id int, c *UpdateMenuModel) error {
	updateMenuDTO := dto.Menu{
		MenuID:       c.MenuID,
		Name:         c.Name,
		Category:     c.Category,
		Price:        c.Price,
		Description:  c.Description,
		MenuImageUrl: c.MenuImageUrl,
		UpdatedAt:    c.UpdatedAt,
	}

	if err := m.repo.Menu().UpdateMenu(ctx, id, &updateMenuDTO); err != nil {
		return fmt.Errorf("menu srv / update menu : %w", err)
	}

	return nil
}
