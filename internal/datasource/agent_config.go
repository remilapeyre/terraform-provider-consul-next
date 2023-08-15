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

func NewAgentConfig() datasource.DataSource {
	return NewDataSource(
		"agent_config",
		agentConfigSchema(),
		&AgentConfig{},
	)
}

type AgentConfig struct{}

func (d *AgentConfig) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	config, err := client.Agent().Self()
	if err != nil {
		resp.Diagnostics.AddError("failed to read agent configuration", err.Error())
		return
	}

	res := &structs.AgentConfig{
		ID:                "agent_config",
		Datacenter:        config["Config"]["Datacenter"].(string),
		PrimaryDatacenter: config["Config"]["PrimaryDatacenter"].(string),
		NodeName:          config["Config"]["NodeName"].(string),
		NodeID:            config["Config"]["NodeID"].(string),
		Revision:          config["Config"]["Revision"].(string),
		Server:            config["Config"]["Server"].(bool),
		Version:           config["Config"]["Version"].(string),
		BuildDate:         config["Config"]["BuildDate"].(string),
	}
	if config["Config"]["Partition"] != nil {
		partition := config["Config"]["Partition"].(string)
		res.Partition = &partition
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, res)...)
}
