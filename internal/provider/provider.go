// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/consul/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	consuldatasource "github.com/remilapeyre/terraform-provider-consul-next/internal/datasource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	consulresource "github.com/remilapeyre/terraform-provider-consul-next/internal/resource"
)

// Ensure ConsulProvider satisfies various provider interfaces.
var _ provider.Provider = &ConsulProvider{}

// ConsulProvider defines the provider implementation.
type ConsulProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *ConsulProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "consul"
	resp.Version = p.version
}

func (p *ConsulProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = configSchema()
}

func (p *ConsulProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	config := api.DefaultConfig()
	resp.Diagnostics.Append(models.DecodeConfig(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := api.NewClient(config)
	if err != nil {
		resp.Diagnostics.AddError("failed to create Consul API client", err.Error())
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *ConsulProvider) Resources(ctx context.Context) []func() resource.Resource {
	resources := []func() resource.Resource{
		consulresource.NewACLAuthMethod,
		consulresource.NewACLBindingRule,
		consulresource.NewACLPolicy,
		consulresource.NewACLRole,
		consulresource.NewACLToken,
		consulresource.NewACLTokenPolicyAttachment,
		consulresource.NewACLTokenRoleAttachment,
		consulresource.NewAdminPartition,
		consulresource.NewAutopilotConfig,
		consulresource.NewCatalogService,
		consulresource.NewCertificateAuthority,
		consulresource.NewKeyPrefix,
		consulresource.NewNamespace,
		consulresource.NewNamespacePolicyAttachment,
		consulresource.NewNamespaceRoleAttachment,
		consulresource.NewNetworkArea,
		consulresource.NewNode,
		consulresource.NewPeering,
		consulresource.NewPeeringToken,
		consulresource.NewPreparedQuery,
	}

	resources = append(resources, consulresource.ConfigEntriesResources()...)
	return resources
}

func (p *ConsulProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		consuldatasource.NewACLAuthMethod,
		consuldatasource.NewACLBindingRule,
		consuldatasource.NewACLPolicy,
		consuldatasource.NewACLRole,
		consuldatasource.NewACLToken,
		consuldatasource.NewACLTokenSecretID,
		consuldatasource.NewAdminPartition,
		consuldatasource.NewAgentConfig,
		consuldatasource.NewAutopilotHealth,
		consuldatasource.NewCertificateAuthority,
		consuldatasource.NewConfigEntry,
		consuldatasource.NewDatacenters,
		consuldatasource.NewKeyPrefix,
		consuldatasource.NewKeys,
		consuldatasource.NewNamespace,
		consuldatasource.NewNetworkAreaMembers,
		consuldatasource.NewNetworkSegments,
		consuldatasource.NewNode,
		consuldatasource.NewNodes,
		consuldatasource.NewPeering,
		consuldatasource.NewPeerings,
		consuldatasource.NewService,
		consuldatasource.NewServiceHealth,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ConsulProvider{
			version: version,
		}
	}
}
