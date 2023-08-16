# Check: consul_config_entry_service.test
resource "consul_config_entry_service" "test" {
  name     = "foo"
  protocol = "http"
}

# Check: consul_config_entry_proxy.test
resource "consul_config_entry_proxy" "test" {
  name = "global"
  config = jsonencode({
    foo = "bar"
  })
}

# Check: consul_config_entry_service_router.test
resource "consul_config_entry_service_router" "test" {
  name = consul_config_entry_service.test.name
  routes = [
    {
      match = {
        http = {
          pathprefix = "/admin"
        }
      }

      destination = {
        namespace = "default"
        partition = "default"
        service   = consul_config_entry_service.test.name
      }
    }
  ]
}
