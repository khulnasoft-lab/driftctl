package azurerm

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const AzurePrivateDNSAAAARecordResourceType = "azurerm_private_dns_aaaa_record"

func initAzurePrivateDNSAAAARecordMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AzurePrivateDNSAAAARecordResourceType, func(res *resource.Resource) {
		res.Attributes().SafeDelete([]string{"timeouts"})
	})
	resourceSchemaRepository.SetHumanReadableAttributesFunc(AzurePrivateDNSAAAARecordResourceType, func(res *resource.Resource) map[string]string {
		val := res.Attrs
		attrs := make(map[string]string)
		if name := val.GetString("name"); name != nil && *name != "" {
			attrs["Name"] = *name
		}
		if zone := val.GetString("zone_name"); zone != nil && *zone != "" {
			attrs["Zone"] = *zone
		}
		return attrs
	})
}
