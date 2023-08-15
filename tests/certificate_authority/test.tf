# Check: consul_certificate_authority.test
resource "consul_certificate_authority" "test" {
  connect_provider            = "consul"
  force_without_cross_signing = false

  config = jsonencode({
    IntermediateCertTTL = "216h"
    LeafCertTTL         = "72h"
    RootCertTTL         = "216h"
  })
}
