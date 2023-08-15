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

func NewACLBindingRule() datasource.DataSource {
	return NewDataSource(
		"acl_binding_rule",
		aclBindingRuleSchema(),
		&ACLBindingRule{},
	)
}

type ACLBindingRule struct{}

func (d *ACLBindingRule) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *api.ACLBindingRule
	resp.Diagnostics.Append(models.DecodeACLBindingRule(ctx, req.Config, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().BindingRuleRead(data.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to read ACL binding rule %q", data.ID), err.Error())
		return
	}

	resp.Diagnostics.Append(models.Set(ctx, &resp.State, token)...)
}
