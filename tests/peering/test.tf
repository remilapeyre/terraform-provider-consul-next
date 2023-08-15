# ExpectError: no peering named "test" could be found
data "consul_peering" "test" {
  name = "test"
}
