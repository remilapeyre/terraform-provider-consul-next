// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	test "github.com/remilapeyre/terraform-plugin-test"
	"github.com/sethvargo/go-retry"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"consul": providerserver.NewProtocol6WithError(New("test")()),
	}

	testClient *api.Client = nil

	serverConfiguration = `
		limits = {
			http_max_conns_per_client = -1
		}

		acl = {
			enabled        = true
			default_policy = "allow"
			down_policy    = "extend-cache"

			tokens = {
			initial_management = "12345678-1234-1234-1234-1234567890ab"
			}
		}`
)

func testAccPreCheck(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "consul.hcl")
	f, err := os.Create(configPath)
	if err != nil {
		t.Fatalf("failed to create configuration file: %v", err)
	}
	_, err = f.WriteString(serverConfiguration)
	if err != nil {
		t.Fatalf("failed to write configuration: %v", err)
	}
	f.Close()

	path := os.Getenv("CONSUL_TEST_BINARY")
	if path == "" {
		path = "consul"
	}
	path, err = exec.LookPath(path)
	if err != nil {
		t.Fatalf("failed to find nomad in PATH: %v", err)
	}

	cmd := exec.Command(path, "agent", "-dev", "-config-file", configPath)

	if os.Getenv("TF_ACC_CONSUL_LOG") != "" {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start Consul: %s", err)
	}

	t.Setenv("CONSUL_HTTP_TOKEN", "12345678-1234-1234-1234-1234567890ab")
	config := api.DefaultConfig()
	testClient, err = api.NewClient(config)
	require.NoError(t, err)

	b := retry.NewConstant(1 * time.Second)
	err = retry.Do(context.Background(), retry.WithMaxRetries(10, b), func(ctx context.Context) error {
		_, err := testClient.Status().Leader()
		return retry.RetryableError(err)
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		cmd.Process.Kill()
		cmd.Process.Wait()
	})
}

func TestResources(t *testing.T) {
	f := func(t *testing.T, tc *resource.TestCase) {
		tc.PreCheck = func() { testAccPreCheck(t) }
		tc.ProtoV6ProviderFactories = testAccProtoV6ProviderFactories
	}
	opts := &test.TestOptions{
		IgnoreChange: func(name, key, value string) bool {
			if ignored := test.DefaultIgnoreChangeFunc(name, key, value); ignored {
				return true
			}

			ignored := []string{
				"consul_acl_token_role_attachment.attachment.id",
				"consul_peering_token.test.peering_token",
				"data.consul_agent_config.test.node_name",
				"data.consul_autopilot_health.test.servers.0.name",
				"data.consul_nodes.test.nodes.0.node",
			}
			name = fmt.Sprintf("%s.%s", name, key)
			return slices.Contains(ignored, name)
		},
	}
	test.Test(t, "../../tests/", f, opts)
}
