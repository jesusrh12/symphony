/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {
  DeleteHyperlinkMutation,
  DeleteHyperlinkMutationResponse,
  DeleteHyperlinkMutationVariables,
} from './__generated__/DeleteHyperlinkMutation.graphql';
import type {MutationCallbacks} from './MutationCallbacks.js';
import type {SelectorStoreUpdater} from 'relay-runtime';

import RelayEnvironment from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';

const mutation = graphql`
  mutation DeleteHyperlinkMutation($id: ID!) {
    deleteHyperlink(id: $id) {
      ...HyperlinkTableRow_hyperlink
    }
  }
`;

export default (
  variables: DeleteHyperlinkMutationVariables,
  callbacks?: MutationCallbacks<DeleteHyperlinkMutationResponse>,
  updater?: SelectorStoreUpdater,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation<DeleteHyperlinkMutation>(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
