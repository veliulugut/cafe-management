// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/order"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOrderID sets the "order_id" field.
func (oc *OrderCreate) SetOrderID(i int) *OrderCreate {
	oc.mutation.SetOrderID(i)
	return oc
}

// SetTableID sets the "table_id" field.
func (oc *OrderCreate) SetTableID(i int) *OrderCreate {
	oc.mutation.SetTableID(i)
	return oc
}

// SetUserID sets the "user_id" field.
func (oc *OrderCreate) SetUserID(i int) *OrderCreate {
	oc.mutation.SetUserID(i)
	return oc
}

// SetOrderType sets the "order_type" field.
func (oc *OrderCreate) SetOrderType(i int) *OrderCreate {
	oc.mutation.SetOrderType(i)
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(s string) *OrderCreate {
	oc.mutation.SetStatus(s)
	return oc
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrderCreate) SetUpdatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := order.DefaultCreatedAt
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := order.DefaultUpdatedAt
		oc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "Order.order_id"`)}
	}
	if _, ok := oc.mutation.TableID(); !ok {
		return &ValidationError{Name: "table_id", err: errors.New(`ent: missing required field "Order.table_id"`)}
	}
	if _, ok := oc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Order.user_id"`)}
	}
	if _, ok := oc.mutation.OrderType(); !ok {
		return &ValidationError{Name: "order_type", err: errors.New(`ent: missing required field "Order.order_type"`)}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Order.status"`)}
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Order.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Order.updated_at"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	)
	_spec.OnConflict = oc.conflict
	if value, ok := oc.mutation.OrderID(); ok {
		_spec.SetField(order.FieldOrderID, field.TypeInt, value)
		_node.OrderID = value
	}
	if value, ok := oc.mutation.TableID(); ok {
		_spec.SetField(order.FieldTableID, field.TypeInt, value)
		_node.TableID = value
	}
	if value, ok := oc.mutation.UserID(); ok {
		_spec.SetField(order.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := oc.mutation.OrderType(); ok {
		_spec.SetField(order.FieldOrderType, field.TypeInt, value)
		_node.OrderType = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Order.Create().
//		SetOrderID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
func (oc *OrderCreate) OnConflict(opts ...sql.ConflictOption) *OrderUpsertOne {
	oc.conflict = opts
	return &OrderUpsertOne{
		create: oc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oc *OrderCreate) OnConflictColumns(columns ...string) *OrderUpsertOne {
	oc.conflict = append(oc.conflict, sql.ConflictColumns(columns...))
	return &OrderUpsertOne{
		create: oc,
	}
}

type (
	// OrderUpsertOne is the builder for "upsert"-ing
	//  one Order node.
	OrderUpsertOne struct {
		create *OrderCreate
	}

	// OrderUpsert is the "OnConflict" setter.
	OrderUpsert struct {
		*sql.UpdateSet
	}
)

// SetOrderID sets the "order_id" field.
func (u *OrderUpsert) SetOrderID(v int) *OrderUpsert {
	u.Set(order.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderUpsert) UpdateOrderID() *OrderUpsert {
	u.SetExcluded(order.FieldOrderID)
	return u
}

// AddOrderID adds v to the "order_id" field.
func (u *OrderUpsert) AddOrderID(v int) *OrderUpsert {
	u.Add(order.FieldOrderID, v)
	return u
}

// SetTableID sets the "table_id" field.
func (u *OrderUpsert) SetTableID(v int) *OrderUpsert {
	u.Set(order.FieldTableID, v)
	return u
}

// UpdateTableID sets the "table_id" field to the value that was provided on create.
func (u *OrderUpsert) UpdateTableID() *OrderUpsert {
	u.SetExcluded(order.FieldTableID)
	return u
}

// AddTableID adds v to the "table_id" field.
func (u *OrderUpsert) AddTableID(v int) *OrderUpsert {
	u.Add(order.FieldTableID, v)
	return u
}

// SetUserID sets the "user_id" field.
func (u *OrderUpsert) SetUserID(v int) *OrderUpsert {
	u.Set(order.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrderUpsert) UpdateUserID() *OrderUpsert {
	u.SetExcluded(order.FieldUserID)
	return u
}

// AddUserID adds v to the "user_id" field.
func (u *OrderUpsert) AddUserID(v int) *OrderUpsert {
	u.Add(order.FieldUserID, v)
	return u
}

// SetOrderType sets the "order_type" field.
func (u *OrderUpsert) SetOrderType(v int) *OrderUpsert {
	u.Set(order.FieldOrderType, v)
	return u
}

// UpdateOrderType sets the "order_type" field to the value that was provided on create.
func (u *OrderUpsert) UpdateOrderType() *OrderUpsert {
	u.SetExcluded(order.FieldOrderType)
	return u
}

// AddOrderType adds v to the "order_type" field.
func (u *OrderUpsert) AddOrderType(v int) *OrderUpsert {
	u.Add(order.FieldOrderType, v)
	return u
}

// SetStatus sets the "status" field.
func (u *OrderUpsert) SetStatus(v string) *OrderUpsert {
	u.Set(order.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OrderUpsert) UpdateStatus() *OrderUpsert {
	u.SetExcluded(order.FieldStatus)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderUpsert) SetCreatedAt(v time.Time) *OrderUpsert {
	u.Set(order.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderUpsert) UpdateCreatedAt() *OrderUpsert {
	u.SetExcluded(order.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderUpsert) SetUpdatedAt(v time.Time) *OrderUpsert {
	u.Set(order.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderUpsert) UpdateUpdatedAt() *OrderUpsert {
	u.SetExcluded(order.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OrderUpsertOne) UpdateNewValues() *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Order.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrderUpsertOne) Ignore() *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderUpsertOne) DoNothing() *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderCreate.OnConflict
// documentation for more info.
func (u *OrderUpsertOne) Update(set func(*OrderUpsert)) *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *OrderUpsertOne) SetOrderID(v int) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetOrderID(v)
	})
}

// AddOrderID adds v to the "order_id" field.
func (u *OrderUpsertOne) AddOrderID(v int) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.AddOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateOrderID() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateOrderID()
	})
}

// SetTableID sets the "table_id" field.
func (u *OrderUpsertOne) SetTableID(v int) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetTableID(v)
	})
}

// AddTableID adds v to the "table_id" field.
func (u *OrderUpsertOne) AddTableID(v int) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.AddTableID(v)
	})
}

// UpdateTableID sets the "table_id" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateTableID() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateTableID()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrderUpsertOne) SetUserID(v int) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *OrderUpsertOne) AddUserID(v int) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateUserID() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateUserID()
	})
}

// SetOrderType sets the "order_type" field.
func (u *OrderUpsertOne) SetOrderType(v int) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetOrderType(v)
	})
}

// AddOrderType adds v to the "order_type" field.
func (u *OrderUpsertOne) AddOrderType(v int) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.AddOrderType(v)
	})
}

// UpdateOrderType sets the "order_type" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateOrderType() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateOrderType()
	})
}

// SetStatus sets the "status" field.
func (u *OrderUpsertOne) SetStatus(v string) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateStatus() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateStatus()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderUpsertOne) SetCreatedAt(v time.Time) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateCreatedAt() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderUpsertOne) SetUpdatedAt(v time.Time) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateUpdatedAt() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *OrderUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	err      error
	builders []*OrderCreate
	conflict []sql.ConflictOption
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ocb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Order.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
func (ocb *OrderCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderUpsertBulk {
	ocb.conflict = opts
	return &OrderUpsertBulk{
		create: ocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ocb *OrderCreateBulk) OnConflictColumns(columns ...string) *OrderUpsertBulk {
	ocb.conflict = append(ocb.conflict, sql.ConflictColumns(columns...))
	return &OrderUpsertBulk{
		create: ocb,
	}
}

// OrderUpsertBulk is the builder for "upsert"-ing
// a bulk of Order nodes.
type OrderUpsertBulk struct {
	create *OrderCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OrderUpsertBulk) UpdateNewValues() *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrderUpsertBulk) Ignore() *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderUpsertBulk) DoNothing() *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderCreateBulk.OnConflict
// documentation for more info.
func (u *OrderUpsertBulk) Update(set func(*OrderUpsert)) *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *OrderUpsertBulk) SetOrderID(v int) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetOrderID(v)
	})
}

// AddOrderID adds v to the "order_id" field.
func (u *OrderUpsertBulk) AddOrderID(v int) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.AddOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateOrderID() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateOrderID()
	})
}

// SetTableID sets the "table_id" field.
func (u *OrderUpsertBulk) SetTableID(v int) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetTableID(v)
	})
}

// AddTableID adds v to the "table_id" field.
func (u *OrderUpsertBulk) AddTableID(v int) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.AddTableID(v)
	})
}

// UpdateTableID sets the "table_id" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateTableID() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateTableID()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrderUpsertBulk) SetUserID(v int) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *OrderUpsertBulk) AddUserID(v int) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateUserID() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateUserID()
	})
}

// SetOrderType sets the "order_type" field.
func (u *OrderUpsertBulk) SetOrderType(v int) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetOrderType(v)
	})
}

// AddOrderType adds v to the "order_type" field.
func (u *OrderUpsertBulk) AddOrderType(v int) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.AddOrderType(v)
	})
}

// UpdateOrderType sets the "order_type" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateOrderType() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateOrderType()
	})
}

// SetStatus sets the "status" field.
func (u *OrderUpsertBulk) SetStatus(v string) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateStatus() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateStatus()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderUpsertBulk) SetCreatedAt(v time.Time) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateCreatedAt() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderUpsertBulk) SetUpdatedAt(v time.Time) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateUpdatedAt() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *OrderUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrderCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
