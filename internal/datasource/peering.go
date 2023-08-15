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

func NewPeering() datasource.DataSource {
	return NewDataSource(
		"peering",
		peeringSchema(),
		&Peering{},
	)
}

type Peering struct{}

func (d *Peering) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config *api.Peering
	resp.Diagnostics.Append(models.DecodePeering(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	peering, _, err := client.Peerings().Read(ctx, config.Name, nil)
	summary := fmt.Sprintf("failed to read peering %q", config.Name)
	if err != nil {
		resp.Diagnostics.AddError(summary, err.Error())
		return
	}
	if peering == nil {
		resp.Diagnostics.AddError(summary, fmt.Sprintf("no peering named %q could be found", config.Name))
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, peering)...)
}
