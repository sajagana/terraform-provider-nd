package resource_multi_cluster_connectivity_aci

import (
	"context"
	"encoding/json"
	"fmt"
	"terraform-provider-nd/internal/infra/api"

	"log"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// setModelId sets the Id field on the model based on ClusterName.
// This is kept outside resource_codec_gen.go to avoid conflicts with the internal generator.
func setModelId(model *MultiClusterConnectivityAciModel) {
	if !model.ClusterName.IsNull() && !model.ClusterName.IsUnknown() {
		model.Id = types.StringValue(model.ClusterName.ValueString())
	} else {
		model.Id = types.StringNull()
	}
}

// rscCreateMultiClusterConnectivityAci creates a multi cluster connectivity aci resource
func (r *multiClusterConnectivityAciResource) rscCreateMultiClusterConnectivityAci(ctx context.Context, dg *diag.Diagnostics, input *MultiClusterConnectivityAciModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	inData := input.GetModelData()

	// Create multi cluster connectivity aci API client
	clusterAPI := api.NewClusterACIAPI(nil, r.infraClient.ApiClient)

	// Convert model data to JSON
	clusterPayload, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Multi Cluster Connectivity ACI",
			fmt.Sprintf("Could not create multi cluster connectivity aci, Data Marshall error: %v", err),
		)
		return
	}

	// Call the API to create the multi cluster connectivity aci
	res, err := clusterAPI.Post(clusterPayload)
	if err != nil {
		dg.AddError(
			"Error Creating Multi Cluster Connectivity ACI",
			fmt.Sprintf("Could not create multi cluster connectivity aci, unexpected error: %v %v", err, res),
		)
		return
	}

	r.rscGetMultiClusterConnectivityAci(ctx, dg, input)

	// Set Id from ClusterName (logic kept outside generated codec)
	setModelId(input)
}

// rscGetMultiClusterConnectivityAci retrieves multi cluster connectivity aci information by name
func (r *multiClusterConnectivityAciResource) rscGetMultiClusterConnectivityAci(ctx context.Context, dg *diag.Diagnostics, in *MultiClusterConnectivityAciModel) {

	// Preserve sensitive fields that are not returned by the API
	preservedUserName := in.UserName
	preservedPassword := in.Password
	preservedLoginDomain := in.LoginDomain

	clusterAPI := api.NewClusterACIAPI(nil, r.infraClient.ApiClient)
	clusterAPI.ClusterName = in.ClusterName.ValueString()
	respData, err := clusterAPI.Get()

	if err != nil {
		dg.AddError(
			"Error Reading Multi Cluster Connectivity ACI",
			fmt.Sprintf("Could not read multi cluster connectivity aci, unexpected error: %v %v", err, respData),
		)
		return
	}

	if clusterAPI.ClusterName == "" {
		var clustersResp map[string][]NDFCMultiClusterConnectivityAciModel
		err = json.Unmarshal(respData, &clustersResp)
		if err != nil {
			dg.AddError(
				"Error Reading Multi Cluster Connectivity ACI",
				fmt.Sprintf("Could not unmarshal multi cluster connectivity aci response, unexpected error: %v", err),
			)
			return
		}

		hostName := in.HostName.ValueString()
		clusterType := in.ClusterType.ValueString()
		for _, cluster := range clustersResp["clusters"] {
			if cluster.Spec.HostName == hostName && cluster.Spec.ClusterType == clusterType {
				in.SetModelData(&cluster)
				in.UserName = preservedUserName
				in.Password = preservedPassword
				in.LoginDomain = preservedLoginDomain
				setModelId(in)
				return
			}
		}

		dg.AddError(
			"Error Reading Multi Cluster Connectivity ACI",
			fmt.Sprintf("Could not find cluster with onboardUrl %q and clusterType %q in the response", hostName, clusterType),
		)
		return
	}

	var clusterResp NDFCMultiClusterConnectivityAciModel
	err = json.Unmarshal(respData, &clusterResp)
	if err != nil {
		dg.AddError(
			"Error Reading Multi Cluster Connectivity ACI",
			fmt.Sprintf("Could not unmarshal multi cluster connectivity aci response, unexpected error: %v", err),
		)
		return
	}

	in.SetModelData(&clusterResp)

	// Restore sensitive fields after SetModelData (API does not return them)
	in.UserName = preservedUserName
	in.Password = preservedPassword
	in.LoginDomain = preservedLoginDomain

	setModelId(in)
}

// updateSpecValue extends NDFCSpecValue with fields only needed during update.
type updateSpecValue struct {
	NDFCSpecValue
	ReRegister *bool `json:"reRegister,omitempty"`
}

// updatePayload wraps the update spec for JSON marshalling.
type updatePayload struct {
	Spec updateSpecValue `json:"spec,omitempty"`
}

// rscUpdateMultiClusterConnectivityAci updates a multi cluster connectivity aci with the provided payload
func (r *multiClusterConnectivityAciResource) rscUpdateMultiClusterConnectivityAci(ctx context.Context, dg *diag.Diagnostics, clusterModel *MultiClusterConnectivityAciModel) {
	inData := clusterModel.GetModelData()

	clusterAPI := api.NewClusterACIAPI(nil, r.infraClient.ApiClient)
	clusterAPI.ClusterName = clusterModel.ClusterName.ValueString()

	// This is only used for the update operation and not for create, as create will register the cluster for the first time.
	// For update, we want to ensure that the changes are applied by re-registering the cluster with the new details.
	reRegister := true
	payload := updatePayload{
		Spec: updateSpecValue{
			NDFCSpecValue: inData.Spec,
			ReRegister:    &reRegister,
		},
	}

	inDataBytes, err := json.Marshal(payload)
	if err != nil {
		dg.AddError(
			"Error Updating Multi Cluster Connectivity ACI",
			fmt.Sprintf("Could not update multi cluster connectivity aci, Data Marshall error: %v", err),
		)
		log.Printf("[ERROR] Error Updating Multi Cluster Connectivity ACI: error=%s", err.Error())
		return
	}
	res, err := clusterAPI.Put(inDataBytes)

	if err != nil {
		dg.AddError(
			"Error Updating Multi Cluster Connectivity ACI",
			fmt.Sprintf("Could not update multi cluster connectivity aci, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Updating Multi Cluster Connectivity ACI: error=%s", err.Error())
		return
	}
	// Read the updated multi cluster connectivity aci
	r.rscGetMultiClusterConnectivityAci(ctx, dg, clusterModel)

	// Set Id from ClusterName (logic kept outside generated codec)
	setModelId(clusterModel)

}

// rscDeleteMultiClusterConnectivityAci deletes a multi cluster connectivity aci by name
func (r *multiClusterConnectivityAciResource) rscDeleteMultiClusterConnectivityAci(ctx context.Context, dg *diag.Diagnostics, state *MultiClusterConnectivityAciModel) {
	clusterAPI := api.NewClusterACIAPI(nil, r.infraClient.ApiClient)
	clusterAPI.ClusterName = state.ClusterName.ValueString()

	// Build the remove payload with credentials and force flag
	removePayload := api.ClusterRemovePayload{
		Credentials: api.ClusterRemoveCredentials{
			User:     state.UserName.ValueString(),
			Password: state.Password.ValueString(),
		},
	}

	if !state.LoginDomain.IsNull() && !state.LoginDomain.IsUnknown() {
		removePayload.Credentials.LoginDomain = state.LoginDomain.ValueString()
	}

	payload, err := json.Marshal(removePayload)
	if err != nil {
		dg.AddError(
			"Error Deleting Multi Cluster Connectivity ACI",
			fmt.Sprintf("Could not delete multi cluster connectivity aci, Data Marshall error: %v", err),
		)
		return
	}

	res, err := clusterAPI.PostDelete(payload)
	if err != nil {
		dg.AddError(
			"Error Deleting Multi Cluster Connectivity ACI",
			fmt.Sprintf("Could not delete multi cluster connectivity aci, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting Multi Cluster Connectivity ACI: error=%s", err.Error())
		return
	}
}
