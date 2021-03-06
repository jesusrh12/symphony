// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/symphony/pkg/ent/kqi"
	"github.com/facebookincubator/symphony/pkg/ent/kqisource"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
)

// KqiSourceQuery is the builder for querying KqiSource entities.
type KqiSourceQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.KqiSource
	// eager-loading edges.
	withKqiSourceFk *KqiQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ksq *KqiSourceQuery) Where(ps ...predicate.KqiSource) *KqiSourceQuery {
	ksq.predicates = append(ksq.predicates, ps...)
	return ksq
}

// Limit adds a limit step to the query.
func (ksq *KqiSourceQuery) Limit(limit int) *KqiSourceQuery {
	ksq.limit = &limit
	return ksq
}

// Offset adds an offset step to the query.
func (ksq *KqiSourceQuery) Offset(offset int) *KqiSourceQuery {
	ksq.offset = &offset
	return ksq
}

// Order adds an order step to the query.
func (ksq *KqiSourceQuery) Order(o ...OrderFunc) *KqiSourceQuery {
	ksq.order = append(ksq.order, o...)
	return ksq
}

// QueryKqiSourceFk chains the current query on the kqiSourceFk edge.
func (ksq *KqiSourceQuery) QueryKqiSourceFk() *KqiQuery {
	query := &KqiQuery{config: ksq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ksq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kqisource.Table, kqisource.FieldID, selector),
			sqlgraph.To(kqi.Table, kqi.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, kqisource.KqiSourceFkTable, kqisource.KqiSourceFkColumn),
		)
		fromU = sqlgraph.SetNeighbors(ksq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first KqiSource entity in the query. Returns *NotFoundError when no kqisource was found.
func (ksq *KqiSourceQuery) First(ctx context.Context) (*KqiSource, error) {
	nodes, err := ksq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{kqisource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ksq *KqiSourceQuery) FirstX(ctx context.Context) *KqiSource {
	node, err := ksq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KqiSource id in the query. Returns *NotFoundError when no id was found.
func (ksq *KqiSourceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ksq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{kqisource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ksq *KqiSourceQuery) FirstIDX(ctx context.Context) int {
	id, err := ksq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only KqiSource entity in the query, returns an error if not exactly one entity was returned.
func (ksq *KqiSourceQuery) Only(ctx context.Context) (*KqiSource, error) {
	nodes, err := ksq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{kqisource.Label}
	default:
		return nil, &NotSingularError{kqisource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ksq *KqiSourceQuery) OnlyX(ctx context.Context) *KqiSource {
	node, err := ksq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only KqiSource id in the query, returns an error if not exactly one id was returned.
func (ksq *KqiSourceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ksq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{kqisource.Label}
	default:
		err = &NotSingularError{kqisource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ksq *KqiSourceQuery) OnlyIDX(ctx context.Context) int {
	id, err := ksq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KqiSources.
func (ksq *KqiSourceQuery) All(ctx context.Context) ([]*KqiSource, error) {
	if err := ksq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ksq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ksq *KqiSourceQuery) AllX(ctx context.Context) []*KqiSource {
	nodes, err := ksq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KqiSource ids.
func (ksq *KqiSourceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ksq.Select(kqisource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ksq *KqiSourceQuery) IDsX(ctx context.Context) []int {
	ids, err := ksq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ksq *KqiSourceQuery) Count(ctx context.Context) (int, error) {
	if err := ksq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ksq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ksq *KqiSourceQuery) CountX(ctx context.Context) int {
	count, err := ksq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ksq *KqiSourceQuery) Exist(ctx context.Context) (bool, error) {
	if err := ksq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ksq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ksq *KqiSourceQuery) ExistX(ctx context.Context) bool {
	exist, err := ksq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ksq *KqiSourceQuery) Clone() *KqiSourceQuery {
	if ksq == nil {
		return nil
	}
	return &KqiSourceQuery{
		config:          ksq.config,
		limit:           ksq.limit,
		offset:          ksq.offset,
		order:           append([]OrderFunc{}, ksq.order...),
		unique:          append([]string{}, ksq.unique...),
		predicates:      append([]predicate.KqiSource{}, ksq.predicates...),
		withKqiSourceFk: ksq.withKqiSourceFk.Clone(),
		// clone intermediate query.
		sql:  ksq.sql.Clone(),
		path: ksq.path,
	}
}

//  WithKqiSourceFk tells the query-builder to eager-loads the nodes that are connected to
// the "kqiSourceFk" edge. The optional arguments used to configure the query builder of the edge.
func (ksq *KqiSourceQuery) WithKqiSourceFk(opts ...func(*KqiQuery)) *KqiSourceQuery {
	query := &KqiQuery{config: ksq.config}
	for _, opt := range opts {
		opt(query)
	}
	ksq.withKqiSourceFk = query
	return ksq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.KqiSource.Query().
//		GroupBy(kqisource.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ksq *KqiSourceQuery) GroupBy(field string, fields ...string) *KqiSourceGroupBy {
	group := &KqiSourceGroupBy{config: ksq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ksq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.KqiSource.Query().
//		Select(kqisource.FieldCreateTime).
//		Scan(ctx, &v)
//
func (ksq *KqiSourceQuery) Select(field string, fields ...string) *KqiSourceSelect {
	selector := &KqiSourceSelect{config: ksq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ksq.sqlQuery(), nil
	}
	return selector
}

func (ksq *KqiSourceQuery) prepareQuery(ctx context.Context) error {
	if ksq.path != nil {
		prev, err := ksq.path(ctx)
		if err != nil {
			return err
		}
		ksq.sql = prev
	}
	if err := kqisource.Policy.EvalQuery(ctx, ksq); err != nil {
		return err
	}
	return nil
}

func (ksq *KqiSourceQuery) sqlAll(ctx context.Context) ([]*KqiSource, error) {
	var (
		nodes       = []*KqiSource{}
		_spec       = ksq.querySpec()
		loadedTypes = [1]bool{
			ksq.withKqiSourceFk != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &KqiSource{config: ksq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, ksq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ksq.withKqiSourceFk; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*KqiSource)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.KqiSourceFk = []*Kqi{}
		}
		query.withFKs = true
		query.Where(predicate.Kqi(func(s *sql.Selector) {
			s.Where(sql.InValues(kqisource.KqiSourceFkColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.kqi_source_kqi_source_fk
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "kqi_source_kqi_source_fk" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "kqi_source_kqi_source_fk" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.KqiSourceFk = append(node.Edges.KqiSourceFk, n)
		}
	}

	return nodes, nil
}

func (ksq *KqiSourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ksq.querySpec()
	return sqlgraph.CountNodes(ctx, ksq.driver, _spec)
}

func (ksq *KqiSourceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ksq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ksq *KqiSourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kqisource.Table,
			Columns: kqisource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kqisource.FieldID,
			},
		},
		From:   ksq.sql,
		Unique: true,
	}
	if ps := ksq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ksq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ksq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ksq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, kqisource.ValidColumn)
			}
		}
	}
	return _spec
}

func (ksq *KqiSourceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ksq.driver.Dialect())
	t1 := builder.Table(kqisource.Table)
	selector := builder.Select(t1.Columns(kqisource.Columns...)...).From(t1)
	if ksq.sql != nil {
		selector = ksq.sql
		selector.Select(selector.Columns(kqisource.Columns...)...)
	}
	for _, p := range ksq.predicates {
		p(selector)
	}
	for _, p := range ksq.order {
		p(selector, kqisource.ValidColumn)
	}
	if offset := ksq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ksq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KqiSourceGroupBy is the builder for group-by KqiSource entities.
type KqiSourceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ksgb *KqiSourceGroupBy) Aggregate(fns ...AggregateFunc) *KqiSourceGroupBy {
	ksgb.fns = append(ksgb.fns, fns...)
	return ksgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ksgb *KqiSourceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ksgb.path(ctx)
	if err != nil {
		return err
	}
	ksgb.sql = query
	return ksgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ksgb *KqiSourceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ksgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ksgb *KqiSourceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ksgb.fields) > 1 {
		return nil, errors.New("ent: KqiSourceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ksgb *KqiSourceGroupBy) StringsX(ctx context.Context) []string {
	v, err := ksgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ksgb *KqiSourceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ksgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqisource.Label}
	default:
		err = fmt.Errorf("ent: KqiSourceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ksgb *KqiSourceGroupBy) StringX(ctx context.Context) string {
	v, err := ksgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ksgb *KqiSourceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ksgb.fields) > 1 {
		return nil, errors.New("ent: KqiSourceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ksgb *KqiSourceGroupBy) IntsX(ctx context.Context) []int {
	v, err := ksgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ksgb *KqiSourceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ksgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqisource.Label}
	default:
		err = fmt.Errorf("ent: KqiSourceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ksgb *KqiSourceGroupBy) IntX(ctx context.Context) int {
	v, err := ksgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ksgb *KqiSourceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ksgb.fields) > 1 {
		return nil, errors.New("ent: KqiSourceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ksgb *KqiSourceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ksgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ksgb *KqiSourceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ksgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqisource.Label}
	default:
		err = fmt.Errorf("ent: KqiSourceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ksgb *KqiSourceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ksgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ksgb *KqiSourceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ksgb.fields) > 1 {
		return nil, errors.New("ent: KqiSourceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ksgb *KqiSourceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ksgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ksgb *KqiSourceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ksgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqisource.Label}
	default:
		err = fmt.Errorf("ent: KqiSourceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ksgb *KqiSourceGroupBy) BoolX(ctx context.Context) bool {
	v, err := ksgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ksgb *KqiSourceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ksgb.fields {
		if !kqisource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ksgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ksgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ksgb *KqiSourceGroupBy) sqlQuery() *sql.Selector {
	selector := ksgb.sql
	columns := make([]string, 0, len(ksgb.fields)+len(ksgb.fns))
	columns = append(columns, ksgb.fields...)
	for _, fn := range ksgb.fns {
		columns = append(columns, fn(selector, kqisource.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ksgb.fields...)
}

// KqiSourceSelect is the builder for select fields of KqiSource entities.
type KqiSourceSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (kss *KqiSourceSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := kss.path(ctx)
	if err != nil {
		return err
	}
	kss.sql = query
	return kss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kss *KqiSourceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := kss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (kss *KqiSourceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(kss.fields) > 1 {
		return nil, errors.New("ent: KqiSourceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := kss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kss *KqiSourceSelect) StringsX(ctx context.Context) []string {
	v, err := kss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (kss *KqiSourceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqisource.Label}
	default:
		err = fmt.Errorf("ent: KqiSourceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kss *KqiSourceSelect) StringX(ctx context.Context) string {
	v, err := kss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (kss *KqiSourceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(kss.fields) > 1 {
		return nil, errors.New("ent: KqiSourceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := kss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kss *KqiSourceSelect) IntsX(ctx context.Context) []int {
	v, err := kss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (kss *KqiSourceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqisource.Label}
	default:
		err = fmt.Errorf("ent: KqiSourceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kss *KqiSourceSelect) IntX(ctx context.Context) int {
	v, err := kss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (kss *KqiSourceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(kss.fields) > 1 {
		return nil, errors.New("ent: KqiSourceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := kss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kss *KqiSourceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := kss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (kss *KqiSourceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqisource.Label}
	default:
		err = fmt.Errorf("ent: KqiSourceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kss *KqiSourceSelect) Float64X(ctx context.Context) float64 {
	v, err := kss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (kss *KqiSourceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(kss.fields) > 1 {
		return nil, errors.New("ent: KqiSourceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := kss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kss *KqiSourceSelect) BoolsX(ctx context.Context) []bool {
	v, err := kss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (kss *KqiSourceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqisource.Label}
	default:
		err = fmt.Errorf("ent: KqiSourceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kss *KqiSourceSelect) BoolX(ctx context.Context) bool {
	v, err := kss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kss *KqiSourceSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range kss.fields {
		if !kqisource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := kss.sqlQuery().Query()
	if err := kss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kss *KqiSourceSelect) sqlQuery() sql.Querier {
	selector := kss.sql
	selector.Select(selector.Columns(kss.fields...)...)
	return selector
}
