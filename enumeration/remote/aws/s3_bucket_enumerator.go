package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/alerter"
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/alerts"
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/aws/repository"
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/common"
	remoteerror "github.com/khulnasoft-lab/driftctl/enumeration/remote/error"
	tf "github.com/khulnasoft-lab/driftctl/enumeration/remote/terraform"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource/aws"
	"github.com/sirupsen/logrus"
)

type S3BucketEnumerator struct {
	repository     repository.S3Repository
	factory        resource.ResourceFactory
	providerConfig tf.TerraformProviderConfig
	alerter        alerter.AlerterInterface
}

func NewS3BucketEnumerator(repo repository.S3Repository, factory resource.ResourceFactory, providerConfig tf.TerraformProviderConfig, alerter alerter.AlerterInterface) *S3BucketEnumerator {
	return &S3BucketEnumerator{
		repository:     repo,
		factory:        factory,
		providerConfig: providerConfig,
		alerter:        alerter,
	}
}

func (e *S3BucketEnumerator) SupportedType() resource.ResourceType {
	return aws.AwsS3BucketResourceType
}

func (e *S3BucketEnumerator) Enumerate() ([]*resource.Resource, error) {
	buckets, err := e.repository.ListAllBuckets()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(buckets))

	for _, bucket := range buckets {
		region, err := e.repository.GetBucketLocation(*bucket.Name)
		if err != nil {
			alerts.SendEnumerationAlert(common.RemoteAWSTerraform, e.alerter, remoteerror.NewResourceScanningError(err, string(e.SupportedType()), *bucket.Name))
			continue
		}
		if region == "" || region != e.providerConfig.DefaultAlias {
			logrus.WithFields(logrus.Fields{
				"region": region,
				"bucket": *bucket.Name,
			}).Debug("Skipped bucket")
			continue
		}
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				*bucket.Name,
				map[string]interface{}{
					"alias": region,
				},
			),
		)
	}

	return results, err
}
