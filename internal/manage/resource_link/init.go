package resource_link

import (
	"terraform-provider-nd/internal/registry"
)

func init() {
	registry.RegisterResource(ModuleKey, NewLinkResource)
}
