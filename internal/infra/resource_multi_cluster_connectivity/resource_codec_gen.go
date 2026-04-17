// Code generated;  DO NOT EDIT.

package resource_multi_cluster_connectivity

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCMultiClusterConnectivityModel struct {
	Id          string          `json:"-"`
	ForceRemove *bool           `json:"force,omitempty"`
	Status      NDFCStatusValue `json:"status,omitempty"`
	Spec        NDFCSpecValue   `json:"spec,omitempty"`
}

type NDFCStatusValue struct {
	State          string `json:"state,omitempty"`
	Version        string `json:"version,omitempty"`
	Local          *bool  `json:"local,omitempty"`
	PrimaryCluster *bool  `json:"primaryCluster,omitempty"`
	Virtual        *bool  `json:"virtual,omitempty"`
}

type NDFCSpecValue struct {
	Hostname    string                   `json:"onboardUrl,omitempty"`
	ClusterType string                   `json:"clusterType,omitempty"`
	ClusterName string                   `json:"name,omitempty"`
	ReRegister  *bool                    `json:"reRegister,omitempty"`
	Credentials NDFCSpecCredentialsValue `json:"credentials,omitempty"`
	Location    NDFCSpecLocationValue    `json:"location,omitempty"`
	Aci         NDFCSpecAciValue         `json:"aci,omitempty"`
	Nd          NDFCSpecNdValue          `json:"nd,omitempty"`
}

type NDFCSpecCredentialsValue struct {
	Username    string `json:"user,omitempty"`
	Password    string `json:"password,omitempty"`
	LoginDomain string `json:"loginDomain,omitempty"`
}

type NDFCSpecLocationValue struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type NDFCSpecAciValue struct {
	LicenseTier    string                    `json:"licenseTier,omitempty"`
	SecurityDomain string                    `json:"securityDomain,omitempty"`
	VerifyCa       *bool                     `json:"verifyCA,omitempty"`
	Orchestration  NDFCAciOrchestrationValue `json:"orchestration,omitempty"`
	Telemetry      NDFCAciTelemetryValue     `json:"telemetry,omitempty"`
}

type NDFCAciOrchestrationValue struct {
	OrchestrationStatus string `json:"status,omitempty"`
}

type NDFCAciTelemetryValue struct {
	TelemetryStatus            string `json:"status,omitempty"`
	TelemetryNetwork           string `json:"network,omitempty"`
	TelemetryEpg               string `json:"epg,omitempty"`
	TelemetryStreamingProtocol string `json:"streamingProtocol,omitempty"`
}

type NDFCSpecNdValue struct {
	MultiClusterLoginDomainName string `json:"multiClusterLoginDomainName,omitempty"`
}

func (v *MultiClusterConnectivityModel) SetModelData(jsonData *NDFCMultiClusterConnectivityModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.Id != "" {
		v.Id = types.StringValue(jsonData.Id)
	} else {
		v.Id = types.StringNull()
	}

	if jsonData.Spec.Hostname != "" {
		v.Hostname = types.StringValue(jsonData.Spec.Hostname)

	} else {
		v.Hostname = types.StringNull()
	}

	if jsonData.Spec.ClusterType != "" {
		v.ClusterType = types.StringValue(jsonData.Spec.ClusterType)

	} else {
		v.ClusterType = types.StringNull()
	}

	if jsonData.Spec.Credentials.Username != "" {
		v.Username = types.StringValue(jsonData.Spec.Credentials.Username)

	} else {
		v.Username = types.StringNull()
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

	if jsonData.Spec.Aci.VerifyCa != nil {
		v.VerifyCa = types.BoolValue(*jsonData.Spec.Aci.VerifyCa)

	} else {
		v.VerifyCa = types.BoolNull()
	}

	if jsonData.Spec.Aci.Orchestration.OrchestrationStatus != "" {
		v.OrchestrationStatus = types.StringValue(jsonData.Spec.Aci.Orchestration.OrchestrationStatus)

	} else {
		v.OrchestrationStatus = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryStatus != "" {
		v.TelemetryStatus = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryStatus)

	} else {
		v.TelemetryStatus = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryNetwork != "" {
		v.TelemetryNetwork = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryNetwork)

	} else {
		v.TelemetryNetwork = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryEpg != "" {
		v.TelemetryEpg = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryEpg)

	} else {
		v.TelemetryEpg = types.StringNull()
	}

	if jsonData.Spec.Aci.Telemetry.TelemetryStreamingProtocol != "" {
		v.TelemetryStreamingProtocol = types.StringValue(jsonData.Spec.Aci.Telemetry.TelemetryStreamingProtocol)

	} else {
		v.TelemetryStreamingProtocol = types.StringNull()
	}

	if jsonData.Spec.Nd.MultiClusterLoginDomainName != "" {
		v.MultiClusterLoginDomainName = types.StringValue(jsonData.Spec.Nd.MultiClusterLoginDomainName)

	} else {
		v.MultiClusterLoginDomainName = types.StringNull()
	}

	if jsonData.Spec.ClusterName != "" {
		v.ClusterName = types.StringValue(jsonData.Spec.ClusterName)

	} else {
		v.ClusterName = types.StringNull()
	}

	if jsonData.Status.State != "" {
		v.State = types.StringValue(jsonData.Status.State)

	} else {
		v.State = types.StringNull()
	}

	if jsonData.Status.Version != "" {
		v.Version = types.StringValue(jsonData.Status.Version)

	} else {
		v.Version = types.StringNull()
	}

	if jsonData.Status.Local != nil {
		v.Local = types.BoolValue(*jsonData.Status.Local)

	} else {
		v.Local = types.BoolNull()
	}

	if jsonData.Status.PrimaryCluster != nil {
		v.PrimaryCluster = types.BoolValue(*jsonData.Status.PrimaryCluster)

	} else {
		v.PrimaryCluster = types.BoolNull()
	}

	if jsonData.Status.Virtual != nil {
		v.Virtual = types.BoolValue(*jsonData.Status.Virtual)

	} else {
		v.Virtual = types.BoolNull()
	}

	if jsonData.ForceRemove != nil {
		v.ForceRemove = types.BoolValue(*jsonData.ForceRemove)

	} else {
		v.ForceRemove = types.BoolNull()
	}

	if jsonData.Spec.ReRegister != nil {
		v.ReRegister = types.BoolValue(*jsonData.Spec.ReRegister)

	} else {
		v.ReRegister = types.BoolNull()
	}

	return err
}

func (v MultiClusterConnectivityModel) GetModelData() *NDFCMultiClusterConnectivityModel {
	var data = new(NDFCMultiClusterConnectivityModel)

	//MARSHAL_BODY

	if !v.Hostname.IsNull() && !v.Hostname.IsUnknown() {
		data.Spec.Hostname = v.Hostname.ValueString()
	} else {
		data.Spec.Hostname = ""
	}

	if !v.ClusterType.IsNull() && !v.ClusterType.IsUnknown() {
		data.Spec.ClusterType = v.ClusterType.ValueString()
	} else {
		data.Spec.ClusterType = ""
	}

	if !v.Username.IsNull() && !v.Username.IsUnknown() {
		data.Spec.Credentials.Username = v.Username.ValueString()
	} else {
		data.Spec.Credentials.Username = ""
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

	if !v.VerifyCa.IsNull() && !v.VerifyCa.IsUnknown() {
		data.Spec.Aci.VerifyCa = new(bool)
		*data.Spec.Aci.VerifyCa = v.VerifyCa.ValueBool()
	} else {
		data.Spec.Aci.VerifyCa = nil
	}

	if !v.OrchestrationStatus.IsNull() && !v.OrchestrationStatus.IsUnknown() {
		data.Spec.Aci.Orchestration.OrchestrationStatus = v.OrchestrationStatus.ValueString()
	} else {
		data.Spec.Aci.Orchestration.OrchestrationStatus = ""
	}

	if !v.TelemetryStatus.IsNull() && !v.TelemetryStatus.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryStatus = v.TelemetryStatus.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryStatus = ""
	}

	if !v.TelemetryNetwork.IsNull() && !v.TelemetryNetwork.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryNetwork = v.TelemetryNetwork.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryNetwork = ""
	}

	if !v.TelemetryEpg.IsNull() && !v.TelemetryEpg.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryEpg = v.TelemetryEpg.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryEpg = ""
	}

	if !v.TelemetryStreamingProtocol.IsNull() && !v.TelemetryStreamingProtocol.IsUnknown() {
		data.Spec.Aci.Telemetry.TelemetryStreamingProtocol = v.TelemetryStreamingProtocol.ValueString()
	} else {
		data.Spec.Aci.Telemetry.TelemetryStreamingProtocol = ""
	}

	if !v.MultiClusterLoginDomainName.IsNull() && !v.MultiClusterLoginDomainName.IsUnknown() {
		data.Spec.Nd.MultiClusterLoginDomainName = v.MultiClusterLoginDomainName.ValueString()
	} else {
		data.Spec.Nd.MultiClusterLoginDomainName = ""
	}

	if !v.ClusterName.IsNull() && !v.ClusterName.IsUnknown() {
		data.Spec.ClusterName = v.ClusterName.ValueString()
	} else {
		data.Spec.ClusterName = ""
	}

	if !v.ForceRemove.IsNull() && !v.ForceRemove.IsUnknown() {
		data.ForceRemove = new(bool)
		*data.ForceRemove = v.ForceRemove.ValueBool()
	} else {
		data.ForceRemove = nil
	}

	if !v.ReRegister.IsNull() && !v.ReRegister.IsUnknown() {
		data.Spec.ReRegister = new(bool)
		*data.Spec.ReRegister = v.ReRegister.ValueBool()
	} else {
		data.Spec.ReRegister = nil
	}

	return data
}
