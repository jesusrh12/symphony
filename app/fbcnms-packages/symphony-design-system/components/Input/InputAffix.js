/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import * as React from 'react';
import Text from '../Text';
import classNames from 'classnames';
import symphony from '@symphony/design-system/theme/symphony';
import {makeStyles} from '@material-ui/styles';
import {useFormElementContext} from '../Form/FormElementContext';
import {useInput} from './InputContext';
import {useMemo} from 'react';

const useStyles = makeStyles(_theme => ({
  root: {
    display: 'flex',
    alignItems: 'center',
    color: symphony.palette.D400,
    alignItems: 'center',
    marginRight: '8px',
    marginLeft: '4px',
  },
  hasValue: {},
  disabled: {
    '&:not($hasValue) $text': {
      color: symphony.palette.disabled,
    },
    pointerEvents: 'none',
    opacity: 0.5,
  },
  text: {
    color: symphony.palette.D400,
  },
  clickable: {
    cursor: 'pointer',
  },
}));

type Props = $ReadOnly<{|
  className?: string,
  children: React.Node,
  onClick?: () => void,
  inheritsDisability?: boolean,
|}>;

const InputAffix = (props: Props) => {
  const {children, className, onClick, inheritsDisability = false} = props;
  const classes = useStyles();
  const {disabled: parentInputDisabled, value} = useInput();

  const {disabled: contextDisabled} = useFormElementContext();

  const disabled = useMemo(
    () => (parentInputDisabled && inheritsDisability) || contextDisabled,
    [parentInputDisabled, inheritsDisability, contextDisabled],
  );

  return (
    <div
      className={classNames(
        classes.root,
        {
          [classes.disabled]: disabled,
          [classes.hasValue]: Boolean(value),
          [classes.clickable]: onClick !== undefined && !disabled,
        },
        className,
      )}
      onClick={disabled ? null : onClick}>
      {typeof children === 'string' ? (
        <Text className={classes.text} variant="body2">
          {children}
        </Text>
      ) : (
        children
      )}
    </div>
  );
};

export default InputAffix;
