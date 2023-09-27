provider "zilliz" {
  cloud_region_id = "gcp-us-west1"
}

data "zilliz_clusters" "test" {}

data "zilliz_cluster" "test" {
  id = data.zilliz_clusters.test.clusters[0].id
}
