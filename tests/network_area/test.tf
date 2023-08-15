# ExpectError: 404
resource "consul_network_area" "test" {
  peer_datacenter = "test"
}
