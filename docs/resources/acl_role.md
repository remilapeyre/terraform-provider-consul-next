---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "consul_acl_role Resource - consul"
subcategory: ""
description: |-
  
---

# consul_acl_role (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of this ACL role.

### Optional

- `description` (String) The description of this ACL role.
- `namespace` (String) The namespace the ACL role is associated with. Namespacing is a Consul Enterprise feature.
- `node_identities` (Attributes List) The nodes associated with this ACL role. (see [below for nested schema](#nestedatt--node_identities))
- `partition` (String) The partition the ACL role is associated with. Partitions are a Consul Enterprise feature.
- `policies` (Attributes List) The policies associated with this ACL role. (see [below for nested schema](#nestedatt--policies))
- `service_identities` (Attributes List) The services associated with this ACL role. (see [below for nested schema](#nestedatt--service_identities))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedatt--node_identities"></a>
### Nested Schema for `node_identities`

Required:

- `node_name` (String) The node associated with this ACL role.

Optional:

- `datacenter` (String)


<a id="nestedatt--policies"></a>
### Nested Schema for `policies`

Optional:

- `id` (String) The ID of the ACL policy associated with this role.
- `name` (String) The name of the ACL policy associated with this role.


<a id="nestedatt--service_identities"></a>
### Nested Schema for `service_identities`

Required:

- `service_name` (String) The service associated with this ACL role.

Optional:

- `datacenters` (List of String)
