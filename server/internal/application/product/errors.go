package product

import (
	"errors"
)

var (
	ErrProductNotFound  = errors.New("product not found")
	ErrSKUAlreadyExists = errors.New("product SKU already exists")
)

func translateRepositoryError(err error) error {
	if errors.Is(err, ErrRepositoryNotFound) {
		return ErrProductNotFound
	}

	return err
}
