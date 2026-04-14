package resource_link_policy_test

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	rl "terraform-provider-nd/internal/manage/resource_link"
)

func newFullyPopulatedTemplateInputs() rl.NDFCConfigDataTemplateInputsValue {
	boolTrue := true
	int64Val := int64(100)
	return rl.NDFCConfigDataTemplateInputsValue{
		SrcInterfaceDescription:              "srcDesc",
		SrcInterfaceConfig:                   "srcConfig",
		DstInterfaceDescription:              "dstDesc",
		DstInterfaceConfig:                   "dstConfig",
		InterfaceAdminState:                  &boolTrue,
		Mtu:                                  &int64Val,
		LinkMtu:                              &int64Val,
		Speed:                                "auto",
		Fec:                                  "auto",
		Macsec:                               &boolTrue,
		SrcIp:                                "1.1.1.1",
		SrcIpv6:                              "::1",
		SrcIpAddress:                         "2.2.2.2",
		SrcIpAddressMask:                     "3.3.3.0/24",
		SrcIpv6AddressMask:                   "2001:db8::/64",
		DstIp:                                "4.4.4.4",
		DstIpv6:                              "::2",
		DstIpAddress:                         "5.5.5.5",
		DstIpAddressMask:                     "6.6.6.0/24",
		DstIpv6Address:                       "::3",
		DstIpv6AddressMask:                   "2001:db8::2/64",
		SrcEbgpAsn:                           "65001",
		DstEbgpAsn:                           "65002",
		EbgpPassword:                         "secret",
		EbgpAuthKeyEncryptionType:            "3des",
		EbgpBfd:                              &boolTrue,
		EbgpLogNeighborChange:                &boolTrue,
		EbgpMaximumPaths:                     &int64Val,
		EbgpMultihop:                         &int64Val,
		EbgpSendComboth:                      &boolTrue,
		EnableEbgpPassword:                   &boolTrue,
		InheritEbgpPasswordMsdSettings:       &boolTrue,
		DefaultVrfEbgpNeighborPassword:       "defaultPw",
		InterfaceVrf:                         "mgmt",
		SrcVrfName:                           "srcVrf",
		DstVrfName:                           "dstVrf",
		VrfNameNxPeerSwitch:                  "peerVrf",
		AutoGenConfigPeer:                    &boolTrue,
		AutoGenConfigDefaultVrf:              &boolTrue,
		AutoGenConfigNxPeerDefaultVrf:        &boolTrue,
		TemplateConfigGenPeer:                "Ext_VRF_Lite_Jython",
		RedistribEbgpRouteMapName:            "routeMap",
		RedistributeRouteServer:              &boolTrue,
		RouteServerRoutingTag:                "tag1",
		RoutingTag:                           "tag2",
		DciTracking:                          &boolTrue,
		DciTrackingEnableFlag:                &boolTrue,
		InheritTtagFabricSetting:             &boolTrue,
		OverrideFabricMacsec:                 &boolTrue,
		MacsecCipherSuite:                    "suite1",
		MacsecPrimaryKeyString:               "key1",
		MacsecPrimaryCryptographicAlgorithm:  "alg1",
		MacsecFallbackKeyString:              "key2",
		MacsecFallbackCryptographicAlgorithm: "alg2",
		Qkd:                                  &boolTrue,
		IgnoreCertificate:                    &boolTrue,
		SrcKmeServerIp:                       "10.0.0.1",
		SrcMacsecKeyChainPrefix:              "srcPrefix",
		SrcQkdProfileName:                    "srcProfile",
		SrcTrustpointLabel:                   "srcTrust",
		DstKmeServerIp:                       "10.0.0.2",
		DstMacsecKeyChainPrefix:              "dstPrefix",
		DstQkdProfileName:                    "dstProfile",
		DstTrustpointLabel:                   "dstTrust",
		DhcpRelayOnSrcInterface:              &boolTrue,
		DhcpRelayOnDstInterface:              &boolTrue,
		BfdEchoOnSrcInterface:                &boolTrue,
		BfdEchoOnDstInterface:                &boolTrue,
		MtuType:                              "jumbo",
		Ipv4Pim:                              &boolTrue,
		Ipv4Trm:                              &boolTrue,
		Ipv6Pim:                              &boolTrue,
		Ipv6Trm:                              &boolTrue,
		IpRedirects:                          &boolTrue,
		SkipConfigGeneration:                 &boolTrue,
		BpduGuard:                            "yes",
		PortTypeFast:                         &boolTrue,
		NativeVlan:                           &int64Val,
		TrunkAllowedVlans:                    "1-100",
		NetflowOnSrcInterface:                &boolTrue,
		NetflowOnDstInterface:                &boolTrue,
		SrcNetflowMonitorName:                "srcMonitor",
		DstNetflowMonitorName:                "dstMonitor",
		TemplateName:                         "myTemplate",
	}
}

func TestFilterByPolicyType_Numbered(t *testing.T) {
	ti := newFullyPopulatedTemplateInputs()
	ti.FilterByPolicyType("numbered")

	if ti.SrcIp == "" {
		t.Error("srcIp should be kept for numbered")
	}
	if ti.DstIp == "" {
		t.Error("dstIp should be kept for numbered")
	}
	if ti.SrcIpv6 == "" {
		t.Error("srcIpv6 should be kept for numbered")
	}
	if ti.DstIpv6 == "" {
		t.Error("dstIpv6 should be kept for numbered")
	}
	if ti.InterfaceAdminState == nil {
		t.Error("interfaceAdminState should be kept for numbered")
	}
	if ti.Mtu == nil {
		t.Error("mtu should be kept for numbered")
	}
	if ti.Speed == "" {
		t.Error("speed should be kept for numbered")
	}
	if ti.Fec == "" {
		t.Error("fec should be kept for numbered")
	}
	if ti.Macsec == nil {
		t.Error("macsec should be kept for numbered")
	}
	if ti.SrcInterfaceDescription == "" {
		t.Error("srcInterfaceDescription should be kept for numbered")
	}
	if ti.DstInterfaceDescription == "" {
		t.Error("dstInterfaceDescription should be kept for numbered")
	}
	if ti.SrcInterfaceConfig == "" {
		t.Error("srcInterfaceConfig should be kept for numbered")
	}
	if ti.DstInterfaceConfig == "" {
		t.Error("dstInterfaceConfig should be kept for numbered")
	}
	if ti.BfdEchoOnSrcInterface == nil {
		t.Error("bfdEchoOnSrcInterface should be kept for numbered")
	}
	if ti.BfdEchoOnDstInterface == nil {
		t.Error("bfdEchoOnDstInterface should be kept for numbered")
	}
	if ti.DhcpRelayOnSrcInterface == nil {
		t.Error("dhcpRelayOnSrcInterface should be kept for numbered")
	}
	if ti.DhcpRelayOnDstInterface == nil {
		t.Error("dhcpRelayOnDstInterface should be kept for numbered")
	}

	// Fields that should be ZEROED for numbered
	if ti.DstNetflowMonitorName != "" {
		t.Errorf("dstNetflowMonitorName should be zeroed for numbered, got %q", ti.DstNetflowMonitorName)
	}
	if ti.TemplateName != "" {
		t.Errorf("templateName should be zeroed for numbered, got %q", ti.TemplateName)
	}
	if ti.DstKmeServerIp != "" {
		t.Errorf("dstKmeServerIp should be zeroed for numbered, got %q", ti.DstKmeServerIp)
	}
	if ti.LinkMtu != nil {
		t.Error("linkMtu should be zeroed for numbered")
	}
	if ti.SrcEbgpAsn != "" {
		t.Errorf("srcEbgpAsn should be zeroed for numbered, got %q", ti.SrcEbgpAsn)
	}
	if ti.EbgpBfd != nil {
		t.Error("ebgpBfd should be zeroed for numbered")
	}
	if ti.NativeVlan != nil {
		t.Error("nativeVlan should be zeroed for numbered")
	}
	if ti.SkipConfigGeneration != nil {
		t.Error("skipConfigGeneration should be zeroed for numbered")
	}
	if ti.Ipv4Trm != nil {
		t.Error("ipv4Trm should be zeroed for numbered")
	}
	if ti.OverrideFabricMacsec != nil {
		t.Error("overrideFabricMacsec should be zeroed for numbered")
	}
}

func TestFilterByPolicyType_UserDefined(t *testing.T) {
	ti := newFullyPopulatedTemplateInputs()
	ti.FilterByPolicyType("userDefined")

	if ti.TemplateName == "" {
		t.Error("templateName should be kept for userDefined")
	}
	if ti.SrcIp != "" {
		t.Errorf("srcIp should be zeroed for userDefined, got %q", ti.SrcIp)
	}
	if ti.Mtu != nil {
		t.Error("mtu should be zeroed for userDefined")
	}
	if ti.InterfaceAdminState != nil {
		t.Error("interfaceAdminState should be zeroed for userDefined")
	}
	if ti.Macsec != nil {
		t.Error("macsec should be zeroed for userDefined")
	}
	if ti.SrcInterfaceDescription != "" {
		t.Errorf("srcInterfaceDescription should be zeroed for userDefined, got %q", ti.SrcInterfaceDescription)
	}
	if ti.DstNetflowMonitorName != "" {
		t.Errorf("dstNetflowMonitorName should be zeroed for userDefined, got %q", ti.DstNetflowMonitorName)
	}
}

func TestFilterByPolicyType_UnknownPolicyNoFiltering(t *testing.T) {
	ti := newFullyPopulatedTemplateInputs()
	ti.FilterByPolicyType("futureUnknownPolicy")

	if ti.SrcIp == "" {
		t.Error("srcIp should remain for unknown policy type")
	}
	if ti.TemplateName == "" {
		t.Error("templateName should remain for unknown policy type")
	}
	if ti.Mtu == nil {
		t.Error("mtu should remain for unknown policy type")
	}
	if ti.DstNetflowMonitorName == "" {
		t.Error("dstNetflowMonitorName should remain for unknown policy type")
	}
}

func TestFilterByPolicyType_EmptyPolicyNoFiltering(t *testing.T) {
	ti := newFullyPopulatedTemplateInputs()
	ti.FilterByPolicyType("")

	if ti.SrcIp == "" {
		t.Error("srcIp should remain for empty policy type")
	}
	if ti.TemplateName == "" {
		t.Error("templateName should remain for empty policy type")
	}
}

func TestFilterByPolicyType_NumberedJSON(t *testing.T) {
	boolFalse := false
	boolTrue := true
	int64Val := int64(9216)
	int64Val100 := int64(100)

	model := rl.NDFCLinkModel{
		SrcFabricName:    "test_fabric",
		SrcSwitchName:    "Spine-1",
		SrcInterfaceName: "Ethernet1/31",
		DstFabricName:    "test_fabric",
		DstSwitchName:    "Leaf-1",
		DstInterfaceName: "Ethernet1/16",
		ConfigData: rl.NDFCConfigDataValue{
			PolicyType: "numbered",
			TemplateInputs: rl.NDFCConfigDataTemplateInputsValue{
				SrcIp:                 "1.1.1.1",
				DstIp:                 "1.1.1.2",
				InterfaceAdminState:   &boolFalse,
				Mtu:                   &int64Val,
				Speed:                 "auto",
				Fec:                   "auto",
				Macsec:                &boolFalse,
				DstNetflowMonitorName: "shouldBeRemoved",
				TemplateName:          "shouldBeRemoved",
				DstKmeServerIp:        "shouldBeRemoved",
				EbgpBfd:               &boolTrue,
				NativeVlan:            &int64Val100,
			},
		},
	}

	// Use BuildFilteredPayload which only includes applicable fields for the policyType
	payload, err := rl.BuildFilteredPayload(&model)
	if err != nil {
		t.Fatalf("BuildFilteredPayload failed: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	configData, ok := result["configData"].(map[string]interface{})
	if !ok {
		t.Fatal("configData not found in JSON")
	}

	templateInputs, ok := configData["templateInputs"].(map[string]interface{})
	if !ok {
		t.Fatal("templateInputs not found in JSON")
	}

	// Applicable fields for "numbered" must be present
	for _, field := range []string{"srcIp", "dstIp", "interfaceAdminState", "mtu", "speed", "fec", "macsec"} {
		if _, exists := templateInputs[field]; !exists {
			t.Errorf("expected field %q in numbered payload but not found", field)
		}
	}

	// Non-applicable fields must NOT be present at all (BuildFilteredPayload excludes them)
	for _, field := range []string{"dstNetflowMonitorName", "templateName", "dstKmeServerIp", "ebgpBfd", "nativeVlan"} {
		if _, exists := templateInputs[field]; exists {
			t.Errorf("field %q should NOT be in numbered payload, but was found with value %v", field, templateInputs[field])
		}
	}
}

func TestFilterByPolicyType_Layer3DciVrfLite(t *testing.T) {
	ti := newFullyPopulatedTemplateInputs()
	ti.FilterByPolicyType("layer3DciVrfLite")

	if ti.DstNetflowMonitorName == "" {
		t.Error("dstNetflowMonitorName should be kept for layer3DciVrfLite")
	}
	if ti.SrcNetflowMonitorName == "" {
		t.Error("srcNetflowMonitorName should be kept for layer3DciVrfLite")
	}
	if ti.IpRedirects == nil {
		t.Error("ipRedirects should be kept for layer3DciVrfLite")
	}
	if ti.TemplateName != "" {
		t.Errorf("templateName should be zeroed for layer3DciVrfLite, got %q", ti.TemplateName)
	}
	if ti.BfdEchoOnSrcInterface != nil {
		t.Error("bfdEchoOnSrcInterface should be zeroed for layer3DciVrfLite")
	}
	if ti.SrcIp != "" {
		t.Errorf("srcIp should be zeroed for layer3DciVrfLite, got %q", ti.SrcIp)
	}
}

func TestFilterByPolicyType_AllPolicyTypesHaveEntries(t *testing.T) {
	expectedPolicies := []string{
		"ipv6LinkLocal", "numbered", "unnumbered", "vpcPeerKeepalive",
		"ipfmNumbered", "preprovision", "iosXeNumbered", "ebgpVrfLite",
		"ospfVrfLite", "layer3DciVrfLite", "csrMultisiteUnderlay",
		"multisiteUnderlay", "vxlanAciMultisiteUnderlay",
		"csrEvpnMultisiteOverlay", "multisiteOverlay", "vxlanAciEvpnOverlay",
		"layer2Dci", "mplsUnderlay", "mplsOverlay", "routedFabric",
		"userDefined", "csrBgpOverIpsec",
	}
	for _, p := range expectedPolicies {
		if _, ok := rl.PolicyTemplateFields[p]; !ok {
			t.Errorf("PolicyTemplateFields missing entry for %q", p)
		}
	}
}

func TestBuildFilteredCreatePayload_Structure(t *testing.T) {
	boolFalse := false
	int64Val := int64(9216)

	model := rl.NDFCLinkModel{
		SrcSwitchId:      "SNOP1",
		DstSwitchId:      "SNOP2",
		SrcFabricName:    "fab1",
		SrcSwitchName:    "Spine-1",
		SrcInterfaceName: "Ethernet1/1",
		DstFabricName:    "fab1",
		DstSwitchName:    "Leaf-1",
		DstInterfaceName: "Ethernet1/1",
		ConfigData: rl.NDFCConfigDataValue{
			PolicyType: "numbered",
			TemplateInputs: rl.NDFCConfigDataTemplateInputsValue{
				SrcIp:               "10.0.0.1",
				DstIp:               "10.0.0.2",
				InterfaceAdminState: &boolFalse,
				Mtu:                 &int64Val,
				Speed:               "auto",
				Fec:                 "auto",
				Macsec:              &boolFalse,
				// Non-applicable field
				TemplateName: "shouldNotAppear",
			},
		},
	}

	payload, err := rl.BuildFilteredCreatePayload(&model)
	if err != nil {
		t.Fatalf("BuildFilteredCreatePayload failed: %v", err)
	}

	var wrapper map[string]interface{}
	if err := json.Unmarshal(payload, &wrapper); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	links, ok := wrapper["links"].([]interface{})
	if !ok || len(links) != 1 {
		t.Fatalf("expected links array with 1 element, got %v", wrapper["links"])
	}

	link := links[0].(map[string]interface{})
	if link["srcSwitchId"] != "SNOP1" {
		t.Errorf("expected srcSwitchId=SNOP1, got %v", link["srcSwitchId"])
	}

	configData := link["configData"].(map[string]interface{})
	ti := configData["templateInputs"].(map[string]interface{})

	if _, exists := ti["srcIp"]; !exists {
		t.Error("srcIp should be present for numbered")
	}
	if _, exists := ti["templateName"]; exists {
		t.Error("templateName should NOT be present for numbered")
	}
}

// --- PreserveNonPolicyDefaults tests ---

func TestPreserveNonPolicyDefaults_NumberedRestoresNonApplicable(t *testing.T) {
	// Simulate the scenario from the log: policy_type = "numbered"
	// After SetModelData, non-applicable fields are null.
	// PreserveNonPolicyDefaults should restore them from the saved (plan) model.

	saved := rl.LinkModel{}
	// Set non-applicable fields to their schema default values (as Terraform plan would)
	saved.EnableEbgpPassword = types.BoolValue(false)
	saved.EbgpBfd = types.BoolValue(false)
	saved.EbgpLogNeighborChange = types.BoolValue(false)
	saved.EbgpMaximumPaths = types.Int64Value(64)
	saved.EbgpMultihop = types.Int64Value(5)
	saved.EbgpSendComboth = types.BoolValue(false)
	saved.InheritEbgpPasswordMsdSettings = types.BoolValue(true)
	saved.EbgpAuthKeyEncryptionType = types.StringValue("3des")
	saved.AutoGenConfigPeer = types.BoolValue(false)
	saved.AutoGenConfigDefaultVrf = types.BoolValue(false)
	saved.AutoGenConfigNxPeerDefaultVrf = types.BoolValue(false)
	saved.TemplateConfigGenPeer = types.StringValue("Ext_VRF_Lite_Jython")
	saved.OverrideFabricMacsec = types.BoolValue(false)
	saved.RedistributeRouteServer = types.BoolValue(false)
	saved.DciTracking = types.BoolValue(false)
	saved.DciTrackingEnableFlag = types.BoolValue(false)
	saved.InheritTtagFabricSetting = types.BoolValue(true)
	saved.Qkd = types.BoolValue(false)
	saved.IgnoreCertificate = types.BoolValue(false)
	saved.Ipv4Pim = types.BoolValue(false)
	saved.Ipv4Trm = types.BoolValue(false)
	saved.Ipv6Pim = types.BoolValue(false)
	saved.Ipv6Trm = types.BoolValue(false)
	saved.IpRedirects = types.BoolValue(false)
	saved.SkipConfigGeneration = types.BoolValue(false)
	saved.PortTypeFast = types.BoolValue(true)
	saved.NetflowOnSrcInterface = types.BoolValue(false)
	saved.NetflowOnDstInterface = types.BoolValue(false)
	saved.LinkMtu = types.Int64Value(9216)

	// current = what SetModelData produced (non-applicable fields are null)
	current := rl.LinkModel{}
	// Applicable fields that the API returned correctly
	current.SrcIp = types.StringValue("1.1.5.1")
	current.DstIp = types.StringValue("1.1.5.2")
	current.Mtu = types.Int64Value(9216)
	current.Fec = types.StringValue("auto")
	current.Speed = types.StringValue("auto")
	current.Macsec = types.BoolValue(false)
	current.InterfaceAdminState = types.BoolValue(false)
	current.PolicyType = types.StringValue("numbered")

	current.PreserveNonPolicyDefaults("numbered", &saved)

	// Non-applicable fields should be restored from saved
	if current.EnableEbgpPassword.IsNull() {
		t.Error("enable_ebgp_password should be restored, not null")
	}
	if current.EnableEbgpPassword.ValueBool() != false {
		t.Error("enable_ebgp_password should be false")
	}
	if current.EbgpMultihop.IsNull() {
		t.Error("ebgp_multihop should be restored, not null")
	}
	if current.EbgpMultihop.ValueInt64() != 5 {
		t.Errorf("ebgp_multihop should be 5, got %d", current.EbgpMultihop.ValueInt64())
	}
	if current.EbgpAuthKeyEncryptionType.IsNull() {
		t.Error("ebgp_auth_key_encryption_type should be restored, not null")
	}
	if current.EbgpAuthKeyEncryptionType.ValueString() != "3des" {
		t.Errorf("ebgp_auth_key_encryption_type should be '3des', got %q", current.EbgpAuthKeyEncryptionType.ValueString())
	}
	if current.TemplateConfigGenPeer.IsNull() {
		t.Error("template_config_gen_peer should be restored, not null")
	}
	if current.PortTypeFast.IsNull() {
		t.Error("port_type_fast should be restored, not null")
	}
	if current.LinkMtu.IsNull() {
		t.Error("link_mtu should be restored, not null")
	}
	if current.InheritTtagFabricSetting.IsNull() {
		t.Error("inherit_ttag_fabric_setting should be restored, not null")
	}
	if current.DciTracking.IsNull() {
		t.Error("dci_tracking should be restored, not null")
	}

	// Applicable fields should NOT be overwritten
	if current.SrcIp.ValueString() != "1.1.5.1" {
		t.Errorf("src_ip should remain '1.1.5.1', got %q", current.SrcIp.ValueString())
	}
	if current.DstIp.ValueString() != "1.1.5.2" {
		t.Errorf("dst_ip should remain '1.1.5.2', got %q", current.DstIp.ValueString())
	}
	if current.Mtu.ValueInt64() != 9216 {
		t.Errorf("mtu should remain 9216, got %d", current.Mtu.ValueInt64())
	}
}

func TestPreserveNonPolicyDefaults_ApplicableEmptyStringPreserved(t *testing.T) {
	// Simulate: applicable string fields where API returns "" and SetModelData sets null,
	// but the plan had explicit empty string.
	saved := rl.LinkModel{}
	saved.SrcIpv6 = types.StringValue("")
	saved.DstIpv6 = types.StringValue("")
	saved.SrcInterfaceDescription = types.StringValue("")
	saved.SrcInterfaceConfig = types.StringValue("")
	saved.DstInterfaceDescription = types.StringValue("")
	saved.DstInterfaceConfig = types.StringValue("")

	current := rl.LinkModel{}
	// SetModelData set these to null because the API returned ""
	// (all fields left at zero value = types.String{} which is null)

	current.PreserveNonPolicyDefaults("numbered", &saved)

	if current.SrcIpv6.IsNull() {
		t.Error("src_ipv6 should be restored from saved empty string, not null")
	}
	if current.SrcIpv6.ValueString() != "" {
		t.Errorf("src_ipv6 should be empty string, got %q", current.SrcIpv6.ValueString())
	}
	if current.DstIpv6.IsNull() {
		t.Error("dst_ipv6 should be restored from saved empty string, not null")
	}
	if current.SrcInterfaceDescription.IsNull() {
		t.Error("src_interface_description should be restored, not null")
	}
	if current.DstInterfaceConfig.IsNull() {
		t.Error("dst_interface_config should be restored, not null")
	}
}

func TestPreserveNonPolicyDefaults_ApplicableFieldWithAPIValueNotOverwritten(t *testing.T) {
	// When the API returns a real value for an applicable field,
	// PreserveNonPolicyDefaults must NOT overwrite it with the saved value.
	saved := rl.LinkModel{}
	saved.SrcIp = types.StringValue("old-ip")
	saved.Mtu = types.Int64Value(1500)
	saved.Macsec = types.BoolValue(true)

	current := rl.LinkModel{}
	current.SrcIp = types.StringValue("new-ip")
	current.Mtu = types.Int64Value(9216)
	current.Macsec = types.BoolValue(false)

	current.PreserveNonPolicyDefaults("numbered", &saved)

	if current.SrcIp.ValueString() != "new-ip" {
		t.Errorf("src_ip should stay 'new-ip', got %q", current.SrcIp.ValueString())
	}
	if current.Mtu.ValueInt64() != 9216 {
		t.Errorf("mtu should stay 9216, got %d", current.Mtu.ValueInt64())
	}
	if current.Macsec.ValueBool() != false {
		t.Error("macsec should stay false")
	}
}

func TestPreserveNonPolicyDefaults_UnknownPolicyNoOp(t *testing.T) {
	saved := rl.LinkModel{}
	saved.EnableEbgpPassword = types.BoolValue(false)

	current := rl.LinkModel{}
	// EnableEbgpPassword is null (zero value)

	current.PreserveNonPolicyDefaults("futureUnknownPolicy", &saved)

	// Unknown policy → no restore, field stays null
	if !current.EnableEbgpPassword.IsNull() {
		t.Error("for unknown policy, fields should not be modified")
	}
}

func TestPreserveNonPolicyDefaults_EmptyPolicyNoOp(t *testing.T) {
	saved := rl.LinkModel{}
	saved.EnableEbgpPassword = types.BoolValue(false)

	current := rl.LinkModel{}

	current.PreserveNonPolicyDefaults("", &saved)

	if !current.EnableEbgpPassword.IsNull() {
		t.Error("for empty policy, fields should not be modified")
	}
}
