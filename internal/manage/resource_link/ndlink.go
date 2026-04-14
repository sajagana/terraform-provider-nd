package resource_link

import (
	"context"
	"encoding/json"
	"fmt"
	"terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RscCreateLink creates a link resource using the Link model
func (r *linkResource) rscCreateLink(ctx context.Context, dg *diag.Diagnostics, input *LinkModel) {
	if input == nil {
		dg.AddError(
			"Invalid Input",
			"The input model is nil",
		)
		return
	}

	inData := input.GetModelData()

	// Create link API client
	linkAPI := api.NewLinkAPI(nil, r.manageClient.ApiClient)

	// Build filtered JSON payload with only templateInputs fields applicable to the policyType
	linkPayload, err := BuildFilteredCreatePayload(inData)
	if err != nil {
		dg.AddError(
			"Error Creating Link",
			fmt.Sprintf("Could not create link, Data Marshall error: %v", err),
		)
		return
	}

	// Call the API to create the link
	res, err := linkAPI.Post(linkPayload)
	if err != nil {
		dg.AddError(
			"Error Creating Link",
			fmt.Sprintf("Could not create link, unexpected error: %v %v", err, res),
		)
		return
	}

	// Extract linkId from POST response: {"links":[{"linkId":"LINK-UUID-...","message":"...","status":"success"}]}
	linkId := res.Get("links.0.linkId").String()
	if linkId == "" {
		dg.AddError(
			"Error Creating Link",
			"Could not extract linkId from API response",
		)
		return
	}
	input.LinkId = types.StringValue(linkId)

	r.rscGetLink(ctx, dg, input)

}

// GetLink retrieves link information by link ID
func (r *linkResource) rscGetLink(ctx context.Context, dg *diag.Diagnostics, in *LinkModel) {

	linkAPI := api.NewLinkAPI(nil, r.manageClient.ApiClient)
	linkAPI.LinkId = in.LinkId.ValueString()
	respData, err := linkAPI.Get()
	if err != nil {
		dg.AddError(
			"Error Creating Link",
			fmt.Sprintf("Could not read link, unexpected error: %v %v", err, respData),
		)
		return
	}
	var outData NDFCLinkModel

	err = json.Unmarshal(respData, &outData)
	if err != nil {
		dg.AddError(
			"Error Creating Link",
			fmt.Sprintf("Could not read link, unexpected error: %v %v", err, respData),
		)
		return
	}
	// log.Printf("Location = %v %v", *outData.Location.Latitude, *outData.Location.Longitude)
	// log.Printf("Netflow = %v", *outData.Management.NetflowSettings.NetflowEnable)
	saved := *in
	in.SetModelData(&outData)
	in.PreserveNonPolicyDefaults(outData.ConfigData.PolicyType, &saved)
	// log.Printf("Location from model=%v,%v", in.Location.Latitude.ValueFloat64(), in.Location.Longitude.ValueFloat64())
}

// UpdateLink updates a link with the provided payload
func (r *linkResource) rscUpdateLink(ctx context.Context, dg *diag.Diagnostics, linkModel *LinkModel) {
	inData := linkModel.GetModelData()

	linkAPI := api.NewLinkAPI(nil, r.manageClient.ApiClient)
	linkAPI.LinkId = inData.LinkId

	// Build filtered JSON payload with only templateInputs fields applicable to the policyType
	inDataBytes, err := BuildFilteredPayload(inData)
	if err != nil {
		dg.AddError(
			"Error Updating Link",
			fmt.Sprintf("Could not update link, Data Marshall error: %v", err),
		)
		tflog.Error(ctx, "Error Updating Link", map[string]interface{}{"error": err.Error()})
		return
	}
	res, err := linkAPI.Put(inDataBytes)
	if err != nil {
		dg.AddError(
			"Error Updating Link",
			fmt.Sprintf("Could not update link, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Updating Link", map[string]interface{}{"error": err.Error()})
		return
	}
	// Read the updated link
	r.rscGetLink(ctx, dg, linkModel)
	// log.Printf("Updated link %s with category %s", inData.LinkId, inData.Category)

}

// DeleteLink deletes a link by link ID
func (r *linkResource) rscDeleteLink(ctx context.Context, dg *diag.Diagnostics, linkId string) {
	linkAPI := api.NewLinkAPI(nil, r.manageClient.ApiClient)
	linkAPI.LinkId = linkId
	res, err := linkAPI.Delete()
	if err != nil {
		dg.AddError(
			"Error Deleting Link",
			fmt.Sprintf("Could not delete link, unexpected error: %v %v", err, res),
		)
		tflog.Error(ctx, "Error Deleting Link", map[string]interface{}{"error": err.Error()})
		return
	}
}
