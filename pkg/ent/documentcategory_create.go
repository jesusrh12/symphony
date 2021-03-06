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
	"github.com/facebookincubator/symphony/pkg/ent/documentcategory"
	"github.com/facebookincubator/symphony/pkg/ent/file"
	"github.com/facebookincubator/symphony/pkg/ent/hyperlink"
	"github.com/facebookincubator/symphony/pkg/ent/locationtype"
)

// DocumentCategoryCreate is the builder for creating a DocumentCategory entity.
type DocumentCategoryCreate struct {
	config
	mutation *DocumentCategoryMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (dcc *DocumentCategoryCreate) SetCreateTime(t time.Time) *DocumentCategoryCreate {
	dcc.mutation.SetCreateTime(t)
	return dcc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (dcc *DocumentCategoryCreate) SetNillableCreateTime(t *time.Time) *DocumentCategoryCreate {
	if t != nil {
		dcc.SetCreateTime(*t)
	}
	return dcc
}

// SetUpdateTime sets the update_time field.
func (dcc *DocumentCategoryCreate) SetUpdateTime(t time.Time) *DocumentCategoryCreate {
	dcc.mutation.SetUpdateTime(t)
	return dcc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (dcc *DocumentCategoryCreate) SetNillableUpdateTime(t *time.Time) *DocumentCategoryCreate {
	if t != nil {
		dcc.SetUpdateTime(*t)
	}
	return dcc
}

// SetName sets the name field.
func (dcc *DocumentCategoryCreate) SetName(s string) *DocumentCategoryCreate {
	dcc.mutation.SetName(s)
	return dcc
}

// SetIndex sets the index field.
func (dcc *DocumentCategoryCreate) SetIndex(i int) *DocumentCategoryCreate {
	dcc.mutation.SetIndex(i)
	return dcc
}

// SetLocationTypeID sets the location_type edge to LocationType by id.
func (dcc *DocumentCategoryCreate) SetLocationTypeID(id int) *DocumentCategoryCreate {
	dcc.mutation.SetLocationTypeID(id)
	return dcc
}

// SetNillableLocationTypeID sets the location_type edge to LocationType by id if the given value is not nil.
func (dcc *DocumentCategoryCreate) SetNillableLocationTypeID(id *int) *DocumentCategoryCreate {
	if id != nil {
		dcc = dcc.SetLocationTypeID(*id)
	}
	return dcc
}

// SetLocationType sets the location_type edge to LocationType.
func (dcc *DocumentCategoryCreate) SetLocationType(l *LocationType) *DocumentCategoryCreate {
	return dcc.SetLocationTypeID(l.ID)
}

// AddFileIDs adds the files edge to File by ids.
func (dcc *DocumentCategoryCreate) AddFileIDs(ids ...int) *DocumentCategoryCreate {
	dcc.mutation.AddFileIDs(ids...)
	return dcc
}

// AddFiles adds the files edges to File.
func (dcc *DocumentCategoryCreate) AddFiles(f ...*File) *DocumentCategoryCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return dcc.AddFileIDs(ids...)
}

// AddHyperlinkIDs adds the hyperlinks edge to Hyperlink by ids.
func (dcc *DocumentCategoryCreate) AddHyperlinkIDs(ids ...int) *DocumentCategoryCreate {
	dcc.mutation.AddHyperlinkIDs(ids...)
	return dcc
}

// AddHyperlinks adds the hyperlinks edges to Hyperlink.
func (dcc *DocumentCategoryCreate) AddHyperlinks(h ...*Hyperlink) *DocumentCategoryCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return dcc.AddHyperlinkIDs(ids...)
}

// Mutation returns the DocumentCategoryMutation object of the builder.
func (dcc *DocumentCategoryCreate) Mutation() *DocumentCategoryMutation {
	return dcc.mutation
}

// Save creates the DocumentCategory in the database.
func (dcc *DocumentCategoryCreate) Save(ctx context.Context) (*DocumentCategory, error) {
	var (
		err  error
		node *DocumentCategory
	)
	dcc.defaults()
	if len(dcc.hooks) == 0 {
		if err = dcc.check(); err != nil {
			return nil, err
		}
		node, err = dcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DocumentCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dcc.check(); err != nil {
				return nil, err
			}
			dcc.mutation = mutation
			node, err = dcc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dcc.hooks) - 1; i >= 0; i-- {
			mut = dcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dcc *DocumentCategoryCreate) SaveX(ctx context.Context) *DocumentCategory {
	v, err := dcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (dcc *DocumentCategoryCreate) defaults() {
	if _, ok := dcc.mutation.CreateTime(); !ok {
		v := documentcategory.DefaultCreateTime()
		dcc.mutation.SetCreateTime(v)
	}
	if _, ok := dcc.mutation.UpdateTime(); !ok {
		v := documentcategory.DefaultUpdateTime()
		dcc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcc *DocumentCategoryCreate) check() error {
	if _, ok := dcc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := dcc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := dcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := dcc.mutation.Name(); ok {
		if err := documentcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := dcc.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New("ent: missing required field \"index\"")}
	}
	return nil
}

func (dcc *DocumentCategoryCreate) sqlSave(ctx context.Context) (*DocumentCategory, error) {
	_node, _spec := dcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dcc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dcc *DocumentCategoryCreate) createSpec() (*DocumentCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &DocumentCategory{config: dcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: documentcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: documentcategory.FieldID,
			},
		}
	)
	if value, ok := dcc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: documentcategory.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := dcc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: documentcategory.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := dcc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: documentcategory.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dcc.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: documentcategory.FieldIndex,
		})
		_node.Index = value
	}
	if nodes := dcc.mutation.LocationTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   documentcategory.LocationTypeTable,
			Columns: []string{documentcategory.LocationTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: locationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dcc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   documentcategory.FilesTable,
			Columns: []string{documentcategory.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dcc.mutation.HyperlinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   documentcategory.HyperlinksTable,
			Columns: []string{documentcategory.HyperlinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hyperlink.FieldID,
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

// DocumentCategoryCreateBulk is the builder for creating a bulk of DocumentCategory entities.
type DocumentCategoryCreateBulk struct {
	config
	builders []*DocumentCategoryCreate
}

// Save creates the DocumentCategory entities in the database.
func (dccb *DocumentCategoryCreateBulk) Save(ctx context.Context) ([]*DocumentCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dccb.builders))
	nodes := make([]*DocumentCategory, len(dccb.builders))
	mutators := make([]Mutator, len(dccb.builders))
	for i := range dccb.builders {
		func(i int, root context.Context) {
			builder := dccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DocumentCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, dccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (dccb *DocumentCategoryCreateBulk) SaveX(ctx context.Context) []*DocumentCategory {
	v, err := dccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
