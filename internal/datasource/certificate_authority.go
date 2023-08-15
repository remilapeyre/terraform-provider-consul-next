// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func NewCertificateAuthority() datasource.DataSource {
	return NewDataSource(
		"certificate_authority",
		certificateAutoritySchema(),
		&CertificateAuthority{},
	)
}

type CertificateAuthority struct{}

func (r *CertificateAuthority) Read(ctx context.Context, client *api.Client, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	config, _, err := client.Connect().CAGetConfig(nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to read Connect CA configuration", err.Error())
		return
	}

	state := &structs.CAConfig{
		ID:     "certificate_autority",
		Config: *config,
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, state)...)
}
