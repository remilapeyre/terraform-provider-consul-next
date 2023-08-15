# ExpectError: no namespace named "default" found
data "consul_namespace" "test" {
  name = "default"
}
