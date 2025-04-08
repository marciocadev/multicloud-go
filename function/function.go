package function

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/marciocadev/multicloud-go/function/handler"
)

func RunFunction(handlerFunc interface{}) {
	myFunc := func(ctx context.Context, event interface{}) (interface{}, error) {
		wrapper := NewWrapper(handlerFunc.(handler.HandlerFunc))
		return wrapper.Handle(ctx, event)
	}

	lambda.Start(myFunc)
}
