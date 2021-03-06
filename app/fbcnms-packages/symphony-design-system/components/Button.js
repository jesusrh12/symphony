/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import type {MouseEventHandler} from '@symphony/design-system/components/Core/Clickable';
import type {SvgIconStyleProps} from '@symphony/design-system/icons/SvgIcon';
import type {TRefFor} from '../types/TRefFor.flow';

import * as React from 'react';
import Clickable from './Core/Clickable';
import Text from './Text';
import classNames from 'classnames';
import symphony from '@symphony/design-system/theme/symphony';
import {joinNullableStrings} from '@fbcnms/util/strings';
import {makeStyles} from '@material-ui/styles';
import {useFormElementContext} from './Form/FormElementContext';
import {useMemo} from 'react';

const useStyles = makeStyles(_theme => ({
  root: {
    border: 0,
    '&:focus': {
      outline: 'none',
    },
    flexShrink: 0,
    display: 'inline-flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    maxWidth: '100%',
  },
  icon: {},
  hasIcon: {
    justifyContent: 'start',
    '& $buttonText': {
      flexGrow: 1,
    },
  },
  rightIcon: {
    marginLeft: '8px',
  },
  leftIcon: {
    marginRight: '8px',
  },
  hasRightIcon: {
    '& $buttonText': {
      textAlign: 'left',
    },
  },
  hasLeftIcon: {
    '& $buttonText': {
      textAlign: 'right',
    },
  },
  primarySkin: {},
  redSkin: {},
  orangeSkin: {},
  greenSkin: {},
  regularSkin: {},
  graySkin: {},
  secondaryGraySkin: {},
  darkGraySkin: {},
  brightGraySkin: {},
  disabled: {},
  containedVariant: {
    height: '36px',
    minWidth: '88px',
    padding: '4px 12px',
    borderRadius: '4px',
    '&$hasRightIcon': {
      padding: '4px 8px 4px 16px',
    },
    '&$hasLeftIcon': {
      padding: '4px 16px 4px 8px',
    },
    '&$primarySkin': {
      backgroundColor: symphony.palette.primary,
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.white,
          fill: symphony.palette.white,
        },
      },
      '&:hover:not($disabled)': {
        backgroundColor: symphony.palette.B700,
      },
      '&:active:not($disabled)': {
        backgroundColor: symphony.palette.B800,
      },
    },
    '&$redSkin': {
      backgroundColor: symphony.palette.R600,
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.white,
          fill: symphony.palette.white,
        },
      },
      '&:hover:not($disabled)': {
        backgroundColor: symphony.palette.R700,
      },
      '&:active:not($disabled)': {
        backgroundColor: symphony.palette.R800,
      },
    },
    '&$orangeSkin': {
      backgroundColor: symphony.palette.Y600,
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.white,
          fill: symphony.palette.white,
        },
      },
      '&:hover:not($disabled)': {
        backgroundColor: symphony.palette.Y700,
      },
      '&:active:not($disabled)': {
        backgroundColor: symphony.palette.Y800,
      },
    },
    '&$greenSkin': {
      backgroundColor: symphony.palette.G600,
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.white,
          fill: symphony.palette.white,
        },
      },
      '&:hover:not($disabled)': {
        backgroundColor: symphony.palette.G700,
      },
      '&:active:not($disabled)': {
        backgroundColor: symphony.palette.G800,
      },
    },
    '&$regularSkin': {
      backgroundColor: symphony.palette.white,
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.secondary,
          fill: symphony.palette.secondary,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.primary,
          fill: symphony.palette.primary,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.B700,
          fill: symphony.palette.B700,
        },
      },
    },
    '&$graySkin': {
      backgroundColor: symphony.palette.background,
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.secondary,
          fill: symphony.palette.secondary,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.primary,
          fill: symphony.palette.primary,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.B700,
          fill: symphony.palette.B700,
        },
      },
    },
    '&$secondaryGraySkin': {
      backgroundColor: symphony.palette.background,
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.D500,
          fill: symphony.palette.D500,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.D900,
          fill: symphony.palette.D900,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.B700,
          fill: symphony.palette.B700,
        },
      },
    },
    '&$darkGraySkin': {
      backgroundColor: symphony.palette.D700,
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.white,
          fill: symphony.palette.white,
        },
      },
      '&:hover:not($disabled)': {
        backgroundColor: symphony.palette.D800,
      },
      '&:active:not($disabled)': {
        backgroundColor: symphony.palette.D900,
      },
    },
    '&$brightGraySkin': {
      backgroundColor: symphony.palette.D300,
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.white,
          fill: symphony.palette.white,
        },
      },
      '&:hover:not($disabled)': {
        backgroundColor: symphony.palette.D400,
      },
      '&:active:not($disabled)': {
        backgroundColor: symphony.palette.D500,
      },
    },
    '&$disabled': {
      backgroundColor: symphony.palette.disabled,
      '& $buttonText, $icon': {
        color: symphony.palette.white,
        fill: symphony.palette.white,
      },
    },
  },
  buttonText: {
    maxHeight: '100%',
    display: 'flex',
  },
  textVariant: {
    display: 'inline-flex',
    textAlign: 'left',
    background: 'none',
    padding: 0,
    height: '24px',
    maxWidth: '100%',
    '&$primarySkin': {
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.primary,
          fill: symphony.palette.primary,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.B700,
          fill: symphony.palette.B700,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.B700,
          fill: symphony.palette.B700,
        },
      },
    },
    '&$redSkin': {
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.R600,
          fill: symphony.palette.R600,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          opacity: 0.75,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          opacity: 0.75,
        },
      },
    },
    '&$regularSkin': {
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.secondary,
          fill: symphony.palette.secondary,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          opacity: 0.75,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          opacity: 0.75,
        },
      },
    },
    '&$graySkin': {
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.D500,
          fill: symphony.palette.D500,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.primary,
          fill: symphony.palette.primary,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.primary,
          fill: symphony.palette.primary,
        },
      },
    },
    '&$secondaryGraySkin': {
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.D500,
          fill: symphony.palette.D500,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.D900,
          fill: symphony.palette.D900,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.primary,
          fill: symphony.palette.primary,
        },
      },
    },
    '&$darkGraySkin': {
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.D700,
          fill: symphony.palette.D700,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          opacity: 0.75,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          opacity: 0.75,
        },
      },
    },
    '&$brightGraySkin': {
      '&:not($disabled)': {
        '& $buttonText, $icon': {
          color: symphony.palette.D300,
          fill: symphony.palette.D300,
        },
      },
      '&:hover:not($disabled)': {
        '& $buttonText, $icon': {
          opacity: 0.75,
        },
      },
      '&:active:not($disabled)': {
        '& $buttonText, $icon': {
          opacity: 0.75,
        },
      },
    },
    '&$disabled': {
      '& $buttonText, $icon': {
        color: symphony.palette.disabled,
        fill: symphony.palette.disabled,
      },
    },
  },
}));

export type ButtonVariant = 'contained' | 'text';
export type ButtonSkin =
  | 'primary'
  | 'regular'
  | 'red'
  | 'gray'
  | 'secondaryGray'
  | 'darkGray'
  | 'brightGray'
  | 'orange'
  | 'green';
export type SvgIconComponent =
  | React.ComponentType<SvgIconStyleProps>
  | React$ComponentType<SvgIconExports>;

export type ButtonProps = $ReadOnly<{|
  skin?: ButtonSkin,
  variant?: ButtonVariant,
  useEllipsis?: ?boolean,
  disabled?: boolean,
  tooltip?: string,
|}>;

export type Props = $ReadOnly<{|
  className?: string,
  children: React.Node,
  onClick?: ?MouseEventHandler,
  onMouseDown?: ?MouseEventHandler,
  leftIcon?: SvgIconComponent,
  leftIconClass?: string,
  rightIcon?: SvgIconComponent,
  rightIconClass?: string,
  ...ButtonProps,
|}>;

const Button = (props: Props, forwardedRef: TRefFor<HTMLElement>) => {
  const {
    className,
    children,
    skin = 'primary',
    disabled: disabledProp = false,
    variant = 'contained',
    useEllipsis = true,
    onClick,
    onMouseDown,
    leftIcon: LeftIcon = null,
    leftIconClass = null,
    rightIcon: RightIcon = null,
    rightIconClass = null,
    tooltip: tooltipProp,
  } = props;
  const classes = useStyles();

  const {
    disabled: contextDisabled,
    tooltip: contextTooltip,
  } = useFormElementContext();

  const disabled = useMemo(() => disabledProp || contextDisabled, [
    disabledProp,
    contextDisabled,
  ]);

  const tooltip = useMemo(
    () => joinNullableStrings([tooltipProp, contextTooltip]),
    [contextTooltip, tooltipProp],
  );

  return (
    <Clickable
      className={classNames(
        classes.root,
        classes[`${skin}Skin`],
        classes[`${variant}Variant`],
        {
          [classes.disabled]: disabled,
          [classes.hasIcon]: LeftIcon != null || RightIcon != null,
          [classes.hasLeftIcon]: LeftIcon != null,
          [classes.hasRightIcon]: RightIcon != null,
        },
        className,
      )}
      tooltip={tooltip ?? ''}
      disabled={disabled}
      onClick={onClick}
      onMouseDown={onMouseDown}
      ref={forwardedRef}>
      {LeftIcon ? (
        <LeftIcon
          color="inherit"
          className={classNames(classes.icon, classes.leftIcon, leftIconClass)}
        />
      ) : null}
      <Text
        variant="body2"
        weight="medium"
        useEllipsis={useEllipsis}
        className={classes.buttonText}>
        {children}
      </Text>
      {RightIcon ? (
        <RightIcon
          className={classNames(
            classes.icon,
            classes.rightIcon,
            rightIconClass,
          )}
          color="inherit"
        />
      ) : null}
    </Clickable>
  );
};

export default React.forwardRef<Props, HTMLElement>(Button);
