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
  AddServiceTypeMutation,
  AddServiceTypeMutationResponse,
  AddServiceTypeMutationVariables,
} from './__generated__/AddServiceTypeMutation.graphql';
import type {MutationCallbacks} from './MutationCallbacks.js';
import type {SelectorStoreUpdater} from 'relay-runtime';

import RelayEnvironment from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';

const mutation = graphql`
  mutation AddServiceTypeMutation($data: ServiceTypeCreateData!) {
    addServiceType(data: $data) {
      id
      name
      discoveryMethod
      propertyTypes {
        ...PropertyTypeFormField_propertyType @relay(mask: false)
      }
      endpointDefinitions {
        id
        index
        role
        name
        equipmentType {
          id
          name
        }
      }
      numberOfServices
    }
  }
`;

export default (
  variables: AddServiceTypeMutationVariables,
  callbacks?: MutationCallbacks<AddServiceTypeMutationResponse>,
  updater?: SelectorStoreUpdater,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation<AddServiceTypeMutation>(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
