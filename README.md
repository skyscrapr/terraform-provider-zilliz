# Terraform provider for Zilliz Cloud

[![Go Reference](https://pkg.go.dev/badge/github.com/skyscrapr/terraform-provider-zilliz.svg)](https://pkg.go.dev/github.com/skyscrapr/terraform-provider-zilliz)
[![Go Report Card](https://goreportcard.com/badge/github.com/skyscrapr/terraform-provider-zilliz)](https://goreportcard.com/report/github.com/skyscrapr/terraform-provider-zilliz)
[![codecov](https://codecov.io/gh/skyscrapr/terraform-provider-pinecone/graph/badge.svg?token=qobuIzQPuM)](https://codecov.io/gh/skyscrapr/terraform-provider-zilliz)
![Github Actions Workflow](https://github.com/skyscrapr/terraform-provider-zilliz/actions/workflows/test.yml/badge.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/skyscrapr/terraform-provider-zilliz)
![License](https://img.shields.io/dub/l/vibe-d.svg)

This provider allows terraform to manage Zilliz resources.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= v1.4.6
- [Go](https://golang.org/doc/install) >= 1.20 (to build the provider plugin)

## Installing the Provider

<!-- The provider is not yet registered in the official [terraform registry](https://registry.terraform.io/providers/skyscrapr/zilliz/latest) 

This enables the provider to be auto-installed when you run ```terraform init``` -->

You can download the latest binary for your target platform from the [releases](https://github.com/skyscrapr/terraform-provider-zilliz/releases) tab.

## Building the Provider

- Clone the repo:
    ```sh
    $ git clone https://github.com/skyscrapr/terraform-provider-zilliz
    ```

- Build the provider: (NOTE: the install directory will be set accoring to GOPATH environment variable)
    ```sh
    $ go install .
    ```

## Usage

You can enable the provider in your terraform configurtion by add the folowing:
```terraform
terraform {
  required_providers {
    openai = {
      source = "skyscrapr/zilliz"
    }
  }
}
```

You can configure the Pinecone client using environment variables to avoid setting sensitive values in terraform config.
- Set `ZILLIZ_API_KEY` to your Pinecone API Key.

## Documentation

Documentation can be found on the [Terraform Registry](https://registry.terraform.io/providers/skyscrapr/zilliz/latest). 

## Examples

Please see the [examples](https://github.com/skyscrapr/terraform-provider-zilliz/examples) for example usage.

## Support

Please raise an issue for any support related requirements.
