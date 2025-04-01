package queue

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"

	q "github.com/marciocadev/multicloud-go/cloud/aws"
)

type QueueClient interface {
	SendMessage(ctx context.Context, messageBody string) error
}

func GetQueueClient() (QueueClient, error) {
	cloud := os.Getenv("CLOUD_PROVIDER")
	switch cloud {
	case "AWS":
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
		if err != nil {
			return nil, fmt.Errorf("erro ao carregar configuração AWS: %w", err)
		}
		client := sqs.NewFromConfig(cfg)
		return &q.SQSClient{
			Client:   client,
			QueueURL: os.Getenv("QUEUE_URL"),
		}, nil
	case "GCP":
		// GCP Pub/Sub
		return nil, fmt.Errorf("GCP QueueClient not implemented")
	case "Azure":
		// Azure Queue Storage
		return nil, fmt.Errorf("Azure QueueClient not implemented")
	case "OCI":
		// OCI Queue
		return nil, fmt.Errorf("OCI QueueClient not implemented")
	default:
		return nil, fmt.Errorf("unsupported cloud provider: %s", cloud)
	}
}
