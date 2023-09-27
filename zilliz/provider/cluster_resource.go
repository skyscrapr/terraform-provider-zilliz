// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"terraform-provider-zilliz/zilliz"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ClusterResource{}
var _ resource.ResourceWithImportState = &ClusterResource{}

func NewClusterResource() resource.Resource {
	return &ClusterResource{}
}

// ClusterResource defines the resource implementation.
type ClusterResource struct {
	client *zilliz.Client
}

func (r *ClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (r *ClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the cluster.",
				Computed:            true,
			},
			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "The name of the cluster.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "An optional description about the cluster.",
				Computed:            true,
			},
			"region_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the region where the cluster exists.",
				Required:            true,
			},
			"cluster_type": schema.StringAttribute{
				MarkdownDescription: "The type of CU associated with the cluster. Possible values are Performance-optimized and Capacity-optimized.",
				Computed:            true,
			},
			"cu_size": schema.Int64Attribute{
				MarkdownDescription: "The size of the CU associated with the cluster.",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "The current status of the cluster. Possible values are INITIALIZING, RUNNING, SUSPENDING, and RESUMING.",
				Computed:            true,
			},
			"connect_address": schema.StringAttribute{
				MarkdownDescription: "The public endpoint of the cluster. You can connect to the cluster using this endpoint from the public network.",
				Computed:            true,
			},
			"private_link_address": schema.StringAttribute{
				MarkdownDescription: "The private endpoint of the cluster. You can set up a private link to allow your VPS in the same cloud region to access your cluster.",
				Computed:            true,
			},
			"create_time": schema.StringAttribute{
				MarkdownDescription: "The time at which the cluster has been created.",
				Computed:            true,
			},
		},
	}
}

func (r *ClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*zilliz.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusterModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.CreateCluster(zilliz.CreateClusterParams{
		CreateCollection:        false,
		CreateExampleCollection: true,
		InstanceName:            data.ClusterName.ValueString(),
		ProjectId:               0,
		RegionId:                data.RegionId.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to CreateCluster, got error: %s", err))
		return
	}

	// save into the Terraform state.
	data.ClusterId = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusterModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	c, err := r.client.DescribeCluster(data.ClusterId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to DescribeCluster, got error: %s", err))
		return
	}

	// Save data into Terraform state
	data.ClusterId = types.StringValue(c.ClusterId)
	data.ClusterName = types.StringValue(c.ClusterName)
	data.Description = types.StringValue(c.Description)
	data.RegionId = types.StringValue(c.RegionId)
	data.ClusterType = types.StringValue(c.ClusterType)
	data.CuSize = types.Int64Value(c.CuSize)
	data.Status = types.StringValue(c.Status)
	data.ConnectAddress = types.StringValue(c.ConnectAddress)
	data.PrivateLinkAddress = types.StringValue(c.PrivateLinkAddress)
	data.CreateTime = types.StringValue(c.CreateTime)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusterModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusterModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}

func (r *ClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
