# ExpectError: no ACL auth method "not-found" found
data "consul_acl_auth_method" "test" {
  name = "not-found"
}
