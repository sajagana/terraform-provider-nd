package resource_multi_cluster_connectivity

import (
	"context"
	"fmt"

	"terraform-provider-nd/internal/infra"
	"terraform-provider-nd/internal/registry"

	"log"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ModuleKey is the key used to get the infra module from the provider.
const ModuleKey = "infra"

// Ensure the implementation satisfies the expected interfaces
var (
	_ resource.Resource                = &multiClusterConnectivityResource{}
	_ resource.ResourceWithConfigure   = &multiClusterConnectivityResource{}
	_ resource.ResourceWithImportState = &multiClusterConnectivityResource{}
)

// NewMultiClusterConnectivityResource is a helper function to simplify the provider implementation.
func NewMultiClusterConnectivityResource() resource.Resource {
	return &multiClusterConnectivityResource{}
}

// multiClusterConnectivityResource is the resource implementation.
type multiClusterConnectivityResource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the resource type name.
func (r *multiClusterConnectivityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_multi_cluster_connectivity"
}

// Schema defines the schema for the resource.
func (r *multiClusterConnectivityResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = MultiClusterConnectivityResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *multiClusterConnectivityResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	infraModule := client.GetModule(ModuleKey)
	if infraModule == nil {
		resp.Diagnostics.AddError(
			"Infra Module Not Found",
			"The infra module was not registered with the provider.",
		)
		return
	}

	r.infraClient = infraModule.(*infra.NexusDashboardInfra)
}

// Create creates the resource and sets the initial Terraform state.
func (r *multiClusterConnectivityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var in MultiClusterConnectivityModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Creating Multi Cluster Connectivity: cluster_name=%s", in.ClusterName.ValueString())

	r.rscCreateMultiClusterConnectivity(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
}

// Read refreshes the Terraform state with the latest data.
func (r *multiClusterConnectivityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MultiClusterConnectivityModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Reading Multi Cluster Connectivity: cluster_name=%s", state.ClusterName.ValueString())

	// Preserve sensitive fields that are not returned by the API
	preservedUsername := state.Username
	preservedPassword := state.Password
	preservedLoginDomain := state.LoginDomain

	r.rscGetMultiClusterConnectivity(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	// Restore sensitive fields after read
	state.Username = preservedUsername
	state.Password = preservedPassword
	state.LoginDomain = preservedLoginDomain

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *multiClusterConnectivityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan MultiClusterConnectivityModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Updating Multi Cluster Connectivity: cluster_name=%s", plan.ClusterName.ValueString())

	r.rscUpdateMultiClusterConnectivity(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *multiClusterConnectivityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MultiClusterConnectivityModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Deleting Multi Cluster Connectivity: cluster_name=%s", state.ClusterName.ValueString())

	r.rscDeleteMultiClusterConnectivity(ctx, &resp.Diagnostics, &state)
}

// ImportState imports a multi cluster connectivity resource by cluster name.
func (r *multiClusterConnectivityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Importing Multi Cluster Connectivity: id=%s", req.ID)

	var state MultiClusterConnectivityModel
	state.ClusterName = types.StringValue(req.ID)

	r.rscGetMultiClusterConnectivity(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
