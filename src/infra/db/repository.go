package db

import (
	"context"
)

type (
	Repository[T any] interface {
		Fetch(c context.Context) (map[int]T, error)
		FindByFilter(c context.Context, filter map[string][]string) (map[int]T, error)
		Create(c context.Context, item T) (T, error)
	}
)
