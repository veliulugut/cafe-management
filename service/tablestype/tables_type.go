package tablestype

import (
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"
)

var _ ServiceTablestype = (*Tablestype)(nil)

func New(repo entadp.RepositoryInterface) *Tablestype {
	return &Tablestype{
		repo: repo,
	}
}

type Tablestype struct {
	repo entadp.RepositoryInterface
}

func (t *Tablestype) CreateTableType(ctx context.Context, c *CreateTableTypeModel) error {
	tablesTypeDto := dto.TablesType{
		TableID:   c.TablesID,
		TableName: c.Name,
	}

	if err := t.repo.TableTypes().CreateTableType(ctx, &tablesTypeDto); err != nil {
		return fmt.Errorf("table type srv / create table type : %w", err)
	}

	return nil

}

func (t *Tablestype) DeleteTableType(ctx context.Context, id int) error {
	if err := t.repo.TableTypes().DeleteTableType(ctx, id); err != nil {
		return fmt.Errorf("table type srv / delete table type : %w", err)
	}

	return nil
}

func (t *Tablestype) UpdateTableType(ctx context.Context, id int, c *UpdateTableTypeModel) error {
	tablesTypeDto := dto.TablesType{
		TableID:   c.TablesID,
		TableName: c.Name,
	}

	if err := t.repo.TableTypes().UpdateTableType(ctx, id, &tablesTypeDto); err != nil {
		return fmt.Errorf("table type srv / update table type : %w", err)
	}

	return nil
}
