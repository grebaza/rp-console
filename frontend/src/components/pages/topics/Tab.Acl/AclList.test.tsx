/* eslint-disable testing-library/no-wait-for-multiple-assertions */
/**
 * Copyright 2022 Redpanda Data, Inc.
 *
 * Use of this software is governed by the Business Source License
 * included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
 *
 * As of the Change Date specified in that file, in accordance with
 * the Business Source License, use of this software will be governed
 * by the Apache License, Version 2.0
 */

import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import AclList from './AclList';
import { observable } from 'mobx';
import { AclStrOperation, AclStrPermission, AclStrResourceType, GetAclOverviewResponse } from '../../../../state/restInterfaces';

it('renders an empty table when no data is present', async () => {
    const store = observable({
        isAuthorizerEnabled: true,
        aclResources: [],
    });

    render(<AclList acl={store} />);

    await waitFor(() => {
        expect(screen.getByText('No data')).toBeInTheDocument();
    })
});

it('renders a table with one entry', async () => {
    const store = observable({
        isAuthorizerEnabled: true,
        aclResources: [
            {
                resourceType: 'Topic' as AclStrResourceType,
                resourceName: 'Test Topic',
                resourcePatternType: 'Unknown',
                acls: [
                    {
                        principal: 'test principal',
                        host: '*',
                        operation: 'All' as AclStrOperation,
                        permissionType: 'Allow' as AclStrPermission,
                    },
                ],
            },
        ],
    } as GetAclOverviewResponse);

    render(<AclList acl={store} />);

    await waitFor(() => {
        expect(screen.getByText('Topic')).toBeInTheDocument();
        expect(screen.getByText('Allow')).toBeInTheDocument();
        expect(screen.getByText('test principal')).toBeInTheDocument();
        expect(screen.getByText('All')).toBeInTheDocument();
        expect(screen.getByText('Unknown')).toBeInTheDocument();
        expect(screen.getByText('Test Topic')).toBeInTheDocument();
        expect(screen.getByText('1')).toBeInTheDocument();
        expect(screen.getByText('*')).toBeInTheDocument();
    })
});

it('informs user about missing permission to view ACLs', async () => {
    render(<AclList acl={null} />);

    await waitFor(() => {
        expect(screen.getByText('You do not have the necessary permissions to view ACLs')).toBeInTheDocument();
    })

});

it('informs user about missing authorizer config in Kafka cluster', async () => {
    const store = observable({
        isAuthorizerEnabled: false,
        aclResources: [],
    });

    render(<AclList acl={store} />);

    await waitFor(() => {
        expect(screen.getByText('There\'s no authorizer configured in your Kafka cluster')).toBeInTheDocument();
    });
});
