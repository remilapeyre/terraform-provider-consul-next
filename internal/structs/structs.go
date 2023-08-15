// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package structs

import "github.com/hashicorp/consul/api"

// The structs package wraps the types that could not be used directly from the
// Consul SDK, most of the time because they had no ID attribute.

type ACLAuthMethod struct {
	ID     string            `terraform:"id,computed"`
	Method api.ACLAuthMethod `terraform:"-,promoted"`
}

type ACLToken struct {
	ID    string       `terraform:"id,computed"`
	Token api.ACLToken `terraform:"-,promoted"`
}

type ACLTokenPolicyAttachment struct {
	ID      string      `terraform:"id,computed"`
	TokenID string      `terraform:"token_id,required"`
	Policy  api.ACLLink `terraform:"policy,required"`
}

type ACLTokenRoleAttachment struct {
	ID      string      `terraform:"id,computed"`
	TokenID string      `terraform:"token_id,required"`
	Role    api.ACLLink `terraform:"role,required"`
}

type ACLTokenSecretID struct {
	ID         string `terraform:"id"`
	AccessorID string `terraform:"accessor_id"`
	SecretID   string `terraform:"secret_id"`
	Partition  string `terraform:"partition"`
	Namespace  string `terraform:"namespace"`
	PGPKey     string `terraform:"pgp_key"`
}

type AutopilotConfig struct {
	ID     string                     `terraform:"id,computed"`
	Config api.AutopilotConfiguration `terraform:"-,promoted"`
}

type AutopilotHealth struct {
	ID              string                  `terraform:"id"`
	AutopilotHealth api.OperatorHealthReply `terraform:"-,promoted"`
}

type AgentConfig struct {
	ID                string  `terraform:"id"`
	Datacenter        string  `terraform:"datacenter"`
	PrimaryDatacenter string  `terraform:"primary_datacenter"`
	NodeName          string  `terraform:"node_name"`
	NodeID            string  `terraform:"node_id"`
	Partition         *string `terraform:"partition"`
	Revision          string  `terraform:"revision"`
	Server            bool    `terraform:"server"`
	Version           string  `terraform:"version"`
	BuildDate         string  `terraform:"build_date"`
}

type CAConfig struct {
	ID     string       `terraform:"id,computed"`
	Config api.CAConfig `terraform:"-,promoted"`
}

type ConfigEntry struct {
	ID        string                 `terraform:"id,computed"`
	Kind      string                 `terraform:"kind"`
	Name      string                 `terraform:"name"`
	Namespace string                 `terraform:"namespace"`
	Partition string                 `terraform:"partition"`
	Config    map[string]interface{} `terraform:"config"`
	Meta      map[string]string      `terraform:"meta"`
}

type Datacenters struct {
	ID          string   `terraform:"id"`
	Datacenters []string `terraform:"datacenters"`
}

type Keys struct {
	ID   string        `terraform:"id"`
	Keys []*api.KVPair `terraform:"keys"`
}

type KeyPrefix struct {
	ID string `terraform:"key_prefix"`
}

type NamespacePolicyAttachment struct {
	ID        string      `terraform:"id,computed"`
	Namespace string      `terraform:"namespace,required"`
	Policy    api.ACLLink `terraform:"policy,required"`
}

type NamespaceRoleAttachment struct {
	ID        string      `terraform:"id,computed"`
	Namespace string      `terraform:"namespace,required"`
	Role      api.ACLLink `terraform:"role,required"`
}

type NetworkAreaMembers struct {
	ID      string            `terraform:"id"`
	Members []*api.SerfMember `terraform:"members"`
}

type NetworkSegments struct {
	ID       string   `terraform:"id"`
	Segments []string `terraform:"segments"`
}

type Nodes struct {
	ID    string      `terraform:"id"`
	Nodes []*api.Node `terraform:"nodes"`
}

type PeeringResource struct {
	Peering      api.Peering `terraform:"-,promoted"`
	PeeringToken string      `terraform:"peering_token,sensitive"`
}

type Peerings struct {
	ID       string         `terraform:"id"`
	Peerings []*api.Peering `terraform:"peerings"`
}

type PeeringToken struct {
	ID       string                           `terraform:"id,computed"`
	Request  api.PeeringGenerateTokenRequest  `terraform:"-,promoted"`
	Response api.PeeringGenerateTokenResponse `terraform:"-,promoted"`
}

type ServiceHealth struct {
	ID       string              `terraform:"id"`
	Services []*api.ServiceEntry `terraform:"services"`
}
