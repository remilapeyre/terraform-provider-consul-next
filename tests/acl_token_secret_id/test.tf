# Check: consul_acl_token.test
resource "consul_acl_token" "test" {}

# Check: data.consul_acl_token_secret_id.test
data "consul_acl_token_secret_id" "test" {
  accessor_id = consul_acl_token.test.id
}
