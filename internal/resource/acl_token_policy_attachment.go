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

func NewACLTokenPolicyAttachment() resource.Resource {
	return NewResource(
		"acl_token_policy_attachment",
		aclTokenPolicyAttachmentSchema(),
		&ACLTokenPolicyAttachment{},
	)
}

type ACLTokenPolicyAttachment struct{}

func (r *ACLTokenPolicyAttachment) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.ACLTokenPolicyAttachment
	resp.Diagnostics.Append(models.DecodeACLTokenPolicyAttachment(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenRead(config.TokenID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read ACL token", err.Error())
		return
	}

	link := &api.ACLLink{
		ID:   config.Policy.ID,
		Name: config.Policy.Name,
	}
	if attachmentFound(token.Policies, link) != nil {
		ident := link.ID + link.Name
		resp.Diagnostics.AddError("policy already attached to the token", fmt.Sprintf("policy %q is already attached to token %q", ident, config.TokenID))
		return
	}

	token.Policies = append(token.Policies, link)

	token, _, err = client.ACL().TokenUpdate(token, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to attach policy", err.Error())
		return
	}

	config.ID = fmt.Sprintf("%s-%s", config.TokenID, config.Policy.ID+config.Policy.Name)
	link = attachmentFound(token.Policies, link)
	if link == nil {
		resp.Diagnostics.AddError("failed to attach policy", "failed to find policy attached to token")
		return
	}
	config.Policy.ID = link.ID
	config.Policy.Name = link.Name
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *ACLTokenPolicyAttachment) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.ACLTokenPolicyAttachment
	resp.Diagnostics.Append(models.DecodeACLTokenPolicyAttachment(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenRead(state.TokenID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL token", err.Error())
		return
	}

	link := &api.ACLLink{
		ID:   state.Policy.ID,
		Name: state.Policy.Name,
	}
	if attachmentFound(token.Policies, link) == nil {
		resp.State.RemoveResource(ctx)
		return
	}
}

func (r *ACLTokenPolicyAttachment) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "acl_token_policy_attachment cannot be updated")
}

func (r *ACLTokenPolicyAttachment) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *structs.ACLTokenPolicyAttachment
	resp.Diagnostics.Append(models.DecodeACLTokenPolicyAttachment(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, _, err := client.ACL().TokenRead(state.TokenID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL policy", err.Error())
		return
	}
	for i, policy := range token.Policies {
		if state.Policy.ID == policy.ID {
			token.Policies = append(token.Policies[:i], token.Policies[i+1:]...)
			break
		}
	}

	_, _, err = client.ACL().TokenUpdate(token, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to detach ACL policy", err.Error())
		return
	}
}

func attachmentFound(links []*api.ACLLink, link *api.ACLLink) *api.ACLLink {
	for _, l := range links {
		if l.ID == link.ID || l.Name == link.Name {
			return l
		}
	}
	return nil
}
