package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"context"
	"fmt"
)

var _ TablesTypeRepostiory = (*TableType)(nil)

func NewTableTypeRepository(db *ent.Client) *TableType {
	return &TableType{
		dbClient: db,
	}
}

type TableType struct {
	dbClient *ent.Client
}

func (t *TableType) CreateTableType(ctx context.Context, c *dto.TablesType) error {
	_, err := t.dbClient.Tables_type.
		Create().
		SetTablesID(c.TableID).
		SetName(c.TableName).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("table type create :%w", err)
	}

	return nil
}

func (t *TableType) DeleteTableType(ctx context.Context, id int) error {
	err := t.dbClient.Tables_type.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("tables type / delete table type :%w", err)
	}

	return nil
}

func (t *TableType) UpdateTableType(ctx context.Context, id int, c *dto.TablesType) error {
	err := t.dbClient.Tables_type.UpdateOneID(id).SetTablesID(c.TableID).SetName(c.TableName).Exec(ctx)
	if err != nil {
		return fmt.Errorf("tables type / update table type :%w", err)
	}

	return nil
}
