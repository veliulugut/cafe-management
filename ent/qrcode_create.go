// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/qrcode"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QrCodeCreate is the builder for creating a QrCode entity.
type QrCodeCreate struct {
	config
	mutation *QrCodeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetURL sets the "url" field.
func (qcc *QrCodeCreate) SetURL(s string) *QrCodeCreate {
	qcc.mutation.SetURL(s)
	return qcc
}

// SetImage sets the "image" field.
func (qcc *QrCodeCreate) SetImage(s string) *QrCodeCreate {
	qcc.mutation.SetImage(s)
	return qcc
}

// Mutation returns the QrCodeMutation object of the builder.
func (qcc *QrCodeCreate) Mutation() *QrCodeMutation {
	return qcc.mutation
}

// Save creates the QrCode in the database.
func (qcc *QrCodeCreate) Save(ctx context.Context) (*QrCode, error) {
	return withHooks(ctx, qcc.sqlSave, qcc.mutation, qcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (qcc *QrCodeCreate) SaveX(ctx context.Context) *QrCode {
	v, err := qcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qcc *QrCodeCreate) Exec(ctx context.Context) error {
	_, err := qcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcc *QrCodeCreate) ExecX(ctx context.Context) {
	if err := qcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qcc *QrCodeCreate) check() error {
	if _, ok := qcc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "QrCode.url"`)}
	}
	if _, ok := qcc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "QrCode.image"`)}
	}
	return nil
}

func (qcc *QrCodeCreate) sqlSave(ctx context.Context) (*QrCode, error) {
	if err := qcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := qcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	qcc.mutation.id = &_node.ID
	qcc.mutation.done = true
	return _node, nil
}

func (qcc *QrCodeCreate) createSpec() (*QrCode, *sqlgraph.CreateSpec) {
	var (
		_node = &QrCode{config: qcc.config}
		_spec = sqlgraph.NewCreateSpec(qrcode.Table, sqlgraph.NewFieldSpec(qrcode.FieldID, field.TypeInt))
	)
	_spec.OnConflict = qcc.conflict
	if value, ok := qcc.mutation.URL(); ok {
		_spec.SetField(qrcode.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := qcc.mutation.Image(); ok {
		_spec.SetField(qrcode.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.QrCode.Create().
//		SetURL(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.QrCodeUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
func (qcc *QrCodeCreate) OnConflict(opts ...sql.ConflictOption) *QrCodeUpsertOne {
	qcc.conflict = opts
	return &QrCodeUpsertOne{
		create: qcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.QrCode.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (qcc *QrCodeCreate) OnConflictColumns(columns ...string) *QrCodeUpsertOne {
	qcc.conflict = append(qcc.conflict, sql.ConflictColumns(columns...))
	return &QrCodeUpsertOne{
		create: qcc,
	}
}

type (
	// QrCodeUpsertOne is the builder for "upsert"-ing
	//  one QrCode node.
	QrCodeUpsertOne struct {
		create *QrCodeCreate
	}

	// QrCodeUpsert is the "OnConflict" setter.
	QrCodeUpsert struct {
		*sql.UpdateSet
	}
)

// SetURL sets the "url" field.
func (u *QrCodeUpsert) SetURL(v string) *QrCodeUpsert {
	u.Set(qrcode.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *QrCodeUpsert) UpdateURL() *QrCodeUpsert {
	u.SetExcluded(qrcode.FieldURL)
	return u
}

// SetImage sets the "image" field.
func (u *QrCodeUpsert) SetImage(v string) *QrCodeUpsert {
	u.Set(qrcode.FieldImage, v)
	return u
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *QrCodeUpsert) UpdateImage() *QrCodeUpsert {
	u.SetExcluded(qrcode.FieldImage)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.QrCode.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *QrCodeUpsertOne) UpdateNewValues() *QrCodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.QrCode.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *QrCodeUpsertOne) Ignore() *QrCodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *QrCodeUpsertOne) DoNothing() *QrCodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the QrCodeCreate.OnConflict
// documentation for more info.
func (u *QrCodeUpsertOne) Update(set func(*QrCodeUpsert)) *QrCodeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&QrCodeUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *QrCodeUpsertOne) SetURL(v string) *QrCodeUpsertOne {
	return u.Update(func(s *QrCodeUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *QrCodeUpsertOne) UpdateURL() *QrCodeUpsertOne {
	return u.Update(func(s *QrCodeUpsert) {
		s.UpdateURL()
	})
}

// SetImage sets the "image" field.
func (u *QrCodeUpsertOne) SetImage(v string) *QrCodeUpsertOne {
	return u.Update(func(s *QrCodeUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *QrCodeUpsertOne) UpdateImage() *QrCodeUpsertOne {
	return u.Update(func(s *QrCodeUpsert) {
		s.UpdateImage()
	})
}

// Exec executes the query.
func (u *QrCodeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for QrCodeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *QrCodeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *QrCodeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *QrCodeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// QrCodeCreateBulk is the builder for creating many QrCode entities in bulk.
type QrCodeCreateBulk struct {
	config
	err      error
	builders []*QrCodeCreate
	conflict []sql.ConflictOption
}

// Save creates the QrCode entities in the database.
func (qccb *QrCodeCreateBulk) Save(ctx context.Context) ([]*QrCode, error) {
	if qccb.err != nil {
		return nil, qccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(qccb.builders))
	nodes := make([]*QrCode, len(qccb.builders))
	mutators := make([]Mutator, len(qccb.builders))
	for i := range qccb.builders {
		func(i int, root context.Context) {
			builder := qccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QrCodeMutation)
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
					_, err = mutators[i+1].Mutate(root, qccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = qccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, qccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qccb *QrCodeCreateBulk) SaveX(ctx context.Context) []*QrCode {
	v, err := qccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qccb *QrCodeCreateBulk) Exec(ctx context.Context) error {
	_, err := qccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qccb *QrCodeCreateBulk) ExecX(ctx context.Context) {
	if err := qccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.QrCode.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.QrCodeUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
func (qccb *QrCodeCreateBulk) OnConflict(opts ...sql.ConflictOption) *QrCodeUpsertBulk {
	qccb.conflict = opts
	return &QrCodeUpsertBulk{
		create: qccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.QrCode.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (qccb *QrCodeCreateBulk) OnConflictColumns(columns ...string) *QrCodeUpsertBulk {
	qccb.conflict = append(qccb.conflict, sql.ConflictColumns(columns...))
	return &QrCodeUpsertBulk{
		create: qccb,
	}
}

// QrCodeUpsertBulk is the builder for "upsert"-ing
// a bulk of QrCode nodes.
type QrCodeUpsertBulk struct {
	create *QrCodeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.QrCode.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *QrCodeUpsertBulk) UpdateNewValues() *QrCodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.QrCode.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *QrCodeUpsertBulk) Ignore() *QrCodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *QrCodeUpsertBulk) DoNothing() *QrCodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the QrCodeCreateBulk.OnConflict
// documentation for more info.
func (u *QrCodeUpsertBulk) Update(set func(*QrCodeUpsert)) *QrCodeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&QrCodeUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *QrCodeUpsertBulk) SetURL(v string) *QrCodeUpsertBulk {
	return u.Update(func(s *QrCodeUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *QrCodeUpsertBulk) UpdateURL() *QrCodeUpsertBulk {
	return u.Update(func(s *QrCodeUpsert) {
		s.UpdateURL()
	})
}

// SetImage sets the "image" field.
func (u *QrCodeUpsertBulk) SetImage(v string) *QrCodeUpsertBulk {
	return u.Update(func(s *QrCodeUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *QrCodeUpsertBulk) UpdateImage() *QrCodeUpsertBulk {
	return u.Update(func(s *QrCodeUpsert) {
		s.UpdateImage()
	})
}

// Exec executes the query.
func (u *QrCodeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the QrCodeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for QrCodeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *QrCodeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
