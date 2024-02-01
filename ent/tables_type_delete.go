// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/predicate"
	"cafe-management/ent/tables_type"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TablesTypeDelete is the builder for deleting a Tables_type entity.
type TablesTypeDelete struct {
	config
	hooks    []Hook
	mutation *TablesTypeMutation
}

// Where appends a list predicates to the TablesTypeDelete builder.
func (ttd *TablesTypeDelete) Where(ps ...predicate.Tables_type) *TablesTypeDelete {
	ttd.mutation.Where(ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *TablesTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ttd.sqlExec, ttd.mutation, ttd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *TablesTypeDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *TablesTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tables_type.Table, sqlgraph.NewFieldSpec(tables_type.FieldID, field.TypeInt))
	if ps := ttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ttd.mutation.done = true
	return affected, err
}

// TablesTypeDeleteOne is the builder for deleting a single Tables_type entity.
type TablesTypeDeleteOne struct {
	ttd *TablesTypeDelete
}

// Where appends a list predicates to the TablesTypeDelete builder.
func (ttdo *TablesTypeDeleteOne) Where(ps ...predicate.Tables_type) *TablesTypeDeleteOne {
	ttdo.ttd.mutation.Where(ps...)
	return ttdo
}

// Exec executes the deletion query.
func (ttdo *TablesTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tables_type.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *TablesTypeDeleteOne) ExecX(ctx context.Context) {
	if err := ttdo.Exec(ctx); err != nil {
		panic(err)
	}
}
