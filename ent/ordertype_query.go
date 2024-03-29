// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/ordertype"
	"cafe-management/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderTypeQuery is the builder for querying OrderType entities.
type OrderTypeQuery struct {
	config
	ctx        *QueryContext
	order      []ordertype.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderTypeQuery builder.
func (otq *OrderTypeQuery) Where(ps ...predicate.OrderType) *OrderTypeQuery {
	otq.predicates = append(otq.predicates, ps...)
	return otq
}

// Limit the number of records to be returned by this query.
func (otq *OrderTypeQuery) Limit(limit int) *OrderTypeQuery {
	otq.ctx.Limit = &limit
	return otq
}

// Offset to start from.
func (otq *OrderTypeQuery) Offset(offset int) *OrderTypeQuery {
	otq.ctx.Offset = &offset
	return otq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (otq *OrderTypeQuery) Unique(unique bool) *OrderTypeQuery {
	otq.ctx.Unique = &unique
	return otq
}

// Order specifies how the records should be ordered.
func (otq *OrderTypeQuery) Order(o ...ordertype.OrderOption) *OrderTypeQuery {
	otq.order = append(otq.order, o...)
	return otq
}

// First returns the first OrderType entity from the query.
// Returns a *NotFoundError when no OrderType was found.
func (otq *OrderTypeQuery) First(ctx context.Context) (*OrderType, error) {
	nodes, err := otq.Limit(1).All(setContextOp(ctx, otq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ordertype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (otq *OrderTypeQuery) FirstX(ctx context.Context) *OrderType {
	node, err := otq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderType ID from the query.
// Returns a *NotFoundError when no OrderType ID was found.
func (otq *OrderTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(1).IDs(setContextOp(ctx, otq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ordertype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (otq *OrderTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := otq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderType entity is found.
// Returns a *NotFoundError when no OrderType entities are found.
func (otq *OrderTypeQuery) Only(ctx context.Context) (*OrderType, error) {
	nodes, err := otq.Limit(2).All(setContextOp(ctx, otq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ordertype.Label}
	default:
		return nil, &NotSingularError{ordertype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (otq *OrderTypeQuery) OnlyX(ctx context.Context) *OrderType {
	node, err := otq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderType ID in the query.
// Returns a *NotSingularError when more than one OrderType ID is found.
// Returns a *NotFoundError when no entities are found.
func (otq *OrderTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(2).IDs(setContextOp(ctx, otq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ordertype.Label}
	default:
		err = &NotSingularError{ordertype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (otq *OrderTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := otq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderTypes.
func (otq *OrderTypeQuery) All(ctx context.Context) ([]*OrderType, error) {
	ctx = setContextOp(ctx, otq.ctx, "All")
	if err := otq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderType, *OrderTypeQuery]()
	return withInterceptors[[]*OrderType](ctx, otq, qr, otq.inters)
}

// AllX is like All, but panics if an error occurs.
func (otq *OrderTypeQuery) AllX(ctx context.Context) []*OrderType {
	nodes, err := otq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderType IDs.
func (otq *OrderTypeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if otq.ctx.Unique == nil && otq.path != nil {
		otq.Unique(true)
	}
	ctx = setContextOp(ctx, otq.ctx, "IDs")
	if err = otq.Select(ordertype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (otq *OrderTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := otq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (otq *OrderTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, otq.ctx, "Count")
	if err := otq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, otq, querierCount[*OrderTypeQuery](), otq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (otq *OrderTypeQuery) CountX(ctx context.Context) int {
	count, err := otq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (otq *OrderTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, otq.ctx, "Exist")
	switch _, err := otq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (otq *OrderTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := otq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (otq *OrderTypeQuery) Clone() *OrderTypeQuery {
	if otq == nil {
		return nil
	}
	return &OrderTypeQuery{
		config:     otq.config,
		ctx:        otq.ctx.Clone(),
		order:      append([]ordertype.OrderOption{}, otq.order...),
		inters:     append([]Interceptor{}, otq.inters...),
		predicates: append([]predicate.OrderType{}, otq.predicates...),
		// clone intermediate query.
		sql:  otq.sql.Clone(),
		path: otq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderType.Query().
//		GroupBy(ordertype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (otq *OrderTypeQuery) GroupBy(field string, fields ...string) *OrderTypeGroupBy {
	otq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderTypeGroupBy{build: otq}
	grbuild.flds = &otq.ctx.Fields
	grbuild.label = ordertype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.OrderType.Query().
//		Select(ordertype.FieldName).
//		Scan(ctx, &v)
func (otq *OrderTypeQuery) Select(fields ...string) *OrderTypeSelect {
	otq.ctx.Fields = append(otq.ctx.Fields, fields...)
	sbuild := &OrderTypeSelect{OrderTypeQuery: otq}
	sbuild.label = ordertype.Label
	sbuild.flds, sbuild.scan = &otq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderTypeSelect configured with the given aggregations.
func (otq *OrderTypeQuery) Aggregate(fns ...AggregateFunc) *OrderTypeSelect {
	return otq.Select().Aggregate(fns...)
}

func (otq *OrderTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range otq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, otq); err != nil {
				return err
			}
		}
	}
	for _, f := range otq.ctx.Fields {
		if !ordertype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if otq.path != nil {
		prev, err := otq.path(ctx)
		if err != nil {
			return err
		}
		otq.sql = prev
	}
	return nil
}

func (otq *OrderTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderType, error) {
	var (
		nodes = []*OrderType{}
		_spec = otq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderType{config: otq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, otq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (otq *OrderTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := otq.querySpec()
	_spec.Node.Columns = otq.ctx.Fields
	if len(otq.ctx.Fields) > 0 {
		_spec.Unique = otq.ctx.Unique != nil && *otq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, otq.driver, _spec)
}

func (otq *OrderTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ordertype.Table, ordertype.Columns, sqlgraph.NewFieldSpec(ordertype.FieldID, field.TypeInt))
	_spec.From = otq.sql
	if unique := otq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if otq.path != nil {
		_spec.Unique = true
	}
	if fields := otq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ordertype.FieldID)
		for i := range fields {
			if fields[i] != ordertype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := otq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := otq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := otq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := otq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (otq *OrderTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(otq.driver.Dialect())
	t1 := builder.Table(ordertype.Table)
	columns := otq.ctx.Fields
	if len(columns) == 0 {
		columns = ordertype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if otq.sql != nil {
		selector = otq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if otq.ctx.Unique != nil && *otq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range otq.predicates {
		p(selector)
	}
	for _, p := range otq.order {
		p(selector)
	}
	if offset := otq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := otq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderTypeGroupBy is the group-by builder for OrderType entities.
type OrderTypeGroupBy struct {
	selector
	build *OrderTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (otgb *OrderTypeGroupBy) Aggregate(fns ...AggregateFunc) *OrderTypeGroupBy {
	otgb.fns = append(otgb.fns, fns...)
	return otgb
}

// Scan applies the selector query and scans the result into the given value.
func (otgb *OrderTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, otgb.build.ctx, "GroupBy")
	if err := otgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderTypeQuery, *OrderTypeGroupBy](ctx, otgb.build, otgb, otgb.build.inters, v)
}

func (otgb *OrderTypeGroupBy) sqlScan(ctx context.Context, root *OrderTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(otgb.fns))
	for _, fn := range otgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*otgb.flds)+len(otgb.fns))
		for _, f := range *otgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*otgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := otgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderTypeSelect is the builder for selecting fields of OrderType entities.
type OrderTypeSelect struct {
	*OrderTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ots *OrderTypeSelect) Aggregate(fns ...AggregateFunc) *OrderTypeSelect {
	ots.fns = append(ots.fns, fns...)
	return ots
}

// Scan applies the selector query and scans the result into the given value.
func (ots *OrderTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ots.ctx, "Select")
	if err := ots.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderTypeQuery, *OrderTypeSelect](ctx, ots.OrderTypeQuery, ots, ots.inters, v)
}

func (ots *OrderTypeSelect) sqlScan(ctx context.Context, root *OrderTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ots.fns))
	for _, fn := range ots.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ots.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ots.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
