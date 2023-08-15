// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func NewCatalogService() resource.Resource {
	return NewResource(
		"service",
		catalogServiceSchema(),
		&CatalogService{},
	)
}

type CatalogService struct{}

func (r *CatalogService) Create(context.Context, *api.Client, resource.CreateRequest, *resource.CreateResponse) {
}

func (r *CatalogService) Read(context.Context, *api.Client, resource.ReadRequest, *resource.ReadResponse) {
}

func (r *CatalogService) Update(context.Context, *api.Client, resource.UpdateRequest, *resource.UpdateResponse) {
}

func (r *CatalogService) Delete(context.Context, *api.Client, resource.DeleteRequest, *resource.DeleteResponse) {
}
