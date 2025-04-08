package handler

import (
	"context"

	"github.com/marciocadev/multicloud-go/cloud"
)

// HandlerFunc é o tipo de função que processa eventos cloud
type HandlerFunc func(context.Context, *cloud.CloudRequest) (*cloud.CloudResponse, error)
