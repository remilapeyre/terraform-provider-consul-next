// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

func NewService() datasource.DataSource {
	return NewDataSource(
		"service",
		catalogServiceSchema(),
		&Service{},
	)
}

type Service struct{}

func (d *Service) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// var service *api.CatalogService
	// resp.Diagnostics.Append(models.DecodeCatalogService(ctx, req.Config, &service)...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// services, _, err := client.Catalog().Service(service.ID, "", nil)
	// if err != nil {
	// 	resp.Diagnostics.AddError(fmt.Sprintf("failed to read ACL token %q", service.AccessorID), err.Error())
	// 	return
	// }
	// resp.Diagnostics.Append(models.Set(ctx, &resp.State, token)...)
}
