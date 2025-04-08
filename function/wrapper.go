package function

import (
	"context"
)

// Wrapper é a interface para wrappers específicos de cada provedor
type Wrapper interface {
	Handle(ctx context.Context, event interface{}) (interface{}, error)
}

// // NewWrapper cria um novo wrapper baseado no provedor configurado
// func NewWrapper(handler handler.HandlerFunc) Wrapper {
// 	provider := cloud.CloudProvider(os.Getenv("CLOUD_PROVIDER"))

// 	switch provider {
// 	case cloud.AWS:
// 		return &aws.AWSWrapper{Handler: handler}
// 	case cloud.GCP:
// 		return &aws.AWSWrapper{Handler: handler}
// 		// gcp.GCPWrapper{Handler: handler}
// 	default:
// 		// Retorna AWS como padrão se não especificado
// 		return &aws.AWSWrapper{Handler: handler}
// 	}
// }
