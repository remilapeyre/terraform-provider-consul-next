// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
)

func NewAdminPartition() resource.Resource {
	return NewResource(
		"admin_partition",
		adminPartitionSchema(),
		&AdminPartition{},
	)
}

type AdminPartition struct{}

func (r *AdminPartition) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *api.Partition
	resp.Diagnostics.Append(models.DecodePartition(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	partition, _, err := client.Partitions().Create(ctx, config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create admin partition", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, partition)...)
}

func (r *AdminPartition) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *api.Partition
	resp.Diagnostics.Append(models.DecodePartition(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	partition, _, err := client.Partitions().Read(ctx, state.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create admin partition", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, partition)...)
}

func (r *AdminPartition) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config *api.Partition
	resp.Diagnostics.Append(models.DecodePartition(ctx, req.State, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	partition, _, err := client.Partitions().Update(ctx, config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create admin partition", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, partition)...)
}

func (r *AdminPartition) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *api.Partition
	resp.Diagnostics.Append(models.DecodePartition(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.Partitions().Delete(ctx, state.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete admin partition", err.Error())
		return
	}
}
