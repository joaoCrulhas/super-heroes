package presentation

import (
	"context"
)

type (
	Response[T any] struct {
		StatusCode int `json:"statusCode"`
		Data       T   `json:"data"`
	}
	Request[T any] struct {
		Body T `json:"body"`
	}
	Controller[Req any, Res any] interface {
		Handle(ctx context.Context, request Request[Req]) Response[Res]
	}
)
