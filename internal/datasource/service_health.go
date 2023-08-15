// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

func NewServiceHealth() datasource.DataSource {
	return NewDataSource(
		"service_health",
		serviceHealthSchema(),
		&ServiceHealth{},
	)
}

type ServiceHealth struct{}

func (d *ServiceHealth) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// var service *api.CatalogServiceHealth
	// resp.Diagnostics.Append(models.DecodeCatalogServiceHealth(ctx, req.Config, &service)...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	client.Health().Service("", "", true, nil)
	// if err != nil {
	// 	resp.Diagnostics.AddError(fmt.Sprintf("failed to read ACL token %q", service.AccessorID), err.Error())
	// 	return
	// }
	// resp.Diagnostics.Append(models.Set(ctx, &resp.State, token)...)
}
