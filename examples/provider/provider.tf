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
