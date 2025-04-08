package function

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/marciocadev/multicloud-go/function/event"
	"github.com/marciocadev/multicloud-go/function/handler"
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
	// Converter a função do usuário para o tipo HandlerFunc
	handlerFunc := handler.HandlerFunc(h)
	// Criar um wrapper que converte eventos AWS para nosso formato
	wrapper := NewWrapper(handlerFunc)
	// Iniciar a função Lambda
	lambda.Start(wrapper)
}
