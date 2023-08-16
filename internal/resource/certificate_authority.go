// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewCertificateAuthority() resource.Resource {
	return NewResource(
		"certificate_authority",
		certificateAuthoritySchema(),
		&CertificateAuthority{},
	)
}

type CertificateAuthority struct{}

func (r *CertificateAuthority) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.CAConfig
	resp.Diagnostics.Append(models.DecodeCAConfig(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.Connect().CASetConfig(&config.Config, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update Connect CA config", err.Error())
		return
	}

	config.ID = "certificate_authority"
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *CertificateAuthority) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.CAConfig
	resp.Diagnostics.Append(models.DecodeCAConfig(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	config, _, err := client.Connect().CAGetConfig(nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read Connect CA config", err.Error())
		return
	}

	state.Config = *config
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, state)...)
}

func (r *CertificateAuthority) Update(context.Context, *api.Client, resource.UpdateRequest, *resource.UpdateResponse) {
}

func (r *CertificateAuthority) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// this is a no-op for now
}
