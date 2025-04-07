package cloud

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/sqs"

	aws "github.com/marciocadev/multicloud-go/cloud/aws"
)

type QueueEvent interface {
	events.SQSEvent
}

type QueueClient interface {
	SendMessage(ctx context.Context, messageBody string) error
}

func GetQueueClient() (QueueClient, error) {
	cloud := os.Getenv("CLOUD_PROVIDER")
	switch cloud {
	case "AWS":
		// AWS SQS
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
		if err != nil {
			return nil, fmt.Errorf("erro ao carregar configuração AWS: %w", err)
		}
		client := sqs.NewFromConfig(cfg)
		return &aws.SQSClient{
			Client:   client,
			QueueURL: os.Getenv("QUEUE_ID"),
		}, nil
	case "GCP":
		// GCP Pub/Sub
		return nil, fmt.Errorf("GCP QueueClient not implemented")
	case "AZURE":
		// Azure Queue Storage
		return nil, fmt.Errorf("AZURE QueueClient not implemented")
	case "OCI":
		// OCI Queue
		return nil, fmt.Errorf("OCI QueueClient not implemented")
	default:
		return nil, fmt.Errorf("unsupported cloud provider: %s", cloud)
	}
}
