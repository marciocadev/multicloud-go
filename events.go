package multicloud

import (
	"fmt"
	"os"

	"github.com/marciocadev/multicloud-go/cloud"
	aws "github.com/marciocadev/multicloud-go/events/aws"
)

// NewQueueEvent cria um novo CloudEvent baseado no provedor configurado
func NewQueueEvent(rawEvent interface{}) (*aws.QueueEvent, error) {
	provider := cloud.CloudProvider(os.Getenv("CLOUD_PROVIDER"))

	switch provider {
	case cloud.AWS:
		fmt.Printf("Conteúdo do evento: %+v\n", rawEvent)
		
		return aws.ParseAWSEvent(rawEvent)
	case cloud.OCI:
		// TODO: Implement OCI event parsing logic
		return nil, fmt.Errorf("OCI event parsing not implemented")
	default:
		return nil, fmt.Errorf("provedor de nuvem não suportado: %s", provider)
	}
}
