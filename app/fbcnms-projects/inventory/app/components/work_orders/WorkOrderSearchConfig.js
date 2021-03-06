/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import type {EntityConfig} from '../comparison_view/ComparisonViewTypes';

import PowerSearchExternalIDFilter from '../comparison_view/PowerSearchExternalIDFilter';
import PowerSearchWorkOrderAssigneeFilter from './PowerSearchWorkOrderAssigneeFilter';
import PowerSearchWorkOrderNameFilter from './PowerSearchWorkOrderNameFilter';
import PowerSearchWorkOrderOwnerFilter from './PowerSearchWorkOrderOwnerFilter';
import PowerSearchWorkOrderPriorityFilter from './PowerSearchWorkOrderPriorityFilter';
import PowerSearchWorkOrderStatusFilter from './PowerSearchWorkOrderStatusFilter';
import PowerSearchWorkOrderTypeFilter from './PowerSearchWorkOrderTypeFilter';

const WorkOrderSearchConfig: Array<EntityConfig> = [
  {
    type: 'work_order',
    label: 'Work Order',
    filters: [
      {
        key: 'work_order_name',
        name: 'work_order_name',
        entityType: 'work_order',
        label: 'Name',
        component: PowerSearchWorkOrderNameFilter,
        defaultOperator: 'contains',
      },
      {
        key: 'work_order_status',
        name: 'work_order_status',
        entityType: 'work_order',
        label: 'Status',
        component: PowerSearchWorkOrderStatusFilter,
        defaultOperator: 'is_one_of',
      },
      {
        key: 'work_order_owner',
        name: 'work_order_owned_by',
        entityType: 'work_order',
        label: 'Owner',
        component: PowerSearchWorkOrderOwnerFilter,
        defaultOperator: 'is_one_of',
      },
      {
        key: 'work_order_type',
        name: 'work_order_type',
        entityType: 'work_order',
        label: 'Type',
        component: PowerSearchWorkOrderTypeFilter,
        defaultOperator: 'is_one_of',
      },
      {
        key: 'work_order_priority',
        name: 'work_order_priority',
        entityType: 'work_order',
        label: 'Priority',
        component: PowerSearchWorkOrderPriorityFilter,
        defaultOperator: 'is_one_of',
      },
      {
        key: 'work_order_assignee',
        name: 'work_order_assigned_to',
        entityType: 'work_order',
        label: 'Assignee',
        component: PowerSearchWorkOrderAssigneeFilter,
        defaultOperator: 'is_one_of',
      },
    ],
  },
  {
    type: 'locations',
    label: 'Location',
    filters: [
      {
        key: 'location_inst_external_id',
        name: 'location_inst_external_id',
        entityType: 'locations',
        label: 'Location External ID',
        component: PowerSearchExternalIDFilter,
        defaultOperator: 'contains',
      },
    ],
  },
];

const WORK_ORDER_FILTERS = Object.freeze({
  NAME: 'work_order_name',
  STATUS: 'work_order_status',
  OWNER: 'work_order_owner',
  TYPE: 'work_order_type',
  PRIORITY: 'work_order_priority',
  ASSIGNEE: 'work_order_assignee',
});

export {WorkOrderSearchConfig, WORK_ORDER_FILTERS};
