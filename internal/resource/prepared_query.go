// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
)

func NewPreparedQuery() resource.Resource {
	return NewResource(
		"prepared_query",
		preparedQueryDefinitionSchema(),
		&PreparedQuery{},
	)
}

type PreparedQuery struct{}

func (r *PreparedQuery) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *api.PreparedQueryDefinition
	resp.Diagnostics.Append(models.DecodePreparedQueryDefinition(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, _, err := client.PreparedQuery().Create(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create prepared query", err.Error())
		return
	}

	config.ID = id
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *PreparedQuery) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *api.PreparedQueryDefinition
	resp.Diagnostics.Append(models.DecodePreparedQueryDefinition(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	definition, _, err := client.PreparedQuery().Get(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read prepared query", err.Error())
		return
	}
	if len(definition) == 0 {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(models.Set(ctx, &resp.State, definition[0])...)
}

func (r *PreparedQuery) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config, state *api.PreparedQueryDefinition
	resp.Diagnostics.Append(models.DecodePreparedQueryDefinition(ctx, req.Config, &config)...)
	resp.Diagnostics.Append(models.DecodePreparedQueryDefinition(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.ID = state.ID
	_, err := client.PreparedQuery().Update(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update prepared query", err.Error())
		return
	}

	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *PreparedQuery) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *api.PreparedQueryDefinition
	resp.Diagnostics.Append(models.DecodePreparedQueryDefinition(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.PreparedQuery().Delete(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete prepared query", err.Error())
		return
	}
}

func (r *PreparedQuery) ImportState(ctx context.Context, _ *api.Client, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Empty().AtName("id"), req, resp)
}
