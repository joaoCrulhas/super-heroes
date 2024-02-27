package db

import (
	"context"
)

type (
	Repository[T any] interface {
		Fetch(c context.Context) ([]T, error)
		FindByFilter(c context.Context, filter map[string]any) ([]T, error)
	}
)
