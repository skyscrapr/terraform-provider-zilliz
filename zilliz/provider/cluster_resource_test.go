// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccClusterResource_basic(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix("tftest")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccClusterResourceConfig_basic(rName, "1"),
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
			//  - Not supported
			//
			// Update and Read testing
			{
				Config: testAccClusterResourceConfig_basic(rName, "2"),
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

func TestAccClusterResource_serverless(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix("tftest")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccClusterResourceConfig_serverless(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("zilliz_cluster.test", "id"),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "cluster_name", rName),
					resource.TestCheckResourceAttr("zilliz_cluster.test", "cu_size", "0"),
					resource.TestCheckNoResourceAttr("zilliz_cluster.test", "plan"),
					resource.TestCheckNoResourceAttr("zilliz_cluster.test", "cu_type"),
					resource.TestCheckResourceAttrSet("zilliz_cluster.test", "project_id"),
				),
			},
			// ImportState - Not supported
			// Update and Read testing - Not supported
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccClusterResourceConfig_basic(name string, cu_size string) string {
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

func testAccClusterResourceConfig_serverless(name string) string {
	return fmt.Sprintf(`
provider "zilliz" {
	cloud_region_id = "gcp-us-west1"
}

data "zilliz_projects" "test" {
}

resource "zilliz_cluster" "test" {
	cluster_name = %q
	project_id   = data.zilliz_projects.test.projects[0].project_id
}
`, name)
}
