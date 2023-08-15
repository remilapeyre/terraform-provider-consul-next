// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
)

func NewACLPolicy() datasource.DataSource {
	return NewDataSource(
		"acl_policy",
		aclPolicySchema(),
		&ACLPolicy{},
	)
}

type ACLPolicy struct{}

func (d *ACLPolicy) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *api.ACLPolicy
	resp.Diagnostics.Append(models.DecodeACLPolicy(ctx, req.Config, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var summary string
	var err error
	var policy *api.ACLPolicy
	if data.ID != "" {
		summary = fmt.Sprintf("failed to read ACL policy %q", data.ID)
		policy, _, err = client.ACL().PolicyRead(data.ID, nil)
	} else if data.Name != "" {
		summary = fmt.Sprintf("failed to read ACL policy %q", data.Name)
		policy, _, err = client.ACL().PolicyReadByName(data.Name, nil)
	} else {
		resp.Diagnostics.AddError("missing required argument", `at least one of "id" or "name" is required`)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(summary, err.Error())
		return
	}
	if policy == nil {
		resp.Diagnostics.AddError(summary, fmt.Sprintf("no ACL auth method %q found", data.Name))
		return
	}

	resp.Diagnostics.Append(models.Set(ctx, &resp.State, policy)...)
}
