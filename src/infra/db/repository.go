package db

import (
	"context"
)

type (
	Repository[T any, R any] interface {
		Fetch(c context.Context) (T, error)
		FindByFilter(c context.Context, filter map[string][]string) (T, error)
		Create(c context.Context, item R) (R, error)
	}
)
