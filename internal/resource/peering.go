// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewPeering() resource.Resource {
	return NewResource(
		"peering",
		peeringSchema(),
		&Peering{},
	)
}

type Peering struct{}

func (r *Peering) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.PeeringResource
	resp.Diagnostics.Append(models.DecodePeeringResource(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := client.Peerings().Establish(ctx, api.PeeringEstablishRequest{
		PeerName:     config.Peering.Name,
		PeeringToken: config.PeeringToken,
		Partition:    config.Peering.Partition,
		Meta:         config.Peering.Meta,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to establish peering", err.Error())
		return
	}

	peering, _, err := client.Peerings().Read(ctx, config.Peering.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read peering", err.Error())
		return
	}

	config.Peering = *peering
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *Peering) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.PeeringResource
	resp.Diagnostics.Append(models.DecodePeeringResource(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	peering, _, err := client.Peerings().Read(ctx, state.Peering.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read peering", err.Error())
		return
	}
	if peering == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	state.Peering = *peering
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, state)...)

}

func (r *Peering) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("update not implemented", "consul_peering cannot be updated")
}

func (r *Peering) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *structs.PeeringResource
	resp.Diagnostics.Append(models.DecodePeeringResource(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.Peerings().Delete(ctx, state.Peering.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete peering", err.Error())
		return
	}
}
