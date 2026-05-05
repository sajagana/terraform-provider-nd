// Code generated;  DO NOT EDIT.

package resource_multi_cluster_connectivity_aci

import (
	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity_aci"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func MultiClusterConnectivityAciModelHelperStateCheck(RscName string, c resource_multi_cluster_connectivity_aci.NDFCMultiClusterConnectivityAciModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Id != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("id").String(), c.Id))
	}
	if c.Spec.ClusterType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cluster_type").String(), c.Spec.ClusterType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cluster_type").String(), "APIC"))
	}
	if c.Spec.ClusterName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cluster_name").String(), c.Spec.ClusterName))
	}
	if c.Spec.HostName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("host_name").String(), c.Spec.HostName))
	}
	if c.Spec.UserName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("user_name").String(), c.Spec.UserName))
	}
	if c.Spec.Password != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("password").String(), c.Spec.Password))
	}
	if c.Spec.LoginDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("login_domain").String(), c.Spec.LoginDomain))
	}
	if c.Spec.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.Spec.FabricName))
	}
	if c.Spec.LicenseTier != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("license_tier").String(), c.Spec.LicenseTier))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("license_tier").String(), "essentials"))
	}
	if c.Spec.SecurityDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_domain").String(), c.Spec.SecurityDomain))
	}
	if c.Spec.ValidatePeerCertificate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("validate_peer_certificate").String(), c.Spec.ValidatePeerCertificate))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("validate_peer_certificate").String(), "false"))
	}
	if c.Spec.EnableOrchestration != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_orchestration").String(), c.Spec.EnableOrchestration))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_orchestration").String(), "disabled"))
	}
	if c.Spec.EnableTelemetry != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_telemetry").String(), c.Spec.EnableTelemetry))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_telemetry").String(), "disabled"))
	}
	if c.Spec.TelemetryCollection != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_collection").String(), c.Spec.TelemetryCollection))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_collection").String(), "inband"))
	}
	if c.Spec.TelemetryStreaming != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_streaming").String(), c.Spec.TelemetryStreaming))
	}
	if c.Spec.TelemetryEpg != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_epg").String(), c.Spec.TelemetryEpg))
	}
	return ret
}
