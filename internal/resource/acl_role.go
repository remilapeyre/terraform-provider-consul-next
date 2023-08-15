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

func NewACLRole() resource.Resource {
	return NewResource(
		"acl_role",
		aclRoleSchema(),
		&ACLRole{},
	)
}

type ACLRole struct{}

func (r *ACLRole) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *api.ACLRole
	resp.Diagnostics.Append(models.DecodeACLRole(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	role, _, err := client.ACL().RoleCreate(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ACL role", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, role)...)
}

func (r *ACLRole) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *api.ACLRole
	resp.Diagnostics.Append(models.DecodeACLRole(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	role, _, err := client.ACL().RoleRead(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read ACL role", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, role)...)
}

func (r *ACLRole) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config *api.ACLRole
	resp.Diagnostics.Append(models.DecodeACLRole(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	role, _, err := client.ACL().RoleUpdate(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update ACL role", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, role)...)
}

func (r *ACLRole) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *api.ACLRole
	resp.Diagnostics.Append(models.DecodeACLRole(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.ACL().RoleDelete(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete ACL role", err.Error())
		return
	}
}

func (r *ACLRole) ImportState(ctx context.Context, _ *api.Client, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Empty().AtName("id"), req, resp)
}
