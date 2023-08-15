# ExpectError: no role found
data "consul_acl_role" "test" {
  name = "test"
}
