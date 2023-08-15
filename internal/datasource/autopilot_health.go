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

func NewAutopilotHealth() datasource.DataSource {
	return NewDataSource(
		"autopilot_health",
		autopilotHealthSchema(),
		&AutopilotHealth{},
	)
}

type AutopilotHealth struct{}

func (d *AutopilotHealth) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	health, err := client.Operator().AutopilotServerHealth(nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read autopilot health status", err.Error())
		return
	}

	res := &structs.AutopilotHealth{
		ID:              "autopilot_health",
		AutopilotHealth: *health,
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, res)...)
}
