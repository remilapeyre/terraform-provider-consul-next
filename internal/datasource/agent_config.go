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
		ID: "agent_config",
	}
	if datacenter, ok := config["Config"]["Datacenter"].(string); ok {
		res.Datacenter = datacenter
	}
	if primaryDatacenter, ok := config["Config"]["PrimaryDatacenter"].(string); ok {
		res.PrimaryDatacenter = primaryDatacenter
	}
	if nodeName, ok := config["Config"]["NodeName"].(string); ok {
		res.NodeName = nodeName
	}
	if nodeID, ok := config["Config"]["NodeID"].(string); ok {
		res.NodeID = nodeID
	}
	if revision, ok := config["Config"]["Revision"].(string); ok {
		res.Revision = revision
	}
	if server, ok := config["Config"]["Server"].(bool); ok {
		res.Server = server
	}
	if version, ok := config["Config"]["Version"].(string); ok {
		res.Version = version
	}
	if buildDate, ok := config["Config"]["BuildDate"].(string); ok {
		res.BuildDate = buildDate
	}
	if config["Config"]["Partition"] != nil {
		partition, ok := config["Config"]["Partition"].(string)
		if ok {
			res.Partition = &partition
		}
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, res)...)
}
