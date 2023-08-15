# ExpectError: no ACL auth method "not-found" found
data "consul_acl_policy" "test" {
  name = "not-found"
}
