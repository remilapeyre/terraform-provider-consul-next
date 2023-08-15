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

func NewNode() datasource.DataSource {
	return NewDataSource(
		"node",
		nodeSchema(),
		&Node{},
	)
}

type Node struct{}

func (d *Node) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config *api.CatalogNode
	resp.Diagnostics.Append(models.DecodeCatalogNode(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	node, _, err := client.Catalog().Node(config.Node.ID, nil)
	summary := fmt.Sprintf("failed to read node %q", config.Node.ID)
	if err != nil {
		resp.Diagnostics.AddError(summary, err.Error())
		return
	}
	if node == nil {
		resp.Diagnostics.AddError(summary, fmt.Sprintf("no node named %q could be found", config.Node.ID))
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, node)...)
}
