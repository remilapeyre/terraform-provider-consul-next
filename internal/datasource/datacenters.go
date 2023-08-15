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

func NewDatacenters() datasource.DataSource {
	return NewDataSource(
		"datacenters",
		datacentersSchema(),
		&Datacenters{},
	)
}

type Datacenters struct{}

func (d *Datacenters) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	datacenters, err := client.Catalog().Datacenters()
	if err != nil {
		resp.Diagnostics.AddError("failed to read datacenters", err.Error())
		return
	}

	config := &structs.Datacenters{
		ID:          "datacenters",
		Datacenters: datacenters,
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}
