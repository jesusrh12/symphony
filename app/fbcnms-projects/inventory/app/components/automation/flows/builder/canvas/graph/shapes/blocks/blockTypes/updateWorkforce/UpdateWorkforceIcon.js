/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import BaseIcon from '../BaseIcon';
import React from 'react';

type Props = $ReadOnly<{|
  className?: ?string,
|}>;

const UpdateWorkforceIcon = (props: Props) => {
  return (
    <BaseIcon shape="square" color="blue" {...props}>
      <g transform="translate(9.000000, 9.000000)">
        <path d="M20.9851,14.0002982 C22.4321,13.9944 23.8681,14.5144 24.9971,15.5174 L24.9971,15.5174 L24.9971,14.0002982 L26.9971,14.0002982 L26.9971,18.0004 C26.9971,18.5524 26.5501,19.0004 25.9961,19.0004 L25.9961,19.0004 L22.0001,19.0004 L22.0001,17.0004 L23.6451,17.0004 C22.1881,15.7274 19.9801,15.6484 18.4451,16.9234 C17.6211,17.6054 17.1141,18.5684 17.0171,19.6344 C16.9191,20.6954 17.2431,21.7324 17.9271,22.5544 C19.3401,24.2504 21.8751,24.4814 23.5761,23.0704 C24.2311,22.5284 24.6751,21.8054 24.8831,20.9964 L24.8831,20.9964 L26.9321,20.9964 C26.6951,22.4124 25.9711,23.6844 24.8541,24.6094 C23.7331,25.5404 22.3701,25.9934 21.0161,25.9934 C19.2901,25.9934 17.5781,25.2594 16.3911,23.8354 C15.4931,22.7574 15.0091,21.4324 15.0001,20.0454 C14.9991,19.8484 15.0071,19.6494 15.0251,19.4504 C15.1711,17.8544 15.9321,16.4094 17.1671,15.3844 C18.2821,14.4594 19.6351,14.0054 20.9821,14.0002982 L20.9821,14.0002982 Z M17,5 C17.5522847,5 18,5.44771525 18,6 L18,8 L22,8 C22.5522847,8 23,8.44771525 23,9 L22.999,11 L21,11 L21,10 L8,10 L8,20 L12,20 L12,22 L7,22 C6.44771525,22 6,21.5522847 6,21 L6,9 C6,8.44771525 6.44771525,8 7,8 L11,8 L11,6 C11,5.44771525 11.4477153,5 12,5 L17,5 Z M16,7 L13,7 L13,8 L16,8 L16,7 Z" />
      </g>
    </BaseIcon>
  );
};

export default UpdateWorkforceIcon;
