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
)

// KqiCategoryCreate is the builder for creating a KqiCategory entity.
type KqiCategoryCreate struct {
	config
	mutation *KqiCategoryMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (kcc *KqiCategoryCreate) SetCreateTime(t time.Time) *KqiCategoryCreate {
	kcc.mutation.SetCreateTime(t)
	return kcc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (kcc *KqiCategoryCreate) SetNillableCreateTime(t *time.Time) *KqiCategoryCreate {
	if t != nil {
		kcc.SetCreateTime(*t)
	}
	return kcc
}

// SetUpdateTime sets the update_time field.
func (kcc *KqiCategoryCreate) SetUpdateTime(t time.Time) *KqiCategoryCreate {
	kcc.mutation.SetUpdateTime(t)
	return kcc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (kcc *KqiCategoryCreate) SetNillableUpdateTime(t *time.Time) *KqiCategoryCreate {
	if t != nil {
		kcc.SetUpdateTime(*t)
	}
	return kcc
}

// SetName sets the name field.
func (kcc *KqiCategoryCreate) SetName(s string) *KqiCategoryCreate {
	kcc.mutation.SetName(s)
	return kcc
}

// AddKqiCategoryFkIDs adds the kqiCategoryFk edge to Kqi by ids.
func (kcc *KqiCategoryCreate) AddKqiCategoryFkIDs(ids ...int) *KqiCategoryCreate {
	kcc.mutation.AddKqiCategoryFkIDs(ids...)
	return kcc
}

// AddKqiCategoryFk adds the kqiCategoryFk edges to Kqi.
func (kcc *KqiCategoryCreate) AddKqiCategoryFk(k ...*Kqi) *KqiCategoryCreate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kcc.AddKqiCategoryFkIDs(ids...)
}

// Mutation returns the KqiCategoryMutation object of the builder.
func (kcc *KqiCategoryCreate) Mutation() *KqiCategoryMutation {
	return kcc.mutation
}

// Save creates the KqiCategory in the database.
func (kcc *KqiCategoryCreate) Save(ctx context.Context) (*KqiCategory, error) {
	var (
		err  error
		node *KqiCategory
	)
	kcc.defaults()
	if len(kcc.hooks) == 0 {
		if err = kcc.check(); err != nil {
			return nil, err
		}
		node, err = kcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KqiCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kcc.check(); err != nil {
				return nil, err
			}
			kcc.mutation = mutation
			node, err = kcc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kcc.hooks) - 1; i >= 0; i-- {
			mut = kcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kcc *KqiCategoryCreate) SaveX(ctx context.Context) *KqiCategory {
	v, err := kcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (kcc *KqiCategoryCreate) defaults() {
	if _, ok := kcc.mutation.CreateTime(); !ok {
		v := kqicategory.DefaultCreateTime()
		kcc.mutation.SetCreateTime(v)
	}
	if _, ok := kcc.mutation.UpdateTime(); !ok {
		v := kqicategory.DefaultUpdateTime()
		kcc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kcc *KqiCategoryCreate) check() error {
	if _, ok := kcc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := kcc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := kcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := kcc.mutation.Name(); ok {
		if err := kqicategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (kcc *KqiCategoryCreate) sqlSave(ctx context.Context) (*KqiCategory, error) {
	_node, _spec := kcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kcc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kcc *KqiCategoryCreate) createSpec() (*KqiCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &KqiCategory{config: kcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kqicategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kqicategory.FieldID,
			},
		}
	)
	if value, ok := kcc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kqicategory.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := kcc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kqicategory.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := kcc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kqicategory.FieldName,
		})
		_node.Name = value
	}
	if nodes := kcc.mutation.KqiCategoryFkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kqicategory.KqiCategoryFkTable,
			Columns: []string{kqicategory.KqiCategoryFkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kqi.FieldID,
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

// KqiCategoryCreateBulk is the builder for creating a bulk of KqiCategory entities.
type KqiCategoryCreateBulk struct {
	config
	builders []*KqiCategoryCreate
}

// Save creates the KqiCategory entities in the database.
func (kccb *KqiCategoryCreateBulk) Save(ctx context.Context) ([]*KqiCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kccb.builders))
	nodes := make([]*KqiCategory, len(kccb.builders))
	mutators := make([]Mutator, len(kccb.builders))
	for i := range kccb.builders {
		func(i int, root context.Context) {
			builder := kccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KqiCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, kccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (kccb *KqiCategoryCreateBulk) SaveX(ctx context.Context) []*KqiCategory {
	v, err := kccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
