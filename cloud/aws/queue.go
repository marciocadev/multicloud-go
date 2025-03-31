package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSClient struct {
	QueueURL string
	Client   *sqs.Client
}

func (s *SQSClient) SendMessage(ctx context.Context, messageBody string) error {
	_, err := s.Client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(s.QueueURL),
		MessageBody: aws.String(messageBody),
	})
	if err != nil {
		return err
	}
	return nil
}
