// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewConfigEntry() resource.Resource {
	return NewResource(
		"config_entry",
		configEntrySchema(),
		&ConfigEntry{},
	)
}

type ConfigEntry struct{}

func (r *ConfigEntry) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.ConfigEntry
	resp.Diagnostics.Append(models.DecodeConfigEntry(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	raw := map[string]interface{}{}
	for k, v := range config.Config {
		raw[k] = v
	}
	raw["Kind"] = config.Kind
	raw["Name"] = config.Name
	raw["Namespace"] = config.Namespace
	raw["Partition"] = config.Partition
	raw["Meta"] = config.Meta

	entry, err := api.DecodeConfigEntry(raw)
	if err != nil {
		resp.Diagnostics.AddError("failed to decode config entry", err.Error())
		return
	}

	_, _, err = client.ConfigEntries().Set(entry, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create config entry", err.Error())
		return
	}

	config.ID = fmt.Sprintf("%s/%s", config.Kind, config.Name)
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *ConfigEntry) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.ConfigEntry
	resp.Diagnostics.Append(models.DecodeConfigEntry(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	entry, _, err := client.ConfigEntries().Get(state.Kind, state.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ConfigEntry", err.Error())
		return
	}

	data, err := json.Marshal(entry)
	if err != nil {
		resp.Diagnostics.AddError("failed to marshal config entry", err.Error())
		return
	}

	m := map[string]interface{}{}
	if err := json.Unmarshal(data, &m); err != nil {
		resp.Diagnostics.AddError("failed to unmarshal config entry", err.Error())
		return
	}

	delete(m, "Kind")
	delete(m, "Name")
	delete(m, "Namespace")
	delete(m, "Partition")
	delete(m, "Meta")
	delete(m, "CreateIndex")
	delete(m, "ModifyIndex")

	state.Config = m
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, state)...)
}

func (r *ConfigEntry) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config *structs.ConfigEntry
	resp.Diagnostics.Append(models.DecodeConfigEntry(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	entry, err := api.DecodeConfigEntry(config.Config)
	if err != nil {
		resp.Diagnostics.AddError("failed to decode config entry", err.Error())
		return
	}

	_, _, err = client.ConfigEntries().Set(entry, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update config entry", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *ConfigEntry) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *structs.ConfigEntry
	resp.Diagnostics.Append(models.DecodeConfigEntry(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.ConfigEntries().Delete(state.Kind, state.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete config entry", err.Error())
		return
	}
}
