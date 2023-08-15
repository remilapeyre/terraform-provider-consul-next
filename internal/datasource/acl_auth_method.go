// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewACLAuthMethod() datasource.DataSource {
	return NewDataSource(
		"acl_auth_method",
		aclAuthMethodSchema(),
		&ACLAuthMethod{},
	)
}

type ACLAuthMethod struct{}

func (d *ACLAuthMethod) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *structs.ACLAuthMethod
	resp.Diagnostics.Append(models.DecodeACLAuthMethod(ctx, req.Config, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	method, _, err := client.ACL().AuthMethodRead(data.Method.Name, nil)
	summary := fmt.Sprintf("failed to read ACL auth method %q", data.Method.Name)
	if err != nil {
		resp.Diagnostics.AddError(summary, err.Error())
		return
	}
	if method == nil {
		resp.Diagnostics.AddError(summary, fmt.Sprintf("no ACL auth method %q found", data.Method.Name))
		return
	}

	data.ID = "acl_auth_method"
	data.Method = *method
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, data)...)
}
