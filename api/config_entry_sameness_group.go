// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package api

type SamenessGroupConfigEntry struct {
	Kind               string
	Name               string                `terraform:"name"`
	Partition          string                `json:",omitempty" terraform:"partition"`
	DefaultForFailover bool                  `json:",omitempty" alias:"default_for_failover" terraform:"default_for_failover"`
	IncludeLocal       bool                  `json:",omitempty" alias:"include_local" terraform:"include_local"`
	Members            []SamenessGroupMember `terraform:"members"`
	Meta               map[string]string     `json:",omitempty" terraform:"meta"`
	CreateIndex        uint64
	ModifyIndex        uint64
}

type SamenessGroupMember struct {
	Partition string `json:",omitempty" terraform:"partition"`
	Peer      string `json:",omitempty" terraform:"peer"`
}

func (s *SamenessGroupConfigEntry) GetKind() string            { return s.Kind }
func (s *SamenessGroupConfigEntry) GetName() string            { return s.Name }
func (s *SamenessGroupConfigEntry) GetPartition() string       { return s.Partition }
func (s *SamenessGroupConfigEntry) GetNamespace() string       { return "" }
func (s *SamenessGroupConfigEntry) GetCreateIndex() uint64     { return s.CreateIndex }
func (s *SamenessGroupConfigEntry) GetModifyIndex() uint64     { return s.ModifyIndex }
func (s *SamenessGroupConfigEntry) GetMeta() map[string]string { return s.Meta }
