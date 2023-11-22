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
