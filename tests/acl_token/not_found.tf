# ExpectError: failed to read ACL token "604c945c-1b35-462a-ad1e-7a1b621878fb"
data "consul_acl_token" "test" {
  accessor_id = "604c945c-1b35-462a-ad1e-7a1b621878fb"
}
