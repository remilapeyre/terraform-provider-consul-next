// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewACLToken() resource.Resource {
	return NewResource(
		"acl_token",
		aclTokenSchema(),
		&ACLToken{},
	)
}

type ACLToken struct{}

func (r *ACLToken) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.ACLToken
	resp.Diagnostics.Append(models.DecodeACLToken(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenCreate(&config.Token, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ACL token", err.Error())
		return
	}

	config.ID = token.AccessorID
	config.Token = *token
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *ACLToken) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.ACLToken
	resp.Diagnostics.Append(models.DecodeACLToken(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenRead(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create namespace", err.Error())
		return
	}

	state.Token = *token
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, state)...)
}

func (r *ACLToken) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config *structs.ACLToken
	resp.Diagnostics.Append(models.DecodeACLToken(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenUpdate(&config.Token, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update ACL token", err.Error())
		return
	}

	config.Token = *token
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *ACLToken) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *structs.ACLToken
	resp.Diagnostics.Append(models.DecodeACLToken(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.ACL().TokenDelete(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete ACL token", err.Error())
		return
	}
}
