package topic

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"

	cloud "github.com/marciocadev/multicloud-go/cloud"
	aws "github.com/marciocadev/multicloud-go/aws"
)

type TopicClient interface {
	Publish(ctx context.Context, messageBody string) error
}

func GetTopicClient() (TopicClient, error) {
	provider := cloud.CloudProvider(os.Getenv("CLOUD_PROVIDER"))
	switch provider {
	case cloud.AWS:
		// AWS SNS
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
		if err != nil {
			return nil, fmt.Errorf("erro ao carregar configuração AWS: %w", err)
		}
		client := sns.NewFromConfig(cfg)
		return &aws.SNSClient{
			Client:   client,
			TopicARN: os.Getenv("TOPIC_ID"),
		}, nil
	case cloud.GCP:
		// GCP Pub/Sub
		return nil, fmt.Errorf("GCP TopicClient not implemented")
	case cloud.AZURE:
		// Azure Service Bus
		return nil, fmt.Errorf("AZURE TopicClient not implemented")
	case cloud.OCI:
		// OCI Topic
		return nil, fmt.Errorf("OCI TopicClient not implemented")
	default:
		return nil, fmt.Errorf("unsupported cloud provider: %s", cloud)
	}
}
