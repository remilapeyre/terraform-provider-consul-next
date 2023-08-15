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

func NewACLRole() datasource.DataSource {
	return NewDataSource(
		"acl_role",
		aclRoleSchema(),
		&ACLRole{},
	)
}

type ACLRole struct{}

func (d *ACLRole) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *api.ACLRole
	resp.Diagnostics.Append(models.DecodeACLRole(ctx, req.Config, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var role *api.ACLRole
	var err error
	var summary string
	if data.ID != "" {
		role, _, err = client.ACL().RoleRead(data.ID, nil)
		summary = fmt.Sprintf("failed to read ACL role %q", data.ID)

	} else if data.Name != "" {
		role, _, err = client.ACL().RoleReadByName(data.Name, nil)
		summary = fmt.Sprintf("failed to read ACL role %q", data.Name)

	} else {
		resp.Diagnostics.AddError("missing required attribute", `at least one of "id" or "name" is required`)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(summary, err.Error())
		return
	}
	if role == nil {
		resp.Diagnostics.AddError(summary, "no role found")
		return
	}

	resp.Diagnostics.Append(models.Set(ctx, &resp.State, role)...)
}
