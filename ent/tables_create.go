// Code generated by ent, DO NOT EDIT.

package ent

import (
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

// TablesCreate is the builder for creating a Tables entity.
type TablesCreate struct {
	config
	mutation *TablesMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNumberOfGuests sets the "number_of_guests" field.
func (tc *TablesCreate) SetNumberOfGuests(i int) *TablesCreate {
	tc.mutation.SetNumberOfGuests(i)
	return tc
}

// SetTableNumber sets the "table_number" field.
func (tc *TablesCreate) SetTableNumber(i int) *TablesCreate {
	tc.mutation.SetTableNumber(i)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TablesCreate) SetDescription(s string) *TablesCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetTableType sets the "table_type" field.
func (tc *TablesCreate) SetTableType(s string) *TablesCreate {
	tc.mutation.SetTableType(s)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TablesCreate) SetCreatedAt(t time.Time) *TablesCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TablesCreate) SetNillableCreatedAt(t *time.Time) *TablesCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TablesCreate) SetUpdatedAt(t time.Time) *TablesCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TablesCreate) SetNillableUpdatedAt(t *time.Time) *TablesCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// AddReservationIDs adds the "reservation" edge to the Reservation entity by IDs.
func (tc *TablesCreate) AddReservationIDs(ids ...int) *TablesCreate {
	tc.mutation.AddReservationIDs(ids...)
	return tc
}

// AddReservation adds the "reservation" edges to the Reservation entity.
func (tc *TablesCreate) AddReservation(r ...*Reservation) *TablesCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tc.AddReservationIDs(ids...)
}

// Mutation returns the TablesMutation object of the builder.
func (tc *TablesCreate) Mutation() *TablesMutation {
	return tc.mutation
}

// Save creates the Tables in the database.
func (tc *TablesCreate) Save(ctx context.Context) (*Tables, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TablesCreate) SaveX(ctx context.Context) *Tables {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TablesCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TablesCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TablesCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := tables.DefaultCreatedAt
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := tables.DefaultUpdatedAt
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TablesCreate) check() error {
	if _, ok := tc.mutation.NumberOfGuests(); !ok {
		return &ValidationError{Name: "number_of_guests", err: errors.New(`ent: missing required field "Tables.number_of_guests"`)}
	}
	if _, ok := tc.mutation.TableNumber(); !ok {
		return &ValidationError{Name: "table_number", err: errors.New(`ent: missing required field "Tables.table_number"`)}
	}
	if _, ok := tc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Tables.description"`)}
	}
	if _, ok := tc.mutation.TableType(); !ok {
		return &ValidationError{Name: "table_type", err: errors.New(`ent: missing required field "Tables.table_type"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Tables.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Tables.updated_at"`)}
	}
	return nil
}

func (tc *TablesCreate) sqlSave(ctx context.Context) (*Tables, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TablesCreate) createSpec() (*Tables, *sqlgraph.CreateSpec) {
	var (
		_node = &Tables{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tables.Table, sqlgraph.NewFieldSpec(tables.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.NumberOfGuests(); ok {
		_spec.SetField(tables.FieldNumberOfGuests, field.TypeInt, value)
		_node.NumberOfGuests = value
	}
	if value, ok := tc.mutation.TableNumber(); ok {
		_spec.SetField(tables.FieldTableNumber, field.TypeInt, value)
		_node.TableNumber = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(tables.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.TableType(); ok {
		_spec.SetField(tables.FieldTableType, field.TypeString, value)
		_node.TableType = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(tables.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(tables.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.ReservationIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tables.Create().
//		SetNumberOfGuests(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TablesUpsert) {
//			SetNumberOfGuests(v+v).
//		}).
//		Exec(ctx)
func (tc *TablesCreate) OnConflict(opts ...sql.ConflictOption) *TablesUpsertOne {
	tc.conflict = opts
	return &TablesUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tables.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TablesCreate) OnConflictColumns(columns ...string) *TablesUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TablesUpsertOne{
		create: tc,
	}
}

type (
	// TablesUpsertOne is the builder for "upsert"-ing
	//  one Tables node.
	TablesUpsertOne struct {
		create *TablesCreate
	}

	// TablesUpsert is the "OnConflict" setter.
	TablesUpsert struct {
		*sql.UpdateSet
	}
)

// SetNumberOfGuests sets the "number_of_guests" field.
func (u *TablesUpsert) SetNumberOfGuests(v int) *TablesUpsert {
	u.Set(tables.FieldNumberOfGuests, v)
	return u
}

// UpdateNumberOfGuests sets the "number_of_guests" field to the value that was provided on create.
func (u *TablesUpsert) UpdateNumberOfGuests() *TablesUpsert {
	u.SetExcluded(tables.FieldNumberOfGuests)
	return u
}

// AddNumberOfGuests adds v to the "number_of_guests" field.
func (u *TablesUpsert) AddNumberOfGuests(v int) *TablesUpsert {
	u.Add(tables.FieldNumberOfGuests, v)
	return u
}

// SetTableNumber sets the "table_number" field.
func (u *TablesUpsert) SetTableNumber(v int) *TablesUpsert {
	u.Set(tables.FieldTableNumber, v)
	return u
}

// UpdateTableNumber sets the "table_number" field to the value that was provided on create.
func (u *TablesUpsert) UpdateTableNumber() *TablesUpsert {
	u.SetExcluded(tables.FieldTableNumber)
	return u
}

// AddTableNumber adds v to the "table_number" field.
func (u *TablesUpsert) AddTableNumber(v int) *TablesUpsert {
	u.Add(tables.FieldTableNumber, v)
	return u
}

// SetDescription sets the "description" field.
func (u *TablesUpsert) SetDescription(v string) *TablesUpsert {
	u.Set(tables.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TablesUpsert) UpdateDescription() *TablesUpsert {
	u.SetExcluded(tables.FieldDescription)
	return u
}

// SetTableType sets the "table_type" field.
func (u *TablesUpsert) SetTableType(v string) *TablesUpsert {
	u.Set(tables.FieldTableType, v)
	return u
}

// UpdateTableType sets the "table_type" field to the value that was provided on create.
func (u *TablesUpsert) UpdateTableType() *TablesUpsert {
	u.SetExcluded(tables.FieldTableType)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TablesUpsert) SetCreatedAt(v time.Time) *TablesUpsert {
	u.Set(tables.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TablesUpsert) UpdateCreatedAt() *TablesUpsert {
	u.SetExcluded(tables.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TablesUpsert) SetUpdatedAt(v time.Time) *TablesUpsert {
	u.Set(tables.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TablesUpsert) UpdateUpdatedAt() *TablesUpsert {
	u.SetExcluded(tables.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Tables.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TablesUpsertOne) UpdateNewValues() *TablesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tables.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TablesUpsertOne) Ignore() *TablesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TablesUpsertOne) DoNothing() *TablesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TablesCreate.OnConflict
// documentation for more info.
func (u *TablesUpsertOne) Update(set func(*TablesUpsert)) *TablesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TablesUpsert{UpdateSet: update})
	}))
	return u
}

// SetNumberOfGuests sets the "number_of_guests" field.
func (u *TablesUpsertOne) SetNumberOfGuests(v int) *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.SetNumberOfGuests(v)
	})
}

// AddNumberOfGuests adds v to the "number_of_guests" field.
func (u *TablesUpsertOne) AddNumberOfGuests(v int) *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.AddNumberOfGuests(v)
	})
}

// UpdateNumberOfGuests sets the "number_of_guests" field to the value that was provided on create.
func (u *TablesUpsertOne) UpdateNumberOfGuests() *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateNumberOfGuests()
	})
}

// SetTableNumber sets the "table_number" field.
func (u *TablesUpsertOne) SetTableNumber(v int) *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.SetTableNumber(v)
	})
}

// AddTableNumber adds v to the "table_number" field.
func (u *TablesUpsertOne) AddTableNumber(v int) *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.AddTableNumber(v)
	})
}

// UpdateTableNumber sets the "table_number" field to the value that was provided on create.
func (u *TablesUpsertOne) UpdateTableNumber() *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateTableNumber()
	})
}

// SetDescription sets the "description" field.
func (u *TablesUpsertOne) SetDescription(v string) *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TablesUpsertOne) UpdateDescription() *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateDescription()
	})
}

// SetTableType sets the "table_type" field.
func (u *TablesUpsertOne) SetTableType(v string) *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.SetTableType(v)
	})
}

// UpdateTableType sets the "table_type" field to the value that was provided on create.
func (u *TablesUpsertOne) UpdateTableType() *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateTableType()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TablesUpsertOne) SetCreatedAt(v time.Time) *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TablesUpsertOne) UpdateCreatedAt() *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TablesUpsertOne) SetUpdatedAt(v time.Time) *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TablesUpsertOne) UpdateUpdatedAt() *TablesUpsertOne {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TablesUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TablesCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TablesUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TablesUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TablesUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TablesCreateBulk is the builder for creating many Tables entities in bulk.
type TablesCreateBulk struct {
	config
	err      error
	builders []*TablesCreate
	conflict []sql.ConflictOption
}

// Save creates the Tables entities in the database.
func (tcb *TablesCreateBulk) Save(ctx context.Context) ([]*Tables, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tables, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TablesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TablesCreateBulk) SaveX(ctx context.Context) []*Tables {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TablesCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TablesCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tables.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TablesUpsert) {
//			SetNumberOfGuests(v+v).
//		}).
//		Exec(ctx)
func (tcb *TablesCreateBulk) OnConflict(opts ...sql.ConflictOption) *TablesUpsertBulk {
	tcb.conflict = opts
	return &TablesUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tables.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TablesCreateBulk) OnConflictColumns(columns ...string) *TablesUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TablesUpsertBulk{
		create: tcb,
	}
}

// TablesUpsertBulk is the builder for "upsert"-ing
// a bulk of Tables nodes.
type TablesUpsertBulk struct {
	create *TablesCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Tables.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TablesUpsertBulk) UpdateNewValues() *TablesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tables.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TablesUpsertBulk) Ignore() *TablesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TablesUpsertBulk) DoNothing() *TablesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TablesCreateBulk.OnConflict
// documentation for more info.
func (u *TablesUpsertBulk) Update(set func(*TablesUpsert)) *TablesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TablesUpsert{UpdateSet: update})
	}))
	return u
}

// SetNumberOfGuests sets the "number_of_guests" field.
func (u *TablesUpsertBulk) SetNumberOfGuests(v int) *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.SetNumberOfGuests(v)
	})
}

// AddNumberOfGuests adds v to the "number_of_guests" field.
func (u *TablesUpsertBulk) AddNumberOfGuests(v int) *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.AddNumberOfGuests(v)
	})
}

// UpdateNumberOfGuests sets the "number_of_guests" field to the value that was provided on create.
func (u *TablesUpsertBulk) UpdateNumberOfGuests() *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateNumberOfGuests()
	})
}

// SetTableNumber sets the "table_number" field.
func (u *TablesUpsertBulk) SetTableNumber(v int) *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.SetTableNumber(v)
	})
}

// AddTableNumber adds v to the "table_number" field.
func (u *TablesUpsertBulk) AddTableNumber(v int) *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.AddTableNumber(v)
	})
}

// UpdateTableNumber sets the "table_number" field to the value that was provided on create.
func (u *TablesUpsertBulk) UpdateTableNumber() *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateTableNumber()
	})
}

// SetDescription sets the "description" field.
func (u *TablesUpsertBulk) SetDescription(v string) *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TablesUpsertBulk) UpdateDescription() *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateDescription()
	})
}

// SetTableType sets the "table_type" field.
func (u *TablesUpsertBulk) SetTableType(v string) *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.SetTableType(v)
	})
}

// UpdateTableType sets the "table_type" field to the value that was provided on create.
func (u *TablesUpsertBulk) UpdateTableType() *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateTableType()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TablesUpsertBulk) SetCreatedAt(v time.Time) *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TablesUpsertBulk) UpdateCreatedAt() *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TablesUpsertBulk) SetUpdatedAt(v time.Time) *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TablesUpsertBulk) UpdateUpdatedAt() *TablesUpsertBulk {
	return u.Update(func(s *TablesUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TablesUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TablesCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TablesCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TablesUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
