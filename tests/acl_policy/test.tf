# Check: consul_acl_policy.test
# Import: consul_acl_policy.test
resource "consul_acl_policy" "test" {
  name  = "test"
  rules = <<-RULE
    node_prefix "" {
      policy = "read"
    }
  RULE
}

# Check: data.consul_acl_policy.test
data "consul_acl_policy" "test" {
  id = consul_acl_policy.test.id
}
