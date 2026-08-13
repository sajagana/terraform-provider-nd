package resource_fabric_aci

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/common/utils"
	"terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type aciClusterResponse struct {
	Spec   aciClusterResponseSpec   `json:"spec,omitempty"`
	Status aciClusterResponseStatus `json:"status,omitempty"`
}

type aciClusterResponseSpec struct {
	NDFCSpecValue
	FabricName string `json:"name,omitempty"`
}

type aciClusterResponseStatus struct {
	LastUpdate aciClusterResponseLastUpdate `json:"lastUpdate,omitempty"`
}

type aciClusterResponseLastUpdate struct {
	Message string `json:"message,omitempty"`
}

type aciClusterCreatePayload struct {
	Spec NDFCSpecValue `json:"spec,omitempty"`
}

type aciClusterRemovePayload struct {
	Credentials NDFCSpecCredentialsValue `json:"credentials"`
	Force       bool                     `json:"force"`
}

type aciClusterDeRegisterPayload struct {
	Credentials NDFCSpecCredentialsValue `json:"credentials"`
}

type aciFabricReRegisterPayload struct {
	ReRegister aciFabricReRegisterValue `json:"reRegister"`
}

type aciFabricReRegisterValue struct {
	Type        string                   `json:"type"`
	OnboardURL  string                   `json:"onboardUrl"`
	Credentials NDFCSpecCredentialsValue `json:"credentials"`
	VerifyCA    bool                     `json:"verifyCA"`
}

const (
	fabricAciReachableMessage     = "Fabric reachable"
	fabricAciRetryInterval        = 10 * time.Second
	fabricAciReachabilityTimeout  = 10 * time.Minute
	fabricAciDefaultUpdateRetries = 10
)

func fabricAciUpdateRetryCount(fabricName string) (int, error) {
	environmentVariable := fabricAciEnvironmentVariableName(fabricName, "UPDATE_RETRY")
	value, ok := os.LookupEnv(environmentVariable)
	if !ok {
		return fabricAciDefaultUpdateRetries, nil
	}

	retryCount, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil || retryCount < 0 {
		return 0, fmt.Errorf(
			"environment variable %s must contain a non-negative integer retry count, got %q",
			environmentVariable,
			value,
		)
	}

	return retryCount, nil
}

func (r aciClusterResponse) modelData() NDFCFabricAciModel {
	data := NDFCFabricAciModel{
		Spec: r.Spec.NDFCSpecValue,
	}
	if data.Spec.Aci.FabricName == "" {
		data.Spec.Aci.FabricName = r.Spec.FabricName
	}
	return data
}

// rscCreateFabricAci creates a fabric ACI resource.
func (r *fabricAciResource) rscCreateFabricAci(ctx context.Context, dg *diag.Diagnostics, fabricAciModel *FabricAciModel) {
	id := fabricAciModel.Id.ValueString()
	log.Printf("[INFO] Create nd_fabric_aci id=%s", id)

	inData := fabricAciModel.GetModelData()
	inData.Spec.ClusterType = "APIC"
	createPayload := aciClusterCreatePayload{Spec: inData.Spec}

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)

	clusterPayload, err := json.Marshal(createPayload)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric ACI",
			fmt.Sprintf("Could not create fabric ACI, Data Marshall error: %v", err),
		)
		return
	}

	res, err := clusterAPI.Post(clusterPayload, &ndapi.APIOptions{DisablePayloadLog: false})
	if err != nil {
		dg.AddError(
			"Error Creating Fabric ACI",
			fmt.Sprintf("Could not create fabric ACI, unexpected error: %v %v", err, res),
		)
		return
	}

	if r.rscGetFabricAci(ctx, dg, fabricAciModel) && !dg.HasError() {
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not read nd_fabric_aci %q after create: resource not found", id),
		)
	}
}

// rscGetFabricAci retrieves fabric ACI information by id.
// It returns true when the remote object was not found.
func (r *fabricAciResource) rscGetFabricAci(ctx context.Context, dg *diag.Diagnostics, fabricAciModel *FabricAciModel) bool {
	id := fabricAciModel.Id.ValueString()
	log.Printf("[INFO] Read nd_fabric_aci id=%s", id)

	preservedUsername := fabricAciModel.Username
	preservedPassword := fabricAciModel.Password
	preservedLoginDomain := fabricAciModel.LoginDomain
	preservedDeRegister := fabricAciModel.DeRegister
	preservedReRegister := fabricAciModel.ReRegister
	if preservedLoginDomain.IsUnknown() {
		preservedLoginDomain = types.StringNull()
	}
	if preservedDeRegister.IsUnknown() {
		preservedDeRegister = types.BoolNull()
	}
	if preservedReRegister.IsUnknown() {
		preservedReRegister = types.BoolNull()
	}

	preservedVerifyCa := fabricAciModel.VerifyCa
	if preservedVerifyCa.IsUnknown() {
		preservedVerifyCa = types.BoolNull()
	}

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = id
	respData, err := clusterAPI.Get()

	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return true
		}
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not read fabric ACI, unexpected error: %v %v", err, respData),
		)
		return false
	}
	if respData == nil {
		log.Printf("[WARN] nd_fabric_aci id=%s not found: empty response", id)
		return true
	}

	var clusterResp aciClusterResponse
	err = json.Unmarshal(respData, &clusterResp)
	if err != nil {
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not unmarshal fabric ACI response, unexpected error: %v", err),
		)
		return false
	}

	modelData := clusterResp.modelData()
	dg.Append(fabricAciModel.SetModelData(&modelData)...)
	if dg.HasError() {
		return false
	}
	fabricAciModel.NormalizeTelemetryNetworkState()

	fabricAciModel.Username = preservedUsername
	fabricAciModel.Password = preservedPassword
	fabricAciModel.LoginDomain = preservedLoginDomain
	fabricAciModel.VerifyCa = preservedVerifyCa
	fabricAciModel.DeRegister = preservedDeRegister
	fabricAciModel.ReRegister = preservedReRegister
	fabricAciModel.Id = types.StringValue(id)
	return false
}

// rscUpdateFabricAci updates a fabric ACI resource.
func (r *fabricAciResource) rscUpdateFabricAci(ctx context.Context, dg *diag.Diagnostics, plan *FabricAciModel, state *FabricAciModel) {
	id := plan.Id.ValueString()
	log.Printf("[INFO] Update nd_fabric_aci id=%s", id)

	// Compare the endpoint-specific payloads so action-only changes do not also
	// send an unchanged PUT /manage/fabrics/{id} request.
	normalUpdateRequired := !reflect.DeepEqual(plan.GetManageModelData(), state.GetManageModelData())
	deRegisterRequested := plan.DeRegister.ValueBool()
	reRegisterRequested := plan.ReRegister.ValueBool()

	// Ignore the action flags when determining whether any ordinary Terraform
	// attribute changed. De-registration must be the only operation in an update.
	planWithoutActions := *plan
	stateWithoutActions := *state
	planWithoutActions.DeRegister = types.BoolNull()
	stateWithoutActions.DeRegister = types.BoolNull()
	planWithoutActions.ReRegister = types.BoolNull()
	stateWithoutActions.ReRegister = types.BoolNull()
	otherAttributesChanged := !reflect.DeepEqual(planWithoutActions, stateWithoutActions)

	if deRegisterRequested {
		if otherAttributesChanged {
			dg.AddError(
				"Invalid Fabric ACI De-registration Update",
				"The de_register action cannot be combined with changes to other nd_fabric_aci attributes. Apply the other changes separately.",
			)
			return
		}

		r.rscDeRegisterFabricAci(dg, plan)
		return
	}

	if reRegisterRequested {
		r.rscReRegisterFabricAci(dg, plan)
		if dg.HasError() {
			return
		}

		if !r.waitForFabricAciReachable(ctx, dg, plan) {
			return
		}

		if !normalUpdateRequired {
			log.Printf("[DEBUG] Skipping normal update for nd_fabric_aci id=%s: Manage payload is unchanged", id)
			if r.rscGetFabricAci(ctx, dg, plan) && !dg.HasError() {
				dg.AddError(
					"Error Reading Fabric ACI",
					fmt.Sprintf("Could not read nd_fabric_aci %q after re-registration: resource not found", id),
				)
			}
			return
		}
	}

	actionResetOnly := !reRegisterRequested && !otherAttributesChanged &&
		(state.DeRegister.ValueBool() || state.ReRegister.ValueBool())
	if actionResetOnly {
		log.Printf("[DEBUG] Skipping backend update for nd_fabric_aci id=%s: only an action attribute changed from true to false", id)
		return
	}

	inData := plan.GetManageModelData()

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = id

	inDataBytes, err := json.Marshal(inData)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric ACI",
			fmt.Sprintf("Could not update fabric ACI, Data Marshall error: %v", err),
		)
		log.Printf("[ERROR] Error Updating Fabric ACI id=%s: error=%s", id, err.Error())
		return
	}

	retryCount, err := fabricAciUpdateRetryCount(id)
	if err != nil {
		dg.AddError("Invalid Fabric ACI Update Retry Environment Variable", err.Error())
		return
	}

	// A fabric can report "Fabric reachable" before all ND feature services are
	// ready after re-registration. Retry transient HTTP 500 responses so an
	// immediate orchestration, telemetry, or other fabric update can complete.
	// Sample intermediate error message: "Orchestration failure: Cannot connect to fabric"
	res, err := clusterAPI.Put(inDataBytes, nil)
	retriesPerformed := 0
	for err != nil && strings.Contains(err.Error(), "StatusCode 500") && retriesPerformed < retryCount {
		retriesPerformed++
		log.Printf(
			"[WARN] Update nd_fabric_aci id=%s returned HTTP 500: response_message=%q; retrying in %s (retry %d/%d)",
			id,
			res.Get("message").String(),
			fabricAciRetryInterval,
			retriesPerformed,
			retryCount,
		)

		retryTimer := time.NewTimer(fabricAciRetryInterval)
		select {
		case <-ctx.Done():
			retryTimer.Stop()
			dg.AddError(
				"Error Updating Fabric ACI",
				fmt.Sprintf(
					"Could not update fabric ACI %q: context canceled while waiting for retry %d/%d: %v",
					id,
					retriesPerformed,
					retryCount,
					ctx.Err(),
				),
			)
			return
		case <-retryTimer.C:
		}

		res, err = clusterAPI.Put(inDataBytes, nil)
	}
	if err != nil {
		errorDetail := fmt.Sprintf("Could not update fabric ACI, unexpected error: %v %v", err, res)
		if retriesPerformed > 0 {
			errorDetail = fmt.Sprintf(
				"Could not update fabric ACI after %d retries, unexpected error: %v %v",
				retriesPerformed,
				err,
				res,
			)
		}
		dg.AddError(
			"Error Updating Fabric ACI",
			errorDetail,
		)
		log.Printf("[ERROR] Error Updating Fabric ACI id=%s: error=%s", id, err.Error())
		return
	}
	if retriesPerformed > 0 {
		log.Printf("[INFO] Update nd_fabric_aci id=%s succeeded after %d retries", id, retriesPerformed)
	}

	if r.rscGetFabricAci(ctx, dg, plan) && !dg.HasError() {
		dg.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not read nd_fabric_aci %q after update: resource not found", id),
		)
	}
}

// rscDeRegisterFabricAci de-registers the local Nexus Dashboard cluster from an ACI fabric.
func (r *fabricAciResource) rscDeRegisterFabricAci(dg *diag.Diagnostics, fabricAciModel *FabricAciModel) {
	id := fabricAciModel.Id.ValueString()
	log.Printf("[INFO] De-register nd_fabric_aci id=%s", id)

	modelData := fabricAciModel.GetModelData()
	deRegisterPayload := aciClusterDeRegisterPayload{
		Credentials: modelData.Spec.Credentials,
	}

	payload, err := json.Marshal(deRegisterPayload)
	if err != nil {
		dg.AddError(
			"Error De-registering Fabric ACI",
			fmt.Sprintf("Could not de-register fabric ACI, Data Marshall error: %v", err),
		)
		return
	}

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = id
	clusterAPI.DeRegister = true

	res, err := clusterAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		dg.AddError(
			"Error De-registering Fabric ACI",
			fmt.Sprintf("Could not de-register fabric ACI, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error De-registering Fabric ACI id=%s: error=%s", id, err.Error())
	}
}

// waitForFabricAciReachable prevents a follow-up update while re-registration
// is still reconnecting ND to the APIC. It waits until ND reports
// status.lastUpdate.message as "Fabric reachable"; the update retry handles the
// remaining interval in which individual ND feature services may not be ready.
func (r *fabricAciResource) waitForFabricAciReachable(ctx context.Context, dg *diag.Diagnostics, fabricAciModel *FabricAciModel) bool {
	id := fabricAciModel.Id.ValueString()
	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = id

	err := utils.PollUntil(ctx, fabricAciRetryInterval, fabricAciReachabilityTimeout, func(_ context.Context) (bool, error) {
		respData, err := clusterAPI.Get()
		if err != nil {
			return false, fmt.Errorf("could not read fabric ACI cluster %q while waiting for re-registration: %w", id, err)
		}
		if respData == nil {
			return false, nil
		}

		var clusterResp aciClusterResponse
		if err := json.Unmarshal(respData, &clusterResp); err != nil {
			return false, fmt.Errorf("could not unmarshal fabric ACI cluster %q status while waiting for re-registration: %w", id, err)
		}

		message := clusterResp.Status.LastUpdate.Message
		log.Printf("[DEBUG] Waiting for nd_fabric_aci id=%s to become reachable: last_update_message=%q", id, message)
		return message == fabricAciReachableMessage, nil
	})

	if err == nil {
		return true
	}

	if errors.Is(err, utils.ErrPollTimeout) {
		dg.AddError(
			"Error Re-registering Fabric ACI",
			fmt.Sprintf("Timed out after %s waiting for fabric ACI cluster %q to report status.lastUpdate.message %q.", fabricAciReachabilityTimeout, id, fabricAciReachableMessage),
		)
		return false
	}

	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		dg.AddError(
			"Error Re-registering Fabric ACI",
			fmt.Sprintf("Context canceled while waiting for fabric ACI cluster %q to become reachable: %v", id, err),
		)
		return false
	}

	dg.AddError("Error Re-registering Fabric ACI", err.Error())
	return false
}

// rscReRegisterFabricAci re-registers a fabric ACI resource with Nexus Dashboard.
func (r *fabricAciResource) rscReRegisterFabricAci(dg *diag.Diagnostics, fabricAciModel *FabricAciModel) {
	id := fabricAciModel.Id.ValueString()
	log.Printf("[INFO] Re-register nd_fabric_aci id=%s", id)

	reRegisterPayload := aciFabricReRegisterPayload{
		ReRegister: aciFabricReRegisterValue{
			Type:       "aci",
			OnboardURL: fabricAciModel.Hostname.ValueString(),
			Credentials: NDFCSpecCredentialsValue{
				Username:    fabricAciModel.Username.ValueString(),
				Password:    fabricAciModel.Password.ValueString(),
				LoginDomain: fabricAciModel.LoginDomain.ValueString(),
			},
			VerifyCA: fabricAciModel.VerifyCa.ValueBool(),
		},
	}

	payload, err := json.Marshal(reRegisterPayload)
	if err != nil {
		dg.AddError(
			"Error Re-registering Fabric ACI",
			fmt.Sprintf("Could not re-register fabric ACI, Data Marshall error: %v", err),
		)
		return
	}

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = id
	clusterAPI.ReRegister = true

	res, err := clusterAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: false})
	if err != nil {
		dg.AddError(
			"Error Re-registering Fabric ACI",
			fmt.Sprintf("Could not re-register fabric ACI, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Re-registering Fabric ACI id=%s: error=%s", id, err.Error())
		return
	}
}

// rscDeleteFabricAci deletes a fabric ACI resource by id.
func (r *fabricAciResource) rscDeleteFabricAci(_ context.Context, dg *diag.Diagnostics, fabricAciModel *FabricAciModel, force bool) {
	id := fabricAciModel.Id.ValueString()
	log.Printf("[INFO] Delete nd_fabric_aci id=%s", id)

	clusterAPI := api.NewFabricAciAPI(r.manageClient.ApiClient, ndapi.DefaultFabric)
	clusterAPI.ClusterName = id
	clusterAPI.Delete = true

	modelData := fabricAciModel.GetModelData()
	removePayload := aciClusterRemovePayload{
		Credentials: modelData.Spec.Credentials,
		Force:       force,
	}

	payload, err := json.Marshal(removePayload)
	if err != nil {
		dg.AddError(
			"Error Deleting Fabric ACI",
			fmt.Sprintf("Could not delete fabric ACI, Data Marshall error: %v", err),
		)
		return
	}

	res, err := clusterAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: false})
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		dg.AddError(
			"Error Deleting Fabric ACI",
			fmt.Sprintf("Could not delete fabric ACI, unexpected error: %v %v", err, res),
		)
		log.Printf("[ERROR] Error Deleting Fabric ACI id=%s: error=%s", id, err.Error())
		return
	}
}
