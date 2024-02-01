// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/predicate"
	"cafe-management/ent/reservation"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReservationUpdate is the builder for updating Reservation entities.
type ReservationUpdate struct {
	config
	hooks    []Hook
	mutation *ReservationMutation
}

// Where appends a list predicates to the ReservationUpdate builder.
func (ru *ReservationUpdate) Where(ps ...predicate.Reservation) *ReservationUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetFullName sets the "full_name" field.
func (ru *ReservationUpdate) SetFullName(s string) *ReservationUpdate {
	ru.mutation.SetFullName(s)
	return ru
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableFullName(s *string) *ReservationUpdate {
	if s != nil {
		ru.SetFullName(*s)
	}
	return ru
}

// ClearFullName clears the value of the "full_name" field.
func (ru *ReservationUpdate) ClearFullName() *ReservationUpdate {
	ru.mutation.ClearFullName()
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *ReservationUpdate) SetCreatedAt(t time.Time) *ReservationUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableCreatedAt(t *time.Time) *ReservationUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ReservationUpdate) SetUpdatedAt(t time.Time) *ReservationUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableUpdatedAt(t *time.Time) *ReservationUpdate {
	if t != nil {
		ru.SetUpdatedAt(*t)
	}
	return ru
}

// SetTableID sets the "table_id" field.
func (ru *ReservationUpdate) SetTableID(i int) *ReservationUpdate {
	ru.mutation.ResetTableID()
	ru.mutation.SetTableID(i)
	return ru
}

// SetNillableTableID sets the "table_id" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableTableID(i *int) *ReservationUpdate {
	if i != nil {
		ru.SetTableID(*i)
	}
	return ru
}

// AddTableID adds i to the "table_id" field.
func (ru *ReservationUpdate) AddTableID(i int) *ReservationUpdate {
	ru.mutation.AddTableID(i)
	return ru
}

// SetPhoneNumber sets the "phone_number" field.
func (ru *ReservationUpdate) SetPhoneNumber(s string) *ReservationUpdate {
	ru.mutation.SetPhoneNumber(s)
	return ru
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillablePhoneNumber(s *string) *ReservationUpdate {
	if s != nil {
		ru.SetPhoneNumber(*s)
	}
	return ru
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (ru *ReservationUpdate) ClearPhoneNumber() *ReservationUpdate {
	ru.mutation.ClearPhoneNumber()
	return ru
}

// SetStartTime sets the "start_time" field.
func (ru *ReservationUpdate) SetStartTime(t time.Time) *ReservationUpdate {
	ru.mutation.SetStartTime(t)
	return ru
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableStartTime(t *time.Time) *ReservationUpdate {
	if t != nil {
		ru.SetStartTime(*t)
	}
	return ru
}

// SetEndTime sets the "end_time" field.
func (ru *ReservationUpdate) SetEndTime(t time.Time) *ReservationUpdate {
	ru.mutation.SetEndTime(t)
	return ru
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableEndTime(t *time.Time) *ReservationUpdate {
	if t != nil {
		ru.SetEndTime(*t)
	}
	return ru
}

// Mutation returns the ReservationMutation object of the builder.
func (ru *ReservationUpdate) Mutation() *ReservationMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReservationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReservationUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReservationUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReservationUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ReservationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(reservation.Table, reservation.Columns, sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.FullName(); ok {
		_spec.SetField(reservation.FieldFullName, field.TypeString, value)
	}
	if ru.mutation.FullNameCleared() {
		_spec.ClearField(reservation.FieldFullName, field.TypeString)
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.SetField(reservation.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(reservation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.TableID(); ok {
		_spec.SetField(reservation.FieldTableID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedTableID(); ok {
		_spec.AddField(reservation.FieldTableID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.PhoneNumber(); ok {
		_spec.SetField(reservation.FieldPhoneNumber, field.TypeString, value)
	}
	if ru.mutation.PhoneNumberCleared() {
		_spec.ClearField(reservation.FieldPhoneNumber, field.TypeString)
	}
	if value, ok := ru.mutation.StartTime(); ok {
		_spec.SetField(reservation.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := ru.mutation.EndTime(); ok {
		_spec.SetField(reservation.FieldEndTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reservation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReservationUpdateOne is the builder for updating a single Reservation entity.
type ReservationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReservationMutation
}

// SetFullName sets the "full_name" field.
func (ruo *ReservationUpdateOne) SetFullName(s string) *ReservationUpdateOne {
	ruo.mutation.SetFullName(s)
	return ruo
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableFullName(s *string) *ReservationUpdateOne {
	if s != nil {
		ruo.SetFullName(*s)
	}
	return ruo
}

// ClearFullName clears the value of the "full_name" field.
func (ruo *ReservationUpdateOne) ClearFullName() *ReservationUpdateOne {
	ruo.mutation.ClearFullName()
	return ruo
}

// SetCreatedAt sets the "created_at" field.
func (ruo *ReservationUpdateOne) SetCreatedAt(t time.Time) *ReservationUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableCreatedAt(t *time.Time) *ReservationUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ReservationUpdateOne) SetUpdatedAt(t time.Time) *ReservationUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableUpdatedAt(t *time.Time) *ReservationUpdateOne {
	if t != nil {
		ruo.SetUpdatedAt(*t)
	}
	return ruo
}

// SetTableID sets the "table_id" field.
func (ruo *ReservationUpdateOne) SetTableID(i int) *ReservationUpdateOne {
	ruo.mutation.ResetTableID()
	ruo.mutation.SetTableID(i)
	return ruo
}

// SetNillableTableID sets the "table_id" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableTableID(i *int) *ReservationUpdateOne {
	if i != nil {
		ruo.SetTableID(*i)
	}
	return ruo
}

// AddTableID adds i to the "table_id" field.
func (ruo *ReservationUpdateOne) AddTableID(i int) *ReservationUpdateOne {
	ruo.mutation.AddTableID(i)
	return ruo
}

// SetPhoneNumber sets the "phone_number" field.
func (ruo *ReservationUpdateOne) SetPhoneNumber(s string) *ReservationUpdateOne {
	ruo.mutation.SetPhoneNumber(s)
	return ruo
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillablePhoneNumber(s *string) *ReservationUpdateOne {
	if s != nil {
		ruo.SetPhoneNumber(*s)
	}
	return ruo
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (ruo *ReservationUpdateOne) ClearPhoneNumber() *ReservationUpdateOne {
	ruo.mutation.ClearPhoneNumber()
	return ruo
}

// SetStartTime sets the "start_time" field.
func (ruo *ReservationUpdateOne) SetStartTime(t time.Time) *ReservationUpdateOne {
	ruo.mutation.SetStartTime(t)
	return ruo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableStartTime(t *time.Time) *ReservationUpdateOne {
	if t != nil {
		ruo.SetStartTime(*t)
	}
	return ruo
}

// SetEndTime sets the "end_time" field.
func (ruo *ReservationUpdateOne) SetEndTime(t time.Time) *ReservationUpdateOne {
	ruo.mutation.SetEndTime(t)
	return ruo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableEndTime(t *time.Time) *ReservationUpdateOne {
	if t != nil {
		ruo.SetEndTime(*t)
	}
	return ruo
}

// Mutation returns the ReservationMutation object of the builder.
func (ruo *ReservationUpdateOne) Mutation() *ReservationMutation {
	return ruo.mutation
}

// Where appends a list predicates to the ReservationUpdate builder.
func (ruo *ReservationUpdateOne) Where(ps ...predicate.Reservation) *ReservationUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReservationUpdateOne) Select(field string, fields ...string) *ReservationUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reservation entity.
func (ruo *ReservationUpdateOne) Save(ctx context.Context) (*Reservation, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReservationUpdateOne) SaveX(ctx context.Context) *Reservation {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReservationUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReservationUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ReservationUpdateOne) sqlSave(ctx context.Context) (_node *Reservation, err error) {
	_spec := sqlgraph.NewUpdateSpec(reservation.Table, reservation.Columns, sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Reservation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reservation.FieldID)
		for _, f := range fields {
			if !reservation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reservation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.FullName(); ok {
		_spec.SetField(reservation.FieldFullName, field.TypeString, value)
	}
	if ruo.mutation.FullNameCleared() {
		_spec.ClearField(reservation.FieldFullName, field.TypeString)
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.SetField(reservation.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(reservation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.TableID(); ok {
		_spec.SetField(reservation.FieldTableID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedTableID(); ok {
		_spec.AddField(reservation.FieldTableID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.PhoneNumber(); ok {
		_spec.SetField(reservation.FieldPhoneNumber, field.TypeString, value)
	}
	if ruo.mutation.PhoneNumberCleared() {
		_spec.ClearField(reservation.FieldPhoneNumber, field.TypeString)
	}
	if value, ok := ruo.mutation.StartTime(); ok {
		_spec.SetField(reservation.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.EndTime(); ok {
		_spec.SetField(reservation.FieldEndTime, field.TypeTime, value)
	}
	_node = &Reservation{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reservation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
