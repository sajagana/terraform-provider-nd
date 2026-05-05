// Code generated;  DO NOT EDIT.

package resource_multi_cluster_connectivity_aci

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCMultiClusterConnectivityAciModel struct {
	Id   string        `json:"-"`
	Spec NDFCSpecValue `json:"spec,omitempty"`
}

type NDFCSpecValue struct {
	ClusterType string                   `json:"clusterType,omitempty"`
	ClusterName string                   `json:"name,omitempty"`
	HostName    string                   `json:"onboardUrl,omitempty"`
	Credentials NDFCSpecCredentialsValue `json:"credentials,omitempty"`
	Location    NDFCSpecLocationValue    `json:"location,omitempty"`
	Aci         NDFCSpecAciValue         `json:"aci,omitempty"`
}

type NDFCSpecCredentialsValue struct {
	UserName    string `json:"user,omitempty"`
	Password    string `json:"password,omitempty"`
	LoginDomain string `json:"loginDomain,omitempty"`
}

type NDFCSpecLocationValue struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type NDFCSpecAciValue struct {
	FabricName              string                    `json:"name,omitempty"`
	LicenseTier             string                    `json:"licenseTier,omitempty"`
	SecurityDomain          string                    `json:"securityDomain,omitempty"`
	ValidatePeerCertificate string                    `json:"verifyCA,omitempty"`
	Orchestration           NDFCAciOrchestrationValue `json:"orchestration,omitempty"`
	Telemetry               NDFCAciTelemetryValue     `json:"telemetry,omitempty"`
}

type NDFCAciOrchestrationValue struct {
	EnableOrchestration string `json:"status,omitempty"`
}

type NDFCAciTelemetryValue struct {
	EnableTelemetry     string `json:"status,omitempty"`
	TelemetryCollection string `json:"network,omitempty"`
	TelemetryStreaming  string `json:"streamingProtocol,omitempty"`
	TelemetryEpg        string `json:"epg,omitempty"`
}

func (v *MultiClusterConnectivityAciModel) SetModelData(jsonData *NDFCMultiClusterConnectivityAciModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Id != "" {
		v.Id = types.StringValue(jsonData.Id)
	} else {
		v.Id = types.StringNull()
	}

	if jsonData.Spec.ClusterType != "" {
		v.ClusterType = types.StringValue(jsonData.Spec.ClusterType)

	} else {
		v.ClusterType = types.StringNull()
	}

	if jsonData.Spec.ClusterName != "" {
		v.ClusterName = types.StringValue(jsonData.Spec.ClusterName)

	} else {
		v.ClusterName = types.StringNull()
	}

	if jsonData.Spec.HostName != "" {
		v.HostName = types.StringValue(jsonData.Spec.HostName)

	} else {
		v.HostName = types.StringNull()
	}

	if jsonData.Spec.Credentials.UserName != "" {
		v.UserName = types.StringValue(jsonData.Spec.Credentials.UserName)

	} else {
		v.UserName = types.StringNull()
	}

	if jsonData.Spec.Credentials.Password != "" {
		v.Password = types.StringValue(jsonData.Spec.Credentials.Password)

	} else {
		v.Password = types.StringNull()
	}

	if jsonData.Spec.Credentials.LoginDomain != "" {
		v.LoginDomain = types.StringValue(jsonData.Spec.Credentials.LoginDomain)

	} else {
		v.LoginDomain = types.StringNull()
	}

	if jsonData.Spec.Location.Latitude != nil {
		v.Latitude = types.Float64Value(float64(*jsonData.Spec.Location.Latitude))

	} else {
		v.Latitude = types.Float64Null()
	}

	if jsonData.Spec.Location.Longitude != nil {
		v.Longitude = types.Float64Value(float64(*jsonData.Spec.Location.Longitude))

	} else {
		v.Longitude = types.Float64Null()
	}

	if jsonData.Spec.Aci.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.Spec.Aci.FabricName)

	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.Spec.Aci.LicenseTier != "" {
		v.LicenseTier = types.StringValue(jsonData.Spec.Aci.LicenseTier)

	} else {
		v.LicenseTier = types.StringNull()
	}

	if jsonData.Spec.Aci.SecurityDomain != "" {
		v.SecurityDomain = types.StringValue(jsonData.Spec.Aci.SecurityDomain)

	} else {
		v.SecurityDomain = types.StringNull()
	}

	if jsonData.Spec.Aci.ValidatePeerCertificate != "" {
		x, _ := strconv.ParseBool(jsonData.Spec.Aci.ValidatePeerCertificate)
		v.ValidatePeerCertificate = types.BoolValue(x)

	} else {
		v.ValidatePeerCertificate = types.BoolNull()
	}

	if jsonData.Spec.Aci.Orchestration.EnableOrchestration != "" {
		v.EnableOrchestration = types.StringValue(jsonData.Spec.Aci.Orchestration.EnableOrchestration)

	} else {
		v.EnableOrchestration = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.EnableTelemetry != "" {
		v.EnableTelemetry = types.StringValue(jsonData.Spec.Aci.Telemetry.EnableTelemetry)

	} else {
		v.EnableTelemetry = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryCollection != "" {
		v.TelemetryCollection = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryCollection)

	} else {
		v.TelemetryCollection = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryStreaming != "" {
		v.TelemetryStreaming = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryStreaming)

	} else {
		v.TelemetryStreaming = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryEpg != "" {
		v.TelemetryEpg = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryEpg)

	} else {
		v.TelemetryEpg = types.StringNull()
	}

	return err
}

func (v MultiClusterConnectivityAciModel) GetModelData() *NDFCMultiClusterConnectivityAciModel {
	var data = new(NDFCMultiClusterConnectivityAciModel)

	//MARSHAL_BODY

	if !v.ClusterType.IsNull() && !v.ClusterType.IsUnknown() {
		data.Spec.ClusterType = v.ClusterType.ValueString()
	} else {
		data.Spec.ClusterType = ""
	}

	if !v.ClusterName.IsNull() && !v.ClusterName.IsUnknown() {
		data.Spec.ClusterName = v.ClusterName.ValueString()
	} else {
		data.Spec.ClusterName = ""
	}

	if !v.HostName.IsNull() && !v.HostName.IsUnknown() {
		data.Spec.HostName = v.HostName.ValueString()
	} else {
		data.Spec.HostName = ""
	}

	if !v.UserName.IsNull() && !v.UserName.IsUnknown() {
		data.Spec.Credentials.UserName = v.UserName.ValueString()
	} else {
		data.Spec.Credentials.UserName = ""
	}

	if !v.Password.IsNull() && !v.Password.IsUnknown() {
		data.Spec.Credentials.Password = v.Password.ValueString()
	} else {
		data.Spec.Credentials.Password = ""
	}

	if !v.LoginDomain.IsNull() && !v.LoginDomain.IsUnknown() {
		data.Spec.Credentials.LoginDomain = v.LoginDomain.ValueString()
	} else {
		data.Spec.Credentials.LoginDomain = ""
	}

	if !v.Latitude.IsNull() && !v.Latitude.IsUnknown() {
		data.Spec.Location.Latitude = new(float64)
		*data.Spec.Location.Latitude = v.Latitude.ValueFloat64()
	} else {
		data.Spec.Location.Latitude = nil
	}

	if !v.Longitude.IsNull() && !v.Longitude.IsUnknown() {
		data.Spec.Location.Longitude = new(float64)
		*data.Spec.Location.Longitude = v.Longitude.ValueFloat64()
	} else {
		data.Spec.Location.Longitude = nil
	}

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.Spec.Aci.FabricName = v.FabricName.ValueString()
	} else {
		data.Spec.Aci.FabricName = ""
	}

	if !v.LicenseTier.IsNull() && !v.LicenseTier.IsUnknown() {
		data.Spec.Aci.LicenseTier = v.LicenseTier.ValueString()
	} else {
		data.Spec.Aci.LicenseTier = ""
	}

	if !v.SecurityDomain.IsNull() && !v.SecurityDomain.IsUnknown() {
		data.Spec.Aci.SecurityDomain = v.SecurityDomain.ValueString()
	} else {
		data.Spec.Aci.SecurityDomain = ""
	}

	if !v.ValidatePeerCertificate.IsNull() && !v.ValidatePeerCertificate.IsUnknown() {
		data.Spec.Aci.ValidatePeerCertificate = strconv.FormatBool(v.ValidatePeerCertificate.ValueBool())
	} else {
		data.Spec.Aci.ValidatePeerCertificate = ""
	}

	if !v.EnableOrchestration.IsNull() && !v.EnableOrchestration.IsUnknown() {
		data.Spec.Aci.Orchestration.EnableOrchestration = v.EnableOrchestration.ValueString()
	} else {
		data.Spec.Aci.Orchestration.EnableOrchestration = ""
	}

	if !v.EnableTelemetry.IsNull() && !v.EnableTelemetry.IsUnknown() {
		data.Spec.Aci.Telemetry.EnableTelemetry = v.EnableTelemetry.ValueString()
	} else {
		data.Spec.Aci.Telemetry.EnableTelemetry = ""
	}

	if !v.TelemetryCollection.IsNull() && !v.TelemetryCollection.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryCollection = v.TelemetryCollection.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryCollection = ""
	}

	if !v.TelemetryStreaming.IsNull() && !v.TelemetryStreaming.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryStreaming = v.TelemetryStreaming.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryStreaming = ""
	}

	if !v.TelemetryEpg.IsNull() && !v.TelemetryEpg.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryEpg = v.TelemetryEpg.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryEpg = ""
	}

	return data
}
