# Check: consul_autopilot_config.test
resource "consul_autopilot_config" "test" {
  cleanup_dead_servers      = true
  disable_upgrade_migration = false
  max_trailing_logs         = 2500
}
