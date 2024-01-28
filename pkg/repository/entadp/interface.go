package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, c *dto.User) error
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, id int, c *dto.User) error
	GetById(ctx context.Context, id int) (*dto.User, error)
	ListUser(ctx context.Context) ([]*ent.User, error)
}

type TablesRepository interface {
	CreateTable(ctx context.Context, c *dto.Tables) error
	DeleteTable(ctx context.Context, id int) error
	UpdateTable(ctx context.Context, id int, c *dto.Tables) error
	ListTable(ctx context.Context) ([]*ent.Tables, error)
}
