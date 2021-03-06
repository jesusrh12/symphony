/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {MutationCallbacks} from './MutationCallbacks.js';
import type {
  RemoveLinkMutation,
  RemoveLinkMutationResponse,
  RemoveLinkMutationVariables,
} from './__generated__/RemoveLinkMutation.graphql';
import type {SelectorStoreUpdater} from 'relay-runtime';

import RelayEnvironment from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';

const mutation = graphql`
  mutation RemoveLinkMutation($id: ID!, $workOrderId: ID) {
    removeLink(id: $id, workOrderId: $workOrderId) {
      ...EquipmentPortsTable_link @relay(mask: false)
    }
  }
`;

export default (
  variables: RemoveLinkMutationVariables,
  callbacks?: MutationCallbacks<RemoveLinkMutationResponse>,
  updater?: SelectorStoreUpdater,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation<RemoveLinkMutation>(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
