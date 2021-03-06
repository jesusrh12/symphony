// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"context"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/activity"
	"github.com/facebookincubator/symphony/pkg/ent/workorder"
	"github.com/facebookincubator/symphony/pkg/ev"
	"github.com/facebookincubator/symphony/pkg/event"
	"github.com/facebookincubator/symphony/pkg/log"
)

func updateActivitiesOnWOCreate(ctx context.Context, entry *event.LogEntry) error {
	userID := entry.UserID
	client := ent.FromContext(ctx)

	wo := entry.CurrState
	_, err := client.Activity.Create().
		SetActivityType(activity.ActivityTypeCreationDateChanged).
		SetIsCreate(true).
		SetNewValue(strconv.FormatInt(entry.Time.Unix(), 10)).
		SetNillableAuthorID(userID).
		SetWorkOrderID(wo.ID).
		Save(ctx)
	if err != nil {
		return err
	}

	if assignee, found := event.FindEdge(wo.Edges, workorder.EdgeAssignee); found {
		assgnID := assignee.IDs[0]
		_, err = client.Activity.Create().
			SetActivityType(activity.ActivityTypeAssigneeChanged).
			SetIsCreate(true).
			SetNillableAuthorID(userID).
			SetWorkOrderID(wo.ID).
			SetNewValue(strconv.Itoa(assgnID)).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	if owner, found := event.FindEdge(wo.Edges, workorder.EdgeOwner); found {
		ownerID := owner.IDs[0]

		_, err = client.Activity.Create().
			SetActivityType(activity.ActivityTypeOwnerChanged).
			SetIsCreate(true).
			SetNillableAuthorID(userID).
			SetWorkOrderID(wo.ID).
			SetNewValue(strconv.Itoa(ownerID)).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	if st, found := event.FindField(wo.Fields, workorder.FieldStatus); found {
		_, err = client.Activity.Create().
			SetActivityType(activity.ActivityTypeStatusChanged).
			SetIsCreate(true).
			SetNillableAuthorID(userID).
			SetWorkOrderID(wo.ID).
			SetNewValue(st.MustGetString()).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	if pri, found := event.FindField(wo.Fields, workorder.FieldPriority); found {
		_, err = client.Activity.Create().
			SetActivityType(activity.ActivityTypePriorityChanged).
			SetIsCreate(true).
			SetNillableAuthorID(userID).
			SetWorkOrderID(wo.ID).
			SetNewValue(pri.MustGetString()).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateActivitiesOnWOUpdate(ctx context.Context, entry *event.LogEntry) error {
	userID := entry.UserID
	client := ent.FromContext(ctx)

	newVal, oldVal, shouldUpdate := getDiffOfUniqueEdgeAsString(entry, workorder.EdgeAssignee)
	if shouldUpdate {
		_, err := client.Activity.Create().
			SetActivityType(activity.ActivityTypeAssigneeChanged).
			SetIsCreate(false).
			SetNillableAuthorID(userID).
			SetWorkOrderID(entry.CurrState.ID).
			SetNillableOldValue(oldVal).
			SetNillableNewValue(newVal).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	newVal, oldVal, shouldUpdate = getDiffOfUniqueEdgeAsString(entry, workorder.EdgeOwner)
	if shouldUpdate {
		_, err := client.Activity.Create().
			SetActivityType(activity.ActivityTypeOwnerChanged).
			SetIsCreate(false).
			SetNillableAuthorID(userID).
			SetWorkOrderID(entry.CurrState.ID).
			SetNillableOldValue(oldVal).
			SetNillableNewValue(newVal).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	newVal, oldVal, shouldUpdate = event.GetStringDiffValuesField(entry, workorder.FieldStatus)
	if shouldUpdate {
		_, err := client.Activity.Create().
			SetActivityType(activity.ActivityTypeStatusChanged).
			SetIsCreate(false).
			SetNillableAuthorID(userID).
			SetWorkOrderID(entry.CurrState.ID).
			SetNillableOldValue(oldVal).
			SetNillableNewValue(newVal).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	newVal, oldVal, shouldUpdate = event.GetStringDiffValuesField(entry, workorder.FieldPriority)
	if shouldUpdate {
		_, err := client.Activity.Create().
			SetActivityType(activity.ActivityTypePriorityChanged).
			SetIsCreate(false).
			SetNillableAuthorID(userID).
			SetWorkOrderID(entry.CurrState.ID).
			SetNillableOldValue(oldVal).
			SetNillableNewValue(newVal).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	newVal, oldVal, shouldUpdate = event.GetStringDiffValuesField(entry, workorder.FieldName)
	if shouldUpdate {
		_, err := client.Activity.Create().
			SetActivityType(activity.ActivityTypeNameChanged).
			SetIsCreate(false).
			SetNillableAuthorID(userID).
			SetWorkOrderID(entry.CurrState.ID).
			SetNillableOldValue(oldVal).
			SetNillableNewValue(newVal).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	newVal, oldVal, shouldUpdate = event.GetStringDiffValuesField(entry, workorder.FieldDescription)
	if shouldUpdate {
		_, err := client.Activity.Create().
			SetActivityType(activity.ActivityTypeDescriptionChanged).
			SetIsCreate(false).
			SetNillableAuthorID(userID).
			SetWorkOrderID(entry.CurrState.ID).
			SetNillableOldValue(oldVal).
			SetNillableNewValue(newVal).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func getDiffOfUniqueEdgeAsString(entry *event.LogEntry, edge string) (*string, *string, bool) {
	newIntVal, oldIntVal, shouldUpdate := event.GetDiffOfUniqueEdge(entry, edge)
	var newStrVal, oldStrVal *string
	if newIntVal != nil {
		newStrVal = pointer.ToString(strconv.Itoa(*newIntVal))
	}
	if oldIntVal != nil {
		oldStrVal = pointer.ToString(strconv.Itoa(*oldIntVal))
	}
	return newStrVal, oldStrVal, shouldUpdate
}

func HandleActivityLog(ctx context.Context, _ log.Logger, evt ev.EventObject) error {
	var err error
	entry, ok := evt.(event.LogEntry)
	if !ok || entry.Type != ent.TypeWorkOrder {
		return nil
	}
	if entry.Operation.Is(ent.OpCreate) {
		err = updateActivitiesOnWOCreate(ctx, &entry)
	} else if entry.Operation.Is(ent.OpUpdate | ent.OpUpdateOne) {
		err = updateActivitiesOnWOUpdate(ctx, &entry)
	}
	if err != nil {
		return err
	}
	return nil
}
