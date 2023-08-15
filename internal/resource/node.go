// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func NewNode() resource.Resource {
	return NewResource(
		"node",
		nodeSchema(),
		&Node{},
	)
}

type Node struct{}

func (r *Node) Create(context.Context, *api.Client, resource.CreateRequest, *resource.CreateResponse) {
}

func (r *Node) Read(context.Context, *api.Client, resource.ReadRequest, *resource.ReadResponse) {}

func (r *Node) Update(context.Context, *api.Client, resource.UpdateRequest, *resource.UpdateResponse) {
}

func (r *Node) Delete(context.Context, *api.Client, resource.DeleteRequest, *resource.DeleteResponse) {
}
