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
  RemoveProjectTypeMutation,
  RemoveProjectTypeMutationResponse,
  RemoveProjectTypeMutationVariables,
} from './__generated__/RemoveProjectTypeMutation.graphql';
import type {SelectorStoreUpdater} from 'relay-runtime';

import RelayEnvironment from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';

const mutation = graphql`
  mutation RemoveProjectTypeMutation($id: ID!) {
    deleteProjectType(id: $id)
  }
`;

export default (
  variables: RemoveProjectTypeMutationVariables,
  callbacks?: MutationCallbacks<RemoveProjectTypeMutationResponse>,
  updater?: SelectorStoreUpdater,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation<RemoveProjectTypeMutation>(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
