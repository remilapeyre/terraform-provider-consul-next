resource "consul_acl_role" "test" {
  name = "test"
}


# Check: consul_acl_token_role_attachment.attachment
resource "consul_acl_token_role_attachment" "attachment" {
  token_id = "00000000-0000-0000-0000-000000000002"
  role = {
    id = consul_acl_role.test.id
  }
}
