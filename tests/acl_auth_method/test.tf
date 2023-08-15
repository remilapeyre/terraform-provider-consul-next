# Check: consul_acl_auth_method.test
# Import: consul_acl_auth_method.test
resource "consul_acl_auth_method" "test" {
  name = "test"
  type = "jwt"
  # max_token_ttl = "120m"  // TODO fix this behavior
  max_token_ttl = "2h0m0s"
  description   = "this is a test auth metho"

  config = jsonencode({
    JWTValidationPubKeys = ["-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAERVchfCZng4mmdvQz1+sJHRN40snC\nYt8NjYOnbnScEXMkyoUmASr88gb7jaVAVt3RYASAbgBjB2Z+EUizWkx5Tg==\n-----END PUBLIC KEY-----"]
  })
}

# Check: data.consul_acl_auth_method.test
data "consul_acl_auth_method" "test" {
  name = consul_acl_auth_method.test.name
}
