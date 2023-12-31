package azurerm

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/azurerm/repository"
	remoteerror "github.com/khulnasoft-lab/driftctl/enumeration/remote/error"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource/azurerm"
)

type AzurermPrivateDNSSRVRecordEnumerator struct {
	repository repository.PrivateDNSRepository
	factory    resource.ResourceFactory
}

func NewAzurermPrivateDNSSRVRecordEnumerator(repo repository.PrivateDNSRepository, factory resource.ResourceFactory) *AzurermPrivateDNSSRVRecordEnumerator {
	return &AzurermPrivateDNSSRVRecordEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *AzurermPrivateDNSSRVRecordEnumerator) SupportedType() resource.ResourceType {
	return azurerm.AzurePrivateDNSSRVRecordResourceType
}

func (e *AzurermPrivateDNSSRVRecordEnumerator) Enumerate() ([]*resource.Resource, error) {

	zones, err := e.repository.ListAllPrivateZones()
	if err != nil {
		return nil, remoteerror.NewResourceListingErrorWithType(err, string(e.SupportedType()), azurerm.AzurePrivateDNSZoneResourceType)
	}

	results := make([]*resource.Resource, 0)

	for _, zone := range zones {
		records, err := e.repository.ListAllSRVRecords(zone)
		if err != nil {
			return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
		}
		for _, record := range records {
			results = append(
				results,
				e.factory.CreateAbstractResource(
					string(e.SupportedType()),
					*record.ID,
					map[string]interface{}{
						"name":      *record.Name,
						"zone_name": *zone.Name,
					},
				),
			)
		}

	}

	return results, err
}
