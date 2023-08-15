// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewACLTokenRoleAttachment() resource.Resource {
	return NewResource(
		"acl_token_role_attachment",
		aclTokenRoleAttachmentSchema(),
		&ACLTokenRoleAttachment{},
	)
}

type ACLTokenRoleAttachment struct{}

func (r *ACLTokenRoleAttachment) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.ACLTokenRoleAttachment
	resp.Diagnostics.Append(models.DecodeACLTokenRoleAttachment(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenRead(config.TokenID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read ACL token", err.Error())
		return
	}

	link := &api.ACLLink{
		ID:   config.Role.ID,
		Name: config.Role.Name,
	}
	if attachmentFound(token.Roles, link) != nil {
		ident := link.ID + link.Name
		resp.Diagnostics.AddError("role already attached to the token", fmt.Sprintf("policy %q is already attached to token %q", ident, config.TokenID))
		return
	}

	token.Roles = append(token.Roles, link)

	token, _, err = client.ACL().TokenUpdate(token, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to attach role", err.Error())
		return
	}

	config.ID = fmt.Sprintf("%s-%s", config.TokenID, config.Role.ID+config.Role.Name)
	link = attachmentFound(token.Roles, link)
	if link == nil {
		resp.Diagnostics.AddError("failed to attach role", "failed to find policy attached to token")
		return
	}
	config.Role.ID = link.ID
	config.Role.Name = link.Name
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *ACLTokenRoleAttachment) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.ACLTokenRoleAttachment
	resp.Diagnostics.Append(models.DecodeACLTokenRoleAttachment(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenRead(state.TokenID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL role", err.Error())
		return
	}

	link := &api.ACLLink{
		ID:   state.Role.ID,
		Name: state.Role.Name,
	}
	if attachmentFound(token.Roles, link) == nil {
		resp.State.RemoveResource(ctx)
		return
	}
}

func (r *ACLTokenRoleAttachment) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "acl_token_policy_attachment cannot be updated")
}

func (r *ACLTokenRoleAttachment) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *structs.ACLTokenRoleAttachment
	resp.Diagnostics.Append(models.DecodeACLTokenRoleAttachment(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenRead(state.TokenID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL role", err.Error())
		return
	}
	for i, policy := range token.Roles {
		if state.Role.ID == policy.ID {
			token.Roles = append(token.Roles[:i], token.Roles[i+1:]...)
			break
		}
	}

	_, _, err = client.ACL().TokenUpdate(token, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL role", err.Error())
		return
	}
}
