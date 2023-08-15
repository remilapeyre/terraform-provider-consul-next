// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
)

func NewAdminPartition() datasource.DataSource {
	return NewDataSource(
		"admin_partition",
		adminPartitionSchema(),
		&AdminPartition{},
	)
}

type AdminPartition struct{}

func (d *AdminPartition) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var partition *api.Partition
	resp.Diagnostics.Append(models.DecodePartition(ctx, req.Config, &partition)...)
	if resp.Diagnostics.HasError() {
		return
	}

	partition, _, err := client.Partitions().Read(ctx, partition.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to read admin partition %q", partition.Name), err.Error())
		return
	}

	resp.Diagnostics.Append(models.Set(ctx, &resp.State, partition)...)
}
