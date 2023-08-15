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

func NewACLToken() datasource.DataSource {
	return NewDataSource(
		"acl_token",
		aclTokenSchema(),
		&ACLToken{},
	)
}

type ACLToken struct{}

func (d *ACLToken) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *structs.ACLToken
	resp.Diagnostics.Append(models.DecodeACLToken(ctx, req.Config, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenRead(data.Token.AccessorID, nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to read ACL token %q", data.Token.AccessorID), err.Error())
		return
	}

	data.ID = token.AccessorID
	data.Token = *token
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, data)...)
}
