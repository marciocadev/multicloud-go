package function

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/marciocadev/multicloud-go/cloud"
	"github.com/marciocadev/multicloud-go/function/aws"
	"github.com/marciocadev/multicloud-go/function/event"
)

// func RunFunction(handlerFunc interface{}) {
// 	myFunc := func(ctx context.Context, event interface{}) (interface{}, error) {
// 		wrapper := NewWrapper(handlerFunc.(handler.HandlerFunc))
// 		return wrapper.Handle(ctx, event)
// 	}

// 	lambda.Start(myFunc)
// }

// RunFunction inicia a função serverless
func RunFunction(h func(context.Context, *event.CloudRequest) (*event.CloudResponse, error)) {

	wrapper := func(ctx context.Context, rawEvent interface{}) (interface{}, error) {
		provider := cloud.CloudProvider(os.Getenv("CLOUD_PROVIDER"))

		switch provider {
		case cloud.AWS:
			// Criar um wrapper AWS (padrão)
			wrapper := &aws.AWSWrapper{Handler: h}
			// Processar o evento
			return wrapper.Handle(ctx, rawEvent)
		default:
			// Criar um wrapper AWS (padrão)
			wrapper := &aws.AWSWrapper{Handler: h}
			// Processar o evento
			return wrapper.Handle(ctx, rawEvent)

		}
	}

	lambda.Start(wrapper)
}
