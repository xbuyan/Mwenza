package domain

import "context"

type Repository[T any] interface {
	Save(ctx context.Context, aggregate T) error
}
