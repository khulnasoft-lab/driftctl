package azurerm

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const AzurePrivateDNSZoneResourceType = "azurerm_private_dns_zone"

func initAzurePrivateDNSZoneMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AzurePrivateDNSZoneResourceType, func(res *resource.Resource) {
		res.Attributes().SafeDelete([]string{"number_of_record_sets"})
		res.Attributes().SafeDelete([]string{"timeouts"})
	})
}
