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

func NewACLBindingRule() resource.Resource {
	return NewResource(
		"acl_binding_rule",
		aclBindingRuleSchema(),
		&ACLBindingRule{},
	)
}

type ACLBindingRule struct{}

func (r *ACLBindingRule) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *api.ACLBindingRule
	resp.Diagnostics.Append(models.DecodeACLBindingRule(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rule, _, err := client.ACL().BindingRuleCreate(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ACL binding rule", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, rule)...)
}

func (r *ACLBindingRule) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *api.ACLBindingRule
	resp.Diagnostics.Append(models.DecodeACLBindingRule(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rule, _, err := client.ACL().BindingRuleRead(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ACL binding rule", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, rule)...)
}

func (r *ACLBindingRule) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config *api.ACLBindingRule
	resp.Diagnostics.Append(models.DecodeACLBindingRule(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rule, _, err := client.ACL().BindingRuleUpdate(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update ACL binding rule", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, rule)...)
}

func (r *ACLBindingRule) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *api.ACLBindingRule
	resp.Diagnostics.Append(models.DecodeACLBindingRule(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.ACL().BindingRuleDelete(state.ID, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete ACL binding rule", err.Error())
		return
	}
}

func (r *ACLBindingRule) ImportState(ctx context.Context, _ *api.Client, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Empty().AtName("id"), req, resp)
}
