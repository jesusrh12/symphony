/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import type {MutationCallbacks} from '../../mutations/MutationCallbacks.js';
import type {
  RemoveWorkOrderMutationResponse,
  RemoveWorkOrderMutationVariables,
} from '../../mutations/__generated__/RemoveWorkOrderMutation.graphql';
import type {WithAlert} from '@fbcnms/ui/components/Alert/withAlert';
import type {WithSnackbarProps} from 'notistack';
import type {WithStyles} from '@material-ui/core';

import Button from '@symphony/design-system/components/Button';
import DeleteOutlineIcon from '@material-ui/icons/DeleteOutline';
import FormActionWithPermissions from '../../common/FormActionWithPermissions';
import React from 'react';
import RemoveWorkOrderMutation from '../../mutations/RemoveWorkOrderMutation';
import classNames from 'classnames';
import nullthrows from '@fbcnms/util/nullthrows';
import withAlert from '@fbcnms/ui/components/Alert/withAlert';
import {LogEvents, ServerLogger} from '../../common/LoggingUtils';
import {withSnackbar} from 'notistack';
import {withStyles} from '@material-ui/core/styles';

const styles = () => ({
  deleteButton: {
    minWidth: 'unset',
    paddingTop: '2px',
    height: '36px',
  },
});

type Props = {|
  className?: string,
  workOrderId: string,
  workOrderTypeId: string,
  ignorePermissions?: ?boolean,
  onWorkOrderRemoved: () => void,
|} & WithStyles<typeof styles> &
  WithAlert &
  WithSnackbarProps;

class WorkOrderDeleteButton extends React.Component<Props> {
  render() {
    const {workOrderTypeId, ignorePermissions, className, classes} = this.props;
    return (
      <FormActionWithPermissions
        permissions={{
          entity: 'workorder',
          action: 'delete',
          workOrderTypeId,
          ignorePermissions,
        }}>
        <Button
          className={classNames(className, classes.deleteButton)}
          variant="text"
          skin="gray"
          onClick={this.removeWorkOrder}>
          <DeleteOutlineIcon />
        </Button>
      </FormActionWithPermissions>
    );
  }

  removeWorkOrder = () => {
    ServerLogger.info(LogEvents.DELETE_WORK_ORDER_BUTTON_CLICKED, {
      source: 'work_order_details',
    });
    const {workOrderId} = this.props;
    this.props
      .confirm({
        message: 'Are you sure you want to delete this work order?',
        confirmLabel: 'Delete',
      })
      .then(confirmed => {
        if (!confirmed) {
          return;
        }

        const variables: RemoveWorkOrderMutationVariables = {
          id: nullthrows(workOrderId),
        };

        const updater = store => {
          this.props.onWorkOrderRemoved();
          store.delete(workOrderId);
        };

        const callbacks: MutationCallbacks<RemoveWorkOrderMutationResponse> = {
          onCompleted: (response, errors) => {
            if (errors && errors[0]) {
              this.props.alert('Failed removing work order');
            }
          },
          onError: (_error: Error) => {
            this.props.alert('Failed removing work order');
          },
        };

        RemoveWorkOrderMutation(variables, callbacks, updater);
      });
  };
}

export default withStyles(styles)(
  withAlert(withSnackbar(WorkOrderDeleteButton)),
);
