// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentport"
	"github.com/facebookincubator/symphony/pkg/ent/link"
	"github.com/facebookincubator/symphony/pkg/ent/property"
	"github.com/facebookincubator/symphony/pkg/ent/schema/enum"
	"github.com/facebookincubator/symphony/pkg/ent/service"
	"github.com/facebookincubator/symphony/pkg/ent/workorder"
)

// LinkCreate is the builder for creating a Link entity.
type LinkCreate struct {
	config
	mutation *LinkMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (lc *LinkCreate) SetCreateTime(t time.Time) *LinkCreate {
	lc.mutation.SetCreateTime(t)
	return lc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (lc *LinkCreate) SetNillableCreateTime(t *time.Time) *LinkCreate {
	if t != nil {
		lc.SetCreateTime(*t)
	}
	return lc
}

// SetUpdateTime sets the update_time field.
func (lc *LinkCreate) SetUpdateTime(t time.Time) *LinkCreate {
	lc.mutation.SetUpdateTime(t)
	return lc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (lc *LinkCreate) SetNillableUpdateTime(t *time.Time) *LinkCreate {
	if t != nil {
		lc.SetUpdateTime(*t)
	}
	return lc
}

// SetFutureState sets the future_state field.
func (lc *LinkCreate) SetFutureState(es enum.FutureState) *LinkCreate {
	lc.mutation.SetFutureState(es)
	return lc
}

// SetNillableFutureState sets the future_state field if the given value is not nil.
func (lc *LinkCreate) SetNillableFutureState(es *enum.FutureState) *LinkCreate {
	if es != nil {
		lc.SetFutureState(*es)
	}
	return lc
}

// AddPortIDs adds the ports edge to EquipmentPort by ids.
func (lc *LinkCreate) AddPortIDs(ids ...int) *LinkCreate {
	lc.mutation.AddPortIDs(ids...)
	return lc
}

// AddPorts adds the ports edges to EquipmentPort.
func (lc *LinkCreate) AddPorts(e ...*EquipmentPort) *LinkCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return lc.AddPortIDs(ids...)
}

// SetWorkOrderID sets the work_order edge to WorkOrder by id.
func (lc *LinkCreate) SetWorkOrderID(id int) *LinkCreate {
	lc.mutation.SetWorkOrderID(id)
	return lc
}

// SetNillableWorkOrderID sets the work_order edge to WorkOrder by id if the given value is not nil.
func (lc *LinkCreate) SetNillableWorkOrderID(id *int) *LinkCreate {
	if id != nil {
		lc = lc.SetWorkOrderID(*id)
	}
	return lc
}

// SetWorkOrder sets the work_order edge to WorkOrder.
func (lc *LinkCreate) SetWorkOrder(w *WorkOrder) *LinkCreate {
	return lc.SetWorkOrderID(w.ID)
}

// AddPropertyIDs adds the properties edge to Property by ids.
func (lc *LinkCreate) AddPropertyIDs(ids ...int) *LinkCreate {
	lc.mutation.AddPropertyIDs(ids...)
	return lc
}

// AddProperties adds the properties edges to Property.
func (lc *LinkCreate) AddProperties(p ...*Property) *LinkCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lc.AddPropertyIDs(ids...)
}

// AddServiceIDs adds the service edge to Service by ids.
func (lc *LinkCreate) AddServiceIDs(ids ...int) *LinkCreate {
	lc.mutation.AddServiceIDs(ids...)
	return lc
}

// AddService adds the service edges to Service.
func (lc *LinkCreate) AddService(s ...*Service) *LinkCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return lc.AddServiceIDs(ids...)
}

// Mutation returns the LinkMutation object of the builder.
func (lc *LinkCreate) Mutation() *LinkMutation {
	return lc.mutation
}

// Save creates the Link in the database.
func (lc *LinkCreate) Save(ctx context.Context) (*Link, error) {
	var (
		err  error
		node *Link
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			node, err = lc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LinkCreate) SaveX(ctx context.Context) *Link {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (lc *LinkCreate) defaults() {
	if _, ok := lc.mutation.CreateTime(); !ok {
		v := link.DefaultCreateTime()
		lc.mutation.SetCreateTime(v)
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		v := link.DefaultUpdateTime()
		lc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LinkCreate) check() error {
	if _, ok := lc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := lc.mutation.FutureState(); ok {
		if err := link.FutureStateValidator(v); err != nil {
			return &ValidationError{Name: "future_state", err: fmt.Errorf("ent: validator failed for field \"future_state\": %w", err)}
		}
	}
	return nil
}

func (lc *LinkCreate) sqlSave(ctx context.Context) (*Link, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lc *LinkCreate) createSpec() (*Link, *sqlgraph.CreateSpec) {
	var (
		_node = &Link{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: link.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: link.FieldID,
			},
		}
	)
	if value, ok := lc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: link.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := lc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: link.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := lc.mutation.FutureState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: link.FieldFutureState,
		})
		_node.FutureState = &value
	}
	if nodes := lc.mutation.PortsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   link.PortsTable,
			Columns: []string{link.PortsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.WorkOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   link.WorkOrderTable,
			Columns: []string{link.WorkOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   link.PropertiesTable,
			Columns: []string{link.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: property.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.ServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   link.ServiceTable,
			Columns: link.ServicePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LinkCreateBulk is the builder for creating a bulk of Link entities.
type LinkCreateBulk struct {
	config
	builders []*LinkCreate
}

// Save creates the Link entities in the database.
func (lcb *LinkCreateBulk) Save(ctx context.Context) ([]*Link, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Link, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (lcb *LinkCreateBulk) SaveX(ctx context.Context) []*Link {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
