// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
)

func NewNamespace() resource.Resource {
	return NewResource(
		"namespace",
		namespaceSchema(),
		&Namespace{},
	)
}

type Namespace struct{}

func (r *Namespace) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *api.Namespace
	resp.Diagnostics.Append(models.DecodeNamespace(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace, _, err := client.Namespaces().Create(config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create namespace", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, namespace)...)
}

func (r *Namespace) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *api.Namespace
	resp.Diagnostics.Append(models.DecodeNamespace(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace, _, err := client.Namespaces().Read(state.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create namespace", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, namespace)...)
}

func (r *Namespace) Update(context.Context, *api.Client, resource.UpdateRequest, *resource.UpdateResponse) {
}

func (r *Namespace) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *api.Namespace
	resp.Diagnostics.Append(models.DecodeNamespace(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.Namespaces().Delete(state.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete namespace", err.Error())
		return
	}
}
