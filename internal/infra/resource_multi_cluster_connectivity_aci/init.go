package resource_multi_cluster_connectivity_aci

import (
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterResource(ModuleKey, NewMultiClusterConnectivityACIResource)
}
