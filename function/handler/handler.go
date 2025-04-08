package handler

import (
	"context"

	"github.com/marciocadev/multicloud-go/function/event"
)

// HandlerFunc é o tipo de função que processa eventos cloud
type HandlerFunc func(context.Context, *event.CloudRequest) (*event.CloudResponse, error)
