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
	"github.com/facebookincubator/symphony/pkg/ent/eventseverity"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/rule"
)

// EventSeverityQuery is the builder for querying EventSeverity entities.
type EventSeverityQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.EventSeverity
	// eager-loading edges.
	withEventseverityrule *RuleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (esq *EventSeverityQuery) Where(ps ...predicate.EventSeverity) *EventSeverityQuery {
	esq.predicates = append(esq.predicates, ps...)
	return esq
}

// Limit adds a limit step to the query.
func (esq *EventSeverityQuery) Limit(limit int) *EventSeverityQuery {
	esq.limit = &limit
	return esq
}

// Offset adds an offset step to the query.
func (esq *EventSeverityQuery) Offset(offset int) *EventSeverityQuery {
	esq.offset = &offset
	return esq
}

// Order adds an order step to the query.
func (esq *EventSeverityQuery) Order(o ...OrderFunc) *EventSeverityQuery {
	esq.order = append(esq.order, o...)
	return esq
}

// QueryEventseverityrule chains the current query on the eventseverityrule edge.
func (esq *EventSeverityQuery) QueryEventseverityrule() *RuleQuery {
	query := &RuleQuery{config: esq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := esq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := esq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eventseverity.Table, eventseverity.FieldID, selector),
			sqlgraph.To(rule.Table, rule.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, eventseverity.EventseverityruleTable, eventseverity.EventseverityruleColumn),
		)
		fromU = sqlgraph.SetNeighbors(esq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EventSeverity entity in the query. Returns *NotFoundError when no eventseverity was found.
func (esq *EventSeverityQuery) First(ctx context.Context) (*EventSeverity, error) {
	nodes, err := esq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{eventseverity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (esq *EventSeverityQuery) FirstX(ctx context.Context) *EventSeverity {
	node, err := esq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EventSeverity id in the query. Returns *NotFoundError when no id was found.
func (esq *EventSeverityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = esq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{eventseverity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (esq *EventSeverityQuery) FirstIDX(ctx context.Context) int {
	id, err := esq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only EventSeverity entity in the query, returns an error if not exactly one entity was returned.
func (esq *EventSeverityQuery) Only(ctx context.Context) (*EventSeverity, error) {
	nodes, err := esq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{eventseverity.Label}
	default:
		return nil, &NotSingularError{eventseverity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (esq *EventSeverityQuery) OnlyX(ctx context.Context) *EventSeverity {
	node, err := esq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only EventSeverity id in the query, returns an error if not exactly one id was returned.
func (esq *EventSeverityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = esq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{eventseverity.Label}
	default:
		err = &NotSingularError{eventseverity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (esq *EventSeverityQuery) OnlyIDX(ctx context.Context) int {
	id, err := esq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EventSeverities.
func (esq *EventSeverityQuery) All(ctx context.Context) ([]*EventSeverity, error) {
	if err := esq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return esq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (esq *EventSeverityQuery) AllX(ctx context.Context) []*EventSeverity {
	nodes, err := esq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EventSeverity ids.
func (esq *EventSeverityQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := esq.Select(eventseverity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (esq *EventSeverityQuery) IDsX(ctx context.Context) []int {
	ids, err := esq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (esq *EventSeverityQuery) Count(ctx context.Context) (int, error) {
	if err := esq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return esq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (esq *EventSeverityQuery) CountX(ctx context.Context) int {
	count, err := esq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (esq *EventSeverityQuery) Exist(ctx context.Context) (bool, error) {
	if err := esq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return esq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (esq *EventSeverityQuery) ExistX(ctx context.Context) bool {
	exist, err := esq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (esq *EventSeverityQuery) Clone() *EventSeverityQuery {
	if esq == nil {
		return nil
	}
	return &EventSeverityQuery{
		config:                esq.config,
		limit:                 esq.limit,
		offset:                esq.offset,
		order:                 append([]OrderFunc{}, esq.order...),
		unique:                append([]string{}, esq.unique...),
		predicates:            append([]predicate.EventSeverity{}, esq.predicates...),
		withEventseverityrule: esq.withEventseverityrule.Clone(),
		// clone intermediate query.
		sql:  esq.sql.Clone(),
		path: esq.path,
	}
}

//  WithEventseverityrule tells the query-builder to eager-loads the nodes that are connected to
// the "eventseverityrule" edge. The optional arguments used to configure the query builder of the edge.
func (esq *EventSeverityQuery) WithEventseverityrule(opts ...func(*RuleQuery)) *EventSeverityQuery {
	query := &RuleQuery{config: esq.config}
	for _, opt := range opts {
		opt(query)
	}
	esq.withEventseverityrule = query
	return esq
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
//	client.EventSeverity.Query().
//		GroupBy(eventseverity.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (esq *EventSeverityQuery) GroupBy(field string, fields ...string) *EventSeverityGroupBy {
	group := &EventSeverityGroupBy{config: esq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := esq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return esq.sqlQuery(), nil
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
//	client.EventSeverity.Query().
//		Select(eventseverity.FieldCreateTime).
//		Scan(ctx, &v)
//
func (esq *EventSeverityQuery) Select(field string, fields ...string) *EventSeveritySelect {
	selector := &EventSeveritySelect{config: esq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := esq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return esq.sqlQuery(), nil
	}
	return selector
}

func (esq *EventSeverityQuery) prepareQuery(ctx context.Context) error {
	if esq.path != nil {
		prev, err := esq.path(ctx)
		if err != nil {
			return err
		}
		esq.sql = prev
	}
	if err := eventseverity.Policy.EvalQuery(ctx, esq); err != nil {
		return err
	}
	return nil
}

func (esq *EventSeverityQuery) sqlAll(ctx context.Context) ([]*EventSeverity, error) {
	var (
		nodes       = []*EventSeverity{}
		_spec       = esq.querySpec()
		loadedTypes = [1]bool{
			esq.withEventseverityrule != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &EventSeverity{config: esq.config}
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
	if err := sqlgraph.QueryNodes(ctx, esq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := esq.withEventseverityrule; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*EventSeverity)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Eventseverityrule = []*Rule{}
		}
		query.withFKs = true
		query.Where(predicate.Rule(func(s *sql.Selector) {
			s.Where(sql.InValues(eventseverity.EventseverityruleColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.event_severity_eventseverityrule
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "event_severity_eventseverityrule" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "event_severity_eventseverityrule" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Eventseverityrule = append(node.Edges.Eventseverityrule, n)
		}
	}

	return nodes, nil
}

func (esq *EventSeverityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := esq.querySpec()
	return sqlgraph.CountNodes(ctx, esq.driver, _spec)
}

func (esq *EventSeverityQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := esq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (esq *EventSeverityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventseverity.Table,
			Columns: eventseverity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventseverity.FieldID,
			},
		},
		From:   esq.sql,
		Unique: true,
	}
	if ps := esq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := esq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := esq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := esq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, eventseverity.ValidColumn)
			}
		}
	}
	return _spec
}

func (esq *EventSeverityQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(esq.driver.Dialect())
	t1 := builder.Table(eventseverity.Table)
	selector := builder.Select(t1.Columns(eventseverity.Columns...)...).From(t1)
	if esq.sql != nil {
		selector = esq.sql
		selector.Select(selector.Columns(eventseverity.Columns...)...)
	}
	for _, p := range esq.predicates {
		p(selector)
	}
	for _, p := range esq.order {
		p(selector, eventseverity.ValidColumn)
	}
	if offset := esq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := esq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EventSeverityGroupBy is the builder for group-by EventSeverity entities.
type EventSeverityGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (esgb *EventSeverityGroupBy) Aggregate(fns ...AggregateFunc) *EventSeverityGroupBy {
	esgb.fns = append(esgb.fns, fns...)
	return esgb
}

// Scan applies the group-by query and scan the result into the given value.
func (esgb *EventSeverityGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := esgb.path(ctx)
	if err != nil {
		return err
	}
	esgb.sql = query
	return esgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (esgb *EventSeverityGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := esgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (esgb *EventSeverityGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(esgb.fields) > 1 {
		return nil, errors.New("ent: EventSeverityGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := esgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (esgb *EventSeverityGroupBy) StringsX(ctx context.Context) []string {
	v, err := esgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (esgb *EventSeverityGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = esgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventseverity.Label}
	default:
		err = fmt.Errorf("ent: EventSeverityGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (esgb *EventSeverityGroupBy) StringX(ctx context.Context) string {
	v, err := esgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (esgb *EventSeverityGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(esgb.fields) > 1 {
		return nil, errors.New("ent: EventSeverityGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := esgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (esgb *EventSeverityGroupBy) IntsX(ctx context.Context) []int {
	v, err := esgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (esgb *EventSeverityGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = esgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventseverity.Label}
	default:
		err = fmt.Errorf("ent: EventSeverityGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (esgb *EventSeverityGroupBy) IntX(ctx context.Context) int {
	v, err := esgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (esgb *EventSeverityGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(esgb.fields) > 1 {
		return nil, errors.New("ent: EventSeverityGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := esgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (esgb *EventSeverityGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := esgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (esgb *EventSeverityGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = esgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventseverity.Label}
	default:
		err = fmt.Errorf("ent: EventSeverityGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (esgb *EventSeverityGroupBy) Float64X(ctx context.Context) float64 {
	v, err := esgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (esgb *EventSeverityGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(esgb.fields) > 1 {
		return nil, errors.New("ent: EventSeverityGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := esgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (esgb *EventSeverityGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := esgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (esgb *EventSeverityGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = esgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventseverity.Label}
	default:
		err = fmt.Errorf("ent: EventSeverityGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (esgb *EventSeverityGroupBy) BoolX(ctx context.Context) bool {
	v, err := esgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (esgb *EventSeverityGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range esgb.fields {
		if !eventseverity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := esgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := esgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (esgb *EventSeverityGroupBy) sqlQuery() *sql.Selector {
	selector := esgb.sql
	columns := make([]string, 0, len(esgb.fields)+len(esgb.fns))
	columns = append(columns, esgb.fields...)
	for _, fn := range esgb.fns {
		columns = append(columns, fn(selector, eventseverity.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(esgb.fields...)
}

// EventSeveritySelect is the builder for select fields of EventSeverity entities.
type EventSeveritySelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ess *EventSeveritySelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ess.path(ctx)
	if err != nil {
		return err
	}
	ess.sql = query
	return ess.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ess *EventSeveritySelect) ScanX(ctx context.Context, v interface{}) {
	if err := ess.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ess *EventSeveritySelect) Strings(ctx context.Context) ([]string, error) {
	if len(ess.fields) > 1 {
		return nil, errors.New("ent: EventSeveritySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ess.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ess *EventSeveritySelect) StringsX(ctx context.Context) []string {
	v, err := ess.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ess *EventSeveritySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ess.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventseverity.Label}
	default:
		err = fmt.Errorf("ent: EventSeveritySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ess *EventSeveritySelect) StringX(ctx context.Context) string {
	v, err := ess.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ess *EventSeveritySelect) Ints(ctx context.Context) ([]int, error) {
	if len(ess.fields) > 1 {
		return nil, errors.New("ent: EventSeveritySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ess.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ess *EventSeveritySelect) IntsX(ctx context.Context) []int {
	v, err := ess.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ess *EventSeveritySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ess.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventseverity.Label}
	default:
		err = fmt.Errorf("ent: EventSeveritySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ess *EventSeveritySelect) IntX(ctx context.Context) int {
	v, err := ess.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ess *EventSeveritySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ess.fields) > 1 {
		return nil, errors.New("ent: EventSeveritySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ess.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ess *EventSeveritySelect) Float64sX(ctx context.Context) []float64 {
	v, err := ess.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ess *EventSeveritySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ess.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventseverity.Label}
	default:
		err = fmt.Errorf("ent: EventSeveritySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ess *EventSeveritySelect) Float64X(ctx context.Context) float64 {
	v, err := ess.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ess *EventSeveritySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ess.fields) > 1 {
		return nil, errors.New("ent: EventSeveritySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ess.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ess *EventSeveritySelect) BoolsX(ctx context.Context) []bool {
	v, err := ess.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ess *EventSeveritySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ess.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventseverity.Label}
	default:
		err = fmt.Errorf("ent: EventSeveritySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ess *EventSeveritySelect) BoolX(ctx context.Context) bool {
	v, err := ess.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ess *EventSeveritySelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ess.fields {
		if !eventseverity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := ess.sqlQuery().Query()
	if err := ess.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ess *EventSeveritySelect) sqlQuery() sql.Querier {
	selector := ess.sql
	selector.Select(selector.Columns(ess.fields...)...)
	return selector
}
