// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exporter

import (
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/equipment"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentport"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentposition"
	"github.com/facebookincubator/symphony/pkg/ent/location"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/schema/enum"
	"github.com/facebookincubator/symphony/pkg/exporter/models"

	"github.com/pkg/errors"
)

func handleEquipmentLocationFilter(q *ent.EquipmentQuery, filter *models.EquipmentFilterInput) (*ent.EquipmentQuery, error) {
	switch filter.FilterType {
	case enum.EquipmentFilterTypeLocationInst:
		return equipmentLocationFilter(q, filter)
	case enum.EquipmentFilterTypeLocationInstExternalID:
		return equipmentLocationExternalIDFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func equipmentLocationExternalIDFilter(q *ent.EquipmentQuery, filter *models.EquipmentFilterInput) (*ent.EquipmentQuery, error) {
	if filter.Operator == enum.FilterOperatorContains {
		return q.Where(equipment.HasLocationWith(location.ExternalIDContainsFold(*filter.StringValue))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func equipmentLocationFilter(q *ent.EquipmentQuery, filter *models.EquipmentFilterInput) (*ent.EquipmentQuery, error) {
	if filter.Operator == enum.FilterOperatorIsOneOf {
		if filter.MaxDepth == nil {
			return nil, errors.New("max depth not supplied to location filter")
		}
		var ps []predicate.Equipment
		for _, lid := range filter.IDSet {
			ps = append(ps, equipment.HasLocationWith(BuildLocationAncestorFilter(lid, 1, *filter.MaxDepth)))
		}
		return q.Where(equipment.Or(ps...)), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

// BuildLocationAncestorFilter returns a joined predicate for location ancestors
func BuildLocationAncestorFilter(locationID, depth, maxDepth int) predicate.Location {
	if depth >= maxDepth {
		return location.ID(locationID)
	}
	return location.Or(
		location.ID(locationID),
		location.HasParentWith(
			BuildLocationAncestorFilter(locationID, depth+1, maxDepth),
		),
	)
}

func LocationFilterPredicate(q *ent.LocationQuery, filter *models.LocationFilterInput) (*ent.LocationQuery, error) {
	if filter.Operator == enum.FilterOperatorIsOneOf {
		if filter.MaxDepth == nil {
			return nil, errors.New("max depth not supplied to location filter")
		}
		var ps []predicate.Location
		for _, lid := range filter.IDSet {
			ps = append(ps, BuildLocationAncestorFilter(lid, 1, *filter.MaxDepth))
		}
		return q.Where(location.Or(ps...)), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

// GetPortLocationPredicate returns a predicate for location ancestors for port
func GetPortLocationPredicate(locationID int, maxDepth *int) predicate.EquipmentPort {
	pred := equipment.HasLocationWith(
		BuildLocationAncestorFilter(locationID, 1, *maxDepth),
	)
	return equipmentport.HasParentWith(
		equipment.Or(
			pred,
			equipment.HasParentPositionWith(
				equipmentposition.HasParentWith(pred),
			),
		),
	)
}
