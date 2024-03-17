// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/predicate"
	"cafe-management/ent/qrcode"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QrCodeDelete is the builder for deleting a QrCode entity.
type QrCodeDelete struct {
	config
	hooks    []Hook
	mutation *QrCodeMutation
}

// Where appends a list predicates to the QrCodeDelete builder.
func (qcd *QrCodeDelete) Where(ps ...predicate.QrCode) *QrCodeDelete {
	qcd.mutation.Where(ps...)
	return qcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qcd *QrCodeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, qcd.sqlExec, qcd.mutation, qcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (qcd *QrCodeDelete) ExecX(ctx context.Context) int {
	n, err := qcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qcd *QrCodeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(qrcode.Table, sqlgraph.NewFieldSpec(qrcode.FieldID, field.TypeInt))
	if ps := qcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, qcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	qcd.mutation.done = true
	return affected, err
}

// QrCodeDeleteOne is the builder for deleting a single QrCode entity.
type QrCodeDeleteOne struct {
	qcd *QrCodeDelete
}

// Where appends a list predicates to the QrCodeDelete builder.
func (qcdo *QrCodeDeleteOne) Where(ps ...predicate.QrCode) *QrCodeDeleteOne {
	qcdo.qcd.mutation.Where(ps...)
	return qcdo
}

// Exec executes the deletion query.
func (qcdo *QrCodeDeleteOne) Exec(ctx context.Context) error {
	n, err := qcdo.qcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{qrcode.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qcdo *QrCodeDeleteOne) ExecX(ctx context.Context) {
	if err := qcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
