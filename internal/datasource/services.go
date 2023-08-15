// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

func NewServices() datasource.DataSource {
	return NewDataSource(
		"services",
		catalogServiceSchema(),
		&Services{},
	)
}

type Services struct{}

func (d *Services) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// var service *api.CatalogServices
	// resp.Diagnostics.Append(models.DecodeCatalogServices(ctx, req.Config, &service)...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	client.Catalog().Services(nil)
	// if err != nil {
	// 	resp.Diagnostics.AddError(fmt.Sprintf("failed to read ACL token %q", service.AccessorID), err.Error())
	// 	return
	// }
	// resp.Diagnostics.Append(models.Set(ctx, &resp.State, token)...)
}
