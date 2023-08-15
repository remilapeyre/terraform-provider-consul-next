resource "consul_acl_policy" "test" {
  name  = "test"
  rules = <<-RULE
    node_prefix "" {
      policy = "read"
    }
  RULE
}

# Check: consul_acl_token.test
resource "consul_acl_token" "test" {
  policies = [
    {
      id = consul_acl_policy.test.id
    }
  ]
}

# Check: data.consul_acl_token.test
data "consul_acl_token" "test" {
  accessor_id = consul_acl_token.test.id
}
