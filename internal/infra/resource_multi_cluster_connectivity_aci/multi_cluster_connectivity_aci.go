package resource_multi_cluster_connectivity_aci

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
	_ resource.Resource                = &multiClusterConnectivityAciResource{}
	_ resource.ResourceWithConfigure   = &multiClusterConnectivityAciResource{}
	_ resource.ResourceWithImportState = &multiClusterConnectivityAciResource{}
)

// NewMultiClusterConnectivityACIResource is a helper function to simplify the provider implementation.
func NewMultiClusterConnectivityACIResource() resource.Resource {
	return &multiClusterConnectivityAciResource{}
}

// multiClusterConnectivityAciResource is the resource implementation.
type multiClusterConnectivityAciResource struct {
	infraClient *infra.NexusDashboardInfra
}

// Metadata returns the resource type name.
func (r *multiClusterConnectivityAciResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_multi_cluster_connectivity_aci"
}

// Schema defines the schema for the resource.
func (r *multiClusterConnectivityAciResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = MultiClusterConnectivityAciResourceSchema(ctx)
}

// Configure adds the provider configured client to the resource.
func (r *multiClusterConnectivityAciResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *multiClusterConnectivityAciResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("[DEBUG] Start create of resource: nd_multi_cluster_connectivity_aci")

	var in MultiClusterConnectivityAciModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &in)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Creating Multi Cluster Connectivity ACI: cluster_name=%s", in.ClusterName.ValueString())

	r.rscCreateMultiClusterConnectivityAci(ctx, &resp.Diagnostics, &in)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &in)...)
	log.Printf("[DEBUG] End create of resource nd_multi_cluster_connectivity_aci with id '%s'", in.Id.ValueString())
}

// Read refreshes the Terraform state with the latest data.
func (r *multiClusterConnectivityAciResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("[DEBUG] Start read of resource: nd_multi_cluster_connectivity_aci")

	var state MultiClusterConnectivityAciModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Reading Multi Cluster Connectivity ACI: cluster_name=%s", state.ClusterName.ValueString())

	// // Preserve sensitive fields that are not returned by the API
	// preservedUserName := state.UserName
	// preservedPassword := state.Password
	// preservedLoginDomain := state.LoginDomain

	r.rscGetMultiClusterConnectivityAci(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	// // Restore sensitive fields after read
	// state.UserName = preservedUserName
	// state.Password = preservedPassword
	// state.LoginDomain = preservedLoginDomain

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End read of resource nd_multi_cluster_connectivity_aci with id '%s'", state.Id.ValueString())
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *multiClusterConnectivityAciResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("[DEBUG] Start update of resource: nd_multi_cluster_connectivity_aci")
	var plan MultiClusterConnectivityAciModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] Updating Multi Cluster Connectivity ACI: cluster_name=%s", plan.ClusterName.ValueString())

	r.rscUpdateMultiClusterConnectivityAci(ctx, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	log.Printf("[DEBUG] End update of resource nd_multi_cluster_connectivity_aci with id '%s'", plan.Id.ValueString())
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *multiClusterConnectivityAciResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("[DEBUG] Start delete of resource: nd_multi_cluster_connectivity_aci")
	var state MultiClusterConnectivityAciModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	r.rscDeleteMultiClusterConnectivityAci(ctx, &resp.Diagnostics, &state)
	log.Printf("[DEBUG] End delete of resource nd_multi_cluster_connectivity_aci with id '%s'", state.Id.ValueString())
}

// ImportState imports a multi cluster connectivity aci resource by cluster name.
func (r *multiClusterConnectivityAciResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("[DEBUG] Start import state of resource: nd_multi_cluster_connectivity_aci")

	var state MultiClusterConnectivityAciModel
	state.ClusterName = types.StringValue(req.ID)

	r.rscGetMultiClusterConnectivityAci(ctx, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	// TODO: The values for `user_name`, `password` and `login_domain` attributes will not be imported when the nd_multi_cluster_connectivity_aci resource imports an already onboarded ACI cluster from Nexus Dashboard.
	// Need to use Environment Variables or CLUSTER_CREDENTIALS_FILE_LOCATION file to import those values during the import of the resource.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	log.Printf("[DEBUG] End import of state resource: nd_multi_cluster_connectivity_aci")
}
