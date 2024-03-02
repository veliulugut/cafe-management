// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/price"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PriceCreate is the builder for creating a Price entity.
type PriceCreate struct {
	config
	mutation *PriceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPrice sets the "price" field.
func (pc *PriceCreate) SetPrice(f float64) *PriceCreate {
	pc.mutation.SetPrice(f)
	return pc
}

// SetPriceName sets the "price_name" field.
func (pc *PriceCreate) SetPriceName(s string) *PriceCreate {
	pc.mutation.SetPriceName(s)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PriceCreate) SetCreatedAt(t time.Time) *PriceCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PriceCreate) SetNillableCreatedAt(t *time.Time) *PriceCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PriceCreate) SetUpdatedAt(t time.Time) *PriceCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PriceCreate) SetNillableUpdatedAt(t *time.Time) *PriceCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// Mutation returns the PriceMutation object of the builder.
func (pc *PriceCreate) Mutation() *PriceMutation {
	return pc.mutation
}

// Save creates the Price in the database.
func (pc *PriceCreate) Save(ctx context.Context) (*Price, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PriceCreate) SaveX(ctx context.Context) *Price {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PriceCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PriceCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PriceCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := price.DefaultCreatedAt
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := price.DefaultUpdatedAt
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PriceCreate) check() error {
	if _, ok := pc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Price.price"`)}
	}
	if _, ok := pc.mutation.PriceName(); !ok {
		return &ValidationError{Name: "price_name", err: errors.New(`ent: missing required field "Price.price_name"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Price.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Price.updated_at"`)}
	}
	return nil
}

func (pc *PriceCreate) sqlSave(ctx context.Context) (*Price, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PriceCreate) createSpec() (*Price, *sqlgraph.CreateSpec) {
	var (
		_node = &Price{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(price.Table, sqlgraph.NewFieldSpec(price.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.Price(); ok {
		_spec.SetField(price.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := pc.mutation.PriceName(); ok {
		_spec.SetField(price.FieldPriceName, field.TypeString, value)
		_node.PriceName = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(price.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(price.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Price.Create().
//		SetPrice(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PriceUpsert) {
//			SetPrice(v+v).
//		}).
//		Exec(ctx)
func (pc *PriceCreate) OnConflict(opts ...sql.ConflictOption) *PriceUpsertOne {
	pc.conflict = opts
	return &PriceUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Price.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PriceCreate) OnConflictColumns(columns ...string) *PriceUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PriceUpsertOne{
		create: pc,
	}
}

type (
	// PriceUpsertOne is the builder for "upsert"-ing
	//  one Price node.
	PriceUpsertOne struct {
		create *PriceCreate
	}

	// PriceUpsert is the "OnConflict" setter.
	PriceUpsert struct {
		*sql.UpdateSet
	}
)

// SetPrice sets the "price" field.
func (u *PriceUpsert) SetPrice(v float64) *PriceUpsert {
	u.Set(price.FieldPrice, v)
	return u
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *PriceUpsert) UpdatePrice() *PriceUpsert {
	u.SetExcluded(price.FieldPrice)
	return u
}

// AddPrice adds v to the "price" field.
func (u *PriceUpsert) AddPrice(v float64) *PriceUpsert {
	u.Add(price.FieldPrice, v)
	return u
}

// SetPriceName sets the "price_name" field.
func (u *PriceUpsert) SetPriceName(v string) *PriceUpsert {
	u.Set(price.FieldPriceName, v)
	return u
}

// UpdatePriceName sets the "price_name" field to the value that was provided on create.
func (u *PriceUpsert) UpdatePriceName() *PriceUpsert {
	u.SetExcluded(price.FieldPriceName)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PriceUpsert) SetCreatedAt(v time.Time) *PriceUpsert {
	u.Set(price.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PriceUpsert) UpdateCreatedAt() *PriceUpsert {
	u.SetExcluded(price.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PriceUpsert) SetUpdatedAt(v time.Time) *PriceUpsert {
	u.Set(price.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PriceUpsert) UpdateUpdatedAt() *PriceUpsert {
	u.SetExcluded(price.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Price.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PriceUpsertOne) UpdateNewValues() *PriceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Price.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PriceUpsertOne) Ignore() *PriceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PriceUpsertOne) DoNothing() *PriceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PriceCreate.OnConflict
// documentation for more info.
func (u *PriceUpsertOne) Update(set func(*PriceUpsert)) *PriceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PriceUpsert{UpdateSet: update})
	}))
	return u
}

// SetPrice sets the "price" field.
func (u *PriceUpsertOne) SetPrice(v float64) *PriceUpsertOne {
	return u.Update(func(s *PriceUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *PriceUpsertOne) AddPrice(v float64) *PriceUpsertOne {
	return u.Update(func(s *PriceUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *PriceUpsertOne) UpdatePrice() *PriceUpsertOne {
	return u.Update(func(s *PriceUpsert) {
		s.UpdatePrice()
	})
}

// SetPriceName sets the "price_name" field.
func (u *PriceUpsertOne) SetPriceName(v string) *PriceUpsertOne {
	return u.Update(func(s *PriceUpsert) {
		s.SetPriceName(v)
	})
}

// UpdatePriceName sets the "price_name" field to the value that was provided on create.
func (u *PriceUpsertOne) UpdatePriceName() *PriceUpsertOne {
	return u.Update(func(s *PriceUpsert) {
		s.UpdatePriceName()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *PriceUpsertOne) SetCreatedAt(v time.Time) *PriceUpsertOne {
	return u.Update(func(s *PriceUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PriceUpsertOne) UpdateCreatedAt() *PriceUpsertOne {
	return u.Update(func(s *PriceUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PriceUpsertOne) SetUpdatedAt(v time.Time) *PriceUpsertOne {
	return u.Update(func(s *PriceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PriceUpsertOne) UpdateUpdatedAt() *PriceUpsertOne {
	return u.Update(func(s *PriceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *PriceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PriceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PriceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PriceUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PriceUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PriceCreateBulk is the builder for creating many Price entities in bulk.
type PriceCreateBulk struct {
	config
	err      error
	builders []*PriceCreate
	conflict []sql.ConflictOption
}

// Save creates the Price entities in the database.
func (pcb *PriceCreateBulk) Save(ctx context.Context) ([]*Price, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Price, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PriceMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PriceCreateBulk) SaveX(ctx context.Context) []*Price {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PriceCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PriceCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Price.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PriceUpsert) {
//			SetPrice(v+v).
//		}).
//		Exec(ctx)
func (pcb *PriceCreateBulk) OnConflict(opts ...sql.ConflictOption) *PriceUpsertBulk {
	pcb.conflict = opts
	return &PriceUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Price.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PriceCreateBulk) OnConflictColumns(columns ...string) *PriceUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PriceUpsertBulk{
		create: pcb,
	}
}

// PriceUpsertBulk is the builder for "upsert"-ing
// a bulk of Price nodes.
type PriceUpsertBulk struct {
	create *PriceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Price.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PriceUpsertBulk) UpdateNewValues() *PriceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Price.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PriceUpsertBulk) Ignore() *PriceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PriceUpsertBulk) DoNothing() *PriceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PriceCreateBulk.OnConflict
// documentation for more info.
func (u *PriceUpsertBulk) Update(set func(*PriceUpsert)) *PriceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PriceUpsert{UpdateSet: update})
	}))
	return u
}

// SetPrice sets the "price" field.
func (u *PriceUpsertBulk) SetPrice(v float64) *PriceUpsertBulk {
	return u.Update(func(s *PriceUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *PriceUpsertBulk) AddPrice(v float64) *PriceUpsertBulk {
	return u.Update(func(s *PriceUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *PriceUpsertBulk) UpdatePrice() *PriceUpsertBulk {
	return u.Update(func(s *PriceUpsert) {
		s.UpdatePrice()
	})
}

// SetPriceName sets the "price_name" field.
func (u *PriceUpsertBulk) SetPriceName(v string) *PriceUpsertBulk {
	return u.Update(func(s *PriceUpsert) {
		s.SetPriceName(v)
	})
}

// UpdatePriceName sets the "price_name" field to the value that was provided on create.
func (u *PriceUpsertBulk) UpdatePriceName() *PriceUpsertBulk {
	return u.Update(func(s *PriceUpsert) {
		s.UpdatePriceName()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *PriceUpsertBulk) SetCreatedAt(v time.Time) *PriceUpsertBulk {
	return u.Update(func(s *PriceUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PriceUpsertBulk) UpdateCreatedAt() *PriceUpsertBulk {
	return u.Update(func(s *PriceUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PriceUpsertBulk) SetUpdatedAt(v time.Time) *PriceUpsertBulk {
	return u.Update(func(s *PriceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PriceUpsertBulk) UpdateUpdatedAt() *PriceUpsertBulk {
	return u.Update(func(s *PriceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *PriceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PriceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PriceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PriceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}