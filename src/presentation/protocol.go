package presentation

import (
	"context"
)

type (
	Response[T any] struct {
		StatusCode int
		Body       T
	}
	Request[T any] struct {
		Body T
	}
	Controller[Req any, Res any] interface {
		Handle(ctx context.Context, request Request[Req]) Response[Res]
	}
)
