package tables

import "context"

type ServiceTables interface {
	CreateTable(ctx context.Context, c *CreateTableModel) error
	UpdateTable(ctx context.Context, id int, c *UpdateTableModel) error
	DeleteTable(ctx context.Context, id int) error
	ListTable(ctx context.Context) ([]*TableModel, error)
}
