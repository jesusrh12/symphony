/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */
import type {EquipmentComparisonViewQueryRendererSearchQueryResponse} from './__generated__/EquipmentComparisonViewQueryRendererSearchQuery.graphql.js';
import type {
  FilterConfig,
  FilterValue,
  FiltersQuery,
} from './ComparisonViewTypes';
import type {PowerSearchEquipmentResultsTable_equipment} from './__generated__/PowerSearchEquipmentResultsTable_equipment.graphql';
import type {PowerSearchLinkFirstEquipmentResultsTable_equipment} from '../services/__generated__/PowerSearchLinkFirstEquipmentResultsTable_equipment.graphql';

import * as React from 'react';
import InventoryQueryRenderer from '../InventoryQueryRenderer';
import PowerSearchBar from '../power_search/PowerSearchBar';
import SearchIcon from '@material-ui/icons/Search';
import Text from '@symphony/design-system/components/Text';
import WizardContext from '@symphony/design-system/components/Wizard/WizardContext';
import useLocationTypes from './hooks/locationTypesHook';
import usePropertyFilters from './hooks/propertiesHook';
import {EquipmentCriteriaConfig} from './EquipmentSearchConfig';
import {LogEvents, ServerLogger} from '../../common/LoggingUtils';
import {buildPropertyFilterConfigs, getSelectedFilter} from './FilterUtils';
import {graphql} from 'relay-runtime';
import {makeStyles} from '@material-ui/styles';
import {useContext, useMemo, useState} from 'react';

const useStyles = makeStyles(theme => ({
  noResultsRoot: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    marginTop: '100px',
  },
  noResultsLabel: {
    color: theme.palette.grey[600],
  },
  searchIcon: {
    color: theme.palette.grey[600],
    marginBottom: '6px',
    fontSize: '36px',
  },
  root: {
    display: 'flex',
    flexDirection: 'column',
    backgroundColor: theme.palette.common.white,
    height: '100%',
  },
  searchResults: {
    flexGrow: 1,
  },
  searchBar: {
    boxShadow: '0px 2px 2px 0px rgba(0, 0, 0, 0.1)',
  },
}));

type Props = $ReadOnly<{|
  limit?: number,
  showExport?: boolean,
  initialFilters?: FiltersQuery,
  children: (props: {
    equipment:
      | PowerSearchEquipmentResultsTable_equipment
      | PowerSearchLinkFirstEquipmentResultsTable_equipment,
  }) => React.Element<*>,
|}>;

const equipmentSearchQuery = graphql`
  query EquipmentComparisonViewQueryRendererSearchQuery(
    $limit: Int
    $filters: [EquipmentFilterInput!]!
  ) {
    equipments(first: $limit, filterBy: $filters) {
      edges {
        node {
          ...PowerSearchEquipmentResultsTable_equipment
          ...PowerSearchLinkFirstEquipmentResultsTable_equipment
        }
      }
      totalCount
    }
  }
`;

const EquipmentComparisonViewQueryRenderer = (props: Props) => {
  const classes = useStyles();
  const {limit, showExport, children, initialFilters} = props;
  const [count, setCount] = useState((0: number));

  const possibleProperties = usePropertyFilters('equipment');
  const equipmentPropertiesFilterConfigs = buildPropertyFilterConfigs(
    possibleProperties,
  );

  const locationTypesFilterConfigs = useLocationTypes();
  const filterConfigs = useMemo(() => {
    return EquipmentCriteriaConfig.map(ent => ent.filters)
      .reduce(
        (allFilters, currentFilter) => allFilters.concat(currentFilter),
        [],
      )
      .concat(equipmentPropertiesFilterConfigs ?? [])
      .concat(locationTypesFilterConfigs ?? []);
  }, [equipmentPropertiesFilterConfigs, locationTypesFilterConfigs]);

  const wizardContext = useContext(WizardContext);

  const filtersContextKey = 'EquipmentComparisonViewQueryRenderer_Filters';

  const updateFilters = newFilters => {
    wizardContext.set(filtersContextKey, newFilters);
  };

  // eslint-disable-next-line no-warning-comments
  // $FlowFixMe
  const filters = (wizardContext.get(filtersContextKey) ||
    initialFilters ||
    []: FiltersQuery);

  return (
    <div className={classes.root}>
      <div className={classes.searchBar}>
        {possibleProperties != null && locationTypesFilterConfigs != null && (
          <PowerSearchBar
            filterValues={filters}
            exportPath={showExport ? '/equipment' : null}
            onFiltersChanged={updateFilters}
            onFilterRemoved={handleFilterRemoved}
            onFilterBlurred={handleFilterBlurred}
            getSelectedFilter={(filterConfig: FilterConfig) =>
              getSelectedFilter(filterConfig, possibleProperties)
            }
            placeholder="Filter equipment"
            searchConfig={EquipmentCriteriaConfig}
            filterConfigs={filterConfigs}
            footer={
              count != null
                ? limit != null && count > limit
                  ? `1 to ${limit} of ${count}`
                  : `1 to ${count}`
                : null
            }
          />
        )}
      </div>
      <InventoryQueryRenderer
        query={equipmentSearchQuery}
        variables={{
          limit: limit,
          filters: filters.map(f => ({
            filterType: f.name.toUpperCase(),
            operator: f.operator.toUpperCase(),
            stringValue: f.stringValue,
            propertyValue: f.propertyValue,
            idSet: f.idSet,
            stringSet: f.stringSet,
          })),
        }}
        render={(
          props: EquipmentComparisonViewQueryRendererSearchQueryResponse,
        ) => {
          const {totalCount, edges} = props.equipments;
          setCount(totalCount);
          if (edges.length === 0) {
            return (
              <div className={classes.noResultsRoot}>
                <SearchIcon className={classes.searchIcon} />
                <Text variant="h6" className={classes.noResultsLabel}>
                  No results found
                </Text>
              </div>
            );
          }
          return children({
            // $FlowFixMe[incompatible-call] $FlowFixMe T74239404 Found via relay types
            equipment: edges.map(edge => edge.node),
          });
        }}
      />
    </div>
  );
};

const handleFilterRemoved = (filter: FilterValue) => {
  ServerLogger.info(LogEvents.EQUIPMENT_COMPARISON_VIEW_FILTER_REMOVED, {
    filterName: filter.name,
  });
};

const handleFilterBlurred = (filter: FilterValue) => {
  ServerLogger.info(LogEvents.EQUIPMENT_COMPARISON_VIEW_FILTER_SET, {
    value: filter,
  });
};

export default EquipmentComparisonViewQueryRenderer;
