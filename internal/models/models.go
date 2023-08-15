/*
Code generated by github-terraform-generator; DO NOT EDIT.
Any modifications will be overwritten
*/

package models

import types "github.com/hashicorp/terraform-plugin-framework/types"

type ACLAuthMethod struct {
	ID             types.String                  `tfsdk:"id"`
	Name           types.String                  `tfsdk:"name"`
	Type           types.String                  `tfsdk:"type"`
	DisplayName    types.String                  `tfsdk:"display_name"`
	Description    types.String                  `tfsdk:"description"`
	MaxTokenTTL    types.String                  `tfsdk:"max_token_ttl"`
	TokenLocality  types.String                  `tfsdk:"token_locality"`
	Config         types.String                  `tfsdk:"config"`
	NamespaceRules []*ACLAuthMethodNamespaceRule `tfsdk:"namespace_rules"`
	Namespace      types.String                  `tfsdk:"namespace"`
	Partition      types.String                  `tfsdk:"partition"`
}

type ACLBindingRule struct {
	ID          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
	AuthMethod  types.String `tfsdk:"auth_method"`
	Selector    types.String `tfsdk:"selector"`
	BindType    types.String `tfsdk:"bind_type"`
	BindName    types.String `tfsdk:"bind_name"`
	Namespace   types.String `tfsdk:"namespace"`
	Partition   types.String `tfsdk:"partition"`
}

type ACLPolicy struct {
	ID          types.String   `tfsdk:"id"`
	Name        types.String   `tfsdk:"name"`
	Description types.String   `tfsdk:"description"`
	Rules       types.String   `tfsdk:"rules"`
	Datacenters []types.String `tfsdk:"datacenters"`
	Namespace   types.String   `tfsdk:"namespace"`
	Partition   types.String   `tfsdk:"partition"`
}

type ACLRole struct {
	ID                types.String          `tfsdk:"id"`
	Name              types.String          `tfsdk:"name"`
	Description       types.String          `tfsdk:"description"`
	Policies          []*ACLLink            `tfsdk:"policies"`
	ServiceIdentities []*ACLServiceIdentity `tfsdk:"service_identities"`
	NodeIdentities    []*ACLNodeIdentity    `tfsdk:"node_identities"`
	Namespace         types.String          `tfsdk:"namespace"`
	Partition         types.String          `tfsdk:"partition"`
}

type ACLToken struct {
	ID                  types.String          `tfsdk:"id"`
	AccessorID          types.String          `tfsdk:"accessor_id"`
	Description         types.String          `tfsdk:"description"`
	Policies            []*ACLLink            `tfsdk:"policies"`
	Roles               []*ACLLink            `tfsdk:"roles"`
	ServiceIdentities   []*ACLServiceIdentity `tfsdk:"service_identities"`
	NodeIdentities      []*ACLNodeIdentity    `tfsdk:"node_identities"`
	Local               types.Bool            `tfsdk:"local"`
	AuthMethod          types.String          `tfsdk:"auth_method"`
	ExpirationTTL       types.String          `tfsdk:"expiration_ttl"`
	ExpirationTime      types.String          `tfsdk:"expiration_time"`
	CreateTime          types.String          `tfsdk:"create_time"`
	Namespace           types.String          `tfsdk:"namespace"`
	Partition           types.String          `tfsdk:"partition"`
	AuthMethodNamespace types.String          `tfsdk:"auth_method_namespace"`
}

type ACLTokenPolicyAttachment struct {
	ID      types.String `tfsdk:"id"`
	TokenID types.String `tfsdk:"token_id"`
	Policy  *ACLLink     `tfsdk:"policy"`
}

type ACLTokenRoleAttachment struct {
	ID      types.String `tfsdk:"id"`
	TokenID types.String `tfsdk:"token_id"`
	Role    *ACLLink     `tfsdk:"role"`
}

type ACLTokenSecretID struct {
	ID         types.String `tfsdk:"id"`
	AccessorID types.String `tfsdk:"accessor_id"`
	SecretID   types.String `tfsdk:"secret_id"`
	Partition  types.String `tfsdk:"partition"`
	Namespace  types.String `tfsdk:"namespace"`
	PGPKey     types.String `tfsdk:"pgp_key"`
}

type AgentConfig struct {
	ID                types.String `tfsdk:"id"`
	Datacenter        types.String `tfsdk:"datacenter"`
	PrimaryDatacenter types.String `tfsdk:"primary_datacenter"`
	NodeName          types.String `tfsdk:"node_name"`
	NodeID            types.String `tfsdk:"node_id"`
	Partition         types.String `tfsdk:"partition"`
	Revision          types.String `tfsdk:"revision"`
	Server            types.Bool   `tfsdk:"server"`
	Version           types.String `tfsdk:"version"`
	BuildDate         types.String `tfsdk:"build_date"`
}

type Area struct {
	ID             types.String   `tfsdk:"id"`
	PeerDatacenter types.String   `tfsdk:"peer_datacenter"`
	RetryJoin      []types.String `tfsdk:"retry_join"`
	UseTLS         types.Bool     `tfsdk:"use_tls"`
}

type AutopilotConfig struct {
	ID                      types.String `tfsdk:"id"`
	CleanupDeadServers      types.Bool   `tfsdk:"cleanup_dead_servers"`
	MaxTrailingLogs         types.Int64  `tfsdk:"max_trailing_logs"`
	MinQuorum               types.Int64  `tfsdk:"min_quorum"`
	RedundancyZoneTag       types.String `tfsdk:"redundancy_zone_tag"`
	DisableUpgradeMigration types.Bool   `tfsdk:"disable_upgrade_migration"`
	UpgradeVersionTag       types.String `tfsdk:"upgrade_version_tag"`
}

type AutopilotHealth struct {
	ID               types.String    `tfsdk:"id"`
	Healthy          types.Bool      `tfsdk:"healthy"`
	FailureTolerance types.Int64     `tfsdk:"failure_tolerance"`
	Servers          []*ServerHealth `tfsdk:"servers"`
}

type CAConfig struct {
	ID                       types.String            `tfsdk:"id"`
	Provider                 types.String            `tfsdk:"connect_provider"`
	Config                   types.String            `tfsdk:"config"`
	State                    map[string]types.String `tfsdk:"state"`
	ForceWithoutCrossSigning types.Bool              `tfsdk:"force_without_cross_signing"`
}

type CatalogNode struct {
	Services        map[string]*AgentService `tfsdk:"services"`
	ID              types.String             `tfsdk:"id"`
	Node            types.String             `tfsdk:"node"`
	Address         types.String             `tfsdk:"address"`
	Datacenter      types.String             `tfsdk:"datacenter"`
	TaggedAddresses map[string]types.String  `tfsdk:"tagged_addresses"`
	Meta            map[string]types.String  `tfsdk:"meta"`
	Partition       types.String             `tfsdk:"partition"`
	PeerName        types.String             `tfsdk:"peer_name"`
	Locality        *Locality                `tfsdk:"locality"`
}

type CatalogService struct {
	ID                       types.String                    `tfsdk:"id"`
	Node                     types.String                    `tfsdk:"node"`
	Address                  types.String                    `tfsdk:"address"`
	Datacenter               types.String                    `tfsdk:"datacenter"`
	TaggedAddresses          map[string]types.String         `tfsdk:"tagged_addresses"`
	NodeMeta                 map[string]types.String         `tfsdk:"node_meta"`
	ServiceID                types.String                    `tfsdk:"service_id"`
	ServiceName              types.String                    `tfsdk:"service_name"`
	ServiceAddress           types.String                    `tfsdk:"service_address"`
	ServiceTaggedAddresses   map[string]*ServiceAddress      `tfsdk:"service_tagged_addresses"`
	ServiceTags              []types.String                  `tfsdk:"service_tags"`
	ServiceMeta              map[string]types.String         `tfsdk:"service_meta"`
	ServicePort              types.Int64                     `tfsdk:"service_port"`
	ServiceWeights           *Weights                        `tfsdk:"service_weights"`
	ServiceEnableTagOverride types.Bool                      `tfsdk:"service_enable_tag_override"`
	ServiceProxy             *AgentServiceConnectProxyConfig `tfsdk:"service_proxy"`
	ServiceLocality          *Locality                       `tfsdk:"service_locality"`
	CreateIndex              types.Int64                     `tfsdk:"create_index"`
	Checks                   []*HealthCheck                  `tfsdk:"checks"`
	ModifyIndex              types.Int64                     `tfsdk:"modify_index"`
	Namespace                types.String                    `tfsdk:"namespace"`
	Partition                types.String                    `tfsdk:"partition"`
}

type Config struct {
	Address    types.String   `tfsdk:"address"`
	Scheme     types.String   `tfsdk:"scheme"`
	PathPrefix types.String   `tfsdk:"path_prefix"`
	Datacenter types.String   `tfsdk:"datacenter"`
	HttpAuth   *HttpBasicAuth `tfsdk:"http_auth"`
	Token      types.String   `tfsdk:"token"`
	TokenFile  types.String   `tfsdk:"token_file"`
	Namespace  types.String   `tfsdk:"namespace"`
	Partition  types.String   `tfsdk:"partition"`
	TLSConfig  *TLSConfig     `tfsdk:"tls_config"`
}

type ConfigEntry struct {
	ID        types.String            `tfsdk:"id"`
	Kind      types.String            `tfsdk:"kind"`
	Name      types.String            `tfsdk:"name"`
	Namespace types.String            `tfsdk:"namespace"`
	Partition types.String            `tfsdk:"partition"`
	Config    types.String            `tfsdk:"config"`
	Meta      map[string]types.String `tfsdk:"meta"`
}

type Datacenters struct {
	ID          types.String   `tfsdk:"id"`
	Datacenters []types.String `tfsdk:"datacenters"`
}

type KeyPrefix struct {
	ID types.String `tfsdk:"key_prefix"`
}

type Keys struct {
	ID   types.String `tfsdk:"id"`
	Keys []*KVPair    `tfsdk:"keys"`
}

type Namespace struct {
	Name        types.String            `tfsdk:"name"`
	Description types.String            `tfsdk:"description"`
	ACLs        *NamespaceACLConfig     `tfsdk:"ac_ls"`
	Meta        map[string]types.String `tfsdk:"meta"`
	DeletedAt   types.String            `tfsdk:"deleted_at"`
	Partition   types.String            `tfsdk:"partition"`
}

type NamespacePolicyAttachment struct {
	ID        types.String `tfsdk:"id"`
	Namespace types.String `tfsdk:"namespace"`
	Policy    *ACLLink     `tfsdk:"policy"`
}

type NamespaceRoleAttachment struct {
	ID        types.String `tfsdk:"id"`
	Namespace types.String `tfsdk:"namespace"`
	Role      *ACLLink     `tfsdk:"role"`
}

type NetworkAreaMembers struct {
	ID      types.String  `tfsdk:"id"`
	Members []*SerfMember `tfsdk:"members"`
}

type NetworkSegments struct {
	ID       types.String   `tfsdk:"id"`
	Segments []types.String `tfsdk:"segments"`
}

type Nodes struct {
	ID    types.String `tfsdk:"id"`
	Nodes []*Node      `tfsdk:"nodes"`
}

type OperatorHealthReply struct {
	Healthy          types.Bool      `tfsdk:"healthy"`
	FailureTolerance types.Int64     `tfsdk:"failure_tolerance"`
	Servers          []*ServerHealth `tfsdk:"servers"`
}

type Partition struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	DeletedAt   types.String `tfsdk:"deleted_at"`
}

type Peering struct {
	ID                  types.String            `tfsdk:"id"`
	Name                types.String            `tfsdk:"name"`
	Partition           types.String            `tfsdk:"partition"`
	DeletedAt           types.String            `tfsdk:"deleted_at"`
	Meta                map[string]types.String `tfsdk:"meta"`
	State               types.String            `tfsdk:"state"`
	PeerID              types.String            `tfsdk:"peer_id"`
	PeerCAPems          []types.String          `tfsdk:"peer_ca_pems"`
	PeerServerName      types.String            `tfsdk:"peer_server_name"`
	PeerServerAddresses []types.String          `tfsdk:"peer_server_addresses"`
	StreamStatus        *PeeringStreamStatus    `tfsdk:"stream_status"`
	Remote              *PeeringRemoteInfo      `tfsdk:"remote"`
}

type PeeringResource struct {
	PeeringToken        types.String            `tfsdk:"peering_token"`
	ID                  types.String            `tfsdk:"id"`
	Name                types.String            `tfsdk:"name"`
	Partition           types.String            `tfsdk:"partition"`
	DeletedAt           types.String            `tfsdk:"deleted_at"`
	Meta                map[string]types.String `tfsdk:"meta"`
	State               types.String            `tfsdk:"state"`
	PeerID              types.String            `tfsdk:"peer_id"`
	PeerCAPems          []types.String          `tfsdk:"peer_ca_pems"`
	PeerServerName      types.String            `tfsdk:"peer_server_name"`
	PeerServerAddresses []types.String          `tfsdk:"peer_server_addresses"`
	StreamStatus        *PeeringStreamStatus    `tfsdk:"stream_status"`
	Remote              *PeeringRemoteInfo      `tfsdk:"remote"`
}

type PeeringToken struct {
	ID                      types.String            `tfsdk:"id"`
	PeerName                types.String            `tfsdk:"peer_name"`
	Partition               types.String            `tfsdk:"partition"`
	Meta                    map[string]types.String `tfsdk:"meta"`
	ServerExternalAddresses []types.String          `tfsdk:"server_external_addresses"`
	PeeringToken            types.String            `tfsdk:"peering_token"`
}

type Peerings struct {
	ID       types.String `tfsdk:"id"`
	Peerings []*Peering   `tfsdk:"peerings"`
}

type PreparedQueryDefinition struct {
	ID       types.String     `tfsdk:"id"`
	Name     types.String     `tfsdk:"name"`
	Session  types.String     `tfsdk:"session"`
	Token    types.String     `tfsdk:"token"`
	Service  *ServiceQuery    `tfsdk:"service"`
	DNS      *QueryDNSOptions `tfsdk:"dns"`
	Template *QueryTemplate   `tfsdk:"template"`
}

type ServiceHealth struct {
	ID       types.String    `tfsdk:"id"`
	Services []*ServiceEntry `tfsdk:"services"`
}

type ACLAuthMethodNamespaceRule struct {
	Selector      types.String `tfsdk:"selector"`
	BindNamespace types.String `tfsdk:"bind_namespace"`
}

type ACLLink struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ACLServiceIdentity struct {
	ServiceName types.String   `tfsdk:"service_name"`
	Datacenters []types.String `tfsdk:"datacenters"`
}

type ACLNodeIdentity struct {
	NodeName   types.String `tfsdk:"node_name"`
	Datacenter types.String `tfsdk:"datacenter"`
}

type ServerHealth struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Address     types.String `tfsdk:"address"`
	SerfStatus  types.String `tfsdk:"serf_status"`
	Version     types.String `tfsdk:"version"`
	Leader      types.Bool   `tfsdk:"leader"`
	LastTerm    types.Int64  `tfsdk:"last_term"`
	LastIndex   types.Int64  `tfsdk:"last_index"`
	Healthy     types.Bool   `tfsdk:"healthy"`
	Voter       types.Bool   `tfsdk:"voter"`
	StableSince types.String `tfsdk:"stable_since"`
}

type AgentService struct {
	ID types.String `tfsdk:"id"`
}

type Locality struct {
	Region types.String `tfsdk:"region"`
	Zone   types.String `tfsdk:"zone"`
}

type ServiceAddress struct {
	Address types.String `tfsdk:"address"`
	Port    types.Int64  `tfsdk:"port"`
}

type Weights struct {
	Passing types.Int64 `tfsdk:"passing"`
	Warning types.Int64 `tfsdk:"warning"`
}

type AgentServiceConnectProxyConfig struct {
	EnvoyExtensions        []*EnvoyExtension       `tfsdk:"envoy_extensions"`
	DestinationServiceName types.String            `tfsdk:"destination_service_name"`
	DestinationServiceID   types.String            `tfsdk:"destination_service_id"`
	LocalServiceAddress    types.String            `tfsdk:"local_service_address"`
	LocalServicePort       types.Int64             `tfsdk:"local_service_port"`
	LocalServiceSocketPath types.String            `tfsdk:"local_service_socket_path"`
	Mode                   types.String            `tfsdk:"mode"`
	TransparentProxy       *TransparentProxyConfig `tfsdk:"transparent_proxy"`
	Config                 types.String            `tfsdk:"config"`
	Upstreams              []*Upstream             `tfsdk:"upstreams"`
	MeshGateway            *MeshGatewayConfig      `tfsdk:"mesh_gateway"`
	Expose                 *ExposeConfig           `tfsdk:"expose"`
	AccessLogs             *AccessLogsConfig       `tfsdk:"access_logs"`
}

type HealthCheck struct {
	Node        types.String           `tfsdk:"node"`
	CheckID     types.String           `tfsdk:"check_id"`
	Name        types.String           `tfsdk:"name"`
	Status      types.String           `tfsdk:"status"`
	Notes       types.String           `tfsdk:"notes"`
	Output      types.String           `tfsdk:"output"`
	ServiceID   types.String           `tfsdk:"service_id"`
	ServiceName types.String           `tfsdk:"service_name"`
	ServiceTags []types.String         `tfsdk:"service_tags"`
	Type        types.String           `tfsdk:"type"`
	Namespace   types.String           `tfsdk:"namespace"`
	Partition   types.String           `tfsdk:"partition"`
	ExposedPort types.Int64            `tfsdk:"exposed_port"`
	PeerName    types.String           `tfsdk:"peer_name"`
	Definition  *HealthCheckDefinition `tfsdk:"definition"`
}

type HttpBasicAuth struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type TLSConfig struct {
	Address            types.String `tfsdk:"address"`
	CAFile             types.String `tfsdk:"ca_file"`
	CAPath             types.String `tfsdk:"ca_path"`
	CAPem              types.String `tfsdk:"ca_pem"`
	CertFile           types.String `tfsdk:"cert_file"`
	CertPEM            types.String `tfsdk:"cert_pem"`
	KeyFile            types.String `tfsdk:"key_file"`
	KeyPEM             types.String `tfsdk:"key_pem"`
	InsecureSkipVerify types.Bool   `tfsdk:"insecure_skip_verify"`
}

type KVPair struct {
	Key       types.String `tfsdk:"key"`
	Flags     types.Int64  `tfsdk:"flags"`
	Value     types.String `tfsdk:"value"`
	Namespace types.String `tfsdk:"namespace"`
	Partition types.String `tfsdk:"partition"`
}

type NamespaceACLConfig struct {
	PolicyDefaults []*ACLLink `tfsdk:"policy_defaults"`
	RoleDefaults   []*ACLLink `tfsdk:"role_defaults"`
}

type SerfMember struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Port       types.Int64  `tfsdk:"port"`
	Datacenter types.String `tfsdk:"datacenter"`
	Role       types.String `tfsdk:"role"`
	Build      types.String `tfsdk:"build"`
	Protocol   types.Int64  `tfsdk:"protocol"`
	Status     types.String `tfsdk:"status"`
	RTT        types.String `tfsdk:"rtt"`
}

type Node struct {
	ID              types.String            `tfsdk:"id"`
	Node            types.String            `tfsdk:"node"`
	Address         types.String            `tfsdk:"address"`
	Datacenter      types.String            `tfsdk:"datacenter"`
	TaggedAddresses map[string]types.String `tfsdk:"tagged_addresses"`
	Meta            map[string]types.String `tfsdk:"meta"`
	Partition       types.String            `tfsdk:"partition"`
	PeerName        types.String            `tfsdk:"peer_name"`
	Locality        *Locality               `tfsdk:"locality"`
}

type PeeringStreamStatus struct {
	ImportedServices []types.String `tfsdk:"imported_services"`
	ExportedServices []types.String `tfsdk:"exported_services"`
	LastHeartbeat    types.String   `tfsdk:"last_heartbeat"`
	LastReceive      types.String   `tfsdk:"last_receive"`
	LastSend         types.String   `tfsdk:"last_send"`
}

type PeeringRemoteInfo struct {
	Partition  types.String `tfsdk:"partition"`
	Datacenter types.String `tfsdk:"datacenter"`
	Locality   *Locality    `tfsdk:"locality"`
}

type ServiceQuery struct {
	Service        types.String            `tfsdk:"name"`
	SamenessGroup  types.String            `tfsdk:"sameness_group"`
	Namespace      types.String            `tfsdk:"namespace"`
	Partition      types.String            `tfsdk:"partition"`
	Near           types.String            `tfsdk:"near"`
	Failover       *QueryFailoverOptions   `tfsdk:"failover"`
	IgnoreCheckIDs []types.String          `tfsdk:"ignore_check_ids"`
	OnlyPassing    types.Bool              `tfsdk:"only_passing"`
	Tags           []types.String          `tfsdk:"tags"`
	NodeMeta       map[string]types.String `tfsdk:"node_meta"`
	ServiceMeta    map[string]types.String `tfsdk:"service_meta"`
	Connect        types.Bool              `tfsdk:"connect"`
}

type QueryDNSOptions struct {
	TTL types.String `tfsdk:"ttl"`
}

type QueryTemplate struct {
	Type            types.String `tfsdk:"type"`
	Regexp          types.String `tfsdk:"regexp"`
	RemoveEmptyTags types.Bool   `tfsdk:"remove_empty_tags"`
}

type ServiceEntry struct{}

type EnvoyExtension struct {
	Name          types.String `tfsdk:"name"`
	Required      types.Bool   `tfsdk:"required"`
	Arguments     types.String `tfsdk:"arguments"`
	ConsulVersion types.String `tfsdk:"consul_version"`
	EnvoyVersion  types.String `tfsdk:"envoy_version"`
}

type TransparentProxyConfig struct {
	OutboundListenerPort types.Int64 `tfsdk:"outbound_listener_port"`
	DialedDirectly       types.Bool  `tfsdk:"dialed_directly"`
}

type Upstream struct {
	DestinationType      types.String       `tfsdk:"destination_type"`
	DestinationPartition types.String       `tfsdk:"destination_partition"`
	DestinationNamespace types.String       `tfsdk:"destination_namespace"`
	DestinationPeer      types.String       `tfsdk:"destination_peer"`
	DestinationName      types.String       `tfsdk:"destination_name"`
	Datacenter           types.String       `tfsdk:"datacenter"`
	LocalBindAddress     types.String       `tfsdk:"local_bind_address"`
	LocalBindPort        types.Int64        `tfsdk:"local_bind_port"`
	LocalBindSocketPath  types.String       `tfsdk:"local_bind_socket_path"`
	LocalBindSocketMode  types.String       `tfsdk:"local_bind_socket_mode"`
	Config               types.String       `tfsdk:"config"`
	MeshGateway          *MeshGatewayConfig `tfsdk:"mesh_gateway"`
	CentrallyConfigured  types.Bool         `tfsdk:"centrally_configured"`
}

type MeshGatewayConfig struct {
	Mode types.String `tfsdk:"mode"`
}

type ExposeConfig struct {
	Checks types.Bool    `tfsdk:"checks"`
	Paths  []*ExposePath `tfsdk:"paths"`
}

type AccessLogsConfig struct {
	Enabled             types.Bool   `tfsdk:"enabled"`
	DisableListenerLogs types.Bool   `tfsdk:"disable_listener_logs"`
	Type                types.String `tfsdk:"type"`
	Path                types.String `tfsdk:"path"`
	JSONFormat          types.String `tfsdk:"json_format"`
	TextFormat          types.String `tfsdk:"text_format"`
}

type HealthCheckDefinition struct {
	HTTP                                   types.String              `tfsdk:"http"`
	Header                                 map[string][]types.String `tfsdk:"header"`
	Method                                 types.String              `tfsdk:"method"`
	Body                                   types.String              `tfsdk:"body"`
	TLSServerName                          types.String              `tfsdk:"tls_server_name"`
	TLSSkipVerify                          types.Bool                `tfsdk:"tls_skip_verify"`
	TCP                                    types.String              `tfsdk:"tcp"`
	UDP                                    types.String              `tfsdk:"udp"`
	GRPC                                   types.String              `tfsdk:"grpc"`
	OSService                              types.String              `tfsdk:"os_service"`
	GRPCUseTLS                             types.Bool                `tfsdk:"grpc_use_tls"`
	IntervalDuration                       types.String              `tfsdk:"interval_duration"`
	TimeoutDuration                        types.String              `tfsdk:"timeout_duration"`
	DeregisterCriticalServiceAfterDuration types.String              `tfsdk:"deregister_critical_service_after_duration"`
}

type QueryFailoverOptions struct {
	NearestN    types.Int64            `tfsdk:"nearest_n"`
	Datacenters []types.String         `tfsdk:"datacenters"`
	Targets     []*QueryFailoverTarget `tfsdk:"targets"`
}

type ExposePath struct {
	ListenerPort    types.Int64  `tfsdk:"listener_port"`
	Path            types.String `tfsdk:"path"`
	LocalPathPort   types.Int64  `tfsdk:"local_path_port"`
	Protocol        types.String `tfsdk:"protocol"`
	ParsedFromCheck types.Bool   `tfsdk:"parsed_from_check"`
}

type QueryFailoverTarget struct {
	Peer       types.String `tfsdk:"peer"`
	Datacenter types.String `tfsdk:"datacenter"`
	Partition  types.String `tfsdk:"partition"`
	Namespace  types.String `tfsdk:"namespace"`
}
