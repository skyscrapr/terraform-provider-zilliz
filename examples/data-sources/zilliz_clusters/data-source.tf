provider "zilliz" {
  cloud_region_id = "gcp-us-west1"
}

data "zilliz_clusters" "example" {}
