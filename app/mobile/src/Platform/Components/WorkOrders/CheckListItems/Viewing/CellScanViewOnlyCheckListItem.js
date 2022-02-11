/**
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

'use strict';

import type {CellScanViewOnlyCheckListItem_item} from './__generated__/CellScanViewOnlyCheckListItem_item.graphql';

import CellScanResults from '@fbcmobile/ui/Components/Views/CellScanResults';
import LabeledTextSection from '@fbcmobile/ui/Components/LabeledTextSection';
import React from 'react';
import {CHECKLIST_ITEM_NOT_APPLICABLE_LABEL} from 'Platform/Components/WorkOrders/CheckListItems/CheckListItemUtils';
import {createFragmentContainer} from 'react-relay-offline';

const graphql = require('babel-plugin-relay/macro');

type Props = $ReadOnly<{|
  +item: CellScanViewOnlyCheckListItem_item,
|}>;

const CellScanViewOnlyCheckListItem = ({item}: Props) => {
  const cellData =
    item.cellData != null ? item.cellData.slice().map(d => ({...d})) : null;
  return (
    <LabeledTextSection
      title={item.title}
      content={
        cellData != null ? (
          <CellScanResults scanResults={cellData} />
        ) : (
          CHECKLIST_ITEM_NOT_APPLICABLE_LABEL
        )
      }
    />
  );
};

export default createFragmentContainer(CellScanViewOnlyCheckListItem, {
  item: graphql`
    fragment CellScanViewOnlyCheckListItem_item on CheckListItem {
      id
      title
      cellData {
        networkType
        signalStrength
        timestamp
        baseStationID
        networkID
        systemID
        cellID
        locationAreaCode
        mobileCountryCode
        mobileNetworkCode
        primaryScramblingCode
        operator
        arfcn
        physicalCellID
        trackingAreaCode
        timingAdvance
        earfcn
        uarfcn
        latitude
        longitude
      }
    }
  `,
});
