package aws

import (
	"github.com/aws/aws-lambda-go/events"

	cloud "github.com/marciocadev/multicloud-go/cloud"
	queue "github.com/marciocadev/multicloud-go/events/queue"
)

func ParseAWSEvent(rawEvent interface{}) (*queue.QueueEvent, error) {
	switch evt := rawEvent.(type) {
	case events.SQSEvent:
		messages := make([]queue.QueueMessage, len(evt.Records))
		for i, record := range evt.Records {
			messages[i] = queue.QueueMessage{
				ID:          record.MessageId,
				Body:        record.Body,
				Attributes:  record.Attributes,
				PublishTime: getAWSTimestamp(record.Attributes["SentTimestamp"]),
				Source:      record.EventSourceARN,
			}
		}
		return &queue.QueueEvent{
			Provider: cloud.AWS,
			Messages: messages,
		}, nil
		
	default:
		return nil, fmt.Errorf("tipo de evento AWS não suportado: %T", rawEvent)
	}
}

// Função auxiliar para converter timestamp AWS para time.Time
func getAWSTimestamp(timestamp string) time.Time {
	if ts, err := strconv.ParseInt(timestamp, 10, 64); err == nil {
		return time.Unix(ts/1000, 0)
	}
	return time.Time{}
}