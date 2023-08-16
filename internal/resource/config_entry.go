// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/models"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
)

func ConfigEntriesResources() []func() resource.Resource {
	configEntries := []func() resource.Resource{
		func() resource.Resource {
			return NewResource(
				"config_entry",
				configEntrySchema(),
				&GenericConfigEntry{},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_tcp_route",
				tcpRouteConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.TCPRouteConfigEntry{
							Kind: api.TCPRoute,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.TCPRouteConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_http_route",
				httpRouteConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.HTTPRouteConfigEntry{
							Kind: api.HTTPRoute,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.HTTPRouteConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_rate_limit_ip",
				rateLimitIpConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.RateLimitIPConfigEntry{
							Kind: api.RateLimitIPConfig,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.RateLimitIPConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_inline_certificate",
				inlineCertificateConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.InlineCertificateConfigEntry{
							Kind: api.InlineCertificate,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.InlineCertificateConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_ingress_gateway",
				ingressGatewayConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.IngressGatewayConfigEntry{
							Kind: api.IngressGateway,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.IngressGatewayConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_terminating_gateway",
				terminatingGatewayConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.TerminatingGatewayConfigEntry{
							Kind: api.TerminatingGateway,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.TerminatingGatewayConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_api_gateway",
				apiGatewayConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.APIGatewayConfigEntry{
							Kind: api.APIGateway,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.APIGatewayConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_service_router",
				serviceRouterConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &structs.ServiceRouterConfigEntry{
							Entry: api.ServiceRouterConfigEntry{
								Kind: api.ServiceRouter,
							},
						}
						diags := models.Decode(ctx, getter, &entry)
						return &entry.Entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						result := &structs.ServiceRouterConfigEntry{
							ID:    fmt.Sprintf("%s/%s", entry.GetKind(), entry.GetName()),
							Entry: *entry.(*api.ServiceRouterConfigEntry),
						}
						return models.Set(ctx, setter, result)
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_service_splitter",
				serviceSplitterConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.ServiceSplitterConfigEntry{
							Kind: api.ServiceSplitter,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.ServiceSplitterConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_service_resolver",
				serviceResolverConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.ServiceResolverConfigEntry{
							Kind: api.ServiceResolver,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.ServiceResolverConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_mesh",
				meshConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.MeshConfigEntry{}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.MeshConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_exported_services",
				exportedServicesConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.ExportedServicesConfigEntry{}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.ExportedServicesConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_sameness_group",
				samenessGroupConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.SamenessGroupConfigEntry{
							Kind: api.SamenessGroup,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.SamenessGroupConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_service",
				serviceConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &structs.ServiceConfigEntry{
							Entry: api.ServiceConfigEntry{
								Kind: api.ServiceDefaults,
							},
						}
						diags := models.Decode(ctx, getter, &entry)
						return &entry.Entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						result := &structs.ServiceConfigEntry{
							ID:    fmt.Sprintf("%s/%s", entry.GetKind(), entry.GetName()),
							Entry: *entry.(*api.ServiceConfigEntry),
						}
						return models.Set(ctx, setter, result)
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_proxy",
				proxyConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &structs.ProxyConfigEntry{
							Entry: api.ProxyConfigEntry{
								Kind: api.ProxyDefaults,
							},
						}
						diags := models.Decode(ctx, getter, &entry)
						return &entry.Entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						result := &structs.ProxyConfigEntry{
							ID:    fmt.Sprintf("%s/%s", entry.GetKind(), entry.GetName()),
							Entry: *entry.(*api.ProxyConfigEntry),
						}
						return models.Set(ctx, setter, result)
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_jwt_provider",
				jwtProviderConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.JWTProviderConfigEntry{
							Kind: api.JWTProvider,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.JWTProviderConfigEntry))
					},
				},
			)
		},
		func() resource.Resource {
			return NewResource(
				"config_entry_service_intentions",
				serviceIntentionsConfigEntrySchema(),
				&ConfigEntry{
					decode: func(ctx context.Context, getter models.Getter) (api.ConfigEntry, diag.Diagnostics) {
						entry := &api.ServiceIntentionsConfigEntry{
							Kind: api.ServiceIntentions,
						}
						diags := models.Decode(ctx, getter, &entry)
						return entry, diags
					},
					set: func(ctx context.Context, setter models.Setter, entry api.ConfigEntry) diag.Diagnostics {
						return models.Set(ctx, setter, entry.(*api.ServiceIntentionsConfigEntry))
					},
				},
			)
		},
	}

	return configEntries
}

type ConfigEntry struct {
	decode func(context.Context, models.Getter) (api.ConfigEntry, diag.Diagnostics)
	set    func(context.Context, models.Setter, api.ConfigEntry) diag.Diagnostics
}

func (r *ConfigEntry) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	entry, diags := r.decode(ctx, req.Config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := client.ConfigEntries().Set(entry, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create config entry", err.Error())
		return
	}

	resp.Diagnostics.Append(r.set(ctx, &resp.State, entry)...)
}

func (r *ConfigEntry) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	entry, diags := r.decode(ctx, req.State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	entry, _, err := client.ConfigEntries().Get(entry.GetKind(), entry.GetName(), nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create config entry", err.Error())
		return
	}

	resp.Diagnostics.Append(r.set(ctx, &resp.State, entry)...)
}

func (r *ConfigEntry) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	entry, diags := r.decode(ctx, req.Config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := client.ConfigEntries().Set(entry, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update config entry", err.Error())
		return
	}

	resp.Diagnostics.Append(r.set(ctx, &resp.State, entry)...)
}

func (r *ConfigEntry) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	entry, diags := r.decode(ctx, req.State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.ConfigEntries().Delete(entry.GetKind(), entry.GetName(), nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete config entry", err.Error())
		return
	}
}

type GenericConfigEntry struct{}

func (r *GenericConfigEntry) Create(ctx context.Context, client *api.Client, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config *structs.ConfigEntry
	resp.Diagnostics.Append(models.Decode(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	raw := map[string]interface{}{}
	for k, v := range config.Config {
		raw[k] = v
	}
	raw["Kind"] = config.Kind
	raw["Name"] = config.Name
	raw["Namespace"] = config.Namespace
	raw["Partition"] = config.Partition
	raw["Meta"] = config.Meta

	entry, err := api.DecodeConfigEntry(raw)
	if err != nil {
		resp.Diagnostics.AddError("failed to decode config entry", err.Error())
		return
	}

	_, _, err = client.ConfigEntries().Set(entry, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create config entry", err.Error())
		return
	}

	config.ID = fmt.Sprintf("%s/%s", config.Kind, config.Name)
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *GenericConfigEntry) Read(ctx context.Context, client *api.Client, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *structs.ConfigEntry
	resp.Diagnostics.Append(models.DecodeConfigEntry(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	entry, _, err := client.ConfigEntries().Get(state.Kind, state.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ConfigEntry", err.Error())
		return
	}

	data, err := json.Marshal(entry)
	if err != nil {
		resp.Diagnostics.AddError("failed to marshal config entry", err.Error())
		return
	}

	m := map[string]interface{}{}
	if err := json.Unmarshal(data, &m); err != nil {
		resp.Diagnostics.AddError("failed to unmarshal config entry", err.Error())
		return
	}

	delete(m, "Kind")
	delete(m, "Name")
	delete(m, "Namespace")
	delete(m, "Partition")
	delete(m, "Meta")
	delete(m, "CreateIndex")
	delete(m, "ModifyIndex")

	state.Config = m
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, state)...)
}

func (r *GenericConfigEntry) Update(ctx context.Context, client *api.Client, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config *structs.ConfigEntry
	resp.Diagnostics.Append(models.DecodeConfigEntry(ctx, req.Config, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	entry, err := api.DecodeConfigEntry(config.Config)
	if err != nil {
		resp.Diagnostics.AddError("failed to decode config entry", err.Error())
		return
	}

	_, _, err = client.ConfigEntries().Set(entry, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to update config entry", err.Error())
		return
	}
	resp.Diagnostics.Append(models.Set(ctx, &resp.State, config)...)
}

func (r *GenericConfigEntry) Delete(ctx context.Context, client *api.Client, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *structs.ConfigEntry
	resp.Diagnostics.Append(models.DecodeConfigEntry(ctx, req.State, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.ConfigEntries().Delete(state.Kind, state.Name, nil)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete config entry", err.Error())
		return
	}
}
