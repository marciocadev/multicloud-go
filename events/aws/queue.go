package aws

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/marciocadev/multicloud-go/cloud"
)

// QueueMessage representa uma mensagem genérica que funciona para ambos os provedores
type QueueMessage struct {
	ID          string            // ID da mensagem
	Body        string            // Corpo da mensagem
	Attributes  map[string]string // Atributos da mensagem
	PublishTime time.Time         // Hora de publicação
	Source      string            // Origem da mensagem (tópico/fila)
}

// QueueEvent representa um evento que pode vir de qualquer provedor
type QueueEvent struct {
	Provider cloud.CloudProvider
	Messages []QueueMessage
}

func ParseAWSEvent(rawEvent interface{}) (*QueueEvent, error) {

	if sqsEvent, ok := rawEvent.(events.SQSEvent); ok {
		fmt.Println("A variável é do tipo SQSEvent:", sqsEvent)
	} else {
		fmt.Println("A variável não é do tipo SQSEvent")
	}

	switch evt := rawEvent.(type) {
	case events.SQSEvent:
		messages := make([]QueueMessage, len(evt.Records))
		for i, record := range evt.Records {
			messages[i] = QueueMessage{
				ID:          record.MessageId,
				Body:        record.Body,
				Attributes:  record.Attributes,
				PublishTime: getAWSTimestamp(record.Attributes["SentTimestamp"]),
				Source:      record.EventSourceARN,
			}
		}
		return &QueueEvent{
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
