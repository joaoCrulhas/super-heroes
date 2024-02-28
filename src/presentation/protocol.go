package presentation

import (
	"context"
)

type (
	Response[T any] struct {
		StatusCode uint  `json:"statusCode"`
		Data       T     `json:"data"`
		Error      error `json:"error"`
	}
	Request[T any] struct {
		Body    T `json:"body"`
		Headers map[string][]string
		Params  map[string]string
	}
	Controller[Req any, Res any] interface {
		Handle(ctx context.Context, request Request[Req]) Response[Res]
	}
)
