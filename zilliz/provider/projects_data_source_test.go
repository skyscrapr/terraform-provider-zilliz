package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProjectsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.zilliz_projects.test", "id"),
				),
			},
		},
	})
}

const testAccProjectsDataSourceConfig = `
provider "zilliz" {
	cloud_region_id = "gcp-us-west1"
}

data "zilliz_projects" "test" {
}
`
