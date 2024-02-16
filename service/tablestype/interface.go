package tablestype

import "context"

type ServiceTablestype interface {
	CreateTableType(ctx context.Context, c *CreateTableTypeModel) error
	DeleteTableType(ctx context.Context, id int) error
	UpdateTableType(ctx context.Context, id int, c *UpdateTableTypeModel) error
}
