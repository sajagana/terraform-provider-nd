// Code generated;  DO NOT EDIT.

package resource_multi_cluster_connectivity

import (
	"strconv"
	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func MultiClusterConnectivityModelHelperStateCheck(RscName string, c resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.Id != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("id").String(), c.Id))
	}
	if c.Spec.Hostname != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("hostname").String(), c.Spec.Hostname))
	}
	if c.Spec.ClusterType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cluster_type").String(), c.Spec.ClusterType))
	}
	if c.Spec.Username != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("username").String(), c.Spec.Username))
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
	}
	if c.Spec.SecurityDomain != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_domain").String(), c.Spec.SecurityDomain))
	}
	if c.VerifyCa != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("verify_ca").String(), strconv.FormatBool(*c.VerifyCa)))
	}
	if c.Spec.OrchestrationStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("orchestration_status").String(), c.Spec.OrchestrationStatus))
	}
	if c.Spec.TelemetryStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_status").String(), c.Spec.TelemetryStatus))
	}
	if c.Spec.TelemetryNetwork != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_network").String(), c.Spec.TelemetryNetwork))
	}
	if c.Spec.TelemetryEpg != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_epg").String(), c.Spec.TelemetryEpg))
	}
	if c.Spec.TelemetryStreamingProtocol != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("telemetry_streaming_protocol").String(), c.Spec.TelemetryStreamingProtocol))
	}
	if c.Spec.MultiClusterLoginDomainName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multi_cluster_login_domain_name").String(), c.Spec.MultiClusterLoginDomainName))
	}
	if c.Spec.ClusterName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cluster_name").String(), c.Spec.ClusterName))
	}
	if c.Status.State != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("state").String(), c.Status.State))
	}
	if c.Status.Version != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("version").String(), c.Status.Version))
	}
	if c.Local != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("local").String(), strconv.FormatBool(*c.Local)))
	}
	if c.PrimaryCluster != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("primary_cluster").String(), strconv.FormatBool(*c.PrimaryCluster)))
	}
	if c.Virtual != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("virtual").String(), strconv.FormatBool(*c.Virtual)))
	}
	if c.ForceRemove != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("force_remove").String(), strconv.FormatBool(*c.ForceRemove)))
	}
	if c.ReRegister != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("re_register").String(), strconv.FormatBool(*c.ReRegister)))
	}
	return ret
}
