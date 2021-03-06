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
	"github.com/facebookincubator/symphony/pkg/ent/equipment"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentposition"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentpositiondefinition"
)

// EquipmentPositionCreate is the builder for creating a EquipmentPosition entity.
type EquipmentPositionCreate struct {
	config
	mutation *EquipmentPositionMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (epc *EquipmentPositionCreate) SetCreateTime(t time.Time) *EquipmentPositionCreate {
	epc.mutation.SetCreateTime(t)
	return epc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (epc *EquipmentPositionCreate) SetNillableCreateTime(t *time.Time) *EquipmentPositionCreate {
	if t != nil {
		epc.SetCreateTime(*t)
	}
	return epc
}

// SetUpdateTime sets the update_time field.
func (epc *EquipmentPositionCreate) SetUpdateTime(t time.Time) *EquipmentPositionCreate {
	epc.mutation.SetUpdateTime(t)
	return epc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (epc *EquipmentPositionCreate) SetNillableUpdateTime(t *time.Time) *EquipmentPositionCreate {
	if t != nil {
		epc.SetUpdateTime(*t)
	}
	return epc
}

// SetDefinitionID sets the definition edge to EquipmentPositionDefinition by id.
func (epc *EquipmentPositionCreate) SetDefinitionID(id int) *EquipmentPositionCreate {
	epc.mutation.SetDefinitionID(id)
	return epc
}

// SetDefinition sets the definition edge to EquipmentPositionDefinition.
func (epc *EquipmentPositionCreate) SetDefinition(e *EquipmentPositionDefinition) *EquipmentPositionCreate {
	return epc.SetDefinitionID(e.ID)
}

// SetParentID sets the parent edge to Equipment by id.
func (epc *EquipmentPositionCreate) SetParentID(id int) *EquipmentPositionCreate {
	epc.mutation.SetParentID(id)
	return epc
}

// SetNillableParentID sets the parent edge to Equipment by id if the given value is not nil.
func (epc *EquipmentPositionCreate) SetNillableParentID(id *int) *EquipmentPositionCreate {
	if id != nil {
		epc = epc.SetParentID(*id)
	}
	return epc
}

// SetParent sets the parent edge to Equipment.
func (epc *EquipmentPositionCreate) SetParent(e *Equipment) *EquipmentPositionCreate {
	return epc.SetParentID(e.ID)
}

// SetAttachmentID sets the attachment edge to Equipment by id.
func (epc *EquipmentPositionCreate) SetAttachmentID(id int) *EquipmentPositionCreate {
	epc.mutation.SetAttachmentID(id)
	return epc
}

// SetNillableAttachmentID sets the attachment edge to Equipment by id if the given value is not nil.
func (epc *EquipmentPositionCreate) SetNillableAttachmentID(id *int) *EquipmentPositionCreate {
	if id != nil {
		epc = epc.SetAttachmentID(*id)
	}
	return epc
}

// SetAttachment sets the attachment edge to Equipment.
func (epc *EquipmentPositionCreate) SetAttachment(e *Equipment) *EquipmentPositionCreate {
	return epc.SetAttachmentID(e.ID)
}

// Mutation returns the EquipmentPositionMutation object of the builder.
func (epc *EquipmentPositionCreate) Mutation() *EquipmentPositionMutation {
	return epc.mutation
}

// Save creates the EquipmentPosition in the database.
func (epc *EquipmentPositionCreate) Save(ctx context.Context) (*EquipmentPosition, error) {
	var (
		err  error
		node *EquipmentPosition
	)
	epc.defaults()
	if len(epc.hooks) == 0 {
		if err = epc.check(); err != nil {
			return nil, err
		}
		node, err = epc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = epc.check(); err != nil {
				return nil, err
			}
			epc.mutation = mutation
			node, err = epc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(epc.hooks) - 1; i >= 0; i-- {
			mut = epc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, epc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (epc *EquipmentPositionCreate) SaveX(ctx context.Context) *EquipmentPosition {
	v, err := epc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (epc *EquipmentPositionCreate) defaults() {
	if _, ok := epc.mutation.CreateTime(); !ok {
		v := equipmentposition.DefaultCreateTime()
		epc.mutation.SetCreateTime(v)
	}
	if _, ok := epc.mutation.UpdateTime(); !ok {
		v := equipmentposition.DefaultUpdateTime()
		epc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (epc *EquipmentPositionCreate) check() error {
	if _, ok := epc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := epc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := epc.mutation.DefinitionID(); !ok {
		return &ValidationError{Name: "definition", err: errors.New("ent: missing required edge \"definition\"")}
	}
	return nil
}

func (epc *EquipmentPositionCreate) sqlSave(ctx context.Context) (*EquipmentPosition, error) {
	_node, _spec := epc.createSpec()
	if err := sqlgraph.CreateNode(ctx, epc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (epc *EquipmentPositionCreate) createSpec() (*EquipmentPosition, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipmentPosition{config: epc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: equipmentposition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmentposition.FieldID,
			},
		}
	)
	if value, ok := epc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: equipmentposition.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := epc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: equipmentposition.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if nodes := epc.mutation.DefinitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   equipmentposition.DefinitionTable,
			Columns: []string{equipmentposition.DefinitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentpositiondefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipmentposition.ParentTable,
			Columns: []string{equipmentposition.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipmentposition.AttachmentTable,
			Columns: []string{equipmentposition.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
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

// EquipmentPositionCreateBulk is the builder for creating a bulk of EquipmentPosition entities.
type EquipmentPositionCreateBulk struct {
	config
	builders []*EquipmentPositionCreate
}

// Save creates the EquipmentPosition entities in the database.
func (epcb *EquipmentPositionCreateBulk) Save(ctx context.Context) ([]*EquipmentPosition, error) {
	specs := make([]*sqlgraph.CreateSpec, len(epcb.builders))
	nodes := make([]*EquipmentPosition, len(epcb.builders))
	mutators := make([]Mutator, len(epcb.builders))
	for i := range epcb.builders {
		func(i int, root context.Context) {
			builder := epcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentPositionMutation)
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
					_, err = mutators[i+1].Mutate(root, epcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, epcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, epcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (epcb *EquipmentPositionCreateBulk) SaveX(ctx context.Context) []*EquipmentPosition {
	v, err := epcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
