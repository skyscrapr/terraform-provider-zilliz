// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccClusterResource(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix("tftest")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccClusterResourceConfig(rName, "1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("zilliz_cluster.test", "id"),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "cluster_name", rName),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "plan", "Standard"),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "cu_size", "1"),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "cu_type", "Performance-optimized"),
					resource.TestCheckResourceAttrSet("zilliz_cluster.test", "project_id"),
				),
			},
			// ImportState testing
			// {
			// 	ResourceName:      "zilliz_cluster.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	// ImportStateVerifyIdentifierAttribute: "name",
			// 	// This is not normally necessary, but is here because this
			// 	// example code does not have an actual upstream service.
			// 	// Once the Read method is able to refresh information from
			// 	// the upstream service, this can be removed.
			// 	// ImportStateVerifyIgnore: []string{"source"},
			// },
			// Update and Read testing
			{
				Config: testAccClusterResourceConfig(rName, "2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("zilliz_cluster.test", "id"),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "cluster_name", rName),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "plan", "Standard"),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "cu_size", "2"),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "cu_type", "Performance-optimized"),
					resource.TestCheckResourceAttrSet("zilliz_cluster.test", "project_id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccClusterResourceConfig(name string, cu_size string) string {
	return fmt.Sprintf(`
provider "zilliz" {
	cloud_region_id = "gcp-us-west1"
}

data "zilliz_projects" "test" {
}

resource "zilliz_cluster" "test" {
	plan         = "Standard"
	cluster_name = %q
	cu_size      = %q
	cu_type      = "Performance-optimized"
	project_id   = data.zilliz_projects.test.projects[0].project_id
}
`, name, cu_size)
}
