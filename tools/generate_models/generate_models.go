// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	generator "github.com/Lenstra/terraform-plugin-generator"
	"github.com/dave/jennifer/jen"
	"github.com/hashicorp/consul/api"
	"github.com/remilapeyre/terraform-provider-consul-next/internal/structs"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

func getFieldInformation(path string, typ reflect.Type, field reflect.StructField) (*generator.FieldInformation, error) {
	info, err := generator.GetFieldInformationFromTerraformTag(path, typ, field)
	if info == nil || err != nil {
		return info, err
	}

	hcl, ok := field.Tag.Lookup("hcl")
	if ok {
		values := strings.Split(hcl, ",")
		hclName := values[0]
		if hclName != "" && hclName != info.Name {
			return nil, fmt.Errorf("the HCL and Terraform names are different: hcl:%q, terraform:%q", hclName, info.Name)
		}

		info.Block = slices.Contains(values, "block")
	}

	return info, nil
}

func datasourceFieldInformation(path string, typ reflect.Type, field reflect.StructField) (*generator.FieldInformation, error) {
	info, err := getFieldInformation(path, typ, field)
	if info == nil || err != nil {
		return info, err
	}

	optionals := []string{
		"acl_policy.id", "acl_policy.name",
		"acl_role.id", "acl_role.name",
	}
	required := []string{
		"acl_auth_method.name",
		"acl_binding_rule.id",
		"acl_token_secret_id.accessor_id",
		"acl_token.accessor_id",
		"namespace.name",
		"node.id",
		"peering.name",
	}

	path += "." + info.Name
	info.Optional = slices.Contains(optionals, path)
	info.Required = slices.Contains(required, path)
	info.Computed = !info.Required

	return info, nil
}

func resourceFieldInformation() (generator.FieldInformationGetter, error) {
	data, err := os.ReadFile("./docs/resources.yaml")
	if err != nil {
		return nil, err
	}
	var docs map[string]map[string]map[string]interface{}
	err = yaml.Unmarshal(data, &docs)
	if err != nil {
		return nil, err
	}

	return func(path string, typ reflect.Type, field reflect.StructField) (*generator.FieldInformation, error) {
		info, err := getFieldInformation(path, typ, field)
		if info == nil || err != nil {
			return info, err
		}

		path += "." + info.Name
		parts := strings.SplitN(path, ".", 2)
		name := parts[0]
		path = parts[1]

		if typ == reflect.TypeOf(api.ACLLink{}) {
			other := "id"
			if info.Name == "id" {
				other = "name"
			}
			info.Validators = jen.Index().Qual("github.com/hashicorp/terraform-plugin-framework/schema/validator", "String").Values(
				jen.Id("NewMutuallyExclusiveValidator").Call(jen.Lit(other)),
			)
		}

		switch default_ := docs[name][path]["default"].(type) {
		case bool:
			info.Default = jen.Qual("github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault", "StaticBool").Call(jen.Lit(default_))
			info.Computed = true
			info.Description += fmt.Sprintf("This defaults to `%v`", default_)
		case string:
			info.Default = jen.Qual("github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault", "StaticString").Call(jen.Lit(default_))
			info.Computed = true
			info.Description += fmt.Sprintf("This defaults to `%q`", default_)
		case int:
			info.Default = jen.Qual("github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default", "StaticInt64").Call(jen.Lit(default_))
			info.Computed = true
			info.Description += fmt.Sprintf("This defaults to %v", default_)
		case nil:
			break
		default:
			return nil, fmt.Errorf("unknown type %T", default_)
		}

		if description, ok := docs[name][path]["description"].(string); ok {
			info.Description = description
		}

		return info, nil
	}, nil
}

func providerFieldInformation() (generator.FieldInformationGetter, error) {
	data, err := os.ReadFile("./docs/provider.yaml")
	if err != nil {
		return nil, err
	}
	var docs map[string]map[string]interface{}
	err = yaml.Unmarshal(data, &docs)
	if err != nil {
		return nil, err
	}

	return func(path string, typ reflect.Type, field reflect.StructField) (*generator.FieldInformation, error) {
		info, err := getFieldInformation(path, typ, field)
		if info == nil || err != nil {
			return info, err
		}

		path += "." + info.Name
		parts := strings.SplitN(path, ".", 2)
		name := parts[0]
		path = parts[1]

		docPath, ok := docs[name][path].(map[string]interface{})
		if !ok {
			docPath = map[string]interface{}{}
		}
		if description, ok := docPath["description"].(string); ok {
			info.Description = description
		}

		return info, nil
	}, nil
}

func main() {
	objects := map[string]interface{}{
		"ACLAuthMethod":             structs.ACLAuthMethod{},
		"ACLBindingRule":            api.ACLBindingRule{},
		"ACLPolicy":                 api.ACLPolicy{},
		"ACLRole":                   api.ACLRole{},
		"ACLToken":                  structs.ACLToken{},
		"ACLTokenPolicyAttachment":  structs.ACLTokenPolicyAttachment{},
		"ACLTokenRoleAttachment":    structs.ACLTokenRoleAttachment{},
		"ACLTokenSecretID":          structs.ACLTokenSecretID{},
		"AgentConfig":               structs.AgentConfig{},
		"Area":                      api.Area{},
		"AutopilotConfig":           structs.AutopilotConfig{},
		"AutopilotHealth":           structs.AutopilotHealth{},
		"CAConfig":                  structs.CAConfig{},
		"CatalogNode":               api.CatalogNode{},
		"CatalogService":            api.CatalogService{},
		"Config":                    api.Config{},
		"Datacenters":               structs.Datacenters{},
		"KeyPrefix":                 structs.KeyPrefix{},
		"Keys":                      structs.Keys{},
		"Namespace":                 api.Namespace{},
		"NetworkAreaMembers":        structs.NetworkAreaMembers{},
		"NetworkSegments":           structs.NetworkSegments{},
		"Nodes":                     structs.Nodes{},
		"OperatorHealthReply":       api.OperatorHealthReply{},
		"Partition":                 api.Partition{},
		"Peering":                   api.Peering{},
		"NamespacePolicyAttachment": structs.NamespacePolicyAttachment{},
		"NamespaceRoleAttachment":   structs.NamespaceRoleAttachment{},
		"PeeringResource":           structs.PeeringResource{},
		"Peerings":                  structs.Peerings{},
		"PeeringToken":              structs.PeeringToken{},
		"PreparedQueryDefinition":   api.PreparedQueryDefinition{},
		"ServiceHealth":             structs.ServiceHealth{},

		// Config entries
		"ConfigEntry":                   structs.ConfigEntry{},
		"TCPRouteConfigEntry":           api.TCPRouteConfigEntry{},
		"HTTPRouteConfigEntry":          api.HTTPRouteConfigEntry{},
		"RateLimitIPConfigEntry":        api.RateLimitIPConfigEntry{},
		"InlineCertificateConfigEntry":  api.InlineCertificateConfigEntry{},
		"IngressGatewayConfigEntry":     api.IngressGatewayConfigEntry{},
		"TerminatingGatewayConfigEntry": api.TerminatingGatewayConfigEntry{},
		"APIGatewayConfigEntry":         api.APIGatewayConfigEntry{},
		"ServiceRouterConfigEntry":      structs.ServiceRouterConfigEntry{},
		"ServiceSplitterConfigEntry":    api.ServiceSplitterConfigEntry{},
		"ServiceResolverConfigEntry":    api.ServiceResolverConfigEntry{},
		"MeshConfigEntry":               api.MeshConfigEntry{},
		"ExportedServicesConfigEntry":   api.ExportedServicesConfigEntry{},
		"SamenessGroupConfigEntry":      api.SamenessGroupConfigEntry{},
		"ServiceConfigEntry":            structs.ServiceConfigEntry{},
		"ProxyConfigEntry":              structs.ProxyConfigEntry{},
		"JWTProviderConfigEntry":        api.JWTProviderConfigEntry{},
		"ServiceIntentionsConfigEntry":  api.ServiceIntentionsConfigEntry{},
	}
	err := generator.GenerateModels("./internal/models", "models", objects, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	objects = map[string]interface{}{
		"acl_auth_method":       structs.ACLAuthMethod{},
		"acl_binding_rule":      api.ACLBindingRule{},
		"acl_policy":            api.ACLPolicy{},
		"acl_role":              api.ACLRole{},
		"acl_token_secret_id":   structs.ACLTokenSecretID{},
		"acl_token":             structs.ACLToken{},
		"admin_partition":       api.Partition{},
		"agent_config":          structs.AgentConfig{},
		"autopilot_health":      structs.AutopilotHealth{},
		"catalog_service":       api.CatalogService{},
		"certificate_authority": structs.CAConfig{},
		"config_entry":          structs.ConfigEntry{},
		"datacenters":           structs.Datacenters{},
		"key_prefix":            structs.KeyPrefix{},
		"keys":                  structs.Keys{},
		"namespace":             api.Namespace{},
		"network_area_members":  structs.NetworkAreaMembers{},
		"network_segments":      structs.NetworkSegments{},
		"node":                  api.CatalogNode{},
		"nodes":                 structs.Nodes{},
		"peering":               api.Peering{},
		"peerings":              structs.Peerings{},
		"service_health":        structs.ServiceHealth{},
	}
	err = generator.GenerateSchema(generator.DataSourceSchema, "./internal/datasource", "datasource", objects, &generator.GeneratorOptions{
		GetFieldInformation: datasourceFieldInformation,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	objects = map[string]interface{}{
		"acl_auth_method":             structs.ACLAuthMethod{},
		"acl_binding_rule":            api.ACLBindingRule{},
		"acl_policy":                  api.ACLPolicy{},
		"acl_role":                    api.ACLRole{},
		"acl_token":                   structs.ACLToken{},
		"acl_token_policy_attachment": structs.ACLTokenPolicyAttachment{},
		"acl_token_role_attachment":   structs.ACLTokenRoleAttachment{},
		"admin_partition":             api.Partition{},
		"autopilot_config":            structs.AutopilotConfig{},
		"catalog_service":             api.CatalogService{},
		"certificate_authority":       structs.CAConfig{},
		"config_entry":                structs.ConfigEntry{},
		"namespace":                   api.Namespace{},
		"node":                        api.CatalogNode{},
		"peering_token":               structs.PeeringToken{},
		"peering":                     structs.PeeringResource{},
		"namespace_policy_attachment": structs.NamespacePolicyAttachment{},
		"namespace_role_attachment":   structs.NamespaceRoleAttachment{},
		"area":                        api.Area{},
		"prepared_query_definition":   api.PreparedQueryDefinition{},

		// Config entries
		"tcp_route_config_entry":           api.TCPRouteConfigEntry{},
		"http_route_config_entry":          api.HTTPRouteConfigEntry{},
		"rate_limit_ip_config_entry":       api.RateLimitIPConfigEntry{},
		"inline_certificate_config_entry":  api.InlineCertificateConfigEntry{},
		"ingress_gateway_config_entry":     api.IngressGatewayConfigEntry{},
		"terminating_gateway_config_entry": api.TerminatingGatewayConfigEntry{},
		"api_gateway_config_entry":         api.APIGatewayConfigEntry{},
		"service_router_config_entry":      structs.ServiceRouterConfigEntry{},
		"service_splitter_config_entry":    api.ServiceSplitterConfigEntry{},
		"service_resolver_config_entry":    api.ServiceResolverConfigEntry{},
		"mesh_config_entry":                api.MeshConfigEntry{},
		"exported_services_config_entry":   api.ExportedServicesConfigEntry{},
		"sameness_group_config_entry":      api.SamenessGroupConfigEntry{},
		"service_config_entry":             structs.ServiceConfigEntry{},
		"proxy_config_entry":               structs.ProxyConfigEntry{},
		"jwt_provider_config_entry":        api.JWTProviderConfigEntry{},
		"service_intentions_config_entry":  api.ServiceIntentionsConfigEntry{},
	}
	f, err := resourceFieldInformation()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = generator.GenerateSchema(generator.ResourceSchema, "./internal/resource", "resource", objects, &generator.GeneratorOptions{
		GetFieldInformation: f,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	objects = map[string]interface{}{
		"config": api.Config{},
	}
	f, err = providerFieldInformation()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = generator.GenerateSchema(generator.ProviderSchema, "./internal/provider", "provider", objects, &generator.GeneratorOptions{
		GetFieldInformation: f,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
