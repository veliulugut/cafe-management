package user

import (
	"context"
)

type ServiceUser interface {
	CreateUser(ctx context.Context, c *CreateUserModel) error
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, id int, c *UpdateUserModel) error
	GetById(ctx context.Context, id int) (*UserModel, error)
	ListUser(ctx context.Context) ([]*UserModel, error)
	GetByUserName(ctx context.Context, name string) (*UserModel, error)
}
