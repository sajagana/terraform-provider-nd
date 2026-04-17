package resource_multi_cluster_connectivity

import (
	"context"
	"encoding/json"
	"fmt"
	"terraform-provider-nd/internal/infra/api"

	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// setModelId sets the Id field on the model based on ClusterName.
// This is kept outside resource_codec_gen.go to avoid conflicts with the internal generator.
func setModelId(model *MultiClusterConnectivityModel) {
	if !model.ClusterName.IsNull() && !model.ClusterName.IsUnknown() {
		model.Id = types.StringValue(model.ClusterName.ValueString())
	} else {
		model.Id = types.StringNull()
	}
}

// RscCreateMultiClusterConnectivity creates a multi cluster connectivity resource
func (r *multiClusterConnectivityResource) rscCreateMultiClusterConnectivity(ctx context.Context, dg *diag.Diagnostics, input *MultiClusterConnectivityModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	inData := input.GetModelData()

	// log.Printf("Creating multi cluster connectivity %s with category %s", inData.ClusterName, inData.Category)

	// Create multi cluster connectivity API client
	clusterAPI := api.NewClusterAPI(nil, r.infraClient.ApiClient)

	// Convert model data to JSON
	clusterPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Multi Cluster Connectivity",
			fmt.Sprintf("Could not create multi cluster connectivity, Data Marshall error: %v", err),
		)
		return
	}

	// Call the API to create the multi cluster connectivity
	res, err := clusterAPI.Post(clusterPayload)
	if err != nil {
		dg.AddError(
			"Error Creating Multi Cluster Connectivity",
			fmt.Sprintf("Could not create multi cluster connectivity, unexpected error: %v %v", err, res),
		)
		return
	}
	// Read the created multi cluster connectivity
	// ND BUG - license_tier is not set in the response. Try delaying the read
	time.Sleep(2 * time.Second)

	// Preserve sensitive fields that are not returned by the API
	preservedUsername := input.Username
	preservedPassword := input.Password
	preservedLoginDomain := input.LoginDomain

	r.rscGetMultiClusterConnectivity(ctx, dg, input)

	// Restore sensitive fields after read
	input.Username = preservedUsername
	input.Password = preservedPassword
	input.LoginDomain = preservedLoginDomain

	// Set Id from ClusterName (logic kept outside generated codec)
	setModelId(input)
}

// GetMultiClusterConnectivity retrieves multi cluster connectivity information by name
func (r *multiClusterConnectivityResource) rscGetMultiClusterConnectivity(ctx context.Context, dg *diag.Diagnostics, in *MultiClusterConnectivityModel) {

	clusterAPI := api.NewClusterAPI(nil, r.infraClient.ApiClient)
	clusterAPI.ClusterName = in.ClusterName.ValueString()
	respData, err := clusterAPI.Get()

	if err != nil {
		dg.AddError(
			"Error Reading Multi Cluster Connectivity",
			fmt.Sprintf("Could not read multi cluster connectivity, unexpected error: %v %v", err, respData),
		)
		return
	}

	if clusterAPI.ClusterName == "" {
		var clustersResp map[string][]NDFCMultiClusterConnectivityModel
		err = json.Unmarshal(respData, &clustersResp)
		if err != nil {
			dg.AddError(
				"Error Reading Multi Cluster Connectivity",
				fmt.Sprintf("Could not unmarshal multi cluster connectivity response, unexpected error: %v", err),
			)
			return
		}

		hostname := in.Hostname.ValueString()
		clusterType := in.ClusterType.ValueString()
		for _, cluster := range clustersResp["clusters"] {
			if cluster.Spec.Hostname == hostname && cluster.Spec.ClusterType == clusterType {
				in.SetModelData(&cluster)
				setModelId(in)
				return
			}
		}

		dg.AddError(
			"Error Reading Multi Cluster Connectivity",
			fmt.Sprintf("Could not find cluster with onboardUrl %q and clusterType %q in the response", hostname, clusterType),
		)
		return
	}

	var clusterResp NDFCMultiClusterConnectivityModel
	err = json.Unmarshal(respData, &clusterResp)
	if err != nil {
		dg.AddError(
			"Error Reading Multi Cluster Connectivity",
			fmt.Sprintf("Could not unmarshal multi cluster connectivity response, unexpected error: %v", err),
		)
		return
	}

	in.SetModelData(&clusterResp)
	setModelId(in)
}

// UpdateMultiClusterConnectivity updates a multi cluster connectivity with the provided payload
func (r *multiClusterConnectivityResource) rscUpdateMultiClusterConnectivity(ctx context.Context, dg *diag.Diagnostics, clusterModel *MultiClusterConnectivityModel) {
	inData := clusterModel.GetModelData()
	// log.Printf("Updating multi cluster connectivity %s with category %s", inData.ClusterName, inData.Category)

	clusterAPI := api.NewClusterAPI(nil, r.infraClient.ApiClient)
	clusterAPI.ClusterName = inData.Spec.ClusterName

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating Multi Cluster Connectivity",
			fmt.Sprintf("Could not update multi cluster connectivity, Data Marshall error: %v", err),
		)
		tflog.Error(ctx, "Error Updating Multi Cluster Connectivity", map[string]interface{}{"error": err.Error()})
		return
	}
	res, err := clusterAPI.Put(inDataBytes)
	if err != nil {
		dg.AddError(
			"Error Updating Multi Cluster Connectivity",
			fmt.Sprintf("Could not update multi cluster connectivity, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Updating Multi Cluster Connectivity", map[string]interface{}{"error": err.Error()})
		return
	}
	// Preserve sensitive fields that are not returned by the API
	preservedUsername := clusterModel.Username
	preservedPassword := clusterModel.Password
	preservedLoginDomain := clusterModel.LoginDomain

	// Read the updated multi cluster connectivity
	r.rscGetMultiClusterConnectivity(ctx, dg, clusterModel)

	// Restore sensitive fields after read
	clusterModel.Username = preservedUsername
	clusterModel.Password = preservedPassword
	clusterModel.LoginDomain = preservedLoginDomain

	// Set Id from ClusterName (logic kept outside generated codec)
	setModelId(clusterModel)
	// log.Printf("Updated multi cluster connectivity %s with category %s", inData.ClusterName, inData.Category)

}

// DeleteMultiClusterConnectivity deletes a multi cluster connectivity by name
func (r *multiClusterConnectivityResource) rscDeleteMultiClusterConnectivity(ctx context.Context, dg *diag.Diagnostics, state *MultiClusterConnectivityModel) {
	clusterAPI := api.NewClusterAPI(nil, r.infraClient.ApiClient)
	clusterAPI.ClusterName = state.ClusterName.ValueString()

	// Build the remove payload with credentials and force flag
	removePayload := api.ClusterRemovePayload{
		Credentials: api.ClusterRemoveCredentials{
			User:     state.Username.ValueString(),
			Password: state.Password.ValueString(),
		},
	}
	if !state.LoginDomain.IsNull() && !state.LoginDomain.IsUnknown() {
		removePayload.Credentials.LoginDomain = state.LoginDomain.ValueString()
	}
	if !state.ForceRemove.IsNull() && !state.ForceRemove.IsUnknown() {
		removePayload.Force = state.ForceRemove.ValueBool()
	}

	payload, err := json.Marshal(removePayload)
	if err != nil {
		dg.AddError(
			"Error Deleting Multi Cluster Connectivity",
			fmt.Sprintf("Could not delete multi cluster connectivity, Data Marshall error: %v", err),
		)
		return
	}

	res, err := clusterAPI.PostDelete(payload)
	if err != nil {
		dg.AddError(
			"Error Deleting Multi Cluster Connectivity",
			fmt.Sprintf("Could not delete multi cluster connectivity, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Deleting Multi Cluster Connectivity", map[string]interface{}{"error": err.Error()})
		return
	}
}
