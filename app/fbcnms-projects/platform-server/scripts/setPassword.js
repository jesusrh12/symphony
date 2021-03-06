/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */
'use strict';

import type {OrganizationType} from '@fbcnms/sequelize-models/models/organization';

import bcrypt from 'bcryptjs';
import {AccessRoles} from '@fbcnms/auth/roles';
import {Organization, User} from '@fbcnms/sequelize-models';
import {createTenant} from '../src/admin/tenant';

const SALT_GEN_ROUNDS = 10;

type UserObject = {
  organization: string,
  email: string,
  password: string,
  role: number,
};

async function updateUser(user: User, userObject: UserObject) {
  const {password, role} = userObject;
  const salt = await bcrypt.genSalt(SALT_GEN_ROUNDS);
  const passwordHash = await bcrypt.hash(password, salt);
  await user.update({
    password: passwordHash,
    role,
  });
}

async function createUser(userObject: UserObject) {
  const {organization, email, password, role} = userObject;
  const salt = await bcrypt.genSalt(SALT_GEN_ROUNDS);
  const passwordHash = await bcrypt.hash(password, salt);
  const org = await createOrFetchOrganization(organization);
  await User.create({
    email: email.toLowerCase(),
    password: passwordHash,
    role,
    networkIDs: [],
    organization: org.name,
    readOnly: false,
  });
}

async function createOrUpdateUser(userObject: UserObject) {
  const user = await User.findOne({
    where: {
      organization: userObject.organization,
      email: userObject.email.toLowerCase(),
    },
  });
  if (!user) {
    await createUser(userObject);
  } else {
    await updateUser(user, userObject);
  }
}

async function createOrFetchOrganization(
  organization: string,
): Promise<OrganizationType> {
  let org = await Organization.findOne({
    where: {
      name: organization,
    },
  });
  if (!org) {
    console.log(`Creating a new Organization: name=${organization} `);
    const [o] = await Promise.all([
      Organization.create({
        name: organization,
        tab: ['inventory', 'nms'],
        networkIDs: [],
        csvCharset: '',
        ssoCert: '',
        ssoIssuer: '',
        ssoEntrypoint: '',
      }),
      createTenant(organization),
    ]);
    org = o;
  }
  return org;
}

function main() {
  const args = process.argv.slice(2);
  if (args.length !== 3) {
    console.log('Usage: setPassword.js <organization> <email> <password>');
    process.exit(1);
  }
  const userObject = {
    organization: args[0],
    email: args[1],
    password: args[2],
    role: AccessRoles.SUPERUSER,
  };
  console.log(
    'Creating a new user: email=' +
      userObject.email +
      ', password=' +
      userObject.password,
  );
  createOrUpdateUser(userObject)
    .then(_res => {
      console.log('Success');
      process.exit();
    })
    .catch(err => {
      console.error(err);
      process.exit(1);
    });
}

main();
