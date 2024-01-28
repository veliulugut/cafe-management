package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"context"
	"fmt"
	"time"
)

var _ TablesRepository = (*Table)(nil)

func NewTablesRepository(db *ent.Client) *Table {
	return &Table{
		dbClient: db,
	}
}

type Table struct {
	dbClient *ent.Client
}

func (t *Table) CreateTable(ctx context.Context, c *dto.Tables) error {
	_, err := t.dbClient.Tables.Create().
		SetNumberOfGuests(c.NumberOfGuests).
		SetTableNumber(c.TableNumber).
		SetDescription(c.Description).
		SetTableType(c.TableType).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *Table) DeleteTable(ctx context.Context, id int) error {
	if err := t.dbClient.Tables.DeleteOneID(id).Exec(ctx); err != nil {
		return fmt.Errorf("delete table/repository:%w", err)
	}

	return nil
}

func (t *Table) UpdateTable(ctx context.Context, id int, c *dto.Tables) error {
	_, err := t.dbClient.Tables.UpdateOneID(id).
		SetNumberOfGuests(c.NumberOfGuests).
		SetTableNumber(c.TableNumber).
		SetDescription(c.Description).
		SetTableType(c.TableType).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("update table/repository:%w", err)
	}

	return nil
}

func (t *Table) ListTable(ctx context.Context) ([]*ent.Tables, error) {
	var (
		tables []*ent.Tables
		err    error
	)

	tables, err = t.dbClient.Tables.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("list table/repository:%w", err)
	}

	return tables, nil
}
