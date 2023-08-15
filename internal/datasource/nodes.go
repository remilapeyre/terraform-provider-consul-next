// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewNodes() datasource.DataSource {
	return NewDataSource(
		"nodes",
		nodesSchema(),
		&Nodes{},
	)
}

type Nodes struct{}

func (d *Nodes) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	nodes, _, err := client.Catalog().Nodes(nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to list nodes", err.Error())
		return
	}

	res := &structs.Nodes{
		ID:    "nodes",
		Nodes: nodes,
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, res)...)
}
