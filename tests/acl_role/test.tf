resource "consul_acl_policy" "test" {
  name  = "test"
  rules = <<-RULE
    node_prefix "" {
      policy = "read"
    }
  RULE
}

# Check: consul_acl_role.test
# Import: consul_acl_role.test
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
