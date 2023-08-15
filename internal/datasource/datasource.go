// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type DataSource struct {
	name   string
	schema schema.Schema
	client *api.Client
	impl   DataSourceImplementation
}

func NewDataSource(name string, schema schema.Schema, impl DataSourceImplementation) *DataSource {
	return &DataSource{
		name:   name,
		schema: schema,
		impl:   impl,
	}
}

type DataSourceImplementation interface {
	Read(context.Context, *api.Client, datasource.ReadRequest, *datasource.ReadResponse)
}

func (d *DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.name
}

func (d *DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = d.schema
}

func (d *DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.impl == nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Missing %s Data Source implementation", d.name),
			fmt.Sprintf("No Data Source implementation found for %s. Please report this issue to the provider developers.", d.name),
		)
		return
	}
	d.impl.Read(ctx, d.client, req, resp)
}
