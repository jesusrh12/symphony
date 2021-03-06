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
	"github.com/facebookincubator/symphony/pkg/ent/checklistitem"
	"github.com/facebookincubator/symphony/pkg/ent/location"
	"github.com/facebookincubator/symphony/pkg/ent/surveyquestion"
	"github.com/facebookincubator/symphony/pkg/ent/surveywifiscan"
)

// SurveyWiFiScanCreate is the builder for creating a SurveyWiFiScan entity.
type SurveyWiFiScanCreate struct {
	config
	mutation *SurveyWiFiScanMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (swfsc *SurveyWiFiScanCreate) SetCreateTime(t time.Time) *SurveyWiFiScanCreate {
	swfsc.mutation.SetCreateTime(t)
	return swfsc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableCreateTime(t *time.Time) *SurveyWiFiScanCreate {
	if t != nil {
		swfsc.SetCreateTime(*t)
	}
	return swfsc
}

// SetUpdateTime sets the update_time field.
func (swfsc *SurveyWiFiScanCreate) SetUpdateTime(t time.Time) *SurveyWiFiScanCreate {
	swfsc.mutation.SetUpdateTime(t)
	return swfsc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableUpdateTime(t *time.Time) *SurveyWiFiScanCreate {
	if t != nil {
		swfsc.SetUpdateTime(*t)
	}
	return swfsc
}

// SetSsid sets the ssid field.
func (swfsc *SurveyWiFiScanCreate) SetSsid(s string) *SurveyWiFiScanCreate {
	swfsc.mutation.SetSsid(s)
	return swfsc
}

// SetNillableSsid sets the ssid field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableSsid(s *string) *SurveyWiFiScanCreate {
	if s != nil {
		swfsc.SetSsid(*s)
	}
	return swfsc
}

// SetBssid sets the bssid field.
func (swfsc *SurveyWiFiScanCreate) SetBssid(s string) *SurveyWiFiScanCreate {
	swfsc.mutation.SetBssid(s)
	return swfsc
}

// SetTimestamp sets the timestamp field.
func (swfsc *SurveyWiFiScanCreate) SetTimestamp(t time.Time) *SurveyWiFiScanCreate {
	swfsc.mutation.SetTimestamp(t)
	return swfsc
}

// SetFrequency sets the frequency field.
func (swfsc *SurveyWiFiScanCreate) SetFrequency(i int) *SurveyWiFiScanCreate {
	swfsc.mutation.SetFrequency(i)
	return swfsc
}

// SetChannel sets the channel field.
func (swfsc *SurveyWiFiScanCreate) SetChannel(i int) *SurveyWiFiScanCreate {
	swfsc.mutation.SetChannel(i)
	return swfsc
}

// SetBand sets the band field.
func (swfsc *SurveyWiFiScanCreate) SetBand(s string) *SurveyWiFiScanCreate {
	swfsc.mutation.SetBand(s)
	return swfsc
}

// SetNillableBand sets the band field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableBand(s *string) *SurveyWiFiScanCreate {
	if s != nil {
		swfsc.SetBand(*s)
	}
	return swfsc
}

// SetChannelWidth sets the channel_width field.
func (swfsc *SurveyWiFiScanCreate) SetChannelWidth(i int) *SurveyWiFiScanCreate {
	swfsc.mutation.SetChannelWidth(i)
	return swfsc
}

// SetNillableChannelWidth sets the channel_width field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableChannelWidth(i *int) *SurveyWiFiScanCreate {
	if i != nil {
		swfsc.SetChannelWidth(*i)
	}
	return swfsc
}

// SetCapabilities sets the capabilities field.
func (swfsc *SurveyWiFiScanCreate) SetCapabilities(s string) *SurveyWiFiScanCreate {
	swfsc.mutation.SetCapabilities(s)
	return swfsc
}

// SetNillableCapabilities sets the capabilities field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableCapabilities(s *string) *SurveyWiFiScanCreate {
	if s != nil {
		swfsc.SetCapabilities(*s)
	}
	return swfsc
}

// SetStrength sets the strength field.
func (swfsc *SurveyWiFiScanCreate) SetStrength(i int) *SurveyWiFiScanCreate {
	swfsc.mutation.SetStrength(i)
	return swfsc
}

// SetLatitude sets the latitude field.
func (swfsc *SurveyWiFiScanCreate) SetLatitude(f float64) *SurveyWiFiScanCreate {
	swfsc.mutation.SetLatitude(f)
	return swfsc
}

// SetNillableLatitude sets the latitude field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableLatitude(f *float64) *SurveyWiFiScanCreate {
	if f != nil {
		swfsc.SetLatitude(*f)
	}
	return swfsc
}

// SetLongitude sets the longitude field.
func (swfsc *SurveyWiFiScanCreate) SetLongitude(f float64) *SurveyWiFiScanCreate {
	swfsc.mutation.SetLongitude(f)
	return swfsc
}

// SetNillableLongitude sets the longitude field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableLongitude(f *float64) *SurveyWiFiScanCreate {
	if f != nil {
		swfsc.SetLongitude(*f)
	}
	return swfsc
}

// SetAltitude sets the altitude field.
func (swfsc *SurveyWiFiScanCreate) SetAltitude(f float64) *SurveyWiFiScanCreate {
	swfsc.mutation.SetAltitude(f)
	return swfsc
}

// SetNillableAltitude sets the altitude field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableAltitude(f *float64) *SurveyWiFiScanCreate {
	if f != nil {
		swfsc.SetAltitude(*f)
	}
	return swfsc
}

// SetHeading sets the heading field.
func (swfsc *SurveyWiFiScanCreate) SetHeading(f float64) *SurveyWiFiScanCreate {
	swfsc.mutation.SetHeading(f)
	return swfsc
}

// SetNillableHeading sets the heading field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableHeading(f *float64) *SurveyWiFiScanCreate {
	if f != nil {
		swfsc.SetHeading(*f)
	}
	return swfsc
}

// SetRssi sets the rssi field.
func (swfsc *SurveyWiFiScanCreate) SetRssi(f float64) *SurveyWiFiScanCreate {
	swfsc.mutation.SetRssi(f)
	return swfsc
}

// SetNillableRssi sets the rssi field if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableRssi(f *float64) *SurveyWiFiScanCreate {
	if f != nil {
		swfsc.SetRssi(*f)
	}
	return swfsc
}

// SetChecklistItemID sets the checklist_item edge to CheckListItem by id.
func (swfsc *SurveyWiFiScanCreate) SetChecklistItemID(id int) *SurveyWiFiScanCreate {
	swfsc.mutation.SetChecklistItemID(id)
	return swfsc
}

// SetNillableChecklistItemID sets the checklist_item edge to CheckListItem by id if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableChecklistItemID(id *int) *SurveyWiFiScanCreate {
	if id != nil {
		swfsc = swfsc.SetChecklistItemID(*id)
	}
	return swfsc
}

// SetChecklistItem sets the checklist_item edge to CheckListItem.
func (swfsc *SurveyWiFiScanCreate) SetChecklistItem(c *CheckListItem) *SurveyWiFiScanCreate {
	return swfsc.SetChecklistItemID(c.ID)
}

// SetSurveyQuestionID sets the survey_question edge to SurveyQuestion by id.
func (swfsc *SurveyWiFiScanCreate) SetSurveyQuestionID(id int) *SurveyWiFiScanCreate {
	swfsc.mutation.SetSurveyQuestionID(id)
	return swfsc
}

// SetNillableSurveyQuestionID sets the survey_question edge to SurveyQuestion by id if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableSurveyQuestionID(id *int) *SurveyWiFiScanCreate {
	if id != nil {
		swfsc = swfsc.SetSurveyQuestionID(*id)
	}
	return swfsc
}

// SetSurveyQuestion sets the survey_question edge to SurveyQuestion.
func (swfsc *SurveyWiFiScanCreate) SetSurveyQuestion(s *SurveyQuestion) *SurveyWiFiScanCreate {
	return swfsc.SetSurveyQuestionID(s.ID)
}

// SetLocationID sets the location edge to Location by id.
func (swfsc *SurveyWiFiScanCreate) SetLocationID(id int) *SurveyWiFiScanCreate {
	swfsc.mutation.SetLocationID(id)
	return swfsc
}

// SetNillableLocationID sets the location edge to Location by id if the given value is not nil.
func (swfsc *SurveyWiFiScanCreate) SetNillableLocationID(id *int) *SurveyWiFiScanCreate {
	if id != nil {
		swfsc = swfsc.SetLocationID(*id)
	}
	return swfsc
}

// SetLocation sets the location edge to Location.
func (swfsc *SurveyWiFiScanCreate) SetLocation(l *Location) *SurveyWiFiScanCreate {
	return swfsc.SetLocationID(l.ID)
}

// Mutation returns the SurveyWiFiScanMutation object of the builder.
func (swfsc *SurveyWiFiScanCreate) Mutation() *SurveyWiFiScanMutation {
	return swfsc.mutation
}

// Save creates the SurveyWiFiScan in the database.
func (swfsc *SurveyWiFiScanCreate) Save(ctx context.Context) (*SurveyWiFiScan, error) {
	var (
		err  error
		node *SurveyWiFiScan
	)
	swfsc.defaults()
	if len(swfsc.hooks) == 0 {
		if err = swfsc.check(); err != nil {
			return nil, err
		}
		node, err = swfsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SurveyWiFiScanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = swfsc.check(); err != nil {
				return nil, err
			}
			swfsc.mutation = mutation
			node, err = swfsc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(swfsc.hooks) - 1; i >= 0; i-- {
			mut = swfsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, swfsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (swfsc *SurveyWiFiScanCreate) SaveX(ctx context.Context) *SurveyWiFiScan {
	v, err := swfsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (swfsc *SurveyWiFiScanCreate) defaults() {
	if _, ok := swfsc.mutation.CreateTime(); !ok {
		v := surveywifiscan.DefaultCreateTime()
		swfsc.mutation.SetCreateTime(v)
	}
	if _, ok := swfsc.mutation.UpdateTime(); !ok {
		v := surveywifiscan.DefaultUpdateTime()
		swfsc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (swfsc *SurveyWiFiScanCreate) check() error {
	if _, ok := swfsc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := swfsc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := swfsc.mutation.Bssid(); !ok {
		return &ValidationError{Name: "bssid", err: errors.New("ent: missing required field \"bssid\"")}
	}
	if _, ok := swfsc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New("ent: missing required field \"timestamp\"")}
	}
	if _, ok := swfsc.mutation.Frequency(); !ok {
		return &ValidationError{Name: "frequency", err: errors.New("ent: missing required field \"frequency\"")}
	}
	if _, ok := swfsc.mutation.Channel(); !ok {
		return &ValidationError{Name: "channel", err: errors.New("ent: missing required field \"channel\"")}
	}
	if _, ok := swfsc.mutation.Strength(); !ok {
		return &ValidationError{Name: "strength", err: errors.New("ent: missing required field \"strength\"")}
	}
	return nil
}

func (swfsc *SurveyWiFiScanCreate) sqlSave(ctx context.Context) (*SurveyWiFiScan, error) {
	_node, _spec := swfsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, swfsc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (swfsc *SurveyWiFiScanCreate) createSpec() (*SurveyWiFiScan, *sqlgraph.CreateSpec) {
	var (
		_node = &SurveyWiFiScan{config: swfsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: surveywifiscan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: surveywifiscan.FieldID,
			},
		}
	)
	if value, ok := swfsc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: surveywifiscan.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := swfsc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: surveywifiscan.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := swfsc.mutation.Ssid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surveywifiscan.FieldSsid,
		})
		_node.Ssid = value
	}
	if value, ok := swfsc.mutation.Bssid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surveywifiscan.FieldBssid,
		})
		_node.Bssid = value
	}
	if value, ok := swfsc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: surveywifiscan.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	if value, ok := swfsc.mutation.Frequency(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: surveywifiscan.FieldFrequency,
		})
		_node.Frequency = value
	}
	if value, ok := swfsc.mutation.Channel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: surveywifiscan.FieldChannel,
		})
		_node.Channel = value
	}
	if value, ok := swfsc.mutation.Band(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surveywifiscan.FieldBand,
		})
		_node.Band = value
	}
	if value, ok := swfsc.mutation.ChannelWidth(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: surveywifiscan.FieldChannelWidth,
		})
		_node.ChannelWidth = value
	}
	if value, ok := swfsc.mutation.Capabilities(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surveywifiscan.FieldCapabilities,
		})
		_node.Capabilities = value
	}
	if value, ok := swfsc.mutation.Strength(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: surveywifiscan.FieldStrength,
		})
		_node.Strength = value
	}
	if value, ok := swfsc.mutation.Latitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: surveywifiscan.FieldLatitude,
		})
		_node.Latitude = value
	}
	if value, ok := swfsc.mutation.Longitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: surveywifiscan.FieldLongitude,
		})
		_node.Longitude = value
	}
	if value, ok := swfsc.mutation.Altitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: surveywifiscan.FieldAltitude,
		})
		_node.Altitude = &value
	}
	if value, ok := swfsc.mutation.Heading(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: surveywifiscan.FieldHeading,
		})
		_node.Heading = &value
	}
	if value, ok := swfsc.mutation.Rssi(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: surveywifiscan.FieldRssi,
		})
		_node.Rssi = &value
	}
	if nodes := swfsc.mutation.ChecklistItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   surveywifiscan.ChecklistItemTable,
			Columns: []string{surveywifiscan.ChecklistItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checklistitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := swfsc.mutation.SurveyQuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   surveywifiscan.SurveyQuestionTable,
			Columns: []string{surveywifiscan.SurveyQuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surveyquestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := swfsc.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   surveywifiscan.LocationTable,
			Columns: []string{surveywifiscan.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
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

// SurveyWiFiScanCreateBulk is the builder for creating a bulk of SurveyWiFiScan entities.
type SurveyWiFiScanCreateBulk struct {
	config
	builders []*SurveyWiFiScanCreate
}

// Save creates the SurveyWiFiScan entities in the database.
func (swfscb *SurveyWiFiScanCreateBulk) Save(ctx context.Context) ([]*SurveyWiFiScan, error) {
	specs := make([]*sqlgraph.CreateSpec, len(swfscb.builders))
	nodes := make([]*SurveyWiFiScan, len(swfscb.builders))
	mutators := make([]Mutator, len(swfscb.builders))
	for i := range swfscb.builders {
		func(i int, root context.Context) {
			builder := swfscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SurveyWiFiScanMutation)
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
					_, err = mutators[i+1].Mutate(root, swfscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, swfscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, swfscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (swfscb *SurveyWiFiScanCreateBulk) SaveX(ctx context.Context) []*SurveyWiFiScan {
	v, err := swfscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
