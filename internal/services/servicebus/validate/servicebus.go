package validate

import (
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

func ServiceBusMaxSizeInMegabytes() pluginsdk.SchemaValidateFunc {
	return validation.IntInSlice([]int{
		1024,
		2048,
		3072,
		4096,
		5120,
		10240,
		20480,
		40960,
		81920,
	})
}
