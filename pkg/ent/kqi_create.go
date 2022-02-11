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
	"github.com/facebookincubator/symphony/pkg/ent/kqi"
	"github.com/facebookincubator/symphony/pkg/ent/kqicategory"
	"github.com/facebookincubator/symphony/pkg/ent/kqiperspective"
	"github.com/facebookincubator/symphony/pkg/ent/kqisource"
	"github.com/facebookincubator/symphony/pkg/ent/kqitarget"
	"github.com/facebookincubator/symphony/pkg/ent/kqitemporalfrequency"
)

// KqiCreate is the builder for creating a Kqi entity.
type KqiCreate struct {
	config
	mutation *KqiMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (kc *KqiCreate) SetCreateTime(t time.Time) *KqiCreate {
	kc.mutation.SetCreateTime(t)
	return kc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (kc *KqiCreate) SetNillableCreateTime(t *time.Time) *KqiCreate {
	if t != nil {
		kc.SetCreateTime(*t)
	}
	return kc
}

// SetUpdateTime sets the update_time field.
func (kc *KqiCreate) SetUpdateTime(t time.Time) *KqiCreate {
	kc.mutation.SetUpdateTime(t)
	return kc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (kc *KqiCreate) SetNillableUpdateTime(t *time.Time) *KqiCreate {
	if t != nil {
		kc.SetUpdateTime(*t)
	}
	return kc
}

// SetName sets the name field.
func (kc *KqiCreate) SetName(s string) *KqiCreate {
	kc.mutation.SetName(s)
	return kc
}

// SetDescription sets the description field.
func (kc *KqiCreate) SetDescription(s string) *KqiCreate {
	kc.mutation.SetDescription(s)
	return kc
}

// SetStartDateTime sets the startDateTime field.
func (kc *KqiCreate) SetStartDateTime(t time.Time) *KqiCreate {
	kc.mutation.SetStartDateTime(t)
	return kc
}

// SetEndDateTime sets the endDateTime field.
func (kc *KqiCreate) SetEndDateTime(t time.Time) *KqiCreate {
	kc.mutation.SetEndDateTime(t)
	return kc
}

// SetFormula sets the formula field.
func (kc *KqiCreate) SetFormula(s string) *KqiCreate {
	kc.mutation.SetFormula(s)
	return kc
}

// SetKqiCategoryFkID sets the kqiCategoryFk edge to KqiCategory by id.
func (kc *KqiCreate) SetKqiCategoryFkID(id int) *KqiCreate {
	kc.mutation.SetKqiCategoryFkID(id)
	return kc
}

// SetNillableKqiCategoryFkID sets the kqiCategoryFk edge to KqiCategory by id if the given value is not nil.
func (kc *KqiCreate) SetNillableKqiCategoryFkID(id *int) *KqiCreate {
	if id != nil {
		kc = kc.SetKqiCategoryFkID(*id)
	}
	return kc
}

// SetKqiCategoryFk sets the kqiCategoryFk edge to KqiCategory.
func (kc *KqiCreate) SetKqiCategoryFk(k *KqiCategory) *KqiCreate {
	return kc.SetKqiCategoryFkID(k.ID)
}

// SetKqiPerspectiveFkID sets the kqiPerspectiveFk edge to KqiPerspective by id.
func (kc *KqiCreate) SetKqiPerspectiveFkID(id int) *KqiCreate {
	kc.mutation.SetKqiPerspectiveFkID(id)
	return kc
}

// SetNillableKqiPerspectiveFkID sets the kqiPerspectiveFk edge to KqiPerspective by id if the given value is not nil.
func (kc *KqiCreate) SetNillableKqiPerspectiveFkID(id *int) *KqiCreate {
	if id != nil {
		kc = kc.SetKqiPerspectiveFkID(*id)
	}
	return kc
}

// SetKqiPerspectiveFk sets the kqiPerspectiveFk edge to KqiPerspective.
func (kc *KqiCreate) SetKqiPerspectiveFk(k *KqiPerspective) *KqiCreate {
	return kc.SetKqiPerspectiveFkID(k.ID)
}

// SetKqiSourceFkID sets the kqiSourceFk edge to KqiSource by id.
func (kc *KqiCreate) SetKqiSourceFkID(id int) *KqiCreate {
	kc.mutation.SetKqiSourceFkID(id)
	return kc
}

// SetNillableKqiSourceFkID sets the kqiSourceFk edge to KqiSource by id if the given value is not nil.
func (kc *KqiCreate) SetNillableKqiSourceFkID(id *int) *KqiCreate {
	if id != nil {
		kc = kc.SetKqiSourceFkID(*id)
	}
	return kc
}

// SetKqiSourceFk sets the kqiSourceFk edge to KqiSource.
func (kc *KqiCreate) SetKqiSourceFk(k *KqiSource) *KqiCreate {
	return kc.SetKqiSourceFkID(k.ID)
}

// SetKqiTemporalFrequencyFkID sets the kqiTemporalFrequencyFk edge to KqiTemporalFrequency by id.
func (kc *KqiCreate) SetKqiTemporalFrequencyFkID(id int) *KqiCreate {
	kc.mutation.SetKqiTemporalFrequencyFkID(id)
	return kc
}

// SetNillableKqiTemporalFrequencyFkID sets the kqiTemporalFrequencyFk edge to KqiTemporalFrequency by id if the given value is not nil.
func (kc *KqiCreate) SetNillableKqiTemporalFrequencyFkID(id *int) *KqiCreate {
	if id != nil {
		kc = kc.SetKqiTemporalFrequencyFkID(*id)
	}
	return kc
}

// SetKqiTemporalFrequencyFk sets the kqiTemporalFrequencyFk edge to KqiTemporalFrequency.
func (kc *KqiCreate) SetKqiTemporalFrequencyFk(k *KqiTemporalFrequency) *KqiCreate {
	return kc.SetKqiTemporalFrequencyFkID(k.ID)
}

// AddKqiTargetFkIDs adds the kqiTargetFk edge to KqiTarget by ids.
func (kc *KqiCreate) AddKqiTargetFkIDs(ids ...int) *KqiCreate {
	kc.mutation.AddKqiTargetFkIDs(ids...)
	return kc
}

// AddKqiTargetFk adds the kqiTargetFk edges to KqiTarget.
func (kc *KqiCreate) AddKqiTargetFk(k ...*KqiTarget) *KqiCreate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kc.AddKqiTargetFkIDs(ids...)
}

// Mutation returns the KqiMutation object of the builder.
func (kc *KqiCreate) Mutation() *KqiMutation {
	return kc.mutation
}

// Save creates the Kqi in the database.
func (kc *KqiCreate) Save(ctx context.Context) (*Kqi, error) {
	var (
		err  error
		node *Kqi
	)
	kc.defaults()
	if len(kc.hooks) == 0 {
		if err = kc.check(); err != nil {
			return nil, err
		}
		node, err = kc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KqiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kc.check(); err != nil {
				return nil, err
			}
			kc.mutation = mutation
			node, err = kc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kc.hooks) - 1; i >= 0; i-- {
			mut = kc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KqiCreate) SaveX(ctx context.Context) *Kqi {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (kc *KqiCreate) defaults() {
	if _, ok := kc.mutation.CreateTime(); !ok {
		v := kqi.DefaultCreateTime()
		kc.mutation.SetCreateTime(v)
	}
	if _, ok := kc.mutation.UpdateTime(); !ok {
		v := kqi.DefaultUpdateTime()
		kc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KqiCreate) check() error {
	if _, ok := kc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := kc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := kc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := kc.mutation.Name(); ok {
		if err := kqi.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := kc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := kc.mutation.StartDateTime(); !ok {
		return &ValidationError{Name: "startDateTime", err: errors.New("ent: missing required field \"startDateTime\"")}
	}
	if _, ok := kc.mutation.EndDateTime(); !ok {
		return &ValidationError{Name: "endDateTime", err: errors.New("ent: missing required field \"endDateTime\"")}
	}
	if _, ok := kc.mutation.Formula(); !ok {
		return &ValidationError{Name: "formula", err: errors.New("ent: missing required field \"formula\"")}
	}
	return nil
}

func (kc *KqiCreate) sqlSave(ctx context.Context) (*Kqi, error) {
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kc *KqiCreate) createSpec() (*Kqi, *sqlgraph.CreateSpec) {
	var (
		_node = &Kqi{config: kc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kqi.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kqi.FieldID,
			},
		}
	)
	if value, ok := kc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kqi.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := kc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kqi.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := kc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kqi.FieldName,
		})
		_node.Name = value
	}
	if value, ok := kc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kqi.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := kc.mutation.StartDateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kqi.FieldStartDateTime,
		})
		_node.StartDateTime = value
	}
	if value, ok := kc.mutation.EndDateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kqi.FieldEndDateTime,
		})
		_node.EndDateTime = value
	}
	if value, ok := kc.mutation.Formula(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kqi.FieldFormula,
		})
		_node.Formula = value
	}
	if nodes := kc.mutation.KqiCategoryFkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kqi.KqiCategoryFkTable,
			Columns: []string{kqi.KqiCategoryFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kqicategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kc.mutation.KqiPerspectiveFkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kqi.KqiPerspectiveFkTable,
			Columns: []string{kqi.KqiPerspectiveFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kqiperspective.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kc.mutation.KqiSourceFkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kqi.KqiSourceFkTable,
			Columns: []string{kqi.KqiSourceFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kqisource.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kc.mutation.KqiTemporalFrequencyFkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kqi.KqiTemporalFrequencyFkTable,
			Columns: []string{kqi.KqiTemporalFrequencyFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kqitemporalfrequency.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kc.mutation.KqiTargetFkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kqi.KqiTargetFkTable,
			Columns: []string{kqi.KqiTargetFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kqitarget.FieldID,
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

// KqiCreateBulk is the builder for creating a bulk of Kqi entities.
type KqiCreateBulk struct {
	config
	builders []*KqiCreate
}

// Save creates the Kqi entities in the database.
func (kcb *KqiCreateBulk) Save(ctx context.Context) ([]*Kqi, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Kqi, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KqiMutation)
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
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (kcb *KqiCreateBulk) SaveX(ctx context.Context) []*Kqi {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
