# Check: consul_peering_token.test
resource "consul_peering_token" "test" {
  peer_name = "eu-cluster"
}
