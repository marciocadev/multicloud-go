package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSClient struct {
	TopicARN string
	Client   *sns.Client
}

func (s *SNSClient) Publish(ctx context.Context, messageBody string) error {
	_, err := s.Client.Publish(ctx, &sns.PublishInput{
		TopicArn: aws.String(s.TopicARN),
		Message:  aws.String(messageBody),
	})
	if err != nil {
		return err
	}
	return nil
}
