// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
)

func NewNetworkArea() resource.Resource {
	return NewResource(
		"network_area",
		areaSchema(),
		&NetworkArea{},
	)
}

type NetworkArea struct{}

func (r *NetworkArea) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *api.Area
	resp.Diagnostics.Append(models.DecodeArea(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, _, err := client.Operator().AreaCreate(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create network area", err.Error())
		return
	}

	config.ID = id
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *NetworkArea) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *api.Area
	resp.Diagnostics.Append(models.DecodeArea(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	areas, _, err := client.Operator().AreaGet(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read network area", err.Error())
		return
	}
	if len(areas) == 0 {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(models.Set(ctx, &resp.State, areas[0])...)
}

func (r *NetworkArea) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, config *api.Area
	resp.Diagnostics.Append(models.DecodeArea(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(models.DecodeArea(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, _, err := client.Operator().AreaUpdate(state.ID, config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create network area", err.Error())
		return
	}

	config.ID = id
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *NetworkArea) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *api.Area
	resp.Diagnostics.Append(models.DecodeArea(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.Operator().AreaDelete(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete network area", err.Error())
		return
	}
}
