// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/menu"
	"cafe-management/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuUpdate is the builder for updating Menu entities.
type MenuUpdate struct {
	config
	hooks    []Hook
	mutation *MenuMutation
}

// Where appends a list predicates to the MenuUpdate builder.
func (mu *MenuUpdate) Where(ps ...predicate.Menu) *MenuUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetMenuID sets the "menu_id" field.
func (mu *MenuUpdate) SetMenuID(i int) *MenuUpdate {
	mu.mutation.ResetMenuID()
	mu.mutation.SetMenuID(i)
	return mu
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableMenuID(i *int) *MenuUpdate {
	if i != nil {
		mu.SetMenuID(*i)
	}
	return mu
}

// AddMenuID adds i to the "menu_id" field.
func (mu *MenuUpdate) AddMenuID(i int) *MenuUpdate {
	mu.mutation.AddMenuID(i)
	return mu
}

// SetName sets the "name" field.
func (mu *MenuUpdate) SetName(s string) *MenuUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableName(s *string) *MenuUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetCategory sets the "category" field.
func (mu *MenuUpdate) SetCategory(s string) *MenuUpdate {
	mu.mutation.SetCategory(s)
	return mu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableCategory(s *string) *MenuUpdate {
	if s != nil {
		mu.SetCategory(*s)
	}
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MenuUpdate) SetCreatedAt(t time.Time) *MenuUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableCreatedAt(t *time.Time) *MenuUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MenuUpdate) SetUpdatedAt(t time.Time) *MenuUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableUpdatedAt(t *time.Time) *MenuUpdate {
	if t != nil {
		mu.SetUpdatedAt(*t)
	}
	return mu
}

// SetDescription sets the "description" field.
func (mu *MenuUpdate) SetDescription(s string) *MenuUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableDescription(s *string) *MenuUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// SetPrice sets the "price" field.
func (mu *MenuUpdate) SetPrice(f float64) *MenuUpdate {
	mu.mutation.ResetPrice()
	mu.mutation.SetPrice(f)
	return mu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (mu *MenuUpdate) SetNillablePrice(f *float64) *MenuUpdate {
	if f != nil {
		mu.SetPrice(*f)
	}
	return mu
}

// AddPrice adds f to the "price" field.
func (mu *MenuUpdate) AddPrice(f float64) *MenuUpdate {
	mu.mutation.AddPrice(f)
	return mu
}

// SetMenuImageURL sets the "menu_image_url" field.
func (mu *MenuUpdate) SetMenuImageURL(s string) *MenuUpdate {
	mu.mutation.SetMenuImageURL(s)
	return mu
}

// SetNillableMenuImageURL sets the "menu_image_url" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableMenuImageURL(s *string) *MenuUpdate {
	if s != nil {
		mu.SetMenuImageURL(*s)
	}
	return mu
}

// Mutation returns the MenuMutation object of the builder.
func (mu *MenuUpdate) Mutation() *MenuMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MenuUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MenuUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MenuUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MenuUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.MenuID(); ok {
		_spec.SetField(menu.FieldMenuID, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedMenuID(); ok {
		_spec.AddField(menu.FieldMenuID, field.TypeInt, value)
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Category(); ok {
		_spec.SetField(menu.FieldCategory, field.TypeString, value)
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(menu.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(menu.FieldDescription, field.TypeString, value)
	}
	if value, ok := mu.mutation.Price(); ok {
		_spec.SetField(menu.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.AddedPrice(); ok {
		_spec.AddField(menu.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.MenuImageURL(); ok {
		_spec.SetField(menu.FieldMenuImageURL, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MenuUpdateOne is the builder for updating a single Menu entity.
type MenuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MenuMutation
}

// SetMenuID sets the "menu_id" field.
func (muo *MenuUpdateOne) SetMenuID(i int) *MenuUpdateOne {
	muo.mutation.ResetMenuID()
	muo.mutation.SetMenuID(i)
	return muo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableMenuID(i *int) *MenuUpdateOne {
	if i != nil {
		muo.SetMenuID(*i)
	}
	return muo
}

// AddMenuID adds i to the "menu_id" field.
func (muo *MenuUpdateOne) AddMenuID(i int) *MenuUpdateOne {
	muo.mutation.AddMenuID(i)
	return muo
}

// SetName sets the "name" field.
func (muo *MenuUpdateOne) SetName(s string) *MenuUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableName(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetCategory sets the "category" field.
func (muo *MenuUpdateOne) SetCategory(s string) *MenuUpdateOne {
	muo.mutation.SetCategory(s)
	return muo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableCategory(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetCategory(*s)
	}
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MenuUpdateOne) SetCreatedAt(t time.Time) *MenuUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableCreatedAt(t *time.Time) *MenuUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MenuUpdateOne) SetUpdatedAt(t time.Time) *MenuUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableUpdatedAt(t *time.Time) *MenuUpdateOne {
	if t != nil {
		muo.SetUpdatedAt(*t)
	}
	return muo
}

// SetDescription sets the "description" field.
func (muo *MenuUpdateOne) SetDescription(s string) *MenuUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableDescription(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// SetPrice sets the "price" field.
func (muo *MenuUpdateOne) SetPrice(f float64) *MenuUpdateOne {
	muo.mutation.ResetPrice()
	muo.mutation.SetPrice(f)
	return muo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillablePrice(f *float64) *MenuUpdateOne {
	if f != nil {
		muo.SetPrice(*f)
	}
	return muo
}

// AddPrice adds f to the "price" field.
func (muo *MenuUpdateOne) AddPrice(f float64) *MenuUpdateOne {
	muo.mutation.AddPrice(f)
	return muo
}

// SetMenuImageURL sets the "menu_image_url" field.
func (muo *MenuUpdateOne) SetMenuImageURL(s string) *MenuUpdateOne {
	muo.mutation.SetMenuImageURL(s)
	return muo
}

// SetNillableMenuImageURL sets the "menu_image_url" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableMenuImageURL(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetMenuImageURL(*s)
	}
	return muo
}

// Mutation returns the MenuMutation object of the builder.
func (muo *MenuUpdateOne) Mutation() *MenuMutation {
	return muo.mutation
}

// Where appends a list predicates to the MenuUpdate builder.
func (muo *MenuUpdateOne) Where(ps ...predicate.Menu) *MenuUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MenuUpdateOne) Select(field string, fields ...string) *MenuUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Menu entity.
func (muo *MenuUpdateOne) Save(ctx context.Context) (*Menu, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MenuUpdateOne) SaveX(ctx context.Context) *Menu {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MenuUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MenuUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MenuUpdateOne) sqlSave(ctx context.Context) (_node *Menu, err error) {
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Menu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menu.FieldID)
		for _, f := range fields {
			if !menu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != menu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.MenuID(); ok {
		_spec.SetField(menu.FieldMenuID, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedMenuID(); ok {
		_spec.AddField(menu.FieldMenuID, field.TypeInt, value)
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Category(); ok {
		_spec.SetField(menu.FieldCategory, field.TypeString, value)
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(menu.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(menu.FieldDescription, field.TypeString, value)
	}
	if value, ok := muo.mutation.Price(); ok {
		_spec.SetField(menu.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.AddedPrice(); ok {
		_spec.AddField(menu.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.MenuImageURL(); ok {
		_spec.SetField(menu.FieldMenuImageURL, field.TypeString, value)
	}
	_node = &Menu{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
