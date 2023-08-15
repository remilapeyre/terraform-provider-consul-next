# ExpectError: policies\[0\].name and policies\[0\].id cannot be both set
resource "consul_acl_token" "test" {
  policies = [
    {
      id   = "1234"
      name = "test"
    }
  ]
}
