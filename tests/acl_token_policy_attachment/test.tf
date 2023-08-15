resource "consul_acl_policy" "test" {
  name  = "test"
  rules = <<-RULE
    node_prefix "" {
      policy = "read"
    }
  RULE
}

# Check: consul_acl_token_policy_attachment.attachment
resource "consul_acl_token_policy_attachment" "attachment" {
  token_id = "00000000-0000-0000-0000-000000000002"
  policy = {
    name = consul_acl_policy.test.name
  }
}
