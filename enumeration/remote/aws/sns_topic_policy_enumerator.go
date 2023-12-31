package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/aws/repository"
	remoteerror "github.com/khulnasoft-lab/driftctl/enumeration/remote/error"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource/aws"
)

type SNSTopicPolicyEnumerator struct {
	repository repository.SNSRepository
	factory    resource.ResourceFactory
}

func NewSNSTopicPolicyEnumerator(repo repository.SNSRepository, factory resource.ResourceFactory) *SNSTopicPolicyEnumerator {
	return &SNSTopicPolicyEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *SNSTopicPolicyEnumerator) SupportedType() resource.ResourceType {
	return aws.AwsSnsTopicPolicyResourceType
}

func (e *SNSTopicPolicyEnumerator) Enumerate() ([]*resource.Resource, error) {
	topics, err := e.repository.ListAllTopics()
	if err != nil {
		return nil, remoteerror.NewResourceListingErrorWithType(err, string(e.SupportedType()), aws.AwsSnsTopicResourceType)
	}

	results := make([]*resource.Resource, 0, len(topics))

	for _, topic := range topics {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				*topic.TopicArn,
				map[string]interface{}{
					"topic_arn": *topic.TopicArn,
				},
			),
		)
	}

	return results, err
}
