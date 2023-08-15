// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func NewKeyPrefix() resource.Resource {
	return NewResource(
		"key_prefix",
		aclPolicySchema(),
		&KeyPrefix{},
	)
}

type KeyPrefix struct{}

func (r *KeyPrefix) Create(context.Context, *api.Client, resource.CreateRequest, *resource.CreateResponse) {
}

func (r *KeyPrefix) Read(context.Context, *api.Client, resource.ReadRequest, *resource.ReadResponse) {
}

func (r *KeyPrefix) Update(context.Context, *api.Client, resource.UpdateRequest, *resource.UpdateResponse) {
}

func (r *KeyPrefix) Delete(context.Context, *api.Client, resource.DeleteRequest, *resource.DeleteResponse) {
}
