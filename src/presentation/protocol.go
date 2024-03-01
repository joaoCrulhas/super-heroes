package presentation

import (
	"context"
)

type (
	Response[T any] struct {
		StatusCode uint  `json:"statusCode,omitempty"`
		Data       T     `json:"data,omitempty"`
		Error      error `json:"error,omitempty"`
	}
	Request[T any] struct {
		Body    T `json:"body"`
		Headers map[string][]string
		Params  map[string]string
		Query   map[string][]string
	}
	Controller[Req any, Res any] interface {
		Handle(ctx context.Context, request Request[Req]) Response[Res]
	}
)

func CreateResponse[T any](statusCode uint, data T, err error) Response[T] {
	return Response[T]{
		StatusCode: statusCode,
		Data:       data,
		Error:      err,
	}
}

const (
	SuperPowerFilter = "superpowers"
)
