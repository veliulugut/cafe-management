// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cafe-management/ent/predicate"
	"cafe-management/ent/qrcode"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QrCodeQuery is the builder for querying QrCode entities.
type QrCodeQuery struct {
	config
	ctx        *QueryContext
	order      []qrcode.OrderOption
	inters     []Interceptor
	predicates []predicate.QrCode
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QrCodeQuery builder.
func (qcq *QrCodeQuery) Where(ps ...predicate.QrCode) *QrCodeQuery {
	qcq.predicates = append(qcq.predicates, ps...)
	return qcq
}

// Limit the number of records to be returned by this query.
func (qcq *QrCodeQuery) Limit(limit int) *QrCodeQuery {
	qcq.ctx.Limit = &limit
	return qcq
}

// Offset to start from.
func (qcq *QrCodeQuery) Offset(offset int) *QrCodeQuery {
	qcq.ctx.Offset = &offset
	return qcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qcq *QrCodeQuery) Unique(unique bool) *QrCodeQuery {
	qcq.ctx.Unique = &unique
	return qcq
}

// Order specifies how the records should be ordered.
func (qcq *QrCodeQuery) Order(o ...qrcode.OrderOption) *QrCodeQuery {
	qcq.order = append(qcq.order, o...)
	return qcq
}

// First returns the first QrCode entity from the query.
// Returns a *NotFoundError when no QrCode was found.
func (qcq *QrCodeQuery) First(ctx context.Context) (*QrCode, error) {
	nodes, err := qcq.Limit(1).All(setContextOp(ctx, qcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{qrcode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qcq *QrCodeQuery) FirstX(ctx context.Context) *QrCode {
	node, err := qcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first QrCode ID from the query.
// Returns a *NotFoundError when no QrCode ID was found.
func (qcq *QrCodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qcq.Limit(1).IDs(setContextOp(ctx, qcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{qrcode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qcq *QrCodeQuery) FirstIDX(ctx context.Context) int {
	id, err := qcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single QrCode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one QrCode entity is found.
// Returns a *NotFoundError when no QrCode entities are found.
func (qcq *QrCodeQuery) Only(ctx context.Context) (*QrCode, error) {
	nodes, err := qcq.Limit(2).All(setContextOp(ctx, qcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{qrcode.Label}
	default:
		return nil, &NotSingularError{qrcode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qcq *QrCodeQuery) OnlyX(ctx context.Context) *QrCode {
	node, err := qcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only QrCode ID in the query.
// Returns a *NotSingularError when more than one QrCode ID is found.
// Returns a *NotFoundError when no entities are found.
func (qcq *QrCodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qcq.Limit(2).IDs(setContextOp(ctx, qcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{qrcode.Label}
	default:
		err = &NotSingularError{qrcode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qcq *QrCodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := qcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of QrCodes.
func (qcq *QrCodeQuery) All(ctx context.Context) ([]*QrCode, error) {
	ctx = setContextOp(ctx, qcq.ctx, "All")
	if err := qcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*QrCode, *QrCodeQuery]()
	return withInterceptors[[]*QrCode](ctx, qcq, qr, qcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (qcq *QrCodeQuery) AllX(ctx context.Context) []*QrCode {
	nodes, err := qcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of QrCode IDs.
func (qcq *QrCodeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if qcq.ctx.Unique == nil && qcq.path != nil {
		qcq.Unique(true)
	}
	ctx = setContextOp(ctx, qcq.ctx, "IDs")
	if err = qcq.Select(qrcode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qcq *QrCodeQuery) IDsX(ctx context.Context) []int {
	ids, err := qcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qcq *QrCodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, qcq.ctx, "Count")
	if err := qcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, qcq, querierCount[*QrCodeQuery](), qcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (qcq *QrCodeQuery) CountX(ctx context.Context) int {
	count, err := qcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qcq *QrCodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, qcq.ctx, "Exist")
	switch _, err := qcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (qcq *QrCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := qcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QrCodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qcq *QrCodeQuery) Clone() *QrCodeQuery {
	if qcq == nil {
		return nil
	}
	return &QrCodeQuery{
		config:     qcq.config,
		ctx:        qcq.ctx.Clone(),
		order:      append([]qrcode.OrderOption{}, qcq.order...),
		inters:     append([]Interceptor{}, qcq.inters...),
		predicates: append([]predicate.QrCode{}, qcq.predicates...),
		// clone intermediate query.
		sql:  qcq.sql.Clone(),
		path: qcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.QrCode.Query().
//		GroupBy(qrcode.FieldURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (qcq *QrCodeQuery) GroupBy(field string, fields ...string) *QrCodeGroupBy {
	qcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &QrCodeGroupBy{build: qcq}
	grbuild.flds = &qcq.ctx.Fields
	grbuild.label = qrcode.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//	}
//
//	client.QrCode.Query().
//		Select(qrcode.FieldURL).
//		Scan(ctx, &v)
func (qcq *QrCodeQuery) Select(fields ...string) *QrCodeSelect {
	qcq.ctx.Fields = append(qcq.ctx.Fields, fields...)
	sbuild := &QrCodeSelect{QrCodeQuery: qcq}
	sbuild.label = qrcode.Label
	sbuild.flds, sbuild.scan = &qcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a QrCodeSelect configured with the given aggregations.
func (qcq *QrCodeQuery) Aggregate(fns ...AggregateFunc) *QrCodeSelect {
	return qcq.Select().Aggregate(fns...)
}

func (qcq *QrCodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range qcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, qcq); err != nil {
				return err
			}
		}
	}
	for _, f := range qcq.ctx.Fields {
		if !qrcode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qcq.path != nil {
		prev, err := qcq.path(ctx)
		if err != nil {
			return err
		}
		qcq.sql = prev
	}
	return nil
}

func (qcq *QrCodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*QrCode, error) {
	var (
		nodes = []*QrCode{}
		_spec = qcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*QrCode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &QrCode{config: qcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (qcq *QrCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qcq.querySpec()
	_spec.Node.Columns = qcq.ctx.Fields
	if len(qcq.ctx.Fields) > 0 {
		_spec.Unique = qcq.ctx.Unique != nil && *qcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, qcq.driver, _spec)
}

func (qcq *QrCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(qrcode.Table, qrcode.Columns, sqlgraph.NewFieldSpec(qrcode.FieldID, field.TypeInt))
	_spec.From = qcq.sql
	if unique := qcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if qcq.path != nil {
		_spec.Unique = true
	}
	if fields := qcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, qrcode.FieldID)
		for i := range fields {
			if fields[i] != qrcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qcq *QrCodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qcq.driver.Dialect())
	t1 := builder.Table(qrcode.Table)
	columns := qcq.ctx.Fields
	if len(columns) == 0 {
		columns = qrcode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qcq.sql != nil {
		selector = qcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qcq.ctx.Unique != nil && *qcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range qcq.predicates {
		p(selector)
	}
	for _, p := range qcq.order {
		p(selector)
	}
	if offset := qcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QrCodeGroupBy is the group-by builder for QrCode entities.
type QrCodeGroupBy struct {
	selector
	build *QrCodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qcgb *QrCodeGroupBy) Aggregate(fns ...AggregateFunc) *QrCodeGroupBy {
	qcgb.fns = append(qcgb.fns, fns...)
	return qcgb
}

// Scan applies the selector query and scans the result into the given value.
func (qcgb *QrCodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qcgb.build.ctx, "GroupBy")
	if err := qcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QrCodeQuery, *QrCodeGroupBy](ctx, qcgb.build, qcgb, qcgb.build.inters, v)
}

func (qcgb *QrCodeGroupBy) sqlScan(ctx context.Context, root *QrCodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(qcgb.fns))
	for _, fn := range qcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*qcgb.flds)+len(qcgb.fns))
		for _, f := range *qcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*qcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// QrCodeSelect is the builder for selecting fields of QrCode entities.
type QrCodeSelect struct {
	*QrCodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (qcs *QrCodeSelect) Aggregate(fns ...AggregateFunc) *QrCodeSelect {
	qcs.fns = append(qcs.fns, fns...)
	return qcs
}

// Scan applies the selector query and scans the result into the given value.
func (qcs *QrCodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qcs.ctx, "Select")
	if err := qcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QrCodeQuery, *QrCodeSelect](ctx, qcs.QrCodeQuery, qcs, qcs.inters, v)
}

func (qcs *QrCodeSelect) sqlScan(ctx context.Context, root *QrCodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(qcs.fns))
	for _, fn := range qcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*qcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}