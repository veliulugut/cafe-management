// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/ordertype"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderTypeCreate is the builder for creating a OrderType entity.
type OrderTypeCreate struct {
	config
	mutation *OrderTypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (otc *OrderTypeCreate) SetName(s string) *OrderTypeCreate {
	otc.mutation.SetName(s)
	return otc
}

// SetID sets the "id" field.
func (otc *OrderTypeCreate) SetID(i int) *OrderTypeCreate {
	otc.mutation.SetID(i)
	return otc
}

// Mutation returns the OrderTypeMutation object of the builder.
func (otc *OrderTypeCreate) Mutation() *OrderTypeMutation {
	return otc.mutation
}

// Save creates the OrderType in the database.
func (otc *OrderTypeCreate) Save(ctx context.Context) (*OrderType, error) {
	return withHooks(ctx, otc.sqlSave, otc.mutation, otc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (otc *OrderTypeCreate) SaveX(ctx context.Context) *OrderType {
	v, err := otc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (otc *OrderTypeCreate) Exec(ctx context.Context) error {
	_, err := otc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otc *OrderTypeCreate) ExecX(ctx context.Context) {
	if err := otc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otc *OrderTypeCreate) check() error {
	if _, ok := otc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OrderType.name"`)}
	}
	return nil
}

func (otc *OrderTypeCreate) sqlSave(ctx context.Context) (*OrderType, error) {
	if err := otc.check(); err != nil {
		return nil, err
	}
	_node, _spec := otc.createSpec()
	if err := sqlgraph.CreateNode(ctx, otc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	otc.mutation.id = &_node.ID
	otc.mutation.done = true
	return _node, nil
}

func (otc *OrderTypeCreate) createSpec() (*OrderType, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderType{config: otc.config}
		_spec = sqlgraph.NewCreateSpec(ordertype.Table, sqlgraph.NewFieldSpec(ordertype.FieldID, field.TypeInt))
	)
	_spec.OnConflict = otc.conflict
	if id, ok := otc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := otc.mutation.Name(); ok {
		_spec.SetField(ordertype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderType.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderTypeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (otc *OrderTypeCreate) OnConflict(opts ...sql.ConflictOption) *OrderTypeUpsertOne {
	otc.conflict = opts
	return &OrderTypeUpsertOne{
		create: otc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (otc *OrderTypeCreate) OnConflictColumns(columns ...string) *OrderTypeUpsertOne {
	otc.conflict = append(otc.conflict, sql.ConflictColumns(columns...))
	return &OrderTypeUpsertOne{
		create: otc,
	}
}

type (
	// OrderTypeUpsertOne is the builder for "upsert"-ing
	//  one OrderType node.
	OrderTypeUpsertOne struct {
		create *OrderTypeCreate
	}

	// OrderTypeUpsert is the "OnConflict" setter.
	OrderTypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *OrderTypeUpsert) SetName(v string) *OrderTypeUpsert {
	u.Set(ordertype.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrderTypeUpsert) UpdateName() *OrderTypeUpsert {
	u.SetExcluded(ordertype.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrderType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ordertype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderTypeUpsertOne) UpdateNewValues() *OrderTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(ordertype.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderType.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrderTypeUpsertOne) Ignore() *OrderTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderTypeUpsertOne) DoNothing() *OrderTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderTypeCreate.OnConflict
// documentation for more info.
func (u *OrderTypeUpsertOne) Update(set func(*OrderTypeUpsert)) *OrderTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *OrderTypeUpsertOne) SetName(v string) *OrderTypeUpsertOne {
	return u.Update(func(s *OrderTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrderTypeUpsertOne) UpdateName() *OrderTypeUpsertOne {
	return u.Update(func(s *OrderTypeUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *OrderTypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderTypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderTypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderTypeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderTypeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderTypeCreateBulk is the builder for creating many OrderType entities in bulk.
type OrderTypeCreateBulk struct {
	config
	err      error
	builders []*OrderTypeCreate
	conflict []sql.ConflictOption
}

// Save creates the OrderType entities in the database.
func (otcb *OrderTypeCreateBulk) Save(ctx context.Context) ([]*OrderType, error) {
	if otcb.err != nil {
		return nil, otcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(otcb.builders))
	nodes := make([]*OrderType, len(otcb.builders))
	mutators := make([]Mutator, len(otcb.builders))
	for i := range otcb.builders {
		func(i int, root context.Context) {
			builder := otcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, otcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = otcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, otcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, otcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (otcb *OrderTypeCreateBulk) SaveX(ctx context.Context) []*OrderType {
	v, err := otcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (otcb *OrderTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := otcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otcb *OrderTypeCreateBulk) ExecX(ctx context.Context) {
	if err := otcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderType.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderTypeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (otcb *OrderTypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderTypeUpsertBulk {
	otcb.conflict = opts
	return &OrderTypeUpsertBulk{
		create: otcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (otcb *OrderTypeCreateBulk) OnConflictColumns(columns ...string) *OrderTypeUpsertBulk {
	otcb.conflict = append(otcb.conflict, sql.ConflictColumns(columns...))
	return &OrderTypeUpsertBulk{
		create: otcb,
	}
}

// OrderTypeUpsertBulk is the builder for "upsert"-ing
// a bulk of OrderType nodes.
type OrderTypeUpsertBulk struct {
	create *OrderTypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrderType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ordertype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderTypeUpsertBulk) UpdateNewValues() *OrderTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(ordertype.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderType.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrderTypeUpsertBulk) Ignore() *OrderTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderTypeUpsertBulk) DoNothing() *OrderTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderTypeCreateBulk.OnConflict
// documentation for more info.
func (u *OrderTypeUpsertBulk) Update(set func(*OrderTypeUpsert)) *OrderTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *OrderTypeUpsertBulk) SetName(v string) *OrderTypeUpsertBulk {
	return u.Update(func(s *OrderTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrderTypeUpsertBulk) UpdateName() *OrderTypeUpsertBulk {
	return u.Update(func(s *OrderTypeUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *OrderTypeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrderTypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderTypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderTypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
