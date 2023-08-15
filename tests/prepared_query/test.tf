# Check: consul_prepared_query.test
# Import: consul_prepared_query.test
resource "consul_prepared_query" "test" {
  name  = "myquery"
  token = "wxyz"

  service = {
    name         = "myapp"
    only_passing = true
    near         = "_agent"

    failover = {
      nearest_n   = 3
      datacenters = ["us-west1", "us-east-2", "asia-east1"]
    }

    tags = ["active", "!standby"]
  }

  dns = {
    ttl = "30s"
  }
}
