package provider

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClusterDataSource(t *testing.T) {
	rName := sdkacctest.RandomWithPrefix("tftest")
	dataSourceResourceName := "data.zilliz_cluster.test"
	resourceName := "zilliz_cluster.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccClusterDataSourceConfig(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceResourceName, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "cluster_name", dataSourceResourceName, "cluster_name"),
					resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceResourceName, "description"),
					resource.TestCheckResourceAttrPair(resourceName, "region_id", dataSourceResourceName, "region_id"),
					resource.TestCheckResourceAttrPair(resourceName, "cluster_type", dataSourceResourceName, "cluster_type"),
					resource.TestCheckResourceAttrPair(resourceName, "cu_size", dataSourceResourceName, "cu_size"),
					resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceResourceName, "status"),
					resource.TestCheckResourceAttrPair(resourceName, "connect_address", dataSourceResourceName, "connect_address"),
					resource.TestCheckResourceAttrPair(resourceName, "private_link_address", dataSourceResourceName, "private_link_address"),
					resource.TestCheckResourceAttrPair(resourceName, "create_time", dataSourceResourceName, "create_time"),
				),
			},
		},
	})
}

func testAccClusterDataSourceConfig(name string) string {
	return fmt.Sprintf(`
provider "zilliz" {
	cloud_region_id = "gcp-us-west1"
}

data "zilliz_projects" "test" {
}

resource "zilliz_cluster" "test" {
	plan         = "Standard"
	cluster_name = %q
	cu_size      = 1
	cu_type      = "Performance-optimized"
	project_id   = data.zilliz_projects.test.projects[0].project_id
}

data "zilliz_cluster" "test" {
	id = zilliz_cluster.test.id
}
`, name)
}
