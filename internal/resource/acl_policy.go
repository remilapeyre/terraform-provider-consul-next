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

func NewACLPolicy() resource.Resource {
	return NewResource(
		"acl_policy",
		aclPolicySchema(),
		&ACLPolicy{},
	)
}

type ACLPolicy struct{}

func (r *ACLPolicy) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *api.ACLPolicy
	resp.Diagnostics.Append(models.DecodeACLPolicy(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, _, err := client.ACL().PolicyCreate(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ACL policy", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, policy)...)
}

func (r *ACLPolicy) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *api.ACLPolicy
	resp.Diagnostics.Append(models.DecodeACLPolicy(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, _, err := client.ACL().PolicyRead(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create namespace", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, policy)...)
}

func (r *ACLPolicy) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config *api.ACLPolicy
	resp.Diagnostics.Append(models.DecodeACLPolicy(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, _, err := client.ACL().PolicyUpdate(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update ACL policy", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, policy)...)
}

func (r *ACLPolicy) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *api.ACLPolicy
	resp.Diagnostics.Append(models.DecodeACLPolicy(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.ACL().PolicyDelete(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete ACL policy", err.Error())
		return
	}
}

func (r *ACLPolicy) ImportState(ctx context.Context, _ *api.Client, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Empty().AtName("id"), req, resp)
}
