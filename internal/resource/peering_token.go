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

func NewPeeringToken() resource.Resource {
	return NewResource(
		"peering_token",
		peeringTokenSchema(),
		&PeeringToken{},
	)
}

type PeeringToken struct{}

func (r *PeeringToken) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.PeeringToken
	resp.Diagnostics.Append(models.DecodePeeringToken(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, _, err := client.Peerings().GenerateToken(ctx, config.Request, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create peering token", err.Error())
		return
	}

	config.ID = "peering_token"
	config.Response = *response
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *PeeringToken) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var config *structs.PeeringToken
	resp.Diagnostics.Append(models.DecodePeeringToken(ctx, req.State, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	peering, _, err := client.Peerings().Read(ctx, config.Request.PeerName, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to find peer", err.Error())
		return
	}
	if peering == nil {
		resp.State.RemoveResource(ctx)
	}
}

func (r *PeeringToken) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not implemented", "peering_token cannot be updated")
}

func (r *PeeringToken) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var config *structs.PeeringToken
	resp.Diagnostics.Append(models.DecodePeeringToken(ctx, req.State, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.Peerings().Delete(ctx, config.Request.PeerName, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete peer", err.Error())
		return
	}
}
