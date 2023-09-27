package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClusterDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccClusterDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.zilliz_cluster.test", "id"),
				),
			},
		},
	})
}

const testAccClusterDataSourceConfig = `
provider "zilliz" {
	cloud_region_id = "gcp-us-west1"
}

data "zilliz_clusters" "test" {}

data "zilliz_cluster" "test" {
	id = data.zilliz_clusters.test.clusters[0].id
}
`
