// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
)

func NewNamespace() datasource.DataSource {
	return NewDataSource(
		"namespace",
		namespaceSchema(),
		&Namespace{},
	)
}

type Namespace struct{}

func (d *Namespace) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config *api.Namespace
	resp.Diagnostics.Append(models.DecodeNamespace(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace, _, err := client.Namespaces().Read(config.Name, nil)
	summary := fmt.Sprintf("failed to read namespace %q", config.Name)
	if err != nil {
		resp.Diagnostics.AddError(summary, err.Error())
		return
	}
	if namespace == nil {
		resp.Diagnostics.AddError(summary, fmt.Sprintf("no namespace named %q found", config.Name))
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, namespace)...)
}
