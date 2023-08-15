// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewConfigEntry() datasource.DataSource {
	return NewDataSource(
		"config_entry",
		configEntrySchema(),
		&ConfigEntry{},
	)
}

type ConfigEntry struct{}

func (d *ConfigEntry) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *structs.ConfigEntry
	resp.Diagnostics.Append(models.DecodeConfigEntry(ctx, req.Config, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	config, _, err := client.ConfigEntries().Get(data.Kind, data.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to read config entry %q", data.Name), err.Error())
		return
	}

	data, err = encodeConfigEntry(config)
	if err != nil {
		resp.Diagnostics.AddError("failed to encode config entry", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, data)...)
}

func encodeConfigEntry(config api.ConfigEntry) (*structs.ConfigEntry, error) {
	res := &structs.ConfigEntry{
		ID:        config.GetName(),
		Name:      config.GetName(),
		Kind:      config.GetKind(),
		Namespace: config.GetNamespace(),
		Partition: config.GetPartition(),
		Meta:      config.GetMeta(),
	}
	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	conf := map[string]interface{}{}
	if err := json.Unmarshal(data, &conf); err != nil {
		return nil, err
	}
	res.Config = conf
	return res, nil
}
