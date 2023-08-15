// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type Resource struct {
	name   string
	schema schema.Schema
	client *api.Client
	impl   ResourceImplementation
}

func NewResource(name string, schema schema.Schema, impl ResourceImplementation) *Resource {
	return &Resource{
		name:   name,
		schema: schema,
		impl:   impl,
	}
}

type ResourceImplementation interface {
	Create(context.Context, *api.Client, resource.CreateRequest, *resource.CreateResponse)
	Read(context.Context, *api.Client, resource.ReadRequest, *resource.ReadResponse)
	Update(context.Context, *api.Client, resource.UpdateRequest, *resource.UpdateResponse)
	Delete(context.Context, *api.Client, resource.DeleteRequest, *resource.DeleteResponse)
}

type ResourceImplementationWithImportState interface {
	ResourceImplementation

	ImportState(context.Context, *api.Client, resource.ImportStateRequest, *resource.ImportStateResponse)
}

type ResourceImplementationWithModifyPlan interface {
	ResourceImplementation

	ModifyPlan(context.Context, *api.Client, resource.ModifyPlanRequest, *resource.ModifyPlanResponse)
}

func (r *Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.name
}

func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.schema
}

func (r *Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	if r.impl == nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Missing %s Resource implementation", r.name),
			fmt.Sprintf("No Resource implementation found for %s. Please report this issue to the provider developers.", r.name),
		)
		return
	}
	r.impl.Create(ctx, r.client, req, resp)
}

func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	if r.impl == nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Missing %s Resource implementation", r.name),
			fmt.Sprintf("No Resource implementation found for %s. Please report this issue to the provider developers.", r.name),
		)
		return
	}
	r.impl.Read(ctx, r.client, req, resp)
}

func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	if r.impl == nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Missing %s Resource implementation", r.name),
			fmt.Sprintf("No Resource implementation found for %s. Please report this issue to the provider developers.", r.name),
		)
		return
	}
	r.impl.Update(ctx, r.client, req, resp)
}

func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	if r.impl == nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Missing %s Resource implementation", r.name),
			fmt.Sprintf("No Resource implementation found for %s. Please report this issue to the provider developers.", r.name),
		)
		return
	}
	r.impl.Delete(ctx, r.client, req, resp)
}

func (r *Resource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if r.impl == nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Missing %s Resource implementation", r.name),
			fmt.Sprintf("No Resource implementation found for %s. Please report this issue to the provider developers.", r.name),
		)
		return
	}
	if resourceWithModifyPlan, ok := r.impl.(ResourceImplementationWithModifyPlan); ok {
		resourceWithModifyPlan.ModifyPlan(ctx, r.client, req, resp)
	}
}

func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resourceWithImportState, ok := r.impl.(ResourceImplementationWithImportState)

	if !ok {
		resp.Diagnostics.AddError(
			"Resource Import Not Implemented",
			"This resource does not support import. Please contact the provider developer for additional information.",
		)
		return
	}

	resourceWithImportState.ImportState(ctx, r.client, req, resp)
}
