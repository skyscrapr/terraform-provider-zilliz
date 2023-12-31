terraform {
  required_providers {
    zilliz = {
      source = "skyscrapr/zilliz"
    }
  }
}

provider "zilliz" {
  cloud_region_id = "gcp-us-west1"
}

data "zilliz_cloud_providers" "example" {}

data "zilliz_cloud_regions" "example" {
  cloud_id = data.zilliz_cloud_providers.example.cloud_providers[0].cloud_id
}
