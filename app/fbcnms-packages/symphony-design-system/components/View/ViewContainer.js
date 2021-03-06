/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {FullViewHeaderProps} from './ViewHeader';
import type {Variant} from './ViewBody';

import * as React from 'react';
import ViewBody, {SIDE_PADDING, VARIANTS} from './ViewBody';
import ViewHeader from './ViewHeader';
import classNames from 'classnames';
import symphony from '../../theme/symphony';
import useVerticalScrollingEffect from '../../hooks/useVerticalScrollingEffect';
import {makeStyles} from '@material-ui/styles';
import {useRef, useState} from 'react';

const useStyles = makeStyles(() => ({
  viewPanel: {
    flexGrow: 1,
    display: 'flex',
    flexDirection: 'column',
    maxHeight: '100%',
    maxWidth: '100%',
  },
  limitedWidth: {
    width: '100%',
    alignSelf: 'center',
    maxWidth: `${symphony.layout.maximalWidth + 2 * SIDE_PADDING}px`,
  },
}));

export type ViewContainerProps = $ReadOnly<{|
  header?: ?FullViewHeaderProps,
  useBodyScrollingEffect?: ?boolean,
  bodyVariant?: ?Variant,
  className?: ?string,
  children: React.Node,
|}>;

export default function ViewContainer(props: ViewContainerProps) {
  const {
    header,
    useBodyScrollingEffect = true,
    bodyVariant,
    className,
    children,
  } = props;
  const classes = useStyles();
  const headerElement = useRef(null);
  const bodyElement = useRef(null);
  const [bodyIsScrolled, setBodyIsScrolled] = useState(false);

  const handleBodyScroll = verticalScrollValues => {
    if (headerElement?.current == null) {
      return;
    }
    setBodyIsScrolled(
      verticalScrollValues.scrollTop > headerElement.current.clientHeight,
    );
  };

  useVerticalScrollingEffect(
    bodyElement,
    handleBodyScroll.bind(this),
    !!useBodyScrollingEffect,
  );

  return (
    <div className={classNames(classes.viewPanel, className)}>
      {!!header && (
        <ViewHeader
          className={classes.limitedWidth}
          ref={headerElement}
          {...header}
          showMinimal={bodyIsScrolled}
        />
      )}
      <ViewBody
        className={bodyVariant === VARIANTS.plain ? null : classes.limitedWidth}
        ref={bodyElement}
        variant={bodyVariant}>
        {children}
      </ViewBody>
    </div>
  );
}
