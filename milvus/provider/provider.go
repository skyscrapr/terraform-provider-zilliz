// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/Mufassa12/milvus-sdk-go/milvus"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure milvusProvider satisfies various provider interfaces.
var _ provider.Provider = &milvusProvider{}

// milvusProvider defines the provider implementation.
type milvusProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// milvusProviderModel describes the provider data model.
type milvusProviderModel struct {
	ApiKey types.String `tfsdk:"api_key"`
	Region types.String `tfsdk:"region"`
}

func (p *milvusProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "milvus"
	resp.Version = p.version
}

func (p *milvusProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				MarkdownDescription: "milvus API Key",
				Optional:            true,
				Sensitive:           true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "milvus Region",
				Required:            true,
			},
		},
	}
}

func (p *milvusProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data milvusProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Default to environment variables, but override
	// with Terraform configuration value if set.
	apiKey := os.Getenv("MILVUS_API_KEY")
	if !data.ApiKey.IsNull() {
		apiKey = data.ApiKey.ValueString()
	}
	client := milvus.NewClient(apiKey, data.Region.ValueString())

	// Example client configuration for data sources and resources
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *milvusProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *milvusProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewExampleDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &milvusProvider{
			version: version,
		}
	}
}
