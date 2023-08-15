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

func NewACLTokenSecretID() datasource.DataSource {
	return NewDataSource(
		"acl_token_secret_id",
		aclTokenSecretIdSchema(),
		&ACLTokenSecretID{},
	)
}

type ACLTokenSecretID struct{}

func (d *ACLTokenSecretID) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *structs.ACLTokenSecretID
	resp.Diagnostics.Append(models.DecodeACLTokenSecretID(ctx, req.Config, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenRead(data.AccessorID, nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to read ACL token %q", data.AccessorID), err.Error())
		return
	}

	data.SecretID = token.SecretID
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, data)...)
}
