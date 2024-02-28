package menu

import (
	"context"
)

type ServiceMenu interface {
	CreateMenu(ctx context.Context, c *CreateMenuModel) error
	DeleteMenu(ctx context.Context, id int) error
	UpdateMenu(ctx context.Context, id int, c *UpdateMenuModel) error
	ListMenu(ctx context.Context) ([]*MenuModel, error)
	GetById(ctx context.Context, id int) (*MenuModel, error)
}
