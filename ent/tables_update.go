// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/predicate"
	"cafe-management/ent/reservation"
	"cafe-management/ent/tables"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TablesUpdate is the builder for updating Tables entities.
type TablesUpdate struct {
	config
	hooks    []Hook
	mutation *TablesMutation
}

// Where appends a list predicates to the TablesUpdate builder.
func (tu *TablesUpdate) Where(ps ...predicate.Tables) *TablesUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetNumberOfGuests sets the "number_of_guests" field.
func (tu *TablesUpdate) SetNumberOfGuests(i int) *TablesUpdate {
	tu.mutation.ResetNumberOfGuests()
	tu.mutation.SetNumberOfGuests(i)
	return tu
}

// SetNillableNumberOfGuests sets the "number_of_guests" field if the given value is not nil.
func (tu *TablesUpdate) SetNillableNumberOfGuests(i *int) *TablesUpdate {
	if i != nil {
		tu.SetNumberOfGuests(*i)
	}
	return tu
}

// AddNumberOfGuests adds i to the "number_of_guests" field.
func (tu *TablesUpdate) AddNumberOfGuests(i int) *TablesUpdate {
	tu.mutation.AddNumberOfGuests(i)
	return tu
}

// SetTableNumber sets the "table_number" field.
func (tu *TablesUpdate) SetTableNumber(i int) *TablesUpdate {
	tu.mutation.ResetTableNumber()
	tu.mutation.SetTableNumber(i)
	return tu
}

// SetNillableTableNumber sets the "table_number" field if the given value is not nil.
func (tu *TablesUpdate) SetNillableTableNumber(i *int) *TablesUpdate {
	if i != nil {
		tu.SetTableNumber(*i)
	}
	return tu
}

// AddTableNumber adds i to the "table_number" field.
func (tu *TablesUpdate) AddTableNumber(i int) *TablesUpdate {
	tu.mutation.AddTableNumber(i)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TablesUpdate) SetDescription(s string) *TablesUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TablesUpdate) SetNillableDescription(s *string) *TablesUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// SetTableType sets the "table_type" field.
func (tu *TablesUpdate) SetTableType(s string) *TablesUpdate {
	tu.mutation.SetTableType(s)
	return tu
}

// SetNillableTableType sets the "table_type" field if the given value is not nil.
func (tu *TablesUpdate) SetNillableTableType(s *string) *TablesUpdate {
	if s != nil {
		tu.SetTableType(*s)
	}
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TablesUpdate) SetCreatedAt(t time.Time) *TablesUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TablesUpdate) SetNillableCreatedAt(t *time.Time) *TablesUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TablesUpdate) SetUpdatedAt(t time.Time) *TablesUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tu *TablesUpdate) SetNillableUpdatedAt(t *time.Time) *TablesUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// AddReservationIDs adds the "reservation" edge to the Reservation entity by IDs.
func (tu *TablesUpdate) AddReservationIDs(ids ...int) *TablesUpdate {
	tu.mutation.AddReservationIDs(ids...)
	return tu
}

// AddReservation adds the "reservation" edges to the Reservation entity.
func (tu *TablesUpdate) AddReservation(r ...*Reservation) *TablesUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.AddReservationIDs(ids...)
}

// Mutation returns the TablesMutation object of the builder.
func (tu *TablesUpdate) Mutation() *TablesMutation {
	return tu.mutation
}

// ClearReservation clears all "reservation" edges to the Reservation entity.
func (tu *TablesUpdate) ClearReservation() *TablesUpdate {
	tu.mutation.ClearReservation()
	return tu
}

// RemoveReservationIDs removes the "reservation" edge to Reservation entities by IDs.
func (tu *TablesUpdate) RemoveReservationIDs(ids ...int) *TablesUpdate {
	tu.mutation.RemoveReservationIDs(ids...)
	return tu
}

// RemoveReservation removes "reservation" edges to Reservation entities.
func (tu *TablesUpdate) RemoveReservation(r ...*Reservation) *TablesUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.RemoveReservationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TablesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TablesUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TablesUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TablesUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TablesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tables.Table, tables.Columns, sqlgraph.NewFieldSpec(tables.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.NumberOfGuests(); ok {
		_spec.SetField(tables.FieldNumberOfGuests, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedNumberOfGuests(); ok {
		_spec.AddField(tables.FieldNumberOfGuests, field.TypeInt, value)
	}
	if value, ok := tu.mutation.TableNumber(); ok {
		_spec.SetField(tables.FieldTableNumber, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedTableNumber(); ok {
		_spec.AddField(tables.FieldTableNumber, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(tables.FieldDescription, field.TypeString, value)
	}
	if value, ok := tu.mutation.TableType(); ok {
		_spec.SetField(tables.FieldTableType, field.TypeString, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(tables.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(tables.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.ReservationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tables.ReservationTable,
			Columns: []string{tables.ReservationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedReservationIDs(); len(nodes) > 0 && !tu.mutation.ReservationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tables.ReservationTable,
			Columns: []string{tables.ReservationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ReservationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tables.ReservationTable,
			Columns: []string{tables.ReservationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tables.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TablesUpdateOne is the builder for updating a single Tables entity.
type TablesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TablesMutation
}

// SetNumberOfGuests sets the "number_of_guests" field.
func (tuo *TablesUpdateOne) SetNumberOfGuests(i int) *TablesUpdateOne {
	tuo.mutation.ResetNumberOfGuests()
	tuo.mutation.SetNumberOfGuests(i)
	return tuo
}

// SetNillableNumberOfGuests sets the "number_of_guests" field if the given value is not nil.
func (tuo *TablesUpdateOne) SetNillableNumberOfGuests(i *int) *TablesUpdateOne {
	if i != nil {
		tuo.SetNumberOfGuests(*i)
	}
	return tuo
}

// AddNumberOfGuests adds i to the "number_of_guests" field.
func (tuo *TablesUpdateOne) AddNumberOfGuests(i int) *TablesUpdateOne {
	tuo.mutation.AddNumberOfGuests(i)
	return tuo
}

// SetTableNumber sets the "table_number" field.
func (tuo *TablesUpdateOne) SetTableNumber(i int) *TablesUpdateOne {
	tuo.mutation.ResetTableNumber()
	tuo.mutation.SetTableNumber(i)
	return tuo
}

// SetNillableTableNumber sets the "table_number" field if the given value is not nil.
func (tuo *TablesUpdateOne) SetNillableTableNumber(i *int) *TablesUpdateOne {
	if i != nil {
		tuo.SetTableNumber(*i)
	}
	return tuo
}

// AddTableNumber adds i to the "table_number" field.
func (tuo *TablesUpdateOne) AddTableNumber(i int) *TablesUpdateOne {
	tuo.mutation.AddTableNumber(i)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TablesUpdateOne) SetDescription(s string) *TablesUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TablesUpdateOne) SetNillableDescription(s *string) *TablesUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// SetTableType sets the "table_type" field.
func (tuo *TablesUpdateOne) SetTableType(s string) *TablesUpdateOne {
	tuo.mutation.SetTableType(s)
	return tuo
}

// SetNillableTableType sets the "table_type" field if the given value is not nil.
func (tuo *TablesUpdateOne) SetNillableTableType(s *string) *TablesUpdateOne {
	if s != nil {
		tuo.SetTableType(*s)
	}
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TablesUpdateOne) SetCreatedAt(t time.Time) *TablesUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TablesUpdateOne) SetNillableCreatedAt(t *time.Time) *TablesUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TablesUpdateOne) SetUpdatedAt(t time.Time) *TablesUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuo *TablesUpdateOne) SetNillableUpdatedAt(t *time.Time) *TablesUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// AddReservationIDs adds the "reservation" edge to the Reservation entity by IDs.
func (tuo *TablesUpdateOne) AddReservationIDs(ids ...int) *TablesUpdateOne {
	tuo.mutation.AddReservationIDs(ids...)
	return tuo
}

// AddReservation adds the "reservation" edges to the Reservation entity.
func (tuo *TablesUpdateOne) AddReservation(r ...*Reservation) *TablesUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.AddReservationIDs(ids...)
}

// Mutation returns the TablesMutation object of the builder.
func (tuo *TablesUpdateOne) Mutation() *TablesMutation {
	return tuo.mutation
}

// ClearReservation clears all "reservation" edges to the Reservation entity.
func (tuo *TablesUpdateOne) ClearReservation() *TablesUpdateOne {
	tuo.mutation.ClearReservation()
	return tuo
}

// RemoveReservationIDs removes the "reservation" edge to Reservation entities by IDs.
func (tuo *TablesUpdateOne) RemoveReservationIDs(ids ...int) *TablesUpdateOne {
	tuo.mutation.RemoveReservationIDs(ids...)
	return tuo
}

// RemoveReservation removes "reservation" edges to Reservation entities.
func (tuo *TablesUpdateOne) RemoveReservation(r ...*Reservation) *TablesUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.RemoveReservationIDs(ids...)
}

// Where appends a list predicates to the TablesUpdate builder.
func (tuo *TablesUpdateOne) Where(ps ...predicate.Tables) *TablesUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TablesUpdateOne) Select(field string, fields ...string) *TablesUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tables entity.
func (tuo *TablesUpdateOne) Save(ctx context.Context) (*Tables, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TablesUpdateOne) SaveX(ctx context.Context) *Tables {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TablesUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TablesUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TablesUpdateOne) sqlSave(ctx context.Context) (_node *Tables, err error) {
	_spec := sqlgraph.NewUpdateSpec(tables.Table, tables.Columns, sqlgraph.NewFieldSpec(tables.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tables.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tables.FieldID)
		for _, f := range fields {
			if !tables.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tables.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.NumberOfGuests(); ok {
		_spec.SetField(tables.FieldNumberOfGuests, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedNumberOfGuests(); ok {
		_spec.AddField(tables.FieldNumberOfGuests, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.TableNumber(); ok {
		_spec.SetField(tables.FieldTableNumber, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedTableNumber(); ok {
		_spec.AddField(tables.FieldTableNumber, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(tables.FieldDescription, field.TypeString, value)
	}
	if value, ok := tuo.mutation.TableType(); ok {
		_spec.SetField(tables.FieldTableType, field.TypeString, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(tables.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(tables.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.ReservationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tables.ReservationTable,
			Columns: []string{tables.ReservationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedReservationIDs(); len(nodes) > 0 && !tuo.mutation.ReservationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tables.ReservationTable,
			Columns: []string{tables.ReservationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ReservationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tables.ReservationTable,
			Columns: []string{tables.ReservationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tables{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tables.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
