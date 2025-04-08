package events

import (
	cloud "github.com/marciocadev/multicloud-go/cloud"

	aws "github.com/marciocadev/multicloud-go/events/queue/aws"
)

// QueueMessage representa uma mensagem genérica que funciona para ambos os provedores
type QueueMessage struct {
	ID          string            // ID da mensagem
	Body        string            // Corpo da mensagem
	Attributes  map[string]string // Atributos da mensagem
	PublishTime time.Time         // Hora de publicação
	Source      string           // Origem da mensagem (tópico/fila)
}

// QueueEvent representa um evento que pode vir de qualquer provedor
type QueueEvent struct {
	Provider cloud.CloudProvider
	Messages []QueueMessage
}

// NewQueueEvent cria um novo CloudEvent baseado no provedor configurado
func NewQueueEvent(rawEvent interface{}) (*QueueEvent, error) {
	provider := cloud.CloudProvider(os.Getenv("CLOUD_PROVIDER"))
	
	switch provider {
	case AWS:
		return aws.ParseAWSEvent(rawEvent)
	case GCP:
		// return parseGCPEvent(rawEvent)
	default:
		return nil, fmt.Errorf("provedor de nuvem não suportado: %s", provider)
	}
}