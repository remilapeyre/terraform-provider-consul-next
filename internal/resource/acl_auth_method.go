// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewACLAuthMethod() resource.Resource {
	return NewResource(
		"acl_auth_method",
		aclAuthMethodSchema(),
		&ACLAuthMethod{},
	)
}

type ACLAuthMethod struct{}

func (r *ACLAuthMethod) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.ACLAuthMethod
	resp.Diagnostics.Append(models.DecodeACLAuthMethod(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	method, _, err := client.ACL().AuthMethodCreate(&config.Method, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ACL auth method", err.Error())
		return
	}

	config.ID = method.Name
	config.Method = *method
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *ACLAuthMethod) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.ACLAuthMethod
	resp.Diagnostics.Append(models.DecodeACLAuthMethod(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	method, _, err := client.ACL().AuthMethodRead(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to ACL auth method", err.Error())
		return
	}

	state.ID = method.Name
	state.Method = *method
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, state)...)
}

func (r *ACLAuthMethod) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state *structs.ACLAuthMethod
	resp.Diagnostics.Append(models.DecodeACLAuthMethod(ctx, req.Config, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	method, _, err := client.ACL().AuthMethodUpdate(&state.Method, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update ACL auth method", err.Error())
		return
	}

	state.Method = *method
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, state)...)
}

func (r *ACLAuthMethod) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *structs.ACLAuthMethod
	resp.Diagnostics.Append(models.DecodeACLAuthMethod(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.ACL().AuthMethodDelete(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ACL auth method", err.Error())
		return
	}
}

func (r *ACLAuthMethod) ImportState(ctx context.Context, _ *api.Client, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Empty().AtName("id"), req, resp)
}
