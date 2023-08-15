// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewAutopilotConfig() resource.Resource {
	return NewResource(
		"autopilot_config",
		autopilotConfigSchema(),
		&AutopilotConfig{},
	)
}

type AutopilotConfig struct{}

func (r *AutopilotConfig) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.Append(r.upsert(ctx, client, req.Config, &resp.State)...)
}

func (r *AutopilotConfig) upsert(ctx context.Context, client *api.Client, getter models.Getter, setter models.Setter) diag.Diagnostics {
	var config *structs.AutopilotConfig
	diags := models.DecodeAutopilotConfig(ctx, getter, &config)
	if diags.HasError() {
		return diags
	}

	err := client.Operator().AutopilotSetConfiguration(&config.Config, nil)
	if err != nil {
		diags.AddError("failed to create admin partition", err.Error())
		return diags
	}

	config.ID = "autopilot_config"
	diags.Append(models.Set(ctx, setter, config)...)
	return diags
}

func (r *AutopilotConfig) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	config, err := client.Operator().AutopilotGetConfiguration(nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create namespace", err.Error())
		return
	}

	state := &structs.AutopilotConfig{
		ID:     "autopilot_config",
		Config: *config,
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, state)...)
}

func (r *AutopilotConfig) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.Append(r.upsert(ctx, client, req.Config, &resp.State)...)
}

func (r *AutopilotConfig) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// this is a no-op for now
}
