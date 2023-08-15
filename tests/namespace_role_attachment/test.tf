resource "consul_acl_policy" "test" {
  name  = "test"
  rules = <<-RULE
    node_prefix "" {
      policy = "read"
    }
  RULE
}

resource "consul_acl_role" "test" {
  name = "test"

  policies = [
    {
      id = consul_acl_policy.test.id
    }
  ]

  service_identities = [
    {
      service_name = "foo"
    }
  ]
}


# ExpectError: namespace not found
resource "consul_namespace_role_attachment" "test" {
  namespace = "default"
  role = {
    id = consul_acl_role.test.name
  }
}
