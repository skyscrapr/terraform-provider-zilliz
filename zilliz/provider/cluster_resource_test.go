// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClusterResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccClusterResourceConfig("one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("zilliz_cluster.test", "configurable_attribute", "one"),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "defaulted", "example value when not configured"),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "id", "example-id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "zilliz_cluster.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				ImportStateVerifyIgnore: []string{"configurable_attribute", "defaulted"},
			},
			// Update and Read testing
			{
				Config: testAccClusterResourceConfig("two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("zilliz_cluster.test", "configurable_attribute", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccClusterResourceConfig(clusterName string) string {
	return fmt.Sprintf(`
provider "zilliz" {
	cloud_region_id = "gcp-us-west1"
}

resource "zilliz_cluster" "test" {
  cluster_name = %[1]q
  region_id = "gcp-us-west1"
}
`, clusterName)
}
