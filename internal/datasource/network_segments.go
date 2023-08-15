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

func NewNetworkSegments() datasource.DataSource {
	return NewDataSource(
		"network_segments",
		networkSegmentsSchema(),
		&NetworkSegments{},
	)
}

type NetworkSegments struct{}

func (d *NetworkSegments) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	segments, _, err := client.Operator().SegmentList(nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read network segments", err.Error())
		return
	}

	data := &structs.NetworkSegments{
		ID:       "network_segments",
		Segments: segments,
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, data)...)
}
