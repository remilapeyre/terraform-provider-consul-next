resource "consul_acl_auth_method" "test" {
  name = "test"
  type = "jwt"

  config = jsonencode({
    JWTValidationPubKeys = ["-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAERVchfCZng4mmdvQz1+sJHRN40snC\nYt8NjYOnbnScEXMkyoUmASr88gb7jaVAVt3RYASAbgBjB2Z+EUizWkx5Tg==\n-----END PUBLIC KEY-----"]
  })
}

# Check: consul_acl_binding_rule.test
# Import: consul_acl_binding_rule.test
resource "consul_acl_binding_rule" "test" {
  auth_method = consul_acl_auth_method.test.id
  bind_type   = "role"
  bind_name   = "foo"
}

# Check: data.consul_acl_binding_rule.test
data "consul_acl_binding_rule" "test" {
  id = consul_acl_binding_rule.test.id
}
