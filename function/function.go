package function

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/marciocadev/multicloud-go/cloud"
	"github.com/marciocadev/multicloud-go/function/aws"
	"github.com/marciocadev/multicloud-go/function/event"
)

// Run inicia a função serverless
func Run(h func(context.Context, *event.CloudRequest) (*event.CloudResponse, error)) {

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
