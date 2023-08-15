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

func NewNetworkAreaMembers() datasource.DataSource {
	return NewDataSource(
		"network_area_members",
		networkAreaMembersSchema(),
		&NetworkAreaMembers{},
	)
}

type NetworkAreaMembers struct{}

func (d *NetworkAreaMembers) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config *structs.NetworkAreaMembers
	resp.Diagnostics.Append(models.DecodeNetworkAreaMembers(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	members, _, err := client.Operator().AreaMembers(config.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to read network area members %q", config.ID), err.Error())
		return
	}

	config.Members = members
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}
