package queue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"

	q "github.com/marciocadev/multicloud-go/cloud/aws"
)

type CloudProvider int

const (
	AWS CloudProvider = iota
	GCP
	Azure
	OCI
)

func (c CloudProvider) String() string {
	switch c {
	case AWS:
		return "aws"
	case GCP:
		return "gcp"
	case Azure:
		return "azure"
	case OCI:
		return "oci"
	default:
		return "unknown"
	}
}

type QueueClient interface {
	SendMessage(ctx context.Context, messageBody string) error
}

func GetQueueClient(cloud CloudProvider, region string, queueUrl string) (QueueClient, error) {
	switch cloud {
	case AWS:
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
		if err != nil {
			return nil, fmt.Errorf("erro ao carregar configuração AWS: %w", err)
		}
		client := sqs.NewFromConfig(cfg)
		return &q.SQSClient{
			Client:   client,
			QueueURL: queueUrl,
		}, nil
	case GCP:
		// GCP Pub/Sub
		return nil, fmt.Errorf("GCP QueueClient not implemented")
	case Azure:
		// Azure Queue Storage
		return nil, fmt.Errorf("Azure QueueClient not implemented")
	case OCI:
		// OCI Queue
		return nil, fmt.Errorf("OCI QueueClient not implemented")
	default:
		return nil, fmt.Errorf("unsupported cloud provider: %s", cloud)
	}
}
