package handler

import (
	"context"

	"github.com/marciocadev/multicloud-go/function/event" // Ensure this path is correct or replace with the actual package path
)

// HandlerFunc é o tipo de função que processa eventos cloud
type HandlerFunc func(context.Context, *event.CloudRequest) (*event.CloudResponse, error)
