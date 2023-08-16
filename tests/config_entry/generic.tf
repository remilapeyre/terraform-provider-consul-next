# Check: consul_config_entry.test
resource "consul_config_entry" "test" {
  kind = "proxy-defaults"
  name = "global"

  config = jsonencode({
    AccessLogs       = {}
    Expose           = {}
    MeshGateway      = {}
    TransparentProxy = {}
    Config = {
      local_connect_timeout_ms = 1000
      handshake_timeout_ms     = 10000
    }
  })
}
