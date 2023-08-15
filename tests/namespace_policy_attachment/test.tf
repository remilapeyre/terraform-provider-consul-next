resource "consul_acl_policy" "agent" {
  name  = "agent"
  rules = <<-RULE
    node_prefix "" {
      policy = "read"
    }
    RULE
}

# ExpectError: namespace not found
resource "consul_namespace_policy_attachment" "test" {
  namespace = "default"
  policy = {
    name = consul_acl_policy.agent.name
  }
}
