package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCloudRegionsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRegionsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.zilliz_cloud_regions.test", "id"),
				),
			},
		},
	})
}

const testAccCloudRegionsDataSourceConfig = `
provider "zilliz" {
	cloud_region_id = "gcp-us-west1"
}

data "zilliz_cloud_regions" "test" {
	cloud_id = "gcp"
}
`
