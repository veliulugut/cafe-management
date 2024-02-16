package tables

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"
)

var _ ServiceTables = (*Tables)(nil)

func New(repo entadp.RepositoryInterface) *Tables {
	return &Tables{
		repo: repo,
	}
}

type Tables struct {
	repo entadp.RepositoryInterface
}

func (t *Tables) CreateTable(ctx context.Context, c *CreateTableModel) error {
	tableDTO := dto.Tables{
		NumberOfGuests: c.NumberOfGuests,
		TableNumber:    c.TableNumber,
		Description:    c.Description,
		TableType:      c.TableType,
	}

	if err := t.repo.Table().CreateTable(ctx, &tableDTO); err != nil {
		return fmt.Errorf("table srv / create table : %w", err)
	}

	return nil
}

func (t *Tables) DeleteTable(ctx context.Context, id int) error {
	if err := t.repo.Table().DeleteTable(ctx, id); err != nil {
		return fmt.Errorf("table srv / delete table : %w", err)
	}

	return nil
}

func (t *Tables) ListTable(ctx context.Context) ([]*TableModel, error) {
	var (
		tables []*ent.Tables
		err    error
	)

	if tables, err = t.repo.Table().ListTable(ctx); err != nil {
		return nil, fmt.Errorf("table srv / list table : %w", err)
	}

	var tableModels []*TableModel
	for _, table := range tables {
		tableModels = append(tableModels, &TableModel{
			NumberOfGuests: table.NumberOfGuests,
			TableNumber:    table.TableNumber,
			Description:    table.Description,
			TableType:      table.TableType,
		})
	}

	return tableModels, nil
}

func (t *Tables) UpdateTable(ctx context.Context, id int, c *UpdateTableModel) error {
	tableDTO := dto.Tables{
		NumberOfGuests: c.NumberOfGuests,
		TableNumber:    c.TableNumber,
		Description:    c.Description,
		TableType:      c.TableType,
	}

	if err := t.repo.Table().UpdateTable(ctx, id, &tableDTO); err != nil {
		return fmt.Errorf("table srv / update table : %w", err)
	}

	return nil
}
