package function

import (
	"context"
)

// Wrapper é a interface para wrappers específicos de cada provedor
type Wrapper interface {
	Handle(ctx context.Context, event interface{}) (interface{}, error)
}
