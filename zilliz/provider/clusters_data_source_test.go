package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClustersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccClustersDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.zilliz_clusters.test", "id"),
				),
			},
		},
	})
}

const testAccClustersDataSourceConfig = `
provider "zilliz" {
	cloud_region_id = "gcp-us-west1"
}

data "zilliz_clusters" "test" {
}
`
