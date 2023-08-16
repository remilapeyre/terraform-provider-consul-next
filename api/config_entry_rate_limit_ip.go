// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package api

type ReadWriteRatesConfig struct {
	ReadRate  float64 `terraform:"read_rate"`
	WriteRate float64 `terraform:"write_rate"`
}

type RateLimitIPConfigEntry struct {
	// Kind of the config entry. This will be set to structs.RateLimitIPConfig
	Kind string
	Name string `terraform:"name"`
	Mode string `terraform:"mode"` // {permissive, enforcing, disabled}

	Meta map[string]string `json:",omitempty" terraform:"meta"`
	// overall limits
	ReadRate  float64 `terraform:"read_rate"`
	WriteRate float64 `terraform:"write_rate"`

	//limits specific to a type of call
	ACL             *ReadWriteRatesConfig `json:",omitempty" terraform:"acl"`              //	OperationCategoryACL             OperationCategory = "ACL"
	Catalog         *ReadWriteRatesConfig `json:",omitempty" terraform:"catalog"`          //   OperationCategoryCatalog         OperationCategory = "Catalog"
	ConfigEntry     *ReadWriteRatesConfig `json:",omitempty" terraform:"config_entry"`     //   OperationCategoryConfigEntry     OperationCategory = "ConfigEntry"
	ConnectCA       *ReadWriteRatesConfig `json:",omitempty" terraform:"connect_ca"`       //   OperationCategoryConnectCA       OperationCategory = "ConnectCA"
	Coordinate      *ReadWriteRatesConfig `json:",omitempty" terraform:"coordinate"`       //   OperationCategoryCoordinate      OperationCategory = "Coordinate"
	DiscoveryChain  *ReadWriteRatesConfig `json:",omitempty" terraform:"discovery_chain"`  //   OperationCategoryDiscoveryChain  OperationCategory = "DiscoveryChain"
	ServerDiscovery *ReadWriteRatesConfig `json:",omitempty" terraform:"server_discovery"` //  OperationCategoryServerDiscovery OperationCategory = "ServerDiscovery"
	Health          *ReadWriteRatesConfig `json:",omitempty" terraform:"health"`           //  OperationCategoryHealth          OperationCategory = "Health"
	Intention       *ReadWriteRatesConfig `json:",omitempty" terraform:"intention"`        //  OperationCategoryIntention       OperationCategory = "Intention"
	KV              *ReadWriteRatesConfig `json:",omitempty" terraform:"kv"`               //  OperationCategoryKV              OperationCategory = "KV"
	Tenancy         *ReadWriteRatesConfig `json:",omitempty" terraform:"tenancy"`          //  OperationCategoryPartition        OperationCategory = "Tenancy"
	PreparedQuery   *ReadWriteRatesConfig `json:",omitempty" terraform:"prepared_query"`   //  OperationCategoryPreparedQuery   OperationCategory = "PreparedQuery"
	Session         *ReadWriteRatesConfig `json:",omitempty" terraform:"session"`          //  OperationCategorySession         OperationCategory = "Session"
	Txn             *ReadWriteRatesConfig `json:",omitempty" terraform:"txn"`              //  OperationCategoryTxn             OperationCategory = "Txn"
	AutoConfig      *ReadWriteRatesConfig `json:",omitempty" terraform:"auto_config"`      //  OperationCategoryAutoConfig      OperationCategory = "AutoConfig"
	FederationState *ReadWriteRatesConfig `json:",omitempty" terraform:"federation_state"` //  OperationCategoryFederationState OperationCategory = "FederationState"
	Internal        *ReadWriteRatesConfig `json:",omitempty" terraform:"internal"`         //  OperationCategoryInternal        OperationCategory = "Internal"
	PeerStream      *ReadWriteRatesConfig `json:",omitempty" terraform:"peer_stream"`      //  OperationCategoryPeerStream      OperationCategory = "PeerStream"
	Peering         *ReadWriteRatesConfig `json:",omitempty" terraform:"peering"`          //  OperationCategoryPeering         OperationCategory = "Peering"
	DataPlane       *ReadWriteRatesConfig `json:",omitempty" terraform:"data_plane"`       //  OperationCategoryDataPlane       OperationCategory = "DataPlane"
	DNS             *ReadWriteRatesConfig `json:",omitempty" terraform:"dns"`              //  OperationCategoryDNS             OperationCategory = "DNS"
	Subscribe       *ReadWriteRatesConfig `json:",omitempty" terraform:"subscribe"`        //  OperationCategorySubscribe       OperationCategory = "Subscribe"
	Resource        *ReadWriteRatesConfig `json:",omitempty" terraform:"resource"`         //  OperationCategoryResource        OperationCategory = "Resource"

	// Partition is the partition the config entry is associated with.
	// Partitioning is a Consul Enterprise feature.
	Partition string `json:",omitempty" terraform:"partition"`

	// Namespace is the namespace the config entry is associated with.
	// Namespacing is a Consul Enterprise feature.
	Namespace string `json:",omitempty" terraform:"namespace"`

	// CreateIndex is the Raft index this entry was created at. This is a
	// read-only field.
	CreateIndex uint64

	// ModifyIndex is used for the Check-And-Set operations and can also be fed
	// back into the WaitIndex of the QueryOptions in order to perform blocking
	// queries.
	ModifyIndex uint64
}

func (r *RateLimitIPConfigEntry) GetKind() string {
	return RateLimitIPConfig
}
func (r *RateLimitIPConfigEntry) GetName() string {
	if r == nil {
		return ""
	}
	return r.Name
}
func (r *RateLimitIPConfigEntry) GetPartition() string {
	return r.Partition
}
func (r *RateLimitIPConfigEntry) GetNamespace() string {
	return r.Namespace
}
func (r *RateLimitIPConfigEntry) GetMeta() map[string]string {
	if r == nil {
		return nil
	}
	return r.Meta
}
func (r *RateLimitIPConfigEntry) GetCreateIndex() uint64 {
	return r.CreateIndex
}
func (r *RateLimitIPConfigEntry) GetModifyIndex() uint64 {
	return r.ModifyIndex
}
